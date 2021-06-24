package apiserver

import (
	"ShortURL/pkg/api"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	router     *mux.Router
	grpcclient api.ShortlinkClient
}

func Start(addr string, client api.ShortlinkClient) error {
	srv := newServer(client)
	srv.configureRouter()

	server := &http.Server{
		Addr:    addr,
		Handler: srv,
	}

	log.Println("Server start:", "https://"+server.Addr)

	return server.ListenAndServe()
}

func newServer(client api.ShortlinkClient) *server {
	s := &server{
		router:     mux.NewRouter(),
		grpcclient: client,
	}

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/home", s.handleHome()) //.Methods("GET")
	s.router.HandleFunc("/{token}", s.handleGet()).Methods("GET")
	s.router.HandleFunc("/create", s.handleCreate()).Methods("POST")
}

func (s *server) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, "Home page")
	}
}

func (s *server) handleCreate() http.HandlerFunc {
	type request struct {
		OriginURL string `json:"url"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
		}

		url := req.OriginURL
		short, err := s.grpcclient.Create(context.Background(), &api.OriginUrl{Url: url})
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}

		short.Url = r.Host + "/" + short.Url

		s.respond(w, r, http.StatusOK, short.Url)
	}
}

func (s *server) handleGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keys := mux.Vars(r)
		short := keys["token"]

		url, err := s.grpcclient.Get(context.Background(), &api.ShortUrl{Url: short})
		if err != nil {
			s.error(w, r, http.StatusNotFound, err)
		}

		http.Redirect(w, r, url.Url, http.StatusSeeOther)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

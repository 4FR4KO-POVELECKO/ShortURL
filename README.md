# ShortURL

##### API Сервис по сокращению ссылок

### Installation
```bash
git clone https://github.com/4FR4KO-POVELECKO/ShortURL.git
cd ./ShortURL
docker-compose up
```
or 


### Сервер
**gRPC** сервер написан на Go, proto3.

**Метов Create:** 
Примает ссылку, валидирует, генерирует токен формата [a-zA-Z0-9_], сохраняет в Redis, возвращает токен. Если ссылка не прошла валидацию возвращает ошибку. 

**Метов Get:**
Принимает токен, валидирует, проверяет в Redis, возвращает оригинальную ссылку. Если токен не найден возвращает ошибку.

### Tech
- Golang
- gRPC
- proto3
- Redis
- Docker

### Структура проекта
- **api**: файлы proto
- **cmd**: main файлы 
- **internal**: бд, модели
- **pkg**: grpc сервер
- **web**: html


### License

MIT

**Free Software, Hell Yeah!**

[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax)

   [dill]: <https://github.com/joemccann/dillinger>
   [git-repo-url]: <https://github.com/joemccann/dillinger.git>
   [john gruber]: <http://daringfireball.net>
   [df1]: <http://daringfireball.net/projects/markdown/>
   [markdown-it]: <https://github.com/markdown-it/markdown-it>
   [Ace Editor]: <http://ace.ajax.org>
   [node.js]: <http://nodejs.org>
   [Twitter Bootstrap]: <http://twitter.github.com/bootstrap/>
   [jQuery]: <http://jquery.com>
   [@tjholowaychuk]: <http://twitter.com/tjholowaychuk>
   [express]: <http://expressjs.com>
   [AngularJS]: <http://angularjs.org>
   [Gulp]: <http://gulpjs.com>

   [PlDb]: <https://github.com/joemccann/dillinger/tree/master/plugins/dropbox/README.md>
   [PlGh]: <https://github.com/joemccann/dillinger/tree/master/plugins/github/README.md>
   [PlGd]: <https://github.com/joemccann/dillinger/tree/master/plugins/googledrive/README.md>
   [PlOd]: <https://github.com/joemccann/dillinger/tree/master/plugins/onedrive/README.md>
   [PlMe]: <https://github.com/joemccann/dillinger/tree/master/plugins/medium/README.md>
   [PlGa]: <https://github.com/RahulHP/dillinger/blob/master/plugins/googleanalytics/README.md>
   
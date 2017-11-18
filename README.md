# purview

## POC for a view rendering microservice

the idea here is to completely decouple the `V` in MVC. Making the rendering and all templats 
independently renderable from disparate services to be composed up as per the clients desires. 
This brings the added benefit of providing Horizontal scalability as well as a clearly defined 
interface for interacting with your views.

well thats the idea anyway

## run the POC service
* this provides a few sample templates to test the rendering and calling the service
* configure it as you see fit if you like.

```
$ git clone github.com/xchapter7x/purview
$ cd purview/cmd/purview
$ VIEW_PATH=samples/views go run main.go
```

on another terminal query the service for rendered templates

```
$ curl -XPOST localhost:8080/v1/render/view/samples/views/index.tmpl -d'{"Testing":"blah", "Another":["something","else","completely"]}'
```

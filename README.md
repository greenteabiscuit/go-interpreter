# go-interpreter

Use docker

```
$ cd go-interpreter/ch01/src
$ docker-compose up build
$ docker-compose up
```

To run tests

```
(In another directory)
$ docker exec -it containername /bin/bash
(Inside docker container)
/go/src# echo $GOPATH
/go

/go/src/monkey# go test ./lexer/
ok  	monkey/lexer	0.006s
```

# go-interpreter

Use docker

```
$ cd go-interpreter/src
$ docker-compose up build
$ docker-compose up
```

To import modules

```
// in go directory of go container
root@87214538eec5:/go# go get github.com/greenteabiscuit/go-interpreter/monkey/
```

To run tests

```
(In another directory)
$ docker exec -it containername /bin/bash
(Inside docker container)
/go/src# echo $GOPATH
/go

/go/monkey# go test ./lexer/
ok  	monkey/lexer	0.006s
```

For REPL:

```
/go/monkey# go run main.go
Hello root! This is the Monkey programming language
Feel free to type in commands
>> let add = fn(x,y) {x+y};
{Type:LET Literal:let}
{Type:IDENT Literal:add}
{Type:= Literal:=}
{Type:FUNCTION Literal:fn}
{Type:( Literal:(}
{Type:IDENT Literal:x}
{Type:, Literal:,}
{Type:IDENT Literal:y}
{Type:) Literal:)}
{Type:{ Literal:{}
{Type:IDENT Literal:x}
{Type:+ Literal:+}
{Type:IDENT Literal:y}
{Type:} Literal:}}
{Type:; Literal:;}

```

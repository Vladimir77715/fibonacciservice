# fibonacci service



## Prerequisites

- **[Docker][]**:  latest version.

## Set up servers

To run the servers use [Docker-compose]

go to the project folder and run 

```console
docker-compose up
```
## Run tests and example

## Prerequisites

- **[Go][]**: any one of the **three latest major** [releases][go-releases].

## Run tests


In main folder run the command
 
 ```console
 go test ./... 
```

## Grpc client

When the server is started go to the exaples/grpc/client
Then run command 
 ```console
  go run client.go 
```

## Rest client

When the server is started go to the exaples/grpc/client
Then run command 
 ```console
  go run client.go 
```



[Docker]: https://docs.docker.com
[Docker-compose]: https://docs.docker.com/compose/


[Go]: https://golang.org
[go-releases]: https://golang.org/doc/devel/release.html

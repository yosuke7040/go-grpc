# go-grpc

```bash
cd api
protoc --go_out=../pkg/grpc --go_opt=paths=source_relative \
       --go-grpc_out=../pkg/grpc --go-grpc_opt=paths=source_relative \
       hello.proto
```


```bash
grpcurl -plaintext -d '{"name": "hsaki"}' localhost:8080 myapp.GreetingService.Hello

grpcurl -plaintext -d '{"name": "hsaki"}{"name": "a-san"}{"name": "b-san"}{"name": "c-san"}{"name": "d-san"}' localhost:8080 myapp.GreetingService.HelloBiStreams

```
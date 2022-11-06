> 例子: https://github.com/grpc/grpc-go/tree/master/examples/helloworld

```shell
cd grpc_hello
   
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./api/hello.proto    
```
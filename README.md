This is just a crappy app I’m going to use for practice, includes - grpc, go, rest, angular etc.
```
fire-grpc/
├── go.mod
├── proto/
│   └── fire.proto
├── server/
│   └── main.go
├── client/
│   └── main.go
├── gen/pb          # generated code
```
```
protoc --go_out=. --go-grpc_out=. proto/fire.proto
go mod tidy
go run server/main.go
go run client/main.go
```
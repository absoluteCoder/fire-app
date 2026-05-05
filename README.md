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
└── gateway/        # NEW
    └── main.go
```
Cmds
```
protoc \
  -I . \
  -I ./googleapis \
  --go_out=. \
  --go-grpc_out=. \
  --grpc-gateway_out=. \
  proto/fire.proto
go mod tidy
go run server/main.go
go run client/main.go
go run gateway/main.go # Gateway converts JSON → gRPC request
```

Request
```
 curl -X POST http://localhost:8080/v1/fire/ignite \
  -H "Content-Type: application/json" \
  -d '{"fuel":"gas"}'
```
Res
```
{"message":"🔥 Fire ignited using gas!"}
```
Структуры \
`protoc --go_out=. --go_opt=paths=source_relative proto/m.proto`

Сервер и методы \
`protoc -I proto proto/m.proto  --go-grpc_out=proto/`

запуск тестов \
`go test -v .\tests\grpc.`
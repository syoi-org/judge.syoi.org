build:
    go build -o build/judgectl .

test:
    go test ./...

generate:
    go generate ./...

migrate *ARGS: build
    build/judgectl migrate {{ARGS}}

server *ARGS: build
    build/judgectl server {{ARGS}}

worker *ARGS: build
    build/judgectl worker {{ARGS}}

build:
    go build -o build/judy .

test:
    go test ./...

generate:
    go generate ./...

migrate *ARGS: build
    build/judy migrate {{ARGS}}

server *ARGS: build
    build/judy server {{ARGS}}

worker *ARGS: build
    build/judy worker {{ARGS}}

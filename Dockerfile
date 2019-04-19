FROM golang:latest

COPY hello-world.go /src/
WORKDIR /src

CMD ["go", "run", "hello-world.go"]
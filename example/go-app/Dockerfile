FROM golang:alpine AS builder
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go mod init github.com/nevermosby/go-hello-arch && go build -o hello-arch .

FROM alpine
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/hello-arch .
CMD ["./hello-arch"]
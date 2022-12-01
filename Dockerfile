FROM golang:1.18-buster

RUN go version
ENV GOPARH=/

RUN mkdir form-to-mail-service
WORKDIR /form-to-mail-service

COPY ./ ./

RUN go mod download
RUN go build -o form-to-mail-service ./cmd/main.go

CMD ["./form-to-mail-service"]
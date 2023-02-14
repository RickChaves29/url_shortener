FROM golang:1.20.0-alpine3.17 AS dev

WORKDIR /app/url_shortener

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY cmd ./cmd
COPY internal ./internal
COPY utils ./utils

RUN go build -o app ./cmd/main.go

ENV CONNECT_DB=${CONNECT_DB}

EXPOSE 3030

CMD [ "go", "run", "cmd/main.go" ]

FROM alpine:3.17.2 AS prod
WORKDIR /app/url_shortener

COPY --from=dev /app/url_shortener/app .

CMD [ "./app" ]


FROM golang:1.17-alpine as builder

RUN apk --no-cache add ca-certificates

WORKDIR /build

COPY main/go.mod main/go.sum ./

RUN go mod download

COPY main/ .

RUN CGO_ENABLED=0 go build -o /app .

EXPOSE 80

ENTRYPOINT ["/app"]
FROM golang:1.16.8-alpine as builder
WORKDIR /app
ADD go.mod go.sum ./
RUN go mod download
ADD . ./

RUN apk update && apk add date
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -a -tags netgo -ldflags "-s -w -extldflags '-static'" \
    -o stonks-server ./cmd/stonks-server

FROM scratch as app
ENV PORT=8000
WORKDIR /app
COPY --from=builder /app/stonks-server /app/stonks-server
ENTRYPOINT ["./stonks-server"]
EXPOSE $PORT
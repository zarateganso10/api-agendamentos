FROM golang:1.21 as builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build cmd/api/main.go

FROM scratch

COPY --from=builder /app/main /main
ENTRYPOINT [ "/main" ]

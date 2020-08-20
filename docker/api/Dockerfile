# builder
FROM golang:latest AS builder
WORKDIR /api
COPY . .
RUN make tidy
RUN env GOARCH=386 GOOS=linux CGO_ENABLED=0 make api

# runner
FROM alpine
WORKDIR /app
COPY --from=builder /api/bin/powder_api .
COPY --from=builder /api/config/api.json .
COPY --from=builder /api/config/server.crt .
COPY --from=builder /api/config/server.key .
ENTRYPOINT ["/app/powder_api", "/app/api.json"]

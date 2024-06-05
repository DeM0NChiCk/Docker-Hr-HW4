FROM golang:1.21-alpine as builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o ./bin/job-service ./cmd/job/main.go

FROM alpine
COPY --from=builder /build/bin/job-service /job-service
COPY --from=builder /build/config/config.yml /config/config.yml

EXPOSE 5000

ENTRYPOINT ["/job-service"]

FROM golang:1.22-alpine as builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o ./bin/cv-service ./cmd/cv/main.go

FROM alpine
COPY --from=builder /build/bin/cv-service /cv-service
COPY --from=builder /build/config/config.yml /config/config.yml

EXPOSE 5001

ENTRYPOINT ["/cv-service"]
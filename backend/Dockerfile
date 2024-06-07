FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go generate ./datastore/ent/ent/.

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/cine-rest/.

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

ENTRYPOINT ["./main"]
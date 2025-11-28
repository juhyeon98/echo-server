# build
FROM golang AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /app/server ./cmd/main.go

# run
FROM alpine:latest
ENV TZ=Asia/Seoul
COPY --from=builder /app/server /usr/local/bin/server
CMD ["server"]
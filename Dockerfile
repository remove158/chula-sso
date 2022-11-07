FROM golang:1.19.3-bullseye AS builder

WORKDIR /app

COPY . /app

RUN go build -o server ./cmd/main.go


FROM debian:bullseye-slim AS prod

COPY --from=builder /app/templates/index.html ./templates/index.html

COPY --from=builder /app/server ./server

CMD ["./server"]



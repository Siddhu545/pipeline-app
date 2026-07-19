# ---- Stage 1: build the app ----
FROM golang:1.26 AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY main.go ./
RUN go build -o server main.go

# ---- Stage 2: run the app ----
FROM debian:stable-slim
WORKDIR /app
COPY --from=build /app/server ./server
EXPOSE 8080
CMD ["./server"]
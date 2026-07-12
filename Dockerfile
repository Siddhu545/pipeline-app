FROM golang:1.26 AS build
WORKDIR /app
COPY go.mod ./
COPY main.go ./
RUN go build -o server main.go

FROM debian:stable-slim
WORKDIR /app
COPY --from=build /app/server ./server
EXPOSE 8080
CMD ["./server"]
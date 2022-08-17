FROM golang:1.18 AS builder
WORKDIR /movies-go
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o movies-go ./cmd/api

FROM scratch
COPY --from=builder ./movies-go/movies-go /usr/bin/movies-go
EXPOSE 4001
ENTRYPOINT ["/usr/bin/movies-go"]
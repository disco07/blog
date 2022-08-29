FROM golang:1.18 AS builder
WORKDIR /blog-go
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o blog-go ./cmd/api

FROM scratch
COPY --from=builder ./blog-go/blog-go /usr/bin/blog-go
EXPOSE 4001
ENTRYPOINT ["/usr/bin/blog-go"]
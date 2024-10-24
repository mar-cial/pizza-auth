FROM golang:1.23 AS builder

WORKDIR /auth

COPY go.mod go.sum ./ 

RUN go mod download 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

FROM gcr.io/distroless/base-debian12

COPY --from=builder /auth/main /

EXPOSE 8080

ENTRYPOINT [ "./main" ]


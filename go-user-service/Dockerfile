FROM golang:1.18-alpine AS builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go mod download
RUN go build 

FROM golang:1.18-alpine
COPY --from=builder /build/go-user-service /app/
ADD config/envs/app.env /app/config/envs/app.env
WORKDIR /app
CMD ["./go-user-service"]
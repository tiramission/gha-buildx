FROM golang:1.25 AS builder
WORKDIR /app
COPY . .
RUN go build -o /bin/main .

FROM ubuntu:24.04
COPY --from=builder /bin/main /bin/main
RUN /bin/main
CMD ["/bin/main"]
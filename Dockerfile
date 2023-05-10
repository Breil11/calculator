FROM golang:1.16-alpine3.13

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /calculator

FROM gcr.io/distroless/base-debian10
COPY --from=0 /calculator /
CMD ["/calculator"]

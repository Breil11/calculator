FROM golang:1.19-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY calculator/ .

RUN go build -o /super_calculator

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /super_calculator /super_calculator

USER nonroot:nonroot

ENTRYPOINT ["/super_calculator"]

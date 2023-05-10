FROM golang:alpine AS build

WORKDIR /app
COPY . .

RUN apk update && \
    apk add git && \
    go mod tidy && \
    go build -o /super_calculator

FROM gcr.io/distroless/base-debian10
COPY --from=build /super_calculator /

CMD ["/super_calculator"]

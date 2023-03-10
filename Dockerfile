# build
FROM golang:1.19-alpine as build

RUN apk update && apk add gcc libc-dev

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY ./contracts ./contracts
COPY ./dydx ./dydx

COPY *go .

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o kwent .

# deploy
FROM scratch

WORKDIR /

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=build /app/kwent ./

ENTRYPOINT [ "./kwent", "-docker" ]

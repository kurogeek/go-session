FROM golang:1.12-alpine as build-env

RUN apk add --update --no-cache ca-certificates git

RUN mkdir /go-session
WORKDIR /go-session

# COPY ./app/go.mod .
# COPY ./app/go.sum .

# RUN go mod download

COPY ./app .

RUN go get -u -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/go-session

FROM scratch
COPY --from=build-env /go/bin/go-session /go/bin/go-session

CMD [ "/go/bin/go-session" ]
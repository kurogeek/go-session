FROM golang:1.12

RUN mkdir -p /app-dir
WORKDIR /app-dir

# COPY ./app .
VOLUME [ "/app-dir" ]

# USER panupong
# RUN go get -u -v
# RUN go build

ENTRYPOINT go get -u -v && go build && ./go-session
# CMD [ "./go-session" ]
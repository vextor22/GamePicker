FROM golang:1.14.2 AS build_base
RUN mkdir /app 
RUN mkdir /app/steamconnector
WORKDIR /app
COPY go.mod .
COPY ./steamconnector/go.mod ./steamconnector/.
COPY go.sum . 
RUN go mod download

FROM build_base AS server_builder
ADD . /app/ 
RUN go build -o app_main .

FROM ubuntu:xenial as app_server
COPY --from=server_builder /app/app_main /bin/server
COPY --from=server_builder /app/config.yml /bin/config.yml
ENTRYPOINT ["/bin/server"]

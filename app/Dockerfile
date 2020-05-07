FROM golang:1.14.2 AS build_base
RUN mkdir /app 
RUN mkdir /app/steam-connector
WORKDIR /app
COPY go.mod .
COPY ./steam-connector/go.mod ./steam-connector/.
COPY go.sum . 
RUN go mod download

FROM build_base AS server_builder
ADD . /app/ 
RUN go build -o app_main .
RUN ls -lh
RUN pwd

FROM ubuntu:xenial as app_server
COPY --from=server_builder /app/app_main /bin/server
COPY --from=server_builder /app/config.yaml /bin/config.yaml
ENTRYPOINT ["/bin/server"]

FROM golang

COPY service.go /opt/
RUN ls /opt
WORKDIR /opt/
RUN go build service.go

EXPOSE 8080/tcp

ENTRYPOINT ["/opt/service"]


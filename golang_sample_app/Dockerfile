FROM golang:1.14.13

ADD ./server_gin /opt/server_gin
RUN chmod +x /opt/server_gin

WORKDIR /opt/

# EXPOSE 8080

CMD ["./server_gin"]

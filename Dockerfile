FROM scratch

WORKDIR /tmp/server
COPY . /tmp/server

EXPOSE 8088
CMD ["./demo_server"]
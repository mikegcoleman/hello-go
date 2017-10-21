FROM scratch

COPY hello_go /

EXPOSE 8080

ENTRYPOINT [ "hello_go" ]
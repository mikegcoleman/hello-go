FROM scratch

COPY ./hello_world /

EXPOSE 8080

ENTRYPOINT [ "hello_world" ]
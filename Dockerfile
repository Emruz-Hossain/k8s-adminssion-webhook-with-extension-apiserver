FROM ubuntu
COPY kbc ./bin/kbc
ENTRYPOINT ["/bin/kbc"]
EXPOSE 8443
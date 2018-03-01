FROM busybox:glibc
COPY kbc ./bin/kbc
ENTRYPOINT ["/bin/kbc"]
EXPOSE 8443
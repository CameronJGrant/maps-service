FROM alpine
COPY build/server /
ENTRYPOINT ["/server"]
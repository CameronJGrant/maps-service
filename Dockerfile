FROM alpine
COPY build/server /
CMD ["/server"]
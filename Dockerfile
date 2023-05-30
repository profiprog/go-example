FROM scratch
COPY go-example /
ENTRYPOINT ["go-example"]

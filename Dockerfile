FROM golang:1.17

WORKDIR /go/src

ENV PATH="/go/bin:${PATH}"

CMD ["tail", "-f", "/dev/null"]
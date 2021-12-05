FROM golang:1.17-alpine as BASE
WORKDIR /src
COPY go/ .
RUN go build -o /bin/aoc

FROM scratch
COPY --from=BASE /bin/aoc /bin/aoc
COPY go/input.txt .
ENTRYPOINT [ "/bin/aoc" ]

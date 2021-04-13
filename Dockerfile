FROM golang:alpine

WORKDIR $GOPATH/src/GoBreakingBad

COPY . .
EXPOSE 8080

CMD ["go", "run", "."]
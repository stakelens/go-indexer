FROM golang:1.22.2-alpine

WORKDIR /app

COPY . /app

RUN go build -o .

RUN chmod +x staking-indexer

CMD ["./staking-indexer"]
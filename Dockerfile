FROM golang:1.21 as builder


WORKDIR /app

COPY go.* ./
 
RUN go mod download

COPY ./infra/ ./infra
COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /rinha-go-fx

CMD ["/rinha-go-fx"]

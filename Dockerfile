FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

ENV PORT=1325

EXPOSE 1325

CMD [ "./main","serve" ]
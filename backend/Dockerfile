FROM golang:1.21.4

WORKDIR /app

COPY . .

RUN go build -o go-app

EXPOSE 3001

CMD ["./go-app"]

FROM golang:1.20

WORKDIR /app

COPY . ./

RUN go mod download

EXPOSE 8080

CMD ["go", "run", "."]
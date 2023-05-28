# bitcoin-api
Golang-based API for getting current Bitcoin price in UAH. To implement RESTful API I used popular framework Gin. API includes '/rate', '/subscribe' and '/sendEmails' endpoints.

# How to run API

## Running in Docker

If you want to manually run the container all you need is type in:
```
docker build -t api .
docker run --name api -d --rm -p 8080:8080 api 
```
This commands will build image and run container on the localhost:8080.

Also you can use `docker compose`:

```
docker compose up
```


## Running using go compiler

You also can run the program by using just:
```
go mod download
go run .
```

# About bitcoin-api

## Storing Emails

To store emails I used go built-in package **os** to create, open and write files. To make code cleaner I implemented all logic in `fileManager.go` file located in folder `./controlers`. Eamils are stored in the file named **emails.db**, but you can easily change the name of the file in the **.env**. When someone tries to subscribe using `/subscribe` endpoint, I am validating the email that was sent and also checking whether the email is already in the database.

## Caching the files

## Sending emails
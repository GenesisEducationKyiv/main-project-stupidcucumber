# bitcoin-api
Golang-based API for getting current Bitcoin price in UAH. To implement RESTful API I used popular framework Gin. API includes '/api/rate', '/api/subscribe' and '/api/sendEmails' endpoints.

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


## Getting Bitcoin price
To get bitcoin Price I used Binance public API. I also cached the price in the file **.cache** along with timestamp when the price were changed. BITCOIN price automatically converts into UAH. If something goes wrong API returns status 400, otherwise it returns price and status 200. 

## Storing Emails

To store emails I used go built-in package **os** to create, open and write files. To make code cleaner I implemented all logic in `fileManager.go` file located in folder `./controlers`. Eamils are stored in the file named **emails.db**, but you can easily change the name of the file in the **.env**. When someone tries to subscribe using `/api/subscribe` endpoint, I am validating the email that was sent and also checking whether the email is already in the database.

API returns status 200 if email is valid, otherwise it returns status 409.

## Sending emails
When `/api/sendEmails` endpoint is used system extracts all subscribed emails from `email.db` and sends them to all of them. Service generating smagnificent template to make API look more serious. I used 'text/template' and Gomail package to generate and send emails. I also created a specific account and APP KEY to get access to the Gmail application (however there is a private version of the service, where the access is implemented using OAuth.v2). All creadentials is stored in `.env` so please do not use them unintentionaly, even though this account is not neccessary to me or other people.

## Caching the files

To make things faster and use Binance API as little as possible I implemented cache system, which tracks when the last time I updated price and updates if price is too old. Each time user requests `/api/rate` or `/api/sendEmails` it checks the price and timestamp.

## Chain responsibility

To make API more stable, I request price from Coingeko API if Binance API is failed.
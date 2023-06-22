# Stage 0: building an executable
FROM golang:1.20 AS building
WORKDIR /app
COPY ./go.mod ./go.mod
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o app

# Copying and running an executable
FROM alpine:3.18
WORKDIR /root/
COPY --from=building /app/app ./ 
COPY --from=building /app/.env ./
COPY --from=building /app/templates ./templates

EXPOSE 8080
CMD [ "./app" ]
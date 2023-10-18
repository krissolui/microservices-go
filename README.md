# Microservices in Go

### Content:

-   front end
    -   Run as a microservice, proxied through a Caddy server
-   broker
    -   Support sending RPC requests to logger service
    -   Support sending gRPC requests to logger service
-   authentication service
    -   Postgres
-   logger service
    -   MongoDB
    -   Accept RPC requests
    -   Accept gRPC requests
-   mail service
    -   mailhog
-   listener service
    -   RabbitMQ
        > Broker will push message to RabbitMQ instead of calling the logger service directly

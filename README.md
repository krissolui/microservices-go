# Microservices in Go

Microservices written in Go following the Udemy course.

### Course: [Working with Microservices in Go (Golang)](https://www.udemy.com/course/working-with-microservices-in-go/)

### Content:

-   front end
-   broker
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

# ms-arch

A project to demonstrate microservie architecture. This project includes features like
- CORS restrictions
- Service discovery
- Ditributed tracing
- Caching

# Running
**To run the project manually**, you'll need the following prerequisites
- MySQL
- Redis
- Zipkin
- Maven

Run all the projects & order of execution should be
`api-gateway` => `auth-service` => `product-service`

**note** : update `application.yml` with local confuguration before running.

**To run this project with** [**Docker**](https://www.docker.com/)

First build the project from `ms-architecture` folder using
`mvn clean install`

Then, 

`docker compose up --build -d`<br />
`-d` is for running in detached mode & it's optional

# To-do
Need to implement
- Kafka streams
- Distributed locks

# RentACar

Yolcu360's Take-home case study

### The case is now complete, to see final words read the end

## How to run

1-) Run docker compose file

```console
docker build -t yolcuapp .
docker-compose -f docker-compose.app.yml up
```

2-) To see all the end-points go to the swagger end-point in the app

```URL
localhost:3000/swagger
```

## How to run tests (a bit tricky :))
* Run the only the database service from the docker-compose.yml file
* Then run this command
```console
go test -v ./...
```

## Tools of Choice

* Language & Framework -- Go & Fiber
* Database -- Postgresql
* ORM -- Gorm
* Api Documentation -- fiber-swagger (yaml file can be found in /docs folder)
* Validation -- Validator (comes with fiber)
* Testing -- testify

## Database Design

![yolcu360_DB_DIAGRAM drawio (1)](https://user-images.githubusercontent.com/29152340/175826289-618e7e8f-a7e2-43de-a07c-f7c1c18f53da.svg)


# Final Words
### What did I do good
* Most of the technologies I have used technologies; I am not familiar with them, I wanted to learn new thing whilst coding this project
* My database design (normalizing) is okay

### What did I do not-good
* Unit-tests lacks quality
* Code organization is kind of sluggish

### What would I add if I had more time
* Write more elegant unit tests, also tests inbound responses
* Create seperate test environment with docker-compose (I have tried this but it would take alot of time to figure ut out)

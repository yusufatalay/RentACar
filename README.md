# RentACar

Yolcu360's Take-home case study

### Updated will come as I complete functional requirements

## How to run

1-) Run docker compose file

` ``` `

```console
docker-compose -f docker-compose.app.yml up
```

2-) To see all the end-points go to the swagger end-point in the app

```URL
localhost:3000/swagger
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

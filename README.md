# api-server-boilerplate
A docker-ready golang api server boilerplate that helps to set up project faster.

We have been developing Golang microservice for a few years now. Thought it will be nice to create a boilerplate from what has been working for us, and things we have learnt so far. This is by no means the "best" or "conventional" layout for all purposes, you should alway evaluate the nature and need of your project before adopting other technologies, but please feel free to take any part of the code that you might find useful.

## Get started
### Development on docker-compose
Go to the root the project folder and
```
docker-compose build
docker-compose up
```
After the docker containers start to run, the server will be in a hot reload mode and listen to the file changes.


### Development on native os
```
cp .env.example .env.local
```
change env variables to your local development enviornment.
```
./scripts/run-migrate up 1
./scripts/run-web
```

### Example endpoints
User sign up
```  
curl --location --request POST 'localhost:8081/api/v1/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "first_name": "test",
    "last_name": "test",
    "email": "test@gmail.com",
    "password": "123",
    "confirm_password": "123"
}'
```
User login
```  
curl --location --request POST 'localhost:8081/api/v1/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email":"test@gmail.com",
    "password": "123"
}'
```
After that you should be able to get the JWT token


### Swtich between development mode and production mode
Change ENV=local to ENV=production in .env.local file
```
./scripts/migrate-run.sh
./scripts/web-run.sh
```

## Database migrate
We are using [golang-migrate](https://github.com/golang-migrate/migrate) to manage the database schema, so when you need to add a new model to the code base, you could use to add a new schema file
```
make migrate-create name=alter_user_table_add_image_url
```

## Handling context
You might find code looks like throughout the codebase.
```
type AuthServiceWithContext func(ctx context.Context) AuthService
```
It is basically a factory function that return an actually object that concerns with the context. And you might ask why couldn't we pass context object as an argument to all the functions, the reason is really about not be repetitive and keep the interface clean by removing context

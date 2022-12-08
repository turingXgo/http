# Simple User Web App

## Usage

- Simply create user by sending `POST` request with user data
```bash
curl http://localhost:8081/ -d '{"name": "Kid", "age": 12}'
```
- Retrieve user information by user `name`
```bash
curl http://localhost:8081/user -d '{"name": "AZ"}'
```

## Run

- Run `postgresql` database using docker or any other options:
```
docker run --name postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres:14.2
```
- Connect to the postgres and create new database with `data` table:
```
psql -h localhost -p 5432 -U demo -W
create database data;
\c data
create table data (name text, age int);
```
- Run application
```
go run cmd/main.go
```
- Send request
- Profit 😎

Facing problems? Contact your mentor.
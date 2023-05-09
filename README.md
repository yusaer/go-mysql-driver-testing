# go-mysql-driver-testing

`go-sql-driver/mysql` for library validation  
Mainly for identifying the cause of `invalid connection` occurrence.

## Start Up 

```zsh
docker-compose up -d
```

## How to log in to the database

```zsh
docker-compose run cli
```

## Settings

### Create Table

```sql
CREATE TABLE users (id int, name varchar(10));
```

### For Error

Since `wait_timeout` is set to 3 seconds, `Sleep(1 * time.Second)`, you can make it throw out an `invalid connection` by setting `1` to `3` or more.

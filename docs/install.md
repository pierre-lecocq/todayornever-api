# Installation guide

## Prerequisites

To compile and run `todayornever-api`, you need to install the following software:

- [The GO programming language](https://go.dev/)
- [SQLite3](https://www.sqlite.org/)

Additionally, for managing database migrations, install `golang-migrate` by following [the instructions in this comment](https://github.com/golang-migrate/migrate/issues/670#issuecomment-1118029997)

## Create the database

The application stores the data in a `sqlite3` database.

To create the database and the schema, use:

```sh
touch mydatabase.db
migrate -path ./app/migrations -database sqlite3://./mydatabase.db up
```

## Create a configuration file

A configuration file is needed to pass to the program some values that may vary according to the running environment.

Simply create a `.env` file with the content described in [the dedicated documentation](./configuration.md)

## Compiling

```sh
make
```

## Running

```sh
./todayornever-api
```

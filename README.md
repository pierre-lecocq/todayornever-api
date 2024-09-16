# Today or never

<p align="center">
Focus on what you can do today!
</p>

---

## What is it?

This program is a service to manage your tasks, focusing on what you can do today.

It is a REST API returning including the following features:

- JSON responses
- JWT authentication
- CORS protection
- Rate limiting
- Request logging

## Install

### Prerequisites

To compile or run it, you need the following software installed:

- [The GO programming language](https://go.dev/)
- [SQLite3](https://www.sqlite.org/)

### Setting up the database

Database migrations are handled by [golang-migrate](https://github.com/golang-migrate/migrate).

First, install the tool by following [the instructions in this comment](https://github.com/golang-migrate/migrate/issues/670#issuecomment-1118029997)

Then, create a database and then run migrations with:

```
migrate -path ./app/migrations -database <DSN> up
```

### Creating a config file

A configuration file is needed to pass to the program some values that may vary according to the running environment.

Simply create a `.env` file with the following content (please adapt the values):

```toml
SERVICE_HOST=localhost
SERVICE_PORT=8080

LOGGER_LEVEL=3

AUTH_ISSUER=todayornever-api
AUTH_SECRET=<REDACTED>
AUTH_EXPIRES=1

DATABASE_ENGINE=sqlite3
DATABASE_DSN=./todayorneverd.db
```

Of course, all these values are overridable when launching the program.

### Compiling the binary

```sh
make
```

## Running the program

```sh
env $(cat .env | xargs) ./todayornever-api
```

---

## Data model

A `task` is represented by the following JSON object:

```json
{
  "id": int,
  "user_id": int,
  "title": string,
  "state": string,
  "due_at": datetime,
  "created_at": datetime,
  "updated_at": datetime,
  "position": int,
  "overdue": boolean
}
```

## Authentication

Authentication is based on JWT tokens. The steps to get on of these are:

- (first time) Register a new user with `POST /signup` giving an `email`, a `username` and a `password` in the JSON body
- Login with an existing user with `POST /login` giving the `email` and `password` in the JSON body
- Get the token from the response, if successful

In the next queries, to use the token, include it the `Authorization` header as a `Bearer`.

## Endpoints

| Method        | Path          | Usage                       | Auth required? | Required parameters                                             |
|---------------|---------------|-----------------------------|----------------|-----------------------------------------------------------------|
| `GET`         | `/`           | Greeting message            | N              |                                                                 |
| `POST`        | `/signup`     | Register a new user         | N              | `email` (valid email), `username` (string), `password` (string) |
| `POST`        | `/login`      | Login with an existing user | N              | `email` (valid email), `password` (string)                      |
| `GET`         | `/tasks`      | List of the tasks           | Y              |                                                                 |
| `POST`        | `/tasks`      | Create a new task           | Y              | `title` (string)                                                |
| `GET`         | `/tasks/{id}` | Fetch a task                | Y              |                                                                 |
| `PATCH`/`PUT` | `/tasks/{id}` | Update a task               | Y              | Any value from the data model task object                       |
| `DELETE`      | `/tasks/{id}` | Delete a task               | Y              |                                                                 |

## Examples

```sh
# Greeting message
curl localhost:8080/

# Signup
curl -X POST -H 'Content-Type: application/json' -d '{"email": "me@example.com", "username": "me", "password": "Sup3rStr0ngP4ass!"}' localhost:8080/signup

# Login
curl -X POST -H 'Content-Type: application/json' -d '{"email": "me@example.com", "password": "Sup3rStr0ngP4ass!"}' localhost:8080/login

# List tasks
curl -H 'Authorization: Bearer <TOKEN>' localhost:8080/tasks

# Create a task
curl -X POST -H 'Authorization: Bearer <TOKEN>' -H 'Content-Type: application/json' -d '{"title": "My new task"}' localhost:8080/tasks

# Fetch a task
curl -H 'Authorization: Bearer <TOKEN>' localhost:8080/tasks/1

# Update a task
curl -X PUT -H 'Authorization: Bearer <TOKEN>' -H 'Content-Type: application/json' -d '{"title": "My edited task"}' localhost:8080/tasks/1

# Delete a task
curl -X DELETE -H 'Authorization: Bearer <TOKEN>' localhost:8080/tasks/1
```

<details>
  <summary>Click here to see a restclient test suite (Emacs required)</summary>

```
# -*- restclient -*-

:host = http://localhost:8080
:token = xxx

# Index
GET :host/

# Signup
POST :host/signup
Content-Type: application/json

{"email": "me@example.com", "username": "me", "password": "Sup3rStr0ngP4ass!"}

# Login
POST :host/login
Content-Type: application/json
-> jq-set-var :token .token

{"email": "me@example.com", "password": "Sup3rStr0ngP4ass!"}

# List tasks
GET :host/tasks
Authorization: Bearer :token

# Fetch task
GET :host/tasks/1
Authorization: Bearer :token

# Create task
POST :host/tasks
Authorization: Bearer :token
Content-Type: application/json

{"title": "My first task"}

# Update task
PATCH :host/tasks/1
Authorization: Bearer :token
Content-Type: application/json

{"title": "My edited task"}

# Delete task
DELETE :host/tasks/1
Authorization: Bearer :token

```
</details>

---

## License

<details>
  <summary>This software is distributed under the MIT license</summary>

```
MIT License

Copyright (c) [year] [fullname]

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
</details>

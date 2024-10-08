# API specs

The API is based on commonly used REST principles.

The authentication is handled by [JWT tokens](https://jwt.io/introduction) passed in the `Authorization` header. The workflow [can be found here](./authentication.md).

## Endpoints

| Method        | Path             | Usage                       | Auth required?             | Required parameters                                             |
|---------------|------------------|-----------------------------|----------------------------|-----------------------------------------------------------------|
| `GET`         | `/`              | Greeting message            | No                         |                                                                 |
| `POST`        | `/signup`        | Register a new user         | No                         | `email` (valid email), `username` (string), `password` (string) |
| `POST`        | `/login`         | Login with an existing user | No                         | `email` (valid email), `password` (string)                      |
| `GET`         | `/projects`      | List of the projects        | [Yes](./authentication.md) |                                                                 |
| `POST`        | `/projects`      | Create a new project        | [Yes](./authentication.md) | `title` (string)                                                |
| `GET`         | `/projects/{id}` | Fetch a project             | [Yes](./authentication.md) |                                                                 |
| `PATCH`/`PUT` | `/projects/{id}` | Update a project            | [Yes](./authentication.md) | Any value from the data model project object                    |
| `DELETE`      | `/projects/{id}` | Delete a project            | [Yes](./authentication.md) |                                                                 |
| `GET`         | `/tasks`         | List of the tasks           | [Yes](./authentication.md) |                                                                 |
| `POST`        | `/tasks`         | Create a new task           | [Yes](./authentication.md) | `title` (string)                                                |
| `GET`         | `/tasks/{id}`    | Fetch a task                | [Yes](./authentication.md) |                                                                 |
| `PATCH`/`PUT` | `/tasks/{id}`    | Update a task               | [Yes](./authentication.md) | Any value from the data model task object                       |
| `DELETE`      | `/tasks/{id}`    | Delete a task               | [Yes](./authentication.md) |                                                                 |

## Queries examples

```sh
# Greeting message
curl localhost:8080/

# Signup
curl -X POST -H 'Content-Type: application/json' -d '{"email": "me@example.com", "username": "me", "password": "Sup3rStr0ngP4ass!"}' localhost:8080/signup

# Login
curl -X POST -H 'Content-Type: application/json' -d '{"email": "me@example.com", "password": "Sup3rStr0ngP4ass!"}' localhost:8080/login

# List projects
curl -H 'Authorization: Bearer <TOKEN>' localhost:8080/projects

# Create a project
curl -X POST -H 'Authorization: Bearer <TOKEN>' -H 'Content-Type: application/json' -d '{"name": "My new project"}' localhost:8080/projects

# Fetch a project
curl -H 'Authorization: Bearer <TOKEN>' localhost:8080/projects/1

# Update a project
curl -X PUT -H 'Authorization: Bearer <TOKEN>' -H 'Content-Type: application/json' -d '{"name": "My edited project"}' localhost:8080/projects/1

# Delete a project
curl -X DELETE -H 'Authorization: Bearer <TOKEN>' localhost:8080/projects/1

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

# List projects
GET :host/projects
Authorization: Bearer :token

# Fetch project
GET :host/projects/1
Authorization: Bearer :token

# Create project
POST :host/projects
Authorization: Bearer :token
Content-Type: application/json

{"name": "My first project"}

# Update project
PATCH :host/projects/1
Authorization: Bearer :token
Content-Type: application/json

{"name": "My first edited project"}

# Delete project
DELETE :host/projects/1
Authorization: Bearer :token

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

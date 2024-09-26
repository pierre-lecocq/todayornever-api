# Data model

## Project object

A `project` is represented by the following JSON object:

```json
{
  "id": int,
  "user_id": int,
  "name": string,
  "description": string,
  "created_at": datetime,
  "updated_at": datetime,
  "position": int,
  "overdue": boolean
}
```

## Task object

A `task` is represented by the following JSON object:

```json
{
  "id": int,
  "user_id": int,
  "project_id": int,
  "parent_task_id": int,
  "title": string,
  "state": string,
  "due_at": datetime,
  "created_at": datetime,
  "updated_at": datetime,
  "position": int,
  "overdue": boolean
}
```

## Paginated list of tasks

## Authentication object

## Error

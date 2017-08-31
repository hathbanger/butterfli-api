# Users

### POST/users

Request

```
{
    username: <string>,
    password: <string>,
}
```

Response

```
{
    <User>
}
```

### GET/users/:id

Response

```
{
    <User>
}
```

### PUT/users/:id

Request

```
{
    username: <string>,
    password: <string>,
}
```

Response

```
{
    <User>
}
```

### DELETE/users/:id

Response

```
{
    "user": <userId>
    "status": deleted"
}
```
# Frontend

## Project setup

```bash
npm install
```

### Compiles and hot-reloads for development

```bash
npm run serve -- --port 8000
```

### Compiles and minifies for production

```bash
npm run build
```

### Lints and fixes files

```bash
npm run lint
```

## Websocket events

### Outgoing event formats

#### `create_room`

```bash
{ "event": "create_room", "room": <room_name> }
```

#### `send_msg`

```bash
{
    "event": "send_msg",
    "room": <room_name>,
    "date": <date_string>,
    "sender": <nickname>,
    "text": <text>
}
```

### Expected ingoing event formats

#### `create_room`

```bash
{ "event": "create_room", "room": <room_name> }
```

#### `send_msg`

```bash
{
    "event": "send_msg",
    "room": <room_name>,
    "date": <date_string>,
    "sender": <nickname>, 
    "text": <text>
}
```

#### `fetch_all`

```bash
{
    "event": "fetch_all",
    "rooms": [{
        "room": <room_name>,
        "msgs": [{
            "text": <text>,
            "sender": <nickname>,
            "date": <date_string>
        }]
    }]
}
```

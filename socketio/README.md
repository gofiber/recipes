### Disclaimer:
__This an example of a basic chat, connecting multiple sockets.__

## Websocket Chat Example
With this example, you can create a simple chatroom using Websockets. This example supports multiple users and allows them to send messages to each other.

### Connect to the websocket
```
ws://localhost:3000/ws/<user-id>
```
### Message object example

```
{
"from": "<user-id>",
"to": "<recipient-user-id>",
"data": "hello"
}
```


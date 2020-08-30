### Example for fiber as a client to gRPC server.

A sample program to showcase fiber as a client to a gRPC server.

#### Endpoints

| Method | URL           | Return value |
| ------ | ------------- | ------------ |
| GET    | "/add/:a/:b"  | a + b        |
| GET    | "/mult/:a/:b" | a \* b       |

#### Output

```bash
-> curl http://localhost:3000/add/33445/443234
{"result":"476679"}
-> curl http://localhost:3000/mult/33445/443234
{"result":"14823961130"}
```

# Testing Different Architectures

Since we have a whole lot of options for different backends/frontends, I propose
we create a test implementation of all options. Then we can benchmark all of them
as well as see how the code looks.

Each frontend/backend should implement it in HTTP and websockets, so
we can see difference. Common code should be abstracted away into libraries.

## App Specification
It should be easy to implement and also representitive of final traffic.

The format to send dimmer info from client to server is JSON like this:

```json
{
    "<dimmer number>": "<level>",
    [...]
}
```

and then the server should respond with

```json
{
    "<dimmer number>": "<level>",
    [...]}
```


### Backend
#### HTTP
client sends POST w/ json

### Websockets
client sends JSON on websocket

### Frontend
Should display a couple sliders that when it moves sends request to backend




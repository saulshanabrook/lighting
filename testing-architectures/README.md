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
    <dimmer number>: <level>,
    [...]
}
```

and then the server should respond with

```json
{
    <dimmer number>: <level>,
    [...]
}
```


### Backend
#### HTTP
client sends POST w/ json

### Websockets
client sends JSON on websocket

### Frontend
Should display a couple sliders that when it moves sends request to backend



### Resaults
Both Flask and Falcon seem plenty fast for Python, running on my laptop.

```
test_1         | pythonfalcon
test_1         | average response time: 1.97816ms+-0.4634222420212478ms
test_1         | pythonflask
test_1         | average response time: 2.35102ms+-0.3634197292387963ms
```

To run these tests run `docker-compose up` in this directory.

I have stopped pursueing websockets right now, because the flask-websockets
library I was using wasn't maintained and I was getting this error:

```
pythonflaskws_1 | Error: class uri 'flask_sockets.worker' invalid or not found:
pythonflaskws_1 |
pythonflaskws_1 | [Traceback (most recent call last):
pythonflaskws_1 |   File "/usr/local/lib/python2.7/site-packages/gunicorn/util.py", line 140, in load_class
pythonflaskws_1 |     mod = import_module('.'.join(components))
pythonflaskws_1 |   File "/usr/local/lib/python2.7/importlib/__init__.py", line 37, in import_module
pythonflaskws_1 |     __import__(name)
pythonflaskws_1 |   File "/usr/local/lib/python2.7/site-packages/flask_sockets.py", line 15, in <module>
pythonflaskws_1 |     from geventwebsocket.gunicorn.workers import GeventWebSocketWorker as Worker
pythonflaskws_1 |   File "/usr/local/lib/python2.7/site-packages/geventwebsocket/__init__.py", line 18, in <module>
pythonflaskws_1 |     from .server import WebSocketServer
pythonflaskws_1 |   File "/usr/local/lib/python2.7/site-packages/geventwebsocket/server.py", line 1, in <module>
pythonflaskws_1 |     from gevent.pywsgi import WSGIServer
pythonflaskws_1 |   File "/usr/local/lib/python2.7/site-packages/gevent/pywsgi.py", line 12, in <module>
pythonflaskws_1 |     from gevent import socket
pythonflaskws_1 |   File "/usr/local/lib/python2.7/site-packages/gevent/socket.py", line 659, in <module>
pythonflaskws_1 |     from gevent.ssl import sslwrap_simple as ssl, SSLError as sslerror, SSLSocket as SSLType
pythonflaskws_1 |   File "/usr/local/lib/python2.7/site-packages/gevent/ssl.py", line 386, in <module>
pythonflaskws_1 |     def get_server_certificate(addr, ssl_version=PROTOCOL_SSLv3, ca_certs=None):
pythonflaskws_1 | NameError: name 'PROTOCOL_SSLv3' is not defined
pythonflaskws_1 | ]
pythonflaskws_1 |
```


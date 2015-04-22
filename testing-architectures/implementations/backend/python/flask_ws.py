import json


from flask import Flask, jsonify
from flask_sockets import Sockets

from ._base import set_dimmers, load_dimmers


app = Flask(__name__)
sockets = Sockets(app)


@sockets.route('/')
def hello(ws):
    while True:
        message = ws.receive()
        set_dimmers(json.loads(message))
        ws.send(jsonify(**load_dimmers()))

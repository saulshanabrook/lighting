from flask import Flask, request, jsonify

from ._base import set_dimmers, load_dimmers


app = Flask(__name__)


@app.route("/", methods=['POST'])
def hello():
    set_dimmers(request.as_json())
    return jsonify(load_dimmers())

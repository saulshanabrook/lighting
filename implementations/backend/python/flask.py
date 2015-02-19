from flask import Flask, request, jsonify

from ._base import set_dimmers


app = Flask(__name__)


@app.route("/", methods=['POST'])
def hello():
    set_dimmers(request.as_json())
    return jsonify(status="OK")

import requests
import statistics
import collections
import websocket
import json
import ssl
import datetime

Server = collections.namedtuple('Server', ['host', 'type'])

SERVERS = [
    Server('pythonfalcon', 'http'),
    Server('pythonflask', 'http'),
#    Server('pythonflaskws', 'websockets'),
]

TEST_DATA = {
    1: 100,
    2: 0,
    3: 50,
    4: 30
}
NUMBER_TRIES = 100


def convert_to_ms(timedelta):
    return timedelta.total_seconds() * 1000


def hit_http(host, data, times):
    url = 'http://{}:8000'.format(host)
    for _ in range(NUMBER_TRIES):
        r = requests.post(url, json=data)
        yield r.elapsed


def hit_websockets(host, data, times):
    url = "ws://{}:8000".format(host)
    ws = websocket.create_connection(url, sslopt={
        "cert_reqs": ssl.CERT_NONE,
        "check_hostname": False})
    payload = json.puts(data)
    for _ in range(NUMBER_TRIES):
        start_time = datetime.datetime.now()
        ws.send(payload)
        ws.result = ws.recv()
        yield datetime.datetime.now() - start_time


def test(data, servers, number_tries):
    for (host, type) in servers:
        print(host)

        if type == 'http':
            times = hit_http(host, data, number_tries)
        elif type == 'websockets':
            times = hit_websockets(host, data, number_tries)

        times = list(map(convert_to_ms, times))

        print("average response time: {}ms+-{}ms".format(
            statistics.mean(times),
            statistics.pstdev(times)
        ))

test(TEST_DATA, SERVERS, NUMBER_TRIES)

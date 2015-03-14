import os

import requests


HOST = 'web'

TYPE = os.environ['TYPE']
TEST_DATA = {
    1: 100,
    2: 0,
    3: 50,
    4: 30
}

if TYPE == 'http':
    r = requests.post(HOST, json=TEST_DATA)
    assert r.data == TEST_DATA
else:


from __future__ import absolute_import

import json

import falcon

from ._base import set_dimmers, load_dimmers


class RequireJSON(object):

    def process_request(self, req, resp):
        if not req.client_accepts_json:
            raise falcon.HTTPNotAcceptable(
                'This API only supports responses encoded as JSON.',
                href='http://docs.examples.com/api/json')

        if req.method in ('POST', 'PUT'):
            if 'application/json' not in req.content_type:
                raise falcon.HTTPUnsupportedMediaType(
                    'This API only supports requests encoded as JSON.',
                    href='http://docs.examples.com/api/json')


class JSONTranslator(object):

    def process_request(self, req, resp):
        # req.stream corresponds to the WSGI wsgi.input environ variable,
        # and allows you to read bytes from the request body.
        #
        # See also: PEP 3333
        if req.content_length in (None, 0):
            # Nothing to do
            return

        body = req.stream.read()
        if not body:
            raise falcon.HTTPBadRequest('Empty request body',
                                        'A valid JSON document is required.')

        try:
            req.context['doc'] = json.loads(body.decode('utf-8'))

        except (ValueError, UnicodeDecodeError):
            raise falcon.HTTPError(falcon.HTTP_753,
                                   'Malformed JSON',
                                   'Could not decode the request body. The '
                                   'JSON was incorrect or not encoded as '
                                   'UTF-8.')

    def process_response(self, req, resp, resource):
        if 'result' not in req.context:
            return

        resp.body = json.dumps(req.context['result'])


class ThingsResource:
    def on_post(self, req, resp):
        try:
            doc = req.context['doc']
        except KeyError:
            raise falcon.HTTPBadRequest(
                'Missing thing',
                'A thing must be submitted in the request body.')

        set_dimmers(doc)

        resp.status = falcon.HTTP_201
        req.context['result'] = load_dimmers()

# Configure your WSGI server to load "things.app" (app is a WSGI callable)
app = falcon.API(middleware=[
    RequireJSON(),
    JSONTranslator(),
])

things = ThingsResource()
app.add_route('/', things)

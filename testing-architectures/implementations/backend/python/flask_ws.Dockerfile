FROM python:2

RUN pip install Flask-Sockets gunicorn
EXPOSE 8000

ADD . /code/
WORKDIR /code/

CMD gunicorn -c gunicorn_config.py -k flask_sockets.worker lighting.flask_ws:app

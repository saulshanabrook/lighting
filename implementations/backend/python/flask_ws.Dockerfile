FROM python:2

RUN pip install Flask-Sockets gunicorn
CMD gunicorn -k flask_sockets.worker flask_ws:app

FROM python:2

RUN pip install falcon gunicorn
EXPOSE 8000

ADD . /code/
WORKDIR /code/

CMD gunicorn -c gunicorn_config.py lighting.falcon:app

FROM python:2

RUN pip install flask gunicorn
CMD gunicorn flask:app

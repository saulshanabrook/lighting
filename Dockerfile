FROM python:2-onbuild

ADD web.py web.py

CMD [ "python", "./web.py" ]

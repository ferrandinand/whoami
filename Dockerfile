FROM ubuntu:latest
RUN apt-get update -y
RUN apt-get install -y python-pip python-dev build-essential
RUN pip install --upgrade pip
COPY . /app
WORKDIR /app
EXPOSE 5000
RUN pip install flask
ENTRYPOINT ["python"]
CMD ["app.py"]

FROM ubuntu:latest

ENV DEBIAN_FRONTEND=noninteractive
ENV TZ=Europe/Istanbul
ENV POSTGRES_USER=ekinoks 
ENV POSTGRES_PASSWORD=ekinoksgccdb 
ENV POSTGRES_DB=ekinoksdb

RUN apt-get update && \
    apt-get install -y postgresql golang tzdata ca-certificates && \
    rm -rf /var/lib/apt/lists/*

VOLUME ["/var/lib/postgresql/data"]

COPY . /app
WORKDIR /app

RUN /etc/init.d/postgresql start && \
    su - postgres -c "psql -c \"CREATE DATABASE ${POSTGRES_DB}\"" && \
    su - postgres -c "psql -d ${POSTGRES_DB} -f /app/create.sql" && \
    go get -u github.com/lib/pq && \
    go get -u github.com/dgrijalva/jwt-go && \
    go build -o main .

EXPOSE 8080

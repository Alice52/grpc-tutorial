FROM golang:1.13.4

MAINTAINER zack.zhang (zzhang_xz@163.com)

# install dependencies
RUN go get gopkg.in/mgo.v2

# env
ENV MONGO_HOST 101.132.45.28
ENV MONGO_PORT 27017
ENV MONGO_DB DATABASE_NAME
ENV MONGO_USER DATABASE_USER
ENV MONGO_PASS DATABASE_PASSWORD

# copy app
ADD . /app
WORKDIR /app

# build
RUN go build -o build/storj src/*.go

ENTRYPOINT ["/app/docker-entrypoint.sh"]
FROM hobord/golang-dev AS backend

WORKDIR /workspace
COPY . /workspace

RUN make configure \
    && make test \
    && make

FROM node:13 AS frontend

WORKDIR /workspace

RUN git clone https://github.com/hobord/invst-portfolio-frontend.git .
RUN yarn \
    && yarn build


FROM ubuntu:focal

ENV PORT=8080
ENV DB_HOST=mysql:3306
ENV DB_USER=dbuser
ENV DB_PASSWORD=secret
ENV DB_NAME=testdb
ENV MIGRATIONS=/app/migrations
ENV FRONTEND=/app/public

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update \
    && apt-get -y install --no-install-recommends netcat

ENV DEBIAN_FRONTEND=dialog

WORKDIR /app

COPY --from=backend /workspace/bin /app
COPY --from=backend /workspace/infrastructure/mysql/migrations /app/migrations
COPY --from=backend /workspace/deployment/init.sh /app/
COPY --from=frontend /workspace/dist /app/public

ENV HOST 0.0.0.0


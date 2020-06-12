FROM hobord/golang-dev AS build

WORKDIR /src

COPY . /src

RUN make configure \
    && make test \
    && make

FROM ubuntu:focal

WORKDIR /app

COPY --from=build /src/bin /app/bin
COPY --from=build /src/infrastructure/mysql/migrations /app/migrations


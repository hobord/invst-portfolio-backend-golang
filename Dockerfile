FROM hobord/golang-dev AS build

workdir /src

COPY . /src

RUN make configure \
    && make test \
    && make

FROM alpine

workdir /app

COPY --from=build /src/bin /app/bin
COPY --from=build /src/infrastructure/mysql/migrations /app/migrations


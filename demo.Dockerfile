FROM hobord/golang-dev AS backend

workdir /workspace
COPY . /workspace

RUN make configure \
    && make test \
    && make

FROM node:13 AS frontend

workdir /workspace

RUN git clone https://github.com/hobord/invst-portfolio-frontend.git .
RUN yarn \
    && yarn build


FROM alpine

ENV DB_HOST=mysql:3306
ENV DB_USER=dbuser
ENV DB_PASSWORD=secret
ENV DB_NAME=testdb
ENV MIGRATIONS=/app/migrations
ENV FRONTEND=/app/public

workdir /app

COPY --from=backend /workspace/bin /app
COPY --from=backend /workspace/infrastructure/mysql/migrations /app/migrations
COPY --from=frontend /workspace/dist /app/public

ENV HOST 0.0.0.0

ENTRYPOINT ["/app/portfolio-server migrate"]
CMD ["/app/portfolio-server", "serve"]

## Develop
I recommend use Docker!

### with vim

https://github.com/hobord/docker-golang-dev

```
docker run --user $(id -u):$(id -g) -it --rm -w=/workspace -v $(pwd):/workspace hobord/golang-dev:vim

vim
```

### with VSCode
Install the [vscode-remote-extensionpack](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack)
(vscode in docker)

Open in Devcontainer. 

### on Cloud 
[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/hobord/invst-portfolio-backend-golang)

## Test
Run unit tests.
```
make configure
make test
```
I created only one example unittest for http GetInstrumentByID handler, I need more time make the all tests.

I using [testify](https://pkg.go.dev/mod/github.com/stretchr/testify@v1.4.0).

And I using [mockery](github.com/vektra/mockery) to generate mocks from interfaces.

## Build

```
make configure
make
```

It will make "portfolio-server" into the "bin" directory

## Use
```
portfolio-server serve [flags]

Flags:
  -c, --cors stringArray     CORS allowed origins You can use multiply this flag. If it is not set then *
  -H, --db_host string       Database host:port
  -d, --db_name string       Database name
  -P, --db_password string   Database password
  -u, --db_user string       Database user
  -f, --frontend string      Public frontend files direcotry path
  -h, --help                 help for serve
  -l, --port int             Listen on this port, default: 8080

portfolio-server serve -l 8080 -H mysql:3306 -d testdb -u dbuser -P secret -f ./
``` 

### DB init
It is create database tables, and seeding the default data.

It is support migrate database stages.
I using github.com/golang-migrate/migrate library.s
```
portfolio-server migrate -H mysql:3306 -d testdb -u dbuser -P secret -m infrastructure/mysql/migrations 

#down (delete all data)

portfolio-server migrate -H mysql:3306 -d testdb -u dbuser -P secret -m infrastructure/mysql/migrations --down
```

## Deployment

### Docker compose

```
docker-compose -f deployment/docker-compose.yaml up
```

### Kubernetes
- You should modify the deployment/kubernetes.yaml file!
- Update the environment variables because this deployment not contains mysql deployment
- I did not configured namespace!

```
kubectl apply -f deployment/kubernetes.yaml
```
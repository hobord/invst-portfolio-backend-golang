## Develop
I recommend using Docker!

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
Open in Gitpod: https://gitpod.io/#https://github.com/hobord/invst-portfolio-backend-golang

## Build

```
make configure
make
```
It will make the server into the bin directory

## Use
```
portfolio-server serve [flags]

Flags:
  -H, --db_host string       Database host:port
  -d, --db_name string       Database name
  -P, --db_password string   Database password
  -u, --db_user string       Database user
  -f, --frontend string      Public frontend files direcotry path  
  -h, --help                 help for serve
  -l, --port int             8080
  -v, --verbose              log requests into stdout

Global Flags:
      --config string   config file (default is $HOME/.backend.yaml)

bin/portfolio-server serve -l 8080 -H mysql:3306 -d testdb -u dbuser -P secret -f ./
``` 

### DB init

```
bin/portfolio-server migrate -H mysql:3306 -d testdb -u dbuser -P secret -m infrastructure/mysql/migrations 

#down

bin/portfolio-server migrate -H mysql:3306 -d testdb -u dbuser -P secret -m infrastructure/mysql/migrations --down
```
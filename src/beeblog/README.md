# Quick Start

## Create database

To mimic the environment of beeblog, you will first create a database.

```shell
  $ sudo mysql -uroot -p
```

```mysql
    CREATE DATABASE beeblog;
```

```mysql
    GRANT ALL PRIVILEGES ON beeblog.* TO xmai@localhost IDENTIFIED BY "Uf4bGZ53Ds*#";
```


![](https://i.loli.net/2019/06/05/5cf792f847f7d37194.png)

## Configurations parsing

Beego will parse the `conf/app.conf` file by default.

```shell
appname = beeblog
httpport = 8088
runmoe = dev

dbhost = 127.0.0.1
dbport = 3306
dbuser = xmai
dbpassword = Uf4bGZ53Ds*#
dbname = ticket

uname = admin
pwd = 123
```

## Installation

```shell
$ go get github.com/astaxie/beego
$ go get github.com/beego/bee
```

## Run server

```shell
$ go run main.go
```

Visting http://localhost:8088/ with a web browser will display your app running:

![](https://i.loli.net/2019/06/14/5d0305e2e3c4e35887.png)

# Reference

[beego](https://beego.me)

[beegoblog](https://github.com/wyr6512/beegoblog)

[Build web application with Golang](https://legacy.gitbook.com/book/astaxie/build-web-application-with-golang/details)
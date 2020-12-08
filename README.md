## Docker网站部署


### 一.docker安装

```.env
curl -fsSL https://download.docker.com/linux/debian/gpg | sudo apt-key add -
curl -sSL https://get.daocloud.io/docker | sh
```

### 二.Mysql安装

#### 1.拉取最新msql镜像
```.env
docker image pull mysql
```
#### 2.新建容器
```.env
docker run -itd --name mysql-server -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql
```

#### 3.进入容器
```.env
docker exec -it mysql-server bash

mysql -uroot -p

# 添加授权
GRANT ALL ON *.* TO 'root'@'%';
flush privileges;

#查看用户加密方式
use mysql;# 切换到mysql库
mysql> select user,plugin from user where user='root'; # 查看加密方式
+------+-----------------------+
| user | plugin                |
+------+-----------------------+
| root | caching_sha2_password |
| root | mysql_native_password |
+------+-----------------------+

## 修改加密方式
mysql> alter user 'root'@'%' identified with mysql_native_password by '123456';
Query OK, 0 rows affected (0.00 sec)

mysql> select user,plugin from user where user='root';
+------+-----------------------+
| user | plugin                |
+------+-----------------------+
| root | mysql_native_password |
| root | mysql_native_password |
+------+-----------------------+
2 rows in set (0.00 sec)

```


### 三.nginx安装


#### 1.拉取最新nginx镜像

```.env
docker pull quay.io/letsencrypt/letsencrypt:latest
```

#### 2.生成证书
```.env
docker run --rm -p 80:80 -p 443:443 -v /etc/letsencrypt:/etc/letsencrypt quay.io/letsencrypt/letsencrypt auth --standalone -m 1144620122@qq.com --agree-tos -d tv.syrme.top -d app.syrme.top
```

#### 3.nginx的Dockerfile
```.env
FROM nginx
COPY nginx.conf /etc/nginx/conf.d/micro-api.conf
```
#### 4.编译容器
```.env
docker build -t nginx-server .
```


### 四.编译项目

#### 1.项目的Dockerfile
```.env
FROM golang

ENV GO111MODULE=on
ENV GOSUMDB=off
ENV GOPROXY="https://goproxy.cn"
WORKDIR $GOPATH/src/micro-api
ADD . $GOPATH/src/micro-api
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build
EXPOSE 8002
ENTRYPOINT ["./micro-api"]
```

#### 2.编译容器
```.env
docker build -t micro-api .
```

### 五.启动项目
```.env
docker run -d --link mysql-server:mysql -p 8002:8002 micro-api
docker run -v /etc/letsencrypt:/etc/letsencrypt --net=host -p 80:80 -p 443:443 nginx-server
```

micro-api.conf文件
```.env
server {
    charset utf-8;
    listen 443 ssl;
    server_name app.syrme.top;
    ssl_certificate /etc/letsencrypt/live/xxx.syrme.top//fullchain.pem; # managed by Certbot
    ssl_certificate_key /etc/letsencrypt/live/xxx.syrme.top/privkey.pem; # managed by Certbot


    location / {
        try_files /_not_exists_ @backend;
    }

    location @backend {
        proxy_set_header X-Forwarded-For $remote_addr;
        proxy_set_header Host            $http_host;

        proxy_pass http://127.0.0.1:8002;
    }

}


server {
    listen 80;
    server_name app.syrme.top;
    return 301 https://$host$request_uri;
}
```

#### 集成redis
```.env
docker pull redis:latest
docker run -itd --name redis-server -p 6379:6379 redis
```

### 一些常用操作
查看镜像
```.env
docker images
```
查看运行的容器
```.env
docker ps -a
```
删除停止的容器：
```.env
docker rm $(sudo docker ps -qf status=exited)
```
指定环境变量
```.env
docker run -d mysql env LANG=C.UTF-8
```
```.env
docker image prune --force --all或者docker image prune -f -a` : 删除所有不使用的镜像
docker container prune -f: 删除所有停止的容器
```

重启全部容器
```.env
docker start $(docker ps -a | awk '{ print $1}' | tail -n +2)
```

修改支持utf-8

```.env
1、添加到profile
echo "export LANG=C.UTF-8" >>/etc/profile && source /etc/profile
2、直接设置环境变量
LANG=C.UTF-8
3、Docker启动容器时，指定环境变量
docker run -d mysql env LANG=C.UTF-8
```
docker image prune -f


docker run --net=none --privileged=true -v /go/src/HlsServer/video:/go/src/HlsServer/video etcd -name hls-server
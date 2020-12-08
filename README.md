### 安装jaeger

最简单的方法，使用docker

```bash
docker pull jaegertracing/all-in-one:latest

docker run -d -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
-p 5775:5775/udp -p 6831:6831/udp -p 6832:6832/udp \
-p 5778:5778  -p 16686:16686 -p 14268:14268  -p 14269:14269   -p 9411:9411 \
jaegertracing/all-in-one:latest
```

6831: jaeger-agent端口

16686：jaeger-query端口

以上两个端口一定要开，进入host:16686



### 运行

运行gateway、service1（http）、service2（rpc）

浏览器打开：http://127.0.0.1:8080/ping

查看追踪日志：进入jaeger-query

![展示.png](https://github.com/jwrookie/trace-demo/blob/master/img/%E5%B1%95%E7%A4%BA.png?raw=true)
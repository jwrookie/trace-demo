/*
 * @Time    : 2020年12月07日 19:05:52
 * @Author  : user
 * @Project : trace-demo
 * @File    : tools.go
 * @Software: GoLand
 * @Describe:
 */
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"io"
	"io/ioutil"
	"net/http"
	"trace"
)

var (
	tracer opentracing.Tracer
	closer io.Closer
)

func traceMiddle(c *gin.Context)  {
	span := tracer.StartSpan(c.Request.URL.String(),opentracing.Tag{Key: string(ext.Component), Value: "http"})
	c.Set("span",span)
	defer func() {
		ext.SpanKindRPCClient.Set(span)
		span.Finish()
	}()
	c.Next()
}

func main() {
	tracer, closer = tools.InitJaegerClient("trace-demo")
	defer closer.Close()
	router := gin.Default()
	router.GET("/ping",traceMiddle,pong)
	router.Run(":8080")
}

func pong(c *gin.Context){
	span := c.MustGet("span").(opentracing.Span)
	httpClient := &http.Client{}
	httpReq ,_ := http.NewRequest("GET","http://127.0.0.1:8081/ping",nil)
	carrier := opentracing.HTTPHeadersCarrier(httpReq.Header)
	if err := tracer.Inject(span.Context(), opentracing.HTTPHeaders, carrier);err != nil {
		span.LogFields(log.String("err", err.Error()))
		c.JSON(200,gin.H{"err":err.Error()})
		return
	}

	res,err := httpClient.Do(httpReq)

	if err != nil {
		span.LogFields(log.String("err", err.Error()))
		c.JSON(200,gin.H{"err":err.Error()})
		return
	}
	defer res.Body.Close()
	resS , _:= ioutil.ReadAll(res.Body)
	c.JSON(200,gin.H{"data" : string(resS)})
}
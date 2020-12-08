/*
 * @Time    : 2020年12月07日 19:06:02
 * @Author  : user
 * @Project : trace-demo
 * @File    : tools.go
 * @Software: GoLand
 * @Describe:
 */
package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"io"
	"trace"
	pb "trace/cmd/protos"
)

const (
	PORT = "127.0.0.1:8082"
)

var (
	tracer opentracing.Tracer
	closer io.Closer
)

func traceMiddle(c *gin.Context)  {
	refFunc := opentracing.FollowsFrom
	carrier := opentracing.HTTPHeadersCarrier(c.Request.Header)
	clientContext, err := tracer.Extract(opentracing.HTTPHeaders, carrier)
	if err != nil {
		c.AbortWithStatusJSON(200,gin.H{"data":err.Error()})
		return
	}
	span := tracer.StartSpan(c.Request.URL.String(),refFunc(clientContext),opentracing.Tag{Key: string(ext.Component), Value: "http"})
	defer func() {
		ext.SpanKindRPCServer.Set(span)
		span.Finish()
	}()
	c.Set("span",span)
	c.Next()
}

func main() {
	tracer, closer = tools.InitJaegerClient("trace-demo")
	defer closer.Close()
	router := gin.Default()
	router.GET("/ping",traceMiddle,pong)
	router.Run(":8081")
}

func pong(c *gin.Context){
	span := c.MustGet("span").(opentracing.Span)
	span.LogFields(log.String(c.Request.URL.String() + "---NewRpcServerClient ：","rpcService:Ping"))
	conn, err := grpc.Dial(PORT, grpc.WithInsecure(), grpc.WithUnaryInterceptor(OpenTracingClientInterceptor()))
	if err != nil {
		span.LogFields(log.String("err", err.Error()))
		return
	}
	client := pb.NewRpcServerClient(conn)
	parm1 := pb.PingReq{}

	ctx := opentracing.ContextWithSpan(context.TODO(),span)
	r, err := client.Ping(ctx, &parm1)
	if err != nil {
		span.LogFields(log.String("err", err.Error()))
		return
	}
	c.String(200,"%s",r.Res)
}

//OpenTracingClientInterceptor  rewrite client's interceptor with open tracing
func OpenTracingClientInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, resp interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		var parentCtx opentracing.SpanContext
		if parent := opentracing.SpanFromContext(ctx); parent != nil {
			parentCtx = parent.Context()
		}
		cliSpan := tracer.StartSpan(
			method,
			opentracing.ChildOf(parentCtx),
			opentracing.Tag{Key: string(ext.Component), Value: "gRPC"},
			ext.SpanKindRPCClient,
		)
		defer cliSpan.Finish()
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		} else {
			md = md.Copy()
		}
		mdWriter := tools.MDReaderWriter{md}
		err := tracer.Inject(cliSpan.Context(), opentracing.TextMap, mdWriter)
		if err != nil {
			grpclog.Errorf("inject to metadata err %v", err)
		}
		ctx = metadata.NewOutgoingContext(ctx, mdWriter.MD)
		err = invoker(ctx, method, req, resp, cc, opts...)
		if err != nil {
			cliSpan.LogFields(log.String("err", err.Error()))
		}
		return err
	}
}


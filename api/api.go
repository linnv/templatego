package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
	"github.com/linnv/logx"
	"github.com/opentracing/opentracing-go"

	conf "qnmock/config"
)

var config *conf.Configuration
var addressGcables *conf.AddressGcables

func Init() {
	config = conf.Config()
	addressGcables = conf.GetAddressGcables()
}

func Hello(c *gin.Context) {
	r := c.Request
	var span opentracing.Span
	opName := r.URL.Path
	// Attempt to join a trace by getting trace context from the headers.
	carrier := opentracing.HTTPHeadersCarrier(r.Header)
	wireContext, err := opentracing.GlobalTracer().Extract(
		opentracing.HTTPHeaders, carrier)
	if err != nil {
		// If for whatever reason we can't join, go ahead an start a new root span.
		span = opentracing.StartSpan(opName)
	} else {
		span = opentracing.StartSpan(opName, opentracing.ChildOf(wireContext))
	}
	defer span.Finish()

	// Since we have to inject our span into the HTTP headers, we create a request
	asyncReq, _ := http.NewRequest("GET", "http://localhost:8081/helloclient", nil)

	carrierNext := opentracing.HTTPHeadersCarrier(asyncReq.Header)
	err = span.Tracer().Inject(
		span.Context(), opentracing.HTTPHeaders, carrierNext)
	if err != nil {
		log.Fatalf("Could not inject span context into header: %v", err)
	}

	if resp, err := http.DefaultClient.Do(asyncReq); err != nil {
		span.SetTag("error", true)
		span.LogEvent(fmt.Sprintf("GET /async error: %v", err))
	} else {
		bsResp, err := httputil.DumpResponse(resp, true)
		if err != nil {
			logx.Warnf("bsResp: %s\n", bsResp)
		}
		bsReq, err := httputil.DumpRequest(asyncReq, true)
		if err != nil {
			logx.Warnf("bsResp: %s\n", bsResp)
		}
		logx.Debugf("bsReq: %s\n bsResp:%s", bsReq, bsResp)
	}

	// wireContext, err := opentracing.GlobalTracer().Extract(
	// 	opentracing.TextMap,
	// 	opentracing.HTTPHeaderTextMapCarrier(r.Header))
	// err = span.Tracer().Inject(span.Context(),
	// 	opentracing.TextMap,
	// 	opentracing.HTTPHeaderTextMapCarrier(asyncReq.Header))

	ctx := &Ctx{}
	bs, err := ioutil.ReadAll(r.Body)
	logx.Debugfln(ctx.GetCurLogCtx()+" raw req bs [%s]\n", bs)
	if err != nil {
		logx.Warnf(ctx.GetCurLogCtx()+" getReq err %v bs [%s]\n", err, bs)
		c.AbortWithStatusJSON(401, err)
		return
	}
	ctx.ReqBody = bs
	bs, err = hello(ctx)
	logx.Debugfln(ctx.GetCurLogCtx()+" raw req bs [%s]\n", bs)
	if err != nil {
		logx.Warnf(ctx.GetCurLogCtx()+" getReq err %v bs [%s]\n", err, bs)
		c.AbortWithStatusJSON(401, err)
		return
	}
	if n, err := c.Writer.Write(bs); err != nil {
		logx.Errorf(ctx.GetCurLogCtx()+" err: %+v write count:%d\n", err, n)
	}
}

func hello(ctx *Ctx) (bs []byte, err error) {
	return []byte("hello" + config.AppName), nil
}

func HelloClient(c *gin.Context) {
	r := c.Request

	bsResp := ""
	bsReq, err := httputil.DumpRequest(r, true)
	if err != nil {
		logx.Warnf("bsResp: %s\n", bsResp)
		return
	}
	logx.Debugf("bsReq: %s\n bsResp:%s", bsReq, bsResp)

	var span opentracing.Span
	opName := r.URL.Path
	// Attempt to join a trace by getting trace context from the headers.
	carrier := opentracing.HTTPHeadersCarrier(r.Header)
	wireContext, err := opentracing.GlobalTracer().Extract(
		opentracing.HTTPHeaders, carrier)
	if err != nil {
		// If for whatever reason we can't join, go ahead an start a new root span.
		span = opentracing.StartSpan(opName)
	} else {
		span = opentracing.StartSpan(opName, opentracing.ChildOf(wireContext))
	}
	defer span.Finish()

	ctx := &Ctx{}
	bs, err := ioutil.ReadAll(r.Body)
	logx.Debugfln(ctx.GetCurLogCtx()+" raw req bs [%s]\n", bs)
	if err != nil {
		logx.Warnf(ctx.GetCurLogCtx()+" getReq err %v bs [%s]\n", err, bs)
		c.AbortWithStatusJSON(401, err)
		return
	}
	ctx.ReqBody = bs
	bs, err = hello(ctx)
	logx.Debugfln(ctx.GetCurLogCtx()+" raw req bs [%s]\n", bs)
	if err != nil {
		logx.Warnf(ctx.GetCurLogCtx()+" getReq err %v bs [%s]\n", err, bs)
		c.AbortWithStatusJSON(401, err)
		return
	}
	if n, err := c.Writer.Write(bs); err != nil {
		logx.Errorf(ctx.GetCurLogCtx()+" err: %+v write count:%d\n", err, n)
	}
	return
}

var ERR_PARAMETER_EMPTY = "参数[%s]不能为空"

type RespAddrsMatch struct {
	IsNormal      bool `json:"IsNormal"`
	IsNormalMaybe bool `json:"IsNormalMaybe"` //if true, addr should be more precisely
}

func AddrsMatch(c *gin.Context) {
	r := c.Request
	addr := r.FormValue("addr")
	if len(addr) < 1 {
		c.AbortWithStatusJSON(401, fmt.Sprintf(ERR_PARAMETER_EMPTY, "addr"))
		return
	}

	ctx := &Ctx{}
	resp := RespAddrsMatch{}
	if addressGcables.MatchShort(addr) {
		resp.IsNormal = true
	} else if addressGcables.Match(addr) {
		resp.IsNormalMaybe = true
	}

	bs, err := json.Marshal(resp)
	if err != nil {
		c.AbortWithStatusJSON(401, err)
	}

	if n, err := c.Writer.Write(bs); err != nil {
		logx.Errorf(ctx.GetCurLogCtx()+" err: %+v write count:%d\n", err, n)
	}
}

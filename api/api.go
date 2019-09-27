package api

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/linnv/logx"

	conf "github.com/linnv/templatego/config"
)

var config *conf.Configuration

func Init() {
	config = conf.Config()
}

func Hello(c *gin.Context) {
	r := c.Request
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

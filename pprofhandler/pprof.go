package pprofhandler

import (
	"net/http/pprof"
	"strings"

	"github.com/OptiServer/fasthttp"
	"github.com/OptiServer/fasthttp/fasthttpadaptor"
)

var (
	cmdline = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Cmdline)
	profile = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Profile)
	symbol  = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Symbol)
	trace   = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Trace)
	index   = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Index)
)

// PprofHandler serves server runtime profiling data in the format expected by the pprof visualization tool.
//
// See https://golang.org/pkg/net/http/pprof/ for details.
func PprofHandler(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Content-Type", "text/html")
	if strings.HasPrefix(string(ctx.Path()), "/debug/pprof/cmdline") {
		cmdline(ctx)
	} else if strings.HasPrefix(string(ctx.Path()), "/debug/pprof/profile") {
		profile(ctx)
	} else if strings.HasPrefix(string(ctx.Path()), "/debug/pprof/symbol") {
		symbol(ctx)
	} else if strings.HasPrefix(string(ctx.Path()), "/debug/pprof/trace") {
		trace(ctx)
	} else {
		index(ctx)
	}
}

package handler

import (
	"fmt"
	"go-helm-rest/helm"
	"net/http"

	"github.com/valyala/fasthttp"
)

func CreateStorage(ctx *fasthttp.RequestCtx) {
	stdout, err := helm.CreateMinio("aclab", "csie1226", "2G")
	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)
	}
	ctx.SetContentType("text/plain")
	fmt.Fprint(ctx, stdout)
}

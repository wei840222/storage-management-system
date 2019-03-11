package handler

import (
	"fmt"
	"go-helm-rest/helm"
	"go-helm-rest/rancher"
	"net/http"

	"github.com/valyala/fasthttp"
)

func CreateStorage(ctx *fasthttp.RequestCtx) {
	releaseInfo, err := helm.CreateMinio("aclab", "csie1226", "2G")
	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)
	}
	endpoint := rancher.GetServiceEndpoint(releaseInfo["releaseName"].(string) + "-minio")
	ctx.SetContentType("application/json")
	fmt.Fprint(ctx, "{\"releaseName\":\""+releaseInfo["releaseName"].(string)+"\",\"endpoint\":\""+endpoint+"\"}")
}

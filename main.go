package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func execCommand(command string) string {
	cmd := exec.Command("bash", "-c", command)
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return string(stdout)
}

func index(ctx *fasthttp.RequestCtx) {
	stdout := execCommand("helm json status wizened-goose")
	ctx.SetContentType("application/json")
	fmt.Fprint(ctx, stdout)
}

func main() {
	router := fasthttprouter.New()
	router.GET("/", index)

	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}

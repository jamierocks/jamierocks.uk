package controllers

import (
    "gopkg.in/macaron.v1"
    "net/http"
)

func GetHomepage(ctx *macaron.Context) {
    ctx.HTML(http.StatusOK, "index")
}

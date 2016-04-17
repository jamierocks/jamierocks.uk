package controllers

import (
    "gopkg.in/macaron.v1"
    "net/http"
)

func GetTimetable9(ctx *macaron.Context) {
    ctx.HTML(http.StatusOK, "school/timetable/9")
}

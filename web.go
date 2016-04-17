package main

import (
    "github.com/jamierocks/jamierocks.uk/controllers"
    "github.com/go-macaron/pongo2"
    "gopkg.in/macaron.v1"
)

func main() {
    // Initialise Macaron
    m := macaron.Classic()
    m.Use(pongo2.Pongoer())

    // Base Routes
    m.Get("/", controllers.GetHomepage)

    // School Routes
    m.Get("/school/timetable/9", controllers.GetTimetable9)

    // Run jamierocks.uk
    m.Run()
}

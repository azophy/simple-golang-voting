package main

import (
    "github.com/labstack/echo"
)

func main() {
    server_url := ":8000"
    e := echo.New()

    e.File("/", "resources/pages/index.html")
    e.Static("/assets", "resources/assets")

    e.Logger.Fatal(e.Start(server_url))
}

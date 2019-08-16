package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("vim-go")

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/hello", hello)
	e.POST("/v1/image", convertImage)

	e.Logger.Fatal(e.Start(":1323"))
}

// hello is return HelloWorld useing TestApi
// To be closed
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World")
}

// ImageParam is post parameter
// TODO move to File
type ImageParam struct {
	Types string `json:"types" form:"types" query:"types"`
	Image string `json:"image" form:"image" query:"image"`
}

// convertImage return JSON Data that image that according to Types of ImageParam
// TODO creating NOW
func convertImage(c echo.Context) error {
	p := new(ImageParam)
	if err := c.Bind(p); err != nil {
		return err
	}
	decode(p.Image)
	return c.JSON(http.StatusOK, p)
}

// decode is that convert base64 to image that according to Types of ImageParam,
// And output it.
// TODO creating NOW
func decode(base64str string) {
	data, _ := base64.StdEncoding.DecodeString(base64str)

	file, _ := os.Create("test.png")

	defer file.Close()
	file.Write(data)
}

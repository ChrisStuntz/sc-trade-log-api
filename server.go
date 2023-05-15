package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Order struct {
	Timestamp string `json:"timestamp"`
	Commodity string `json:"commodity"`
	Price     int    `json:"price"`
	Stock     int    `json:"stock"`
	Demand    string `json:"demand"`
	Location  string `json:"location"`
}

func newOrder(c echo.Context) error {
	o := new(Order)
	if err := c.Bind(o); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, o)
}

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	e.POST("/orders", newOrder)

	e.Logger.Fatal(e.Start(":8080"))

}

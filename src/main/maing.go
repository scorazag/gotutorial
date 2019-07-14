package main

import(
	"fmt"
	"net/http"
	"github.com/labstack/echo"
)

func yallo(c echo.Context) error{
	return c.String(http.StatusOK,"Hola  desde server")
}
func getCats(c echo.Context) error{
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	dataType := c.Param("data")

	if dataType =="string" {
		return c.String(http.StatusOK,fmt.Sprintf("nombre:%s\n and type:%s\n",catName,catType ))
	}
	if dataType =="json" {
		return c.JSON(http.StatusOK,map[string]string{
			"name":catName,
			"type":catType,
		})
	}
	
	return c.JSON(http.StatusBadRequest,map[string]string{
		"error":"salio mal chavo",
	})
}
func main(){
	fmt.Println("Hello")

	e := echo.New()
	e.GET("/",yallo)
	e.GET("/cats/:data",getCats)

	//start the server
	e.Start(":8000")
}
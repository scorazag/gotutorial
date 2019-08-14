package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
	"github.com/labstack/echo"
)

type Cat struct {
	Name string "json:name"
	Type string "json:type"
}

type Dog struct {
	Name string "json:name"
	Type string "json:type"
}

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

func addCat(c echo.Context) error {
	cat := Cat{}

	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		log.Println("Fallo al leer")
		return c.String(http.StatusInternalServerError,"")
	}

	err = json.Unmarshal(b, &cat)

	if err != nil {
		log.Println("Fallo ")
		return c.String(http.StatusInternalServerError,"")
	}

	log.Println(cat)
	return c.String(http.StatusOK, "Gatitot agregado")

}

func addDog(c echo.Context) error{
	dog := Dog{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&dog)

	if err != nil {
		log.Println("Proceso fallido %s",err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Println("Perrito %#v",dog)
	return c.String(http.StatusOK, "Perrito agregado")
}
func main(){
	fmt.Println("Hello")

	e := echo.New()
	e.GET("/",yallo)
	e.GET("/cats/:data",getCats)
	e.POST("/cats", addCat)
	e.POST("/dogs",addDog)

	//start the server
	e.Start(":8000")
}
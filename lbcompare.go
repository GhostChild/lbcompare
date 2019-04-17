package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"io/ioutil"
	"strconv"
)

type Resultdata struct {
	Naobu_v1 int `json:"naobu_v1"`
	Naobu_v2 int `json:"naobu_v2"`
}

func main() {
	e := echo.New()

	data := &Resultdata{}

	f, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
	} else {
		json.Unmarshal(f, data)
	}

	fmt.Println(data)

	e.GET("/submit", func(context echo.Context) error {

		reqdata := context.QueryParam("radio")

		count, _ := strconv.Atoi(reqdata)
		if count == 1 {
			data.Naobu_v1 += 1
		} else if count == 2 {
			data.Naobu_v2 += 1
		}

		j, _ := json.MarshalIndent(data, "", "\t")

		ioutil.WriteFile("data.json", j, 666)

		return context.String(200, reqdata)
	})
	e.GET("/select", func(context echo.Context) error {

		return context.JSON(200, data)
	})

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8082"},
		//AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType,},
	}))

	e.Start(":8889")
}

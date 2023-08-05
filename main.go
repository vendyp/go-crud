package main

import (
	"crudapp/entities"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Running web server on port 5000")

	//db migrations
	connString := "sqlserver://sa:Password1234@localhost:1433?database=CrudDb"
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		panic("failed to connect db")
	}

	//migrate table user
	db.AutoMigrate(&entities.Product{})
	db.AutoMigrate(&entities.ProductCategory{})

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":5000"))
}

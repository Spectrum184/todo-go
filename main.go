package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"todo-go/modules/items/transport/gints"
)

func main() {
	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", gints.CreateItem(db))
			items.GET("", gints.ListItem(db))
			items.GET("/:id", gints.GetItem(db))
			items.PATCH("/:id", gints.UpdateItem(db))
			items.DELETE("/:id", gints.DeleteItem(db))
		}
	}

	e := r.Run(":5000")

	if e != nil {
		return
	}
}

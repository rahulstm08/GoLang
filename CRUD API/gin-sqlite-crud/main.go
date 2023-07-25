// main.go
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    initDB()

    // CRUD endpoints
    r.POST("/items", createItem)
    r.GET("/items", getItems)
    r.PUT("/items/:id", updateItem)
    r.DELETE("/items/:id", deleteItem)

    r.Run(":8080")
}
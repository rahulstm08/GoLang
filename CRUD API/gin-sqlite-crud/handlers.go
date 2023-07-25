// handlers.go
package main

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type Item struct {
    ID   int    `json:"id"`
    Title string `json:"title"`
    Description string  `json:"description"`
    DueDate int `json:"due date"`
    Status string `json:"status"`
}

func createItem(c *gin.Context) {
    var newItem Item
    if err := c.ShouldBindJSON(&newItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
        return
    }

    result, err := db.Exec("INSERT INTO items (name) VALUES (?)", newItem.Title)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create item"})
        return
    }

    id, _ := result.LastInsertId()
    newItem.ID = int(id)
    c.JSON(http.StatusOK, newItem)
}

func getItems(c *gin.Context) {
    rows, err := db.Query("SELECT id, name FROM items")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch items"})
        return
    }
    defer rows.Close()

    var items []Item
    for rows.Next() {
        var item Item
        rows.Scan(&item.ID, &item.Title)
        items = append(items, item)
    }

    c.JSON(http.StatusOK, items)
}

func updateItem(c *gin.Context) {
    itemID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
        return
    }

    var updatedItem Item
    if err := c.ShouldBindJSON(&updatedItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
        return
    }

    _, err = db.Exec("UPDATE items SET name = ? WHERE id = ?", updatedItem.Title, itemID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item"})
        return
    }

    updatedItem.ID = itemID
    c.JSON(http.StatusOK, updatedItem)
}

func deleteItem(c *gin.Context) {
    itemID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
        return
    }

    _, err = db.Exec("DELETE FROM items WHERE id = ?", itemID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
}

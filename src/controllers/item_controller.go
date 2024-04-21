package controllers

import (
	"net/http"

	"smartm2m/src/helper"
	"smartm2m/src/models"

	"github.com/gin-gonic/gin"
)

type ItemController struct{}

func NewItemController() *ItemController {
	return &ItemController{}
}

func (ic *ItemController) CreateItem(c *gin.Context) {
	var item models.ItemNFT
	var err error
	if err = c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenString := c.GetHeader("Authorization")
	item.UserID, err = helper.GetUserIDFromJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	createdItem, err := models.CreateItem(item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdItem)
}

func (ic *ItemController) GetItems(c *gin.Context) {
	userID := c.GetString("userid")

	items, err := models.GetItems(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

func (ic *ItemController) GetItem(c *gin.Context) {
	id := c.Param("id")

	item, err := models.GetItem(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (ic *ItemController) UpdateItem(c *gin.Context) {
	id := c.Param("id")

	tokenString := c.GetHeader("Authorization")
	userid, err := helper.GetUserIDFromJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	itemOld, err := models.GetItem(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if itemOld.UserID != userid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not Authorize to access this item"})
		return
	}

	var item models.ItemNFT
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedItem, err := models.UpdateItem(id, item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedItem)
}

func (ic *ItemController) DeleteItem(c *gin.Context) {
	id := c.Param("id")

	tokenString := c.GetHeader("Authorization")
	userid, err := helper.GetUserIDFromJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	itemOld, err := models.GetItem(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if itemOld.UserID != userid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not Authorize to access this item"})
		return
	}

	err = models.DeleteItem(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
}

func (ic *ItemController) PurchaseItem(c *gin.Context) {
	id := c.Param("id")

	tokenString := c.GetHeader("Authorization")
	userid, err := helper.GetUserIDFromJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	itemOld, err := models.GetItem(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if itemOld.UserID != userid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot purchase youre own items"})
		return
	}

	err = models.PurchaseItem(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item purchased successfully"})
}

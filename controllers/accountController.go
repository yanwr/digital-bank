package controllers

import "github.com/gin-gonic/gin"

func ShowAllAccounts(c *gin.Context) {
	c.JSON(200, gin.H{"value": "works!"})
}

func IndexBalanceAccount(c *gin.Context) {}

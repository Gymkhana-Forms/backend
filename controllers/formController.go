package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllFormTemplates() gin.HandlerFunc {
	data, err := ioutil.ReadFile("./json/test.json")
	if err != nil {
		log.Fatal("Error while opening file", err)
	}

	var payload map[string]interface{}
	err = json.Unmarshal(data, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal()", err)
	}

	return func(c *gin.Context) {
		c.JSON(200, payload)
	}
}

func SubmittedForms() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_type := c.GetString("User_type")
		if user_type != "ADMIN" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to view this resource"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Here are the submitted forms(to be implemented)"})
		return
	}
}

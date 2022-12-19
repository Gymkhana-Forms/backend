package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Gymkhana-Forms/backend/db"
	"github.com/gin-gonic/gin"
)

type FormTemplate struct {
	Schema   map[string]interface{}
	UISchema map[string]interface{}
	Example  map[string]interface{}
}

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

func SelectForm() gin.HandlerFunc {
	return func(c *gin.Context) {
		formtype := c.Params.ByName("formtype")

		var schema map[string]interface{}
		var uischema map[string]interface{}
		var example map[string]interface{}

		data1, err1 := ioutil.ReadFile("./json/" + formtype + "_schema.json")
		if err1 != nil {
			log.Fatal("Error while opening schema file", err1)
		}
		err1 = json.Unmarshal(data1, &schema)
		if err1 != nil {
			log.Fatal("Error during Unmarshal()", err1)
		}

		data2, err2 := ioutil.ReadFile("./json/" + formtype + "_uischema.json")
		if err2 != nil {
			log.Fatal("Error while opening uischema file", err2)
		}
		err2 = json.Unmarshal(data2, &uischema)
		if err2 != nil {
			log.Fatal("Error during Unmarshal()", err2)
		}

		data3, err3 := ioutil.ReadFile("./json/" + formtype + "_example.json")
		if err3 != nil {
			log.Fatal("Error while opening example file", err3)
		}
		err3 = json.Unmarshal(data3, &example)
		if err3 != nil {
			log.Fatal("Error during Unmarshal()", err3)
		}

		curr_form := FormTemplate{
			Schema:   schema,
			UISchema: uischema,
			Example:  example,
		}
		err := c.BindJSON(&curr_form)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, curr_form)
	}
}

func GetAllForms() gin.HandlerFunc {
	var allforms []db.FormData
	err := db.ExtractAllForms(&allforms)

	return func(c *gin.Context) {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(200, allforms)
		}
	}
}

func GetForm() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		var myform db.FormData
		err := db.ExtractFormByID(&myform, id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(200, myform)
		}
	}
}

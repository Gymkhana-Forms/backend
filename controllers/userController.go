package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Gymkhana-Forms/backend/db"
	"github.com/Gymkhana-Forms/backend/helpers"
	"github.com/Gymkhana-Forms/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = db.OpenCollection(db.Client, "users")
var validate *validator.Validate = validator.New()

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic("Something wrong with hashing", err)
	}
	return string(bytes)
}

func VerifyPassword(userPass string, providedPass string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPass), []byte(userPass))
	var match bool = true
	var msg string = ""
	if err != nil {
		match = false
		msg = "Incorrect email or password"
	}
	return match, msg
}

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User

		err := c.BindJSON(&user)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(user)

		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		emailValid := helpers.VerifyEmailDomain(*(user.Email))
		if !emailValid {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Email"})
			return
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "This email already exists"})
			return
		}

		password := HashPassword(*user.Password)
		user.Password = &password

		user.ID = primitive.NewObjectID()
		token, refreshToken, err := helpers.GenerateAllTokens(*user.Email, *user.Name, *user.User_type, *user.RollNo)
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while generating token"})
			return
		}

		user.Token = &token
		user.Refresh_token = &refreshToken

		result, inserterr := userCollection.InsertOne(ctx, user)
		if inserterr != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during insert"})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		var foundUser models.User

		err := c.BindJSON(&user)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Account does not exist"})
			return
		}

		passValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		defer cancel()
		if !passValid {
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
			return
		}

		token, refresh_token, err := helpers.GenerateAllTokens(*foundUser.Email, *foundUser.Name, *foundUser.User_type, *foundUser.RollNo)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		helpers.UpdateAllTokens(token, refresh_token, *foundUser.RollNo)

		err = userCollection.FindOne(ctx, bson.M{"rollno": foundUser.RollNo}).Decode(&foundUser)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, foundUser)
	}
}

package helpers

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/Gymkhana-Forms/backend/db"
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct {
	Email     string
	Name      string
	RollNo    string
	User_type string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = db.OpenCollection(db.Client, "users")

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string, name string, user_type string, rollno string) (token string, refresh_token string, err error) {
	claims := &SignedDetails{
		Email:     email,
		Name:      name,
		RollNo:    rollno,
		User_type: user_type,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshclaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(600)).Unix(),
		},
	}

	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refresh_token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshclaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, refresh_token, err
}

func UpdateAllTokens(token string, refresh_token string, rollno string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var updateObj primitive.D

	updateObj = append(updateObj, bson.E{Key: "token", Value: token})
	updateObj = append(updateObj, bson.E{Key: "refresh_token", Value: refresh_token})

	upsert := true
	filter := bson.M{"rollno": rollno}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := userCollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{Key: "$set", Value: updateObj},
		},
		&opt,
	)
	defer cancel()

	if err != nil {
		log.Panic(err)
		return
	}
}

func ValidateToken(signedtoken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedtoken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)

	if !ok {
		msg = "Invalid Token"
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "Expired Token"
		return
	}

	return claims, msg
}

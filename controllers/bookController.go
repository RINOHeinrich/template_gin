package controllers

import (
	"context"
	"net/http"

	database "github.com/RINOBE/gestion_de_livre/databases"
	"github.com/RINOBE/gestion_de_livre/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var bookCollection *mongo.Collection = database.BookData(database.Client, "book")

func GetAllBooks() gin.HandlerFunc {
	return func(c *gin.Context) {
		//get all books in mongo collection
		filter := bson.D{{}}
		books, err := bookCollection.Find(context.Background(), filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": books})
	}
}
func InsertBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		//insert a book in mongo collection
		var book models.Book
		err := c.BindJSON(&book)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		book.ID = primitive.NewObjectID()
		_, err = bookCollection.InsertOne(context.Background(), book)
		validator_err := validate.Struct(book)
		if validator_err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validator_err.Error()})
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Book inserted successfully"})
	}
}
func DelBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		//delete a book in mongo collection
		id := c.Param("id")
		_, err := bookCollection.DeleteOne(context.TODO(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
	}
}
func UpdateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		//update a book in mongo collection
		id := c.Param("id")
		var book models.Book
		err := c.BindJSON(&book)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		_, err = bookCollection.UpdateOne(context.TODO(), id, book)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
	}
}

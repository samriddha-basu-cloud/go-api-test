package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID       int    `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Position string `json:"position" bson:"position"`
}

var (
	mongoURI   = "mongodb://localhost:27017" // Update with your MongoDB URI
	dbName     = "mydb"                      // Update with your database name
	colName    = "users"                     // Collection name
	collection *mongo.Collection
)

func init() {

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	db := client.Database(dbName)
	collection = db.Collection(colName)
}

func generateRandomID() int {
	// Seed the random number generator with the current timestamp
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(1000) // You can adjust the range of random IDs as needed
}

func createUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	// Generate a random ID
	user.ID = generateRandomID()

	// Insert the user into MongoDB
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// Find user by ID in MongoDB
	user := new(User)
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(user)
	if err != nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, user)
}

func getAllUsers(c echo.Context) error {
	// Find all users in MongoDB
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer cursor.Close(context.Background())

	var users []User
	for cursor.Next(context.Background()) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		users = append(users, user)
	}

	return c.JSON(http.StatusOK, users)
}

func updateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user := new(User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	// Update user by ID in MongoDB
	_, err := collection.ReplaceOne(context.Background(), bson.M{"_id": id}, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// Find the user by ID in MongoDB
	user := new(User)
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(user)
	if err != nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	// Delete the user by ID in MongoDB
	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Return a message indicating which user has been deleted
	message := fmt.Sprintf("User with ID %d and name '%s' has been deleted", id, user.Name)
	return c.JSON(http.StatusOK, message)
}

func main() {
	e := echo.New()

	// Add a simple "Hello, World!" route at the root ("/")
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)
	e.GET("/users", getAllUsers)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	e.Start(":8080") // Change the port as needed
}

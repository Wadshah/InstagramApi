package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/insta/pkg/handlers"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	fmt.Println("ClientOptopm TYPE:", reflect.Typeof(clientOptions), "\n")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Mongo.connect() ERROR:", err)
		os.Exit(1)

	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	col := client.Database("First_Database").Collection("First Collection")
	fmt.Println("Collection Type:", reflect.TypeOf(col), "\n")
	router := mux.NewRouter()

	router.HandleFunc("/users", handlers.AddUser()).Methods(http.MethodPost)        ///Adds a user
	router.HandleFunc("/user/id", handlers.GetUserId()).Methods(http.MethodGet)     ///GetUserID
	router.HandleFunc("/Post", handlers.AddPost()).Methods(http.MethodPost)         ///Adds Posts
	router.HandleFunc("/Post/user/id", handlers.GetPost()).Methods(http.MethodPost) //Gets the post with the id
	router.HandleFunc("/Post/id", handlers.GetAllPost()).Methods(http.MethodPost)   ///Gets All Posts

	log.Println("Running")
	http.ListenAndServe(":4000", router)
}

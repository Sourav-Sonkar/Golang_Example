package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github/Sourav-Sonkar/mongoApi/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//this file contains database connection,database helper,controllers
const connectionString = "mongodb+srv://testUser:test123@cluster0.0wgjt.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
const dbName = "netflix" 
const colName = "watchlist"

//Most important as connection
var collection *mongo.Collection

//connect with mongo db

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongo DB
	client, err := mongo.Connect(context.TODO(), clientOption)

	errHandler(err)

	fmt.Println("Mongo DB connect success")
	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection instance is ready")
}

/*********MongoDB helpers file start here **********/

//insert 1 record
func insertMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	errHandler(err)
	fmt.Println("Inserted 1 movie in db with id:", inserted.InsertedID)
}

//update 1 record
func updatOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	errHandler(err)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	errHandler(err)
	fmt.Println("modified count", result.ModifiedCount)
}

//delet one record
func deleteOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	errHandler(err)
	filter := bson.M{"_id": id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	errHandler(err)
	fmt.Println("Movie got delete with delete count:", deleteCount)
}

//delete all records from mongo db
func deleteAllMovie() int64 {
	//filter := bson.D{{}}
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	errHandler(err)
	fmt.Println("No of movies deleted:", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

//get all movies from database
func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	errHandler(err)
	var movies []primitive.M
	for cur.Next(context.TODO()) {
		var movie bson.M
		err := cur.Decode(&movie)
		errHandler(err)
		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies
}

/*********MongoDB helpers file end here **********/

/*********MongoDB controllers file(Actual file) start here **********/

func GetAllMoviesController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

func MarkedAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	param := mux.Vars(r)
	updatOneMovie(param["id"])
	json.NewEncoder(w).Encode(param)
}

func DeleteOneMovieController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	param := mux.Vars(r)
	deleteOneMovie(param["id"])
	json.NewEncoder(w).Encode(param["id"])
}

func DeleteAllMovieController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)

}

/*********MongoDB controllers file end here **********/

func errHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

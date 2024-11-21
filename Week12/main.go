package main
import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func enablCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
func saveIPHandler(w http.ResponseWriter, r *http.Request){
	enablCors(&w)
	if r.Method != "POST" {
		http.Error(w,"Only POST method is allowed.", http.StatusMethodNotAllowed)
		return
	}
	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w,"Error parsing JSON request body.", http.StatusBadRequest)
		return
	}
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	data["time"] = currentTime
	collection := mongoClient.Database("portfolio").Collection("ips")
	_, err = collection.InsertOne(context.TODO(), data) 
	if err != nil {
		http.Error(w, "Error saving data to MongoDB", http.StatusInternalServerError)
	}

}
func main() {
	var err error
	mongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	fs := http.FileServer(http.Dir("./templates"))
	http.Handle("/", fs)
	http.HandleFunc("/api/vistors", saveIPHandler)
	log.Println("Server started on port 80")
	log.Fatal(http.ListenAndServe(":80", nil))
}

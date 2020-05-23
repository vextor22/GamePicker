package steamconnector

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User type contains information related to steam UserId
type User struct {
	ID          string    `json:"userid" bson:"_id"`
	GameCount   int       `json:"game_count" bson:"game_count"`
	Name        string    `json:"user_name" bson:"user_name"`
	LastUpdated time.Time `json:"last_updated" bson:"last_updated"`
	Games       []Game    `json:"games" bson:"games"`
}

// Game type contains information about a User's game
type Game struct {
	ID       int    `json:"appid" bson:"appid"`
	Name     string `json:"name" bson:"game_name"`
	PlayTime int    `json:"playtime_forever" bson:"playtime_forever"`
}

// SteamUserGameListResponse is used to unmarshal the response provided by the steam user games API
type SteamUserGameListResponse struct {
	Response struct {
		GameCount int    `json:"game_count"`
		GameList  []Game `json:"games"`
	} `json:"response`
}

// RegisterSteamEndpoints registers steam request endpoints with the router
func RegisterSteamEndpoints(r *mux.Router) {
	r.HandleFunc("/user/{id}", userLookup).Methods("GET")
}

func createUser(user string, resp *http.Response) User {

	userWithGames := User{ID: user}
	var games SteamUserGameListResponse
	err := json.NewDecoder(resp.Body).Decode(&games)
	if err != nil {
		log.Printf("ERROR decoding json: %v", err)
	}
	userWithGames.Games = games.Response.GameList
	userWithGames.GameCount = games.Response.GameCount

	return userWithGames
}

func userLookup(w http.ResponseWriter, r *http.Request) {

	client := &http.Client{}
	params := mux.Vars(r)
	userID := params["id"]
	log.Printf("Request: %v", userID)

	req, err := http.NewRequest("GET", "http://api.steampowered.com/IPlayerService/GetOwnedGames/v1", nil)
	if err != nil {
		log.Printf("Error generating request string: %v", err)
		return
	}
	q := req.URL.Query()
	q.Add("key", AppConfig.APIKey)
	q.Add("include_appinfo", "true")
	q.Add("steamid", userID)
	req.URL.RawQuery = q.Encode()

	log.Printf("Request URL: %v", req.URL.String())

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error doing request: %v", err)
		return
	}
	user := createUser(userID, resp)

	storeUser(r.Context(), &user)

	w.Header().Set("Content-Type", "application/json")
	j, _ := json.Marshal(user)
	w.Write(j)
}

func storeUser(ctx context.Context, user *User) {
	client := ctx.Value(mongoKey).(*mongo.Client)
	database := client.Database("gamepicker")
	steamUserData := database.Collection("steamUsers")
	user.LastUpdated = time.Now()

	opts := options.Replace().SetUpsert(true)
	filter := bson.D{bson.E{"_id", user.ID}}
	insertResult, err := steamUserData.ReplaceOne(ctx, filter, user, opts)
	if err != nil {
		log.Printf("Error inserting user: %v", err)
	}
	log.Printf("User inserted: %v", insertResult.UpsertedID)

}

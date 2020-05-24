package steamconnector

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type steamUserFetch struct {
	userFetch User
	ctx       context.Context
	err       error
}

var (
	fetch   = make(chan steamUserFetch)
	fetched = make(chan steamUserFetch)
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
	ID        int    `json:"appid" bson:"appid"`
	Name      string `json:"name" bson:"game_name"`
	PlayTime  int    `json:"playtime_forever" bson:"playtime_forever"`
	ImageLogo string `json:"img_logo_url" bson:"image_logo_url"`
}

// SteamUserGameListResponse is used to unmarshal the response provided by the steam user games API
type SteamUserGameListResponse struct {
	Response struct {
		GameCount int    `json:"game_count"`
		GameList  []Game `json:"games"`
	} `json:"response"`
}

func init() {
	initUserRequestRoutine()
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

	params := mux.Vars(r)
	userID := params["id"]
	user, err := getUserFromDB(r.Context(), userID)
	if err != nil {
		// User could not be retrieved from database, reach out to Steam API
		fetch <- steamUserFetch{userFetch: user, ctx: r.Context()}
		userResponse := <-fetched
		if userResponse.err != nil {
			log.Printf("ERROR retrieving user from steam: %v", userResponse.err)
		}
		user = userResponse.userFetch
	}

	w.Header().Set("Content-Type", "application/json")
	j, _ := json.Marshal(user)
	w.Write(j)
}

func initUserRequestRoutine() {

	go func() {
		for {

			userRequest := <-fetch
			client := &http.Client{}

			req, err := http.NewRequest("GET", "http://api.steampowered.com/IPlayerService/GetOwnedGames/v1", nil)
			if err != nil {
				fetched <- steamUserFetch{userFetch: userRequest.userFetch, err: err}
			}
			q := req.URL.Query()
			q.Add("key", AppConfig.APIKey)
			q.Add("include_appinfo", "true")
			q.Add("steamid", userRequest.userFetch.ID)
			req.URL.RawQuery = q.Encode()

			resp, err := client.Do(req)
			if err != nil {
				fetched <- steamUserFetch{userFetch: userRequest.userFetch, err: err}
			}

			userRequest.userFetch = createUser(userRequest.userFetch.ID, resp)
			storeUser(userRequest.ctx, &userRequest.userFetch)
			fetched <- steamUserFetch{userFetch: userRequest.userFetch, err: err}
			time.Sleep(time.Second)
		}
	}()
}

func getUserFromDB(ctx context.Context, userID string) (User, error) {

	client := ctx.Value(mongoKey).(*mongo.Client)
	database := client.Database("gamepicker")
	steamUserData := database.Collection("steamUsers")

	user := User{ID: userID}
	err := steamUserData.FindOne(ctx, bson.D{bson.E{"_id", userID}}).Decode(&user)
	if err != nil {
		log.Printf("Error finding user %v", err)
		return user, err
	}
	age := time.Now().Sub(user.LastUpdated)
	if age >= AppConfig.CacheDuration {
		return user, fmt.Errorf("User older than %v, please refresh", AppConfig.CacheDuration)

	}

	return user, nil

}

func storeUser(ctx context.Context, user *User) {
	client := ctx.Value(mongoKey).(*mongo.Client)
	database := client.Database("gamepicker")
	steamUserData := database.Collection("steamUsers")
	user.LastUpdated = time.Now()

	opts := options.Replace().SetUpsert(true)
	filter := bson.D{bson.E{"_id", user.ID}}
	_, err := steamUserData.ReplaceOne(ctx, filter, user, opts)
	if err != nil {
		log.Printf("Error inserting user: %v", err)
	}
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/monaco-io/request"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
// When key miss, time taken = 610ms 
// When key found in Cache, time taken = 22ms
func getPhotos(w http.ResponseWriter, r *http.Request){
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	photos, err := rdb.Get(ctx, "photos").Result()
	if err == redis.Nil {
		// key does not exists
		fmt.Println("Key Miss")
		c := request.Client{
			URL:    "https://jsonplaceholder.typicode.com/photos",
			Method: "GET",
		}
		resp := c.Send()
		if !resp.OK(){
			// handle error
			log.Println(resp.Error())
		}
		photos = resp.String()
		err = rdb.Set(ctx, "photos", photos, 0).Err()
		if err != nil {
			panic(err)
		}
	} else if err != nil {
        panic(err)
    }
	json.NewEncoder(w).Encode(photos)
}

func HandleRequests(){
	http.HandleFunc("/photos", getPhotos)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main(){
	HandleRequests()
}

// go env -w GO111MODULE=on
// go mod init github.com/my/repo
// go get github.com/redis/go-redis/v9
// go get github.com/monaco-io/request
// go run main.go 
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(15).Seconds().Do(CreatePost)

	s.StartBlocking()
}

func CreatePost() {
	data := map[string]interface{}{
		"title":  "Hactiv 8",
		"body":   "Challenge 7 - HTTP Request (POST)",
		"userId": rand.Intn(101),
	}

	requestJson, err := json.Marshal(data)
	client := &http.Client{}
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJson))
	req.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Request JSON:")
	log.Println(string(requestJson))
	fmt.Println("Response JSON:")
	log.Println(string(body))
	ShowWeatherReport()
	fmt.Println("=========================")
}

func ShowWeatherReport() {
	weather := map[string]int{
		"water": rand.Intn(101),
		"wind":  rand.Intn(101),
	}

	weatherJson, err := json.MarshalIndent(weather, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Weather report:")
	fmt.Println(string(weatherJson))
	statusWater := ""
	statusWind := ""

	if weather["water"] <= 5 {
		statusWater = "aman"
	} else if weather["water"] >= 6 && weather["water"] <= 8 {
		statusWater = "siaga"
	} else if weather["water"] > 8 {
		statusWater = "bahaya"
	}

	if weather["wind"] <= 6 {
		statusWind = "aman"
	} else if weather["wind"] >= 7 && weather["wind"] <= 15 {
		statusWind = "siaga"
	} else if weather["wind"] > 15 {
		statusWind = "bahaya"
	}

	fmt.Println("status water :", statusWater)
	fmt.Println("status wind :", statusWind)
}

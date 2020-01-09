package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getEnv(key, def string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return def
}

func getListenAddr() string {
	port := getEnv("PORT", "3003")
	return ":" + port
}

func handleRequest(res http.ResponseWriter, req *http.Request) {
	url := getEnv("TARGET", "http://localhost:8080")

	log.Printf("Target %s%s\n", url, req.URL.Path)

	resp, err := http.Get(url + req.URL.Path)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	res.Write(body)
}

func main() {
	log.Printf("Running on: %s\n", getListenAddr())

	http.HandleFunc("/", handleRequest)
	if err := http.ListenAndServe(getListenAddr(), nil); err != nil {
		panic(err)
	}
}

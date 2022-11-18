package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type App struct {
	Port string
}

func (a *App) Start() {
	http.Handle("/", logreq(index))
	addr := fmt.Sprintf(":%s", a.Port)
	log.Printf("Starting app on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func env(key, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return val
}

func index(w http.ResponseWriter, r *http.Request) {
	var imageName string
	imageName = os.Getenv("IMAGE")
	if imageName == "" {
		imageName = "NO_ENV_IMAGE_SET"
	}
	
	fmt.Fprintf(w, "hello world from image: app:%s\n", imageName)
}

func logreq(f func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("path: %s", r.URL.Path)

		f(w, r)
	})
}

func main() {
	server := App{
		Port: env("PORT", "8080"),
	}
	server.Start()
}

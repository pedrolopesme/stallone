package main

import "github.com/pedrolopesme/stallone/internal/api"

func main() {
	app := api.NewAPI()
	app.Run()
}

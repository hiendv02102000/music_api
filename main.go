package main

import "be_soc/internal/api/router"

func main() {
	r := router.NewRouter()
	r.Engine.Run()

}

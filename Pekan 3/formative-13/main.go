package main

import "f13/routers"

func main() {
	var PORT = ":8080"
	routers.StartServer().Run(PORT)
}

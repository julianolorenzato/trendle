package main

import (
	"github.com/julianolorenzato/choosely/config"
	"github.com/julianolorenzato/choosely/network"
)

//var config = viper.Viper{}
//
//func init() {
//	viper.Set("ADDR", "choosely-redis:5432")
//	fmt.Println(viper.Get("ADDR"), "addr")
//	fmt.Println("first?")
//	// Load environment variables
//	err := godotenv.Load(".env")
//	if err != nil {
//		log.Fatal(err)
//	}
//}

func main() {
	// Start app server
	port := config.Env("PORT")
	network.NewHTTPServer(":" + port).Start()
}

package main

import (
	"github.com/kaepa3/tweet/config"
	"github.com/kaepa3/tweet/tweetapi"
)

func main() {
	conf, _ := config.ReadConfig("conf.toml")

	api := tweetapi.GetTwitterApi(*conf)
	api.Tweet("一旦終わり", "If1oO3.jpg")
}

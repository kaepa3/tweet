package main

import (
	"encoding/base64"
	. "fmt"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/kaepa3/tweet/config"
)

func GetTwitterApi(conf config.TwitterConfig) *anaconda.TwitterApi {
	anaconda.SetConsumerKey(conf.ApiKey)
	anaconda.SetConsumerSecret(conf.ApiKeySecret)
	api := anaconda.NewTwitterApi(conf.AccessToken, conf.AccessTokenSecret)
	return api
}

func main() {
	conf, _ := config.ReadConfig("conf.toml")

	api := GetTwitterApi(*conf)
	str := encode()
	media, _ := api.UploadMedia(str)
	v := url.Values{}
	v.Add("media_ids", media.MediaIDString)

	text := "tweet test"
	tweet, err := api.PostTweet(text, v)
	if err != nil {
		panic(err)
	}

	Print(tweet.Text)
}
func encode() string {

	file, _ := os.Open("If1oO3.jpg")
	defer file.Close()

	fi, _ := file.Stat() //FileInfo interface
	size := fi.Size()    //ファイルサイズ

	data := make([]byte, size)
	file.Read(data)

	return base64.StdEncoding.EncodeToString(data)
}

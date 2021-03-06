package tweetapi

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/kaepa3/tweet/config"
)

type TweetApi struct {
	api *anaconda.TwitterApi
}

func GetTwitterApi(conf config.TwitterConfig) *TweetApi {
	anaconda.SetConsumerKey(conf.ApiKey)
	anaconda.SetConsumerSecret(conf.ApiKeySecret)
	api := &TweetApi{anaconda.NewTwitterApi(conf.AccessToken, conf.AccessTokenSecret)}
	if conf.TimeoutSecond != 0 {
		api.api.HttpClient.Timeout = time.Duration(conf.TimeoutSecond) * time.Second
	}
	return api
}

func (ta *TweetApi) Tweet(text string, imgPath string) {

	str := encode(imgPath)
	media, err := ta.api.UploadMedia(str)
	if err != nil {
		fmt.Println(err)
	}
	v := url.Values{}
	v.Add("media_ids", media.MediaIDString)

	_, err = ta.api.PostTweet(text, v)
	if err != nil {
		panic(err)
	}
	fmt.Println(text)
}

func encode(path string) string {

	if file, err := os.Open(path); err == nil {
		defer file.Close()

		fi, _ := file.Stat() //FileInfo interface
		size := fi.Size()    //ファイルサイズ

		data := make([]byte, size)
		file.Read(data)

		return base64.StdEncoding.EncodeToString(data)
	}
	return ""
}

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-2022/graph/model"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func GetPost() ([]*model.Post, error) {
	var response []*model.Post
	log.Info("get post list")
	res, err := http.Get(fmt.Sprintf("%s/post", os.Getenv("URL_CLIENT")))

	if err != nil {
		log.Error("Get post list call", err.Error())
		return nil, err
	}
	log.Info("success call get post list")

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Error("Parse data comments by post", err.Error())
		return nil, err
	}
	return response, nil
}

func CreatePost(input *model.NewPost) (*model.Post, error) {

	buff := new(bytes.Buffer)
	json.NewEncoder(buff).Encode(input)
	req, _ := http.NewRequest(POST, fmt.Sprintf("%s/post", os.Getenv("URL_CLIENT")), buff)
	client := &http.Client{}
	log.Info("create new post")
	res, err := client.Do(req)
	if err != nil {
		log.Error("create new post")
		return nil, err
	}

	defer res.Body.Close()
	var response *model.Post
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Error("Parse create post", err.Error())
		return nil, err
	}
	return response, err
}

package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/lytics/base62"
)

// urlVal := "http://www.biglongurl.io/?utm_content=content&utm_campaign=campaign"
// 	encodedUrl := base62.StdEncoding.EncodeToString([]byte(urlVal))

// 	// Unencoded it
// 	byteUrl, err := base62.StdEncoding.DecodeString(encodedUrl)

type shortURLSvc struct {
	// hash == base64(origURL)
	// map from hash to orig url
	store map[string]string
}

var ErrURLNotFound = errors.New("URL not found")

var ErrInternal = errors.New("Internal server error")

// shortURL will be https://shortURL.com/v1/url/<shortURL>
func (svc *shortURLSvc) Get(shortURL string) (string, error) {
	s := strings.Split(shortURL, "/")
	hash := s[len(s)-1]
	if origURL, ok := svc.store[hash]; ok {
		return origURL, nil
	}
	return "", ErrURLNotFound
}

func (svc *shortURLSvc) Put(origURL string) (string, error) {
	hash := base62.StdEncoding.EncodeToString([]byte(origURL))
	svc.store[hash] = origURL
	return svc.buildShortURL(hash), nil
}

func (svc *shortURLSvc) buildShortURL(hash string) string {
	return fmt.Sprintf("https://shortURL.com/v1/url/%s", hash)
}

func NewShortURLSvc() *shortURLSvc {
	return &shortURLSvc{store: make(map[string]string, 100)}
}

func main() {
	svc := NewShortURLSvc()
	shortURL, err := svc.Put("https://google.com")
	must(err)

	fmt.Println("shortURL: ", shortURL)

	origURL, err := svc.Get(shortURL)
	must(err)
	fmt.Println("orig: ", origURL)
}

// impl must() for me, it should panic when err is not nil
func must(err error) {
	if err != nil {
		panic(err)
	}
}

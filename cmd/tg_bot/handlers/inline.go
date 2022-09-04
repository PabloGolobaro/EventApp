package handlers

import (
	tele "gopkg.in/telebot.v3"
	"strconv"
)

var InlinePhotos = func(c tele.Context) error {
	urls := []string{
		"https://miro.medium.com/max/1200/0*SoqCeEz9EctJBXKw.png",
		"https://www.freecodecamp.org/news/content/images/2021/10/golang.png",
	}

	results := make(tele.Results, len(urls)) // []tele.Result
	for i, url := range urls {
		result := &tele.PhotoResult{
			URL:      url,
			ThumbURL: url, // required for photos
		}

		results[i] = result
		// needed to set a unique string ID for each result
		results[i].SetResultID(strconv.Itoa(i))
	}

	return c.Answer(&tele.QueryResponse{
		Results:   results,
		CacheTime: 60, // a minute
	})
}

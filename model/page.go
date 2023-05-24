package model

import (
	"encoding/json"
	"log"
)

var pages map[string]map[string]interface{}

func PageInit() {
	pages = make(map[string]map[string]interface{}, 10)
}

func PageUpdate(page string, name string, value interface{}) {
	p, ok := pages[page]
	if !ok {
		newPage := make(map[string]interface{})
		newPage[name] = value
		pages[page] = newPage
	} else {
		p[name] = value
	}
}

func PageUpload() {
	for k, v := range pages {
		data, err := json.Marshal(v)
		if err != nil {
			log.Printf("pageupload json marshal error: %s", err)
			log.Println(v)
		}
		RedisClient.Set(k, data, 0)
	}
}

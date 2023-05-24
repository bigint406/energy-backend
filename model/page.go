package model

import (
	"encoding/json"
	"log"
)

type Pages struct {
	pages map[string](map[string]interface{})
}

func (p *Pages) PageInit() {
	p.pages = make(map[string]map[string]interface{}, 10)
}

func (p *Pages) PageUpdate(page string, name string, value interface{}) {
	pp, ok := p.pages[page]
	if !ok {
		newPage := make(map[string]interface{})
		newPage[name] = value
		p.pages[page] = newPage
	} else {
		pp[name] = value
	}
}

func (p *Pages) PageUpload() {
	for k, v := range p.pages {
		data, err := json.Marshal(v)
		if err != nil {
			log.Printf("pageupload json marshal error: %s", err)
			log.Println(v)
		}
		RedisClient.Set(k, data, 0)
	}
}

package main

import (
	"fmt"
	"math/rand"
	"strings"

	"rushsteve1.us/monolith/shared"
)

type OverseerRpc struct {
	Config shared.Config
}

func (ov *OverseerRpc) Top(_ int, out *string) error {
	*out = TopSup.String()
	return nil
}

func (ov *OverseerRpc) DBPath(_ int, out *string) error {
	*out = ov.Config.Database.String()
	return nil
}

var testingPosts = map[string]string{
	"1": "hello",
}

func (ov *OverseerRpc) ListBlog(_ int, out *map[string]string) error {
	*out = make(map[string]string, len(testingPosts))
	for k, v := range testingPosts {
		if len(v) > 20 {
			v = v[:20] + "…"
		}
		(*out)[k] = strings.TrimSpace(v)
	}
	return nil
}

func (ov *OverseerRpc) GetBlogPost(id string, out *string) error {
	*out = testingPosts[id]
	return nil
}

func (ov *OverseerRpc) SetBlogPost(data struct {
	Id   string
	Body string
}, out *int) error {
	testingPosts[data.Id] = data.Body
	*out = 0
	return nil
}

func (ov *OverseerRpc) NewBlogPost(body string, out *string) error {
	id := fmt.Sprint(rand.Int() % 100)
	testingPosts[id] = body
	*out = id
	return nil
}

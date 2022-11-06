package main

import (
	"context"
	"fmt"
	"strings"

	"rushsteve1.us/monolith/shared"
	"rushsteve1.us/monolith/webserver"
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

func (ov *OverseerRpc) ListBlog(_ int, out *map[int64]string) error {
	ws, ok := ServiceMap["WebServer"].Service.(*webserver.WebServer)
	if !ok {
		return fmt.Errorf("could not cast to WebServer")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := ws.DBConn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	posts, err := webserver.ListPosts(conn, ctx)
	if err != nil {
		return err
	}

	*out = make(map[int64]string, len(posts))
	for _, post := range posts {
		(*out)[post.ID] = post.Title
	}

	return nil
}

func (ov *OverseerRpc) GetBlogPost(id int64, out *string) error {
	ws, ok := ServiceMap["WebServer"].Service.(*webserver.WebServer)
	if !ok {
		return fmt.Errorf("could not cast to WebServer")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := ws.DBConn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	post, err := webserver.GetPost(conn, ctx, id)
	if err != nil {
		return err
	}

	*out = post.Title + "\n---\n" + post.Body
	return nil
}

func (ov *OverseerRpc) SetBlogPost(data struct {
	Id   int64
	Body string
}, out *int) error {
	ws, ok := ServiceMap["WebServer"].Service.(*webserver.WebServer)
	if !ok {
		return fmt.Errorf("could not cast to WebServer")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	v := strings.Split(data.Body, "\n---\n")
	title := v[0]
	body := v[1]

	conn, err := ws.DBConn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	err = webserver.UpdatePost(conn, ctx, data.Id, title, body)
	if err != nil {
		return err
	}

	*out = 0
	return nil
}

func (ov *OverseerRpc) NewBlogPost(body string, out *int) error {
	ws, ok := ServiceMap["WebServer"].Service.(*webserver.WebServer)
	if !ok {
		return fmt.Errorf("could not cast to WebServer")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	v := strings.Split(body, "\n---\n")
	title := v[0]
	body = v[1]
	
	conn, err := ws.DBConn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	err = webserver.InsertPost(conn, ctx, title, body)
	if err != nil {
		return err
	}

	*out = 0
	return nil
}

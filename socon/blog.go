package main

import (
	"fmt"
	"io"
	"net/rpc"
	"os"
	"os/exec"
	"strconv"

	"github.com/urfave/cli/v2"
)

func listBlog(ctx *cli.Context) error {
	client, err := rpc.DialHTTP("tcp", ctx.String("addr"))
	if err != nil {
		return err
	}
	defer client.Close()

	var reply map[int64]string
	err = client.Call("Overseer.ListBlog", 0, &reply)
	if err != nil {
		return err
	}

	fmt.Printf("ID\tText\n---------------------\n")
	for id, short := range reply {
		fmt.Printf("%d\t%s\n", id, short)
	}

	return nil

}

func editBlogPost(ctx *cli.Context) error {
	client, err := rpc.DialHTTP("tcp", ctx.String("addr"))
	if err != nil {
		return err
	}
	defer client.Close()

	id, err := strconv.Atoi(ctx.Args().First())
	if err != nil {
		return err
	}

	var reply string
	err = client.Call("Overseer.GetBlogPost", id, &reply)
	if err != nil {
		return err
	}

	file, err := os.CreateTemp("", "monolith-blog-")
	if err != nil {
		return err
	}
	defer os.Remove(file.Name())

	_, err = file.WriteString(reply)
	if err != nil {
		return err
	}

	path, err := exec.LookPath(os.Getenv("EDITOR"))
	if err != nil {
		return err
	}

	cmd := exec.Command(path, file.Name())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	file.Seek(0, 0)
	body, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	var out int
	return client.Call("Overseer.SetBlogPost", struct {
		Id   int64
		Body string
	}{Id: int64(id), Body: string(body)}, &out)
}

func newBlogPost(ctx *cli.Context) error {
	client, err := rpc.DialHTTP("tcp", ctx.String("addr"))
	if err != nil {
		return err
	}
	defer client.Close()

	file, err := os.CreateTemp("", "monolith-blog-*.html")
	if err != nil {
		return err
	}
	defer os.Remove(file.Name())

	file.WriteString("Title Here\n---\nBody Here")
	file.Seek(0, 0)

	path, err := exec.LookPath(os.Getenv("EDITOR"))
	if err != nil {
		return err
	}

	cmd := exec.Command(path, file.Name())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	body, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	var out int
	return client.Call("Overseer.NewBlogPost", string(body), &out)
}

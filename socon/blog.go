package main

import (
	"fmt"
	"io"
	"net/rpc"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func listBlog(ctx *cli.Context) error {
	client, err := rpc.DialHTTP("tcp", ctx.String("addr"))
	if err != nil {
		return err
	}
	defer client.Close()

	var reply map[string]string
	err = client.Call("Overseer.ListBlog", 0, &reply)
	if err != nil {
		return err
	}

	fmt.Printf("ID\tText\n---------------------\n")
	for id, short := range reply {
		fmt.Printf("%s\t%s\n", id, short)
	}

	return nil

}

func editBlogPost(ctx *cli.Context) error {
	client, err := rpc.DialHTTP("tcp", ctx.String("addr"))
	if err != nil {
		return err
	}
	defer client.Close()

	id := ctx.Args().First()

	var reply string
	err = client.Call("Overseer.GetBlogPost", id, &reply)
	if err != nil {
		return err
	}

	file, err := os.CreateTemp("", "monolith-blog-")
	if err != nil {
		return err
	}
	defer file.Close()

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
		Id   string
		Body string
	}{Id: id, Body: string(body)}, &out)
}

func newBlogPost(ctx *cli.Context) error {
	client, err := rpc.DialHTTP("tcp", ctx.String("addr"))
	if err != nil {
		return err
	}
	defer client.Close()

	file, err := os.CreateTemp("", "monolith-blog-")
	if err != nil {
		return err
	}
	defer file.Close()

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

	var id string
	err = client.Call("Overseer.NewBlogPost", string(body), &id)
	if err != nil {
		return err
	}
	fmt.Printf("New blog post created with ID %s", id)
	return nil
}

package main

import (
	"fmt"
	"net/rpc"
	"os"
	"os/exec"
	"syscall"

	"github.com/urfave/cli/v2"
)

var cliApp = &cli.App{
	EnableBashCompletion:   true,
	UseShortOptionHandling: true,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "addr",
			Value: "localhost:9900",
			Usage: "Address of the Overseer",
		},
	},
	Commands: []*cli.Command{
		{
			Name:   "top",
			Usage:  "Service status of the Overseer",
			Action: listServices,
		},
		{
			Name:   "sql",
			Usage:  "Connect to the SQL database directly",
			Action: execSqlTool,
		},
		{
			Name:  "blog",
			Usage: "Manage blog posts",
			Subcommands: []*cli.Command{
				{
					Name:   "list",
					Usage:  "List blog posts",
					Action: listBlog,
				},
				{
					Name:   "edit",
					Usage:  "Edit the given blog post",
					Action: editBlogPost,
				},
				{
					Name:   "new",
					Usage:  "Create a new blog post",
					Action: newBlogPost,
				},
			},
		},
	},
}

func main() {
	err := cliApp.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func listServices(ctx *cli.Context) error {
	client, err := rpc.DialHTTP("tcp", ctx.String("addr"))
	if err != nil {
		return err
	}
	defer client.Close()

	var reply string
	err = client.Call("Overseer.Top", 0, &reply)
	if err != nil {
		return err
	}

	fmt.Print(reply)
	return nil
}

const sqlTool = "sqlite3"

func execSqlTool(ctx *cli.Context) error {
	client, err := rpc.DialHTTP("tcp", ctx.String("addr"))
	if err != nil {
		return err
	}
	defer client.Close()

	var reply string
	err = client.Call("Overseer.DBPath", 0, &reply)
	if err != nil {
		return err
	}

	path, err := exec.LookPath(sqlTool)
	if err != nil {
		return err
	}

	return syscall.Exec(path, append([]string{sqlTool, reply}, ctx.Args().Slice()...), os.Environ())
}

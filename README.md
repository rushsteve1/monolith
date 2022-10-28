# Steve's Monolith

This repo contains all my personal services as one big Mono-repo.
Everything is written in [Go](https://go.dev)

This code is licensed under the terms of the [Apache 2.0 license](./LICENSE.txt)

## Modules

As a mono-repo there are a number of [Go modules] contained within. Each module
is intended to be self-sufficient, generally only requiring the `shared` module

- [`shared`](./shared/) the code shared across the repository
- [`webserver`](./webserver/) the core webserver powering [rushsteve1.us](https://rushsteve1.us)
- [`swissarmybot`](./swissarmybot/) the n<sup>th</sup> incarnation of my discord bot
- [`overseer`](./overseer/) brings everything together with a supervisor
- [`socon`](./socon/) controller for the Overseer over RPC

Most of these modules can be built and used standalone.

### Other Folders

- [`static`](./static/) contains the static assets that will be served by Caddy
- [`cgi-bin`](./cgi-bin/) contains [CGI scripts](https://en.wikipedia.org/wiki/Common_Gateway_Interface) that the WebServer will run

## Major Dependencies

I have tried to limit the number of direct dependencies that this project relies
on, but there are of course a few anyway:

- `modernc.org/sqlite`
- `github.com/bwmarrin/discordgo`
- `github.com/urfave/cli/v2`
- `github.com/robfig/cron/v3`
- `github.com/sirupsen/logrus`
- `github.com/thejerf/suture/v4`
- `bolt.css`

In addition to these I would like to highlight the *excellent*
[Caddy server](https://caddyserver.com/) which much of this project is built around.

Thank you to the authors of all of these libraries,
and all the others that are not listed here. Y'all are the real heroes.

## Deployment

TBD probably Podman/Docker

## Configuration

All configuration is powered by a `config.*.json` where `*` depends on the deployment
target. An [example config is provided](./config.example.json).
The VSCode config expects a `test` config, and the Dockerfile expects a `docker` config.

All modules load this config file and simply use the portions of it that they need.

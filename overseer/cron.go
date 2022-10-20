package main

import (
	"context"
	"fmt"

	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"rushsteve1.us/monolith/shared"
)

type Cron struct {
	Config shared.Config
}

func (cr *Cron) Serve(ctx context.Context) error {
	c := cron.New(cron.WithLogger(cron.PrintfLogger(log.New())))

	log.Info("Cron started")
	c.Run()
	return fmt.Errorf("%s exited unexpectedly", cr.Name())
}

func (cr Cron) Addr() string {
	return "none"
}

func (cr Cron) Name() string {
	return "Cron"
}

func (cr Cron) UseFcgi() bool {
	return false
}

func (cr Cron) String() string {
	return cr.Name()
}

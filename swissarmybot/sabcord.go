package swissarmybot

import (
	"context"
	"database/sql"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"rushsteve1.us/monolith/shared"
)

type SwissArmyBotDiscord struct {
	config shared.Config
	db     *sql.DB
	session *discordgo.Session
}

func (sab *SwissArmyBotDiscord) Serve(ctx context.Context) error {
	var err error
	setupOnce.Do(func() { err = setup(sab.db, ctx) })
	if err != nil {
		return err
	}

	sab.session, err = discordgo.New("Bot " + sab.config.SwissArmyBot.Token)
	if err != nil {
		return err
	}
	defer sab.session.Close()

	sab.session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Infof("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	sab.session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate){
		// TODO
		log.Tracef("%#v", i)
	})

	err = sab.session.Open()
	if err != nil {
		return err
	}

	if sab.config.SwissArmyBot.Cleanup {
		defer sab.cleanup()
	}

	<-ctx.Done()
	return ctx.Err()
}

func (sab SwissArmyBotDiscord) Addr() string {
	return ""
}

func (sab SwissArmyBotDiscord) Name() string {
	return "SwissArmyBot Discord"
}

func (sab SwissArmyBotDiscord) UseFcgi() bool {
	return false
}

func (sab SwissArmyBotDiscord) String() string {
	return sab.Name()
}

func (sab SwissArmyBotDiscord) cleanup() {
	// TODO	
	log.Info("Cleaning up commands")
}

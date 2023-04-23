package main

import (
	"log"
	//to get all secret tokens
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	
	"assignment4/pkg"
	env"assignment4"
	
	"os"
	"os/signal"
	"context"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()


	b, err := bot.New(env.TELEGRAM_BOT_TOKEN)
	if err != nil {
		log.Fatal(err)
	}
	
	b.RegisterHandler(bot.HandlerTypeMessageText, "/image", bot.MatchTypeExact, withGoR)

	b.Start(ctx)
}

func withGoR(ctx context.Context, b *bot.Bot, update *models.Update) {
	go pkg.SendPhoto(ctx, b, update)
}

package pkg

import (
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"context"
)


func SendPhoto(ctx context.Context, b *bot.Bot, update *models.Update) {
	image := RandomPhoto()// Getting random photo from unsplash, func in usnplash.go
		
	params := &bot.SendPhotoParams{
    ChatID: update.Message.Chat.ID,
    Photo: 	&models.InputFileString{Data: image},
	}
	b.SendPhoto(ctx, params)
}

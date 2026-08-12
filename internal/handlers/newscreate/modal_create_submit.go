package newscreate

import (
	"STVbotGo/internal/backend"
	"STVbotGo/internal/config"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func handleCreateNewsModalSubmit(s *discordgo.Session, i *discordgo.InteractionCreate, cfg config.Env) {
	var err error

	if i.ModalSubmitData().CustomID != createNewsModalID {
		return
	}

	if !checkAccessRole(s, i, cfg) {
		respondEphemeral(s, i, "❌ У вас немає прав на створення новин.")
		return
	}

	if checkUserHasTimeout(s, i, cfg) {
		respondEphemeral(s, i, "❌ У вас є активний Тайм-аут")
		return
	}

	if cfg.ApiKey == "" {
		respondEphemeral(s, i, "❌ API ключ для створення новин не налаштований.")
		return
	}

	if cfg.UserNewsChannel == "" {
		respondEphemeral(s, i, "❌ Канал для новин користувачів не налаштований.")
		return
	}

	authorID := getInteractionUserID(i)
	if authorID == "" {
		respondEphemeral(s, i, "❌ Не вдалося визначити автора новини.")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	news := backend.CreateNewsRequest{
		Title:     sanitizeNewsInput(getModalValue(i, "title")),
		Anons:     sanitizeNewsInput(getModalValue(i, "anons")),
		FullText:  sanitizeNewsInput(getModalValue(i, "fullText")),
		Thumbnail: sanitizeNewsInput(getModalValue(i, "thumbnail")),
		Author:    authorID,
	}

	if news.Title == "" || news.Anons == "" || news.FullText == "" {
		respondEphemeral(s, i, "❌ Заповніть усі обовʼязкові поля.")
		return
	}

	if news.Thumbnail != "" && !isAllowedThumbnailURL(news.Thumbnail) {
		respondEphemeral(s, i, "❌ URL фото має починатися з `https` або `http`.")
		return
	}

	client := backend.NewClient(
		"https://api.stvworlds.win",
		cfg.ApiKey,
	)

	var post *backend.Post

	post, err = client.CreatePost(ctx, news)
	if err != nil {
		log.Printf("failed to create news: %v", err)
		respondEphemeral(s, i, fmt.Sprintf("❌ Не вдалося створити новину: %v", err))
		return
	}

	newsURL := blogPageBaseURL + fmt.Sprintf("%v", post.Id)

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "✅ Новину успішно створено.",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		log.Printf("failed to respond to create news modal: %v", err)
	}

	_, err = s.ChannelMessageSendComplex(cfg.UserNewsChannel, &discordgo.MessageSend{
		Components: buildNewsMessageComponents(news, fmt.Sprintf("%v", post.Id), newsURL),
		Flags:      messageFlagsIsComponentsV2,
	})
	if err != nil {
		log.Printf("failed to send news created component to channel %s: %v", cfg.UserNewsChannel, err)
		return
	}
}

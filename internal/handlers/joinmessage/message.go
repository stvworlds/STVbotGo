package joinmessage

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func buildWelcomeComponents(user *discordgo.User) []discordgo.MessageComponent {
	return []discordgo.MessageComponent{
		discordgo.Container{
			Components: []discordgo.MessageComponent{
				discordgo.Section{
					Components: []discordgo.MessageComponent{
						discordgo.TextDisplay{
							Content: buildWelcomeMessage(user),
						},
					},
					Accessory: discordgo.Thumbnail{
						Media: discordgo.UnfurledMediaItem{
							URL: user.AvatarURL("1024"),
						},
						Description: ptr("User"),
					},
				},
				buildLinksRow(),
				discordgo.Separator{
					Divider: ptr(true),
				},
				discordgo.TextDisplay{
					Content: "**З повагою, команда STEVIOUS**",
				},
			},
		},
	}
}

func buildWelcomeMessage(user *discordgo.User) string {
	return fmt.Sprintf(
		"**👋 Вітаємо на STEVIOUS, %s!**\n\n"+
			"Ви приєдналися до українського медійного ванільного сервера Minecraft, що працює у форматі сезонів. "+
			"Перед початком гри рекомендуємо ознайомитися з основними каналами, правилами та можливостями проєкту.\n\n"+
			"**У нас є два основні напрямки для старту:**\n\n"+
			"1. **🧭 Почати гру на сезонному ванільному світі**\n"+
			"Грайте разом з іншими учасниками, розвивайте спільні бази, долучайтеся до колективних проєктів і подій. "+
			"Наш фокус — стабільність, справедливі правила та якісний геймплей без зайвих модифікацій.\n\n"+
			"2. **🎫 Підтримати сервер і отримати додаткові можливості**\n"+
			"Ви можете задонатити на розвиток платформи та отримати косметичні або зручні, але неігрові переваги. "+
			"Усі внески спрямовуються на підтримку технічної інфраструктури та розвиток спільноти.\n\n"+
			"**🎁 Для нових гравців доступний промокод на знижку**, який можна активувати на сайті.\n\n"+
			"**❗ Важливо**\n"+
			"Будь-які дії на сервері впливають на інших гравців та інфраструктуру сезону. "+
			"Заборонені токсичність, зловмисні дії, порушення правил співпраці та руйнування спільних проєктів. "+
			"За це можливі попередження, штрафи або блокування доступу.\n\n"+
			"**🔗 Посилання проєкту:** ❤️",
		user.Username,
	)
}

func buildLinksRow() discordgo.MessageComponent {
	return discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{
			discordgo.Button{
				Label: "Правила",
				Style: discordgo.LinkButton,
				URL:   rulesURL,
				Emoji: &discordgo.ComponentEmoji{
					Name: "📌",
				},
			},
			discordgo.Button{
				Label: "Новини",
				Style: discordgo.LinkButton,
				URL:   newsURL,
				Emoji: &discordgo.ComponentEmoji{
					Name: "🔔",
				},
			},
			discordgo.Button{
				Label: "Підтримка",
				Style: discordgo.LinkButton,
				URL:   supportURL,
				Emoji: &discordgo.ComponentEmoji{
					Name: "📫",
				},
			},
			discordgo.Button{
				Label: "Відгуки",
				Style: discordgo.LinkButton,
				URL:   reviewsURL,
				Emoji: &discordgo.ComponentEmoji{
					Name: "🤝",
				},
			},
		},
	}
}

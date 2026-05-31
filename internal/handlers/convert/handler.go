package convert

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func commandConvert(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	if i.ApplicationCommandData().Name != convertCommandName {
		return
	}

	options := i.ApplicationCommandData().Options
	if len(options) == 0 {
		respondWithError(s, i, "❌ Передай файл для конвертації.")
		return
	}

	attachmentID, ok := options[0].Value.(string)
	if !ok {
		respondWithError(s, i, "❌ Некоректний аргумент файлу.")
		return
	}

	attachment, ok := i.ApplicationCommandData().Resolved.Attachments[attachmentID]
	if !ok || attachment == nil {
		respondWithError(s, i, "❌ Не вдалося знайти прикріплений файл.")
		return
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "⏳ Конвертую файл у GIF...",
		},
	})
	if err != nil {
		log.Printf("failed to defer /convert response: %v", err)
		return
	}

	gifPath, cleanup, err := convertAttachmentToGIF(attachment)
	if cleanup != nil {
		defer cleanup()
	}

	if err != nil {
		log.Printf("failed to convert file to gif: %v", err)
		sendConvertFollowup(s, i, fmt.Sprintf("❌ Помилка конвертації: `%s`", err.Error()), nil)
		return
	}

	file, err := os.Open(gifPath)
	if err != nil {
		log.Printf("failed to open gif file: %v", err)
		sendConvertFollowup(s, i, "❌ Не вдалося відкрити готовий GIF.", nil)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	sendConvertFollowup(s, i, "✅ Готово! Ось твій GIF:", &discordgo.File{
		Name:   "converted.gif",
		Reader: file,
	})
}

func convertAttachmentToGIF(attachment *discordgo.MessageAttachment) (string, func(), error) {
	tempDir, err := os.MkdirTemp("", "stvbot-convert-*")
	if err != nil {
		return "", nil, fmt.Errorf("не вдалося створити тимчасову папку")
	}

	cleanup := func() {
		if err := os.RemoveAll(tempDir); err != nil {
			log.Printf("failed to cleanup temp dir: %v", err)
		}
	}

	inputExt := strings.ToLower(filepath.Ext(attachment.Filename))
	if inputExt == "" {
		inputExt = ".input"
	}

	inputPath := filepath.Join(tempDir, "input"+inputExt)
	outputPath := filepath.Join(tempDir, fmt.Sprintf("converted-%d.gif", time.Now().UnixNano()))

	if err := downloadFile(attachment.URL, inputPath); err != nil {
		cleanup()
		return "", nil, err
	}

	err = ffmpeg.
		Input(inputPath).
		Output(outputPath, ffmpeg.KwArgs{
			"filter_complex": filterComplex,
			"loop":           "0",
		}).
		OverWriteOutput().
		ErrorToStdOut().
		Run()

	if err != nil {
		cleanup()
		return "", nil, fmt.Errorf("ffmpeg-go не зміг конвертувати файл: %w", err)
	}

	return outputPath, cleanup, nil
}

func downloadFile(url string, destination string) error {
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("не вдалося завантажити файл")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return fmt.Errorf("discord повернув статус %d при завантаженні файлу", response.StatusCode)
	}

	file, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("не вдалося створити тимчасовий файл")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	if _, err := io.Copy(file, response.Body); err != nil {
		return fmt.Errorf("не вдалося зберегти файл")
	}

	return nil
}

func respondWithError(s *discordgo.Session, i *discordgo.InteractionCreate, message string) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		log.Printf("failed to respond with error: %v", err)
	}
}

func sendConvertFollowup(s *discordgo.Session, i *discordgo.InteractionCreate, content string, file *discordgo.File) {
	message := &discordgo.WebhookParams{
		Content: content,
	}

	if file != nil {
		message.Files = []*discordgo.File{file}
	}

	if _, err := s.FollowupMessageCreate(i.Interaction, true, message); err != nil {
		log.Printf("failed to send /convert followup: %v", err)
	}
}

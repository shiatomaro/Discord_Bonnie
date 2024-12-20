package main

import (
	"context"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/sashabaranov/go-openai"
)

func main() {
	// Get bot token from the environment
	token := os.Getenv("DISCORD_BOT_TOKEN")
	apiKey := os.Getenv("OPENAI_API_KEY")
	if token == "" || apiKey == "" {
		log.Fatal("No Token Provided or Required Variables missing. set DISCORD_BOT_TOKEN in environment or OPENAI_API_KEY")
	}

	// Creating new OpenAI client
	openAIClient := openai.NewClient(apiKey)

	// Creating new session
	bot, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("error creating discord session: %v", err)
	}

	// Register message handler
	bot.AddHandler(messageCreateHandler)

	// Opening connection to discord
	err = bot.Open()
	if err != nil {
		log.Fatalf("Error opening connection to discord: %v", err)
	}
	log.Println("Bot is running. Press CTRL+C to exit!")
	select {} // to keep running the bot
}

// Handle message create event
func messageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Respond to messages using "!chat"
	if len(m.Content) > 5 && m.Content[:5] == "!chat" {
		// Get message after "!chat"
		userMessage := m.Content[6:]

		// Query OpenAI
		response, err := queryChatGPT(s, userMessage)
		if err != nil {
			log.Printf("Error querying OpenAI: %v", err)
			s.ChannelMessageSend(m.ChannelID, "Hello, Sorry I couldn't process your request")
			return
		}

		// Send the response to Discord
		s.ChannelMessageSend(m.ChannelID, response)
	}
}

// Query OpenAI API with the user's message
func queryChatGPT(s *discordgo.Session, message string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(apiKey)

	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{Role: "system", Content: "You are a helpful assistant."},
			{Role: "user", Content: message},
		},
	}

	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

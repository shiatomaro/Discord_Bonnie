import (
	"os"
	"log"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// get bot token from the environment
	token := os.Geten("DISCORD_BOT_TOKEN")
	if token == "" {
		log.fatal("No Token Provided. set DISCORD_BOT_TOKEN in environment")
	}
	// Creating new session
	bot, err = discordgoNew("Bot" + token)
	if err != nil {
		log.fatal("error creating discord session: %v", err)
	}
	log.Println("Bot is running. Press CTRL+C to exit.")
	select {}
}

// handle incoming messages
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// ignore messages from bot itself
	IF m.Author.ID == s.State.User.ID {
		return
	}
	// reponse to messages
	if m.Content == !chat {
		s.ChannelMessageSend(m.ChannelID, "Hello, How can I help you?")
	}
}
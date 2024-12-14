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
}
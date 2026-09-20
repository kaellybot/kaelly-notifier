package constants

import amqp "github.com/kaellybot/kaelly-amqp"

const cdnURL = "https://raw.githubusercontent.com/kaellybot/kaelly-cdn/refs/heads/main"

// AnkamaGame holds everything that differs per game. One Discord application per
// game, so the bot identity and its token are resolved from the game of the news
// being handled, never from the service's own identity.
type AnkamaGame struct {
	Name            string
	Icon            string
	AMQPGame        amqp.Game
	BotName         string
	BotAvatarURL    string
	DiscordTokenKey string
}

func GetGames() []AnkamaGame {
	return []AnkamaGame{
		{
			Name:            "DOFUS",
			Icon:            cdnURL + "/common/logos/dofus.webp",
			AMQPGame:        amqp.Game_DOFUS_GAME,
			BotName:         "KaellyBot",
			BotAvatarURL:    cdnURL + "/kaellybot/id/face.webp",
			DiscordTokenKey: DiscordToken,
		},
		{
			Name:            "DOFUS Touch",
			Icon:            cdnURL + "/common/logos/dofus_touch.webp",
			AMQPGame:        amqp.Game_DOFUS_TOUCH,
			BotName:         "KaellyTouch",
			BotAvatarURL:    cdnURL + "/kaellytouch/id/face.webp",
			DiscordTokenKey: DiscordTokenDofusTouch,
		},
		{
			// No Discord application: Retro is not served by any bot, so it has no
			// token key and no session is ever built for it.
			Name:     "DOFUS Retro",
			Icon:     cdnURL + "/common/logos/dofus_retro.webp",
			AMQPGame: amqp.Game_DOFUS_RETRO,
		},
	}
}

// GetServedGames lists the games a bot actually posts for: those with their own
// Discord application. A game with no application has no session and no emojis.
func GetServedGames() []AnkamaGame {
	served := make([]AnkamaGame, 0, len(GetGames()))
	for _, game := range GetGames() {
		if game.DiscordTokenKey != "" {
			served = append(served, game)
		}
	}

	return served
}

// GetGame returns the game's identity, and whether it is a game this service knows.
// It does not fall back on a generic "Ankama" identity: a news message whose game is
// unknown must be dropped, not posted under the wrong branding.
func GetGame(amqpGame amqp.Game) (AnkamaGame, bool) {
	for _, game := range GetGames() {
		if game.AMQPGame == amqpGame {
			return game, true
		}
	}

	return AnkamaGame{}, false
}

package constants

import amqp "github.com/kaellybot/kaelly-amqp"

type AnkamaGame struct {
	Name     string
	Icon     string
	AMQPGame amqp.Game
}

func GetGames() []AnkamaGame {
	return []AnkamaGame{
		{
			Name:     "DOFUS",
			Icon:     "https://raw.githubusercontent.com/KaellyBot/Kaelly-cdn/refs/heads/main/common/logos/dofus.webp",
			AMQPGame: amqp.Game_DOFUS_GAME,
		},
		{
			Name:     "DOFUS Touch",
			Icon:     "https://raw.githubusercontent.com/KaellyBot/Kaelly-cdn/refs/heads/main/common/logos/dofus_touch.webp",
			AMQPGame: amqp.Game_DOFUS_TOUCH,
		},
		{
			Name:     "DOFUS Retro",
			Icon:     "https://raw.githubusercontent.com/KaellyBot/Kaelly-cdn/refs/heads/main/common/logos/dofus_retro.webp",
			AMQPGame: amqp.Game_DOFUS_RETRO,
		},
	}
}

func GetGame(amqpGame amqp.Game) AnkamaGame {
	for _, game := range GetGames() {
		if game.AMQPGame == amqpGame {
			return game
		}
	}

	return AnkamaGame{
		Name:     "Ankama",
		Icon:     "https://raw.githubusercontent.com/KaellyBot/Kaelly-cdn/refs/heads/main/common/logos/ankama.webp",
		AMQPGame: amqp.Game_ANY_GAME,
	}
}

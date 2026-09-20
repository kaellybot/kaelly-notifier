package constants

import (
	"testing"

	amqp "github.com/kaellybot/kaelly-amqp"
)

// An unknown game must be reported as such, not branded with a generic identity:
// posting under the wrong bot is worse than not posting at all.
func TestGetGame(t *testing.T) {
	t.Parallel()

	tests := []struct {
		game      amqp.Game
		wantFound bool
		wantBot   string
	}{
		{game: amqp.Game_DOFUS_GAME, wantFound: true, wantBot: "KaellyBot"},
		{game: amqp.Game_DOFUS_TOUCH, wantFound: true, wantBot: "KaellyTouch"},
		{game: amqp.Game_DOFUS_RETRO, wantFound: true, wantBot: ""},
		{game: amqp.Game_ANY_GAME, wantFound: false, wantBot: ""},
	}

	for _, test := range tests {
		t.Run(test.game.String(), func(t *testing.T) {
			t.Parallel()

			game, found := GetGame(test.game)
			if found != test.wantFound {
				t.Errorf("found = %v, want %v", found, test.wantFound)
			}
			if game.BotName != test.wantBot {
				t.Errorf("BotName = %q, want %q", game.BotName, test.wantBot)
			}
		})
	}
}

// A game with a bot needs a token key, otherwise no session can ever be built for
// it; a game without a bot must have none, so no session is attempted.
func TestGamesWithABotHaveATokenKey(t *testing.T) {
	t.Parallel()

	for _, game := range GetGames() {
		hasBot := game.BotName != ""
		hasKey := game.DiscordTokenKey != ""
		if hasBot != hasKey {
			t.Errorf("%v: BotName=%q but DiscordTokenKey=%q", game.AMQPGame,
				game.BotName, game.DiscordTokenKey)
		}
		if hasBot && game.BotAvatarURL == "" {
			t.Errorf("%v: bot %q has no avatar", game.AMQPGame, game.BotName)
		}
	}
}

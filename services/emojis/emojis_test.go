package emojis

import (
	"testing"

	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/kaellybot/kaelly-notifier/models/entities"
)

// stubRepository answers per game, as the real one does through its join.
type stubRepository struct {
	byGame map[amqp.Game][]entities.Emoji
	err    error
}

func (stub stubRepository) GetEmojis(game amqp.Game) ([]entities.Emoji, error) {
	if stub.err != nil {
		return nil, stub.err
	}
	return stub.byGame[game], nil
}

func itemEmoji(id, name, snowflake string) entities.Emoji {
	return entities.Emoji{
		ID: id, Type: constants.EmojiTypeItem, DiscordName: "resource",
		Name: name, Snowflake: snowflake,
	}
}

// byGame builds the lookup rather than declaring it: a map literal keyed by an
// enum has to name every value of it.
type gameEmojis struct {
	game   amqp.Game
	emojis []entities.Emoji
}

func byGame(pairs ...gameEmojis) map[amqp.Game][]entities.Emoji {
	result := make(map[amqp.Game][]entities.Emoji)
	for _, pair := range pairs {
		result[pair.game] = pair.emojis
	}

	return result
}

// The acceptance case: the same emoji ID has a different snowflake per game, and
// a message must never carry the other application's one.
func TestGetItemTypeStringEmoji_IsScopedToTheGame(t *testing.T) {
	t.Parallel()

	resource := amqp.ItemType_RESOURCE_TYPE.String()
	service, err := New(stubRepository{byGame: byGame(
		gameEmojis{amqp.Game_DOFUS_GAME, []entities.Emoji{itemEmoji(resource, "📦", "dofus-sf")}},
		gameEmojis{amqp.Game_DOFUS_TOUCH, []entities.Emoji{itemEmoji(resource, "📦", "touch-sf")}},
	)})
	if err != nil {
		t.Fatalf("New() = %v, want nil", err)
	}

	tests := []struct {
		game amqp.Game
		want string
	}{
		{amqp.Game_DOFUS_GAME, "<:resource:dofus-sf>"},
		{amqp.Game_DOFUS_TOUCH, "<:resource:touch-sf>"},
	}
	for _, test := range tests {
		got := service.GetItemTypeStringEmoji(test.game, amqp.ItemType_RESOURCE_TYPE)
		if got != test.want {
			t.Errorf("%v = %q, want %q", test.game, got, test.want)
		}
	}
}

// An emoji uploaded to only one application must render as its unicode name for
// the other, never as a snowflake that application does not own.
func TestGetItemTypeStringEmoji_FallsBackOnNameWithoutSnowflake(t *testing.T) {
	t.Parallel()

	resource := amqp.ItemType_RESOURCE_TYPE.String()
	service, err := New(stubRepository{byGame: byGame(
		gameEmojis{amqp.Game_DOFUS_GAME, []entities.Emoji{itemEmoji(resource, "📦", "dofus-sf")}},
		// Loaded, but with no snowflake for this application.
		gameEmojis{amqp.Game_DOFUS_TOUCH, []entities.Emoji{itemEmoji(resource, "📦", "")}},
	)})
	if err != nil {
		t.Fatalf("New() = %v, want nil", err)
	}

	if got := service.GetItemTypeStringEmoji(amqp.Game_DOFUS_TOUCH, amqp.ItemType_RESOURCE_TYPE); got != "📦" {
		t.Errorf("got %q, want the unicode name", got)
	}
}

// Without a name either, nothing is rendered rather than a broken reference.
func TestGetItemTypeStringEmoji_EmptyWhenNeitherSnowflakeNorName(t *testing.T) {
	t.Parallel()

	resource := amqp.ItemType_RESOURCE_TYPE.String()
	service, err := New(stubRepository{byGame: byGame(
		gameEmojis{amqp.Game_DOFUS_GAME, []entities.Emoji{itemEmoji(resource, "", "")}},
	)})
	if err != nil {
		t.Fatalf("New() = %v, want nil", err)
	}

	if got := service.GetItemTypeStringEmoji(amqp.Game_DOFUS_GAME, amqp.ItemType_RESOURCE_TYPE); got != "" {
		t.Errorf("got %q, want empty", got)
	}
}

func TestGetMiscStringEmoji_IsScopedToTheGame(t *testing.T) {
	t.Parallel()

	kama := entities.Emoji{
		ID: string(constants.EmojiIDKama), Type: constants.EmojiTypeMisc,
		DiscordName: "kama", Name: "💰", Snowflake: "dofus-kama",
	}
	touchKama := kama
	touchKama.Snowflake = "touch-kama"

	service, err := New(stubRepository{byGame: byGame(
		gameEmojis{amqp.Game_DOFUS_GAME, []entities.Emoji{kama}},
		gameEmojis{amqp.Game_DOFUS_TOUCH, []entities.Emoji{touchKama}},
	)})
	if err != nil {
		t.Fatalf("New() = %v, want nil", err)
	}

	if got := service.GetMiscStringEmoji(amqp.Game_DOFUS_GAME, constants.EmojiIDKama); got != "<:kama:dofus-kama>" {
		t.Errorf("Dofus = %q", got)
	}
	if got := service.GetMiscStringEmoji(amqp.Game_DOFUS_TOUCH, constants.EmojiIDKama); got != "<:kama:touch-kama>" {
		t.Errorf("Touch = %q", got)
	}
}

// A game with no store at all, such as one no bot serves, renders nothing.
func TestGet_UnknownGameRendersNothing(t *testing.T) {
	t.Parallel()

	service, err := New(stubRepository{byGame: byGame()})
	if err != nil {
		t.Fatalf("New() = %v, want nil", err)
	}

	if got := service.GetMiscStringEmoji(amqp.Game_DOFUS_RETRO, constants.EmojiIDKama); got != "" {
		t.Errorf("got %q, want empty", got)
	}
}

// Every served game is loaded, not only the first.
func TestNew_LoadsEveryServedGame(t *testing.T) {
	t.Parallel()

	loaded := make(map[amqp.Game]bool)
	_, err := New(recordingRepository{loaded: loaded})
	if err != nil {
		t.Fatalf("New() = %v, want nil", err)
	}

	for _, game := range constants.GetServedGames() {
		if !loaded[game.AMQPGame] {
			t.Errorf("%v is served but its emojis were never loaded", game.AMQPGame)
		}
	}
}

type recordingRepository struct {
	loaded map[amqp.Game]bool
}

func (repo recordingRepository) GetEmojis(game amqp.Game) ([]entities.Emoji, error) {
	repo.loaded[game] = true
	return nil, nil
}

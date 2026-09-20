package notifiers

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/entities"
)

// sentMessage records which bot posted where, which is the property that matters:
// a news channel belongs to one bot and only its author may crosspost it.
type sentMessage struct {
	game      amqp.Game
	channelID string
	content   string
}

type stubDiscordService struct {
	sent *[]sentMessage
}

func (stub stubDiscordService) AnnounceMessage(_ string, game amqp.Game, newsChannelID string,
	_ *discordgo.MessageSend) {
	*stub.sent = append(*stub.sent, sentMessage{game: game, channelID: newsChannelID})
}

func (stub stubDiscordService) SendMessage(_ string, game amqp.Game, channelID, content string) {
	*stub.sent = append(*stub.sent, sentMessage{game: game, channelID: channelID, content: content})
}

func (stub stubDiscordService) Shutdown() {}

type stubNewsService struct {
	twitterAccounts []entities.TwitterAccount
}

func (stub stubNewsService) GetAlmanaxNews(_ amqp.Language, _ amqp.Game) *entities.AlmanaxNews {
	return nil
}

func (stub stubNewsService) GetFeedSource(_ string, _ amqp.Language,
	_ amqp.Game) *entities.FeedSource {
	return nil
}

func (stub stubNewsService) GetTwitterAccount(accountID string,
	game amqp.Game) *entities.TwitterAccount {
	for _, account := range stub.twitterAccounts {
		if account.ID == accountID && account.Game == game {
			return &account
		}
	}
	return nil
}

func newTestService(sent *[]sentMessage, accounts []entities.TwitterAccount) *Impl {
	return &Impl{
		reportingChannelID: "reportingChannel",
		discordService:     stubDiscordService{sent: sent},
		newsService:        stubNewsService{twitterAccounts: accounts},
	}
}

// Reports share one channel, but each must be posted by the bot of its own game.
func TestReportsAreSentByTheirGamesBot(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		message func(game amqp.Game) *amqp.RabbitMQMessage
	}{
		{
			name: "set",
			message: func(game amqp.Game) *amqp.RabbitMQMessage {
				return &amqp.RabbitMQMessage{
					Type: amqp.RabbitMQMessage_NEWS_SET, Game: game,
					NewsSetMessage: &amqp.NewsSetMessage{CreatedSetIds: []string{"1"}},
				}
			},
		},
		{
			name: "game",
			message: func(game amqp.Game) *amqp.RabbitMQMessage {
				return &amqp.RabbitMQMessage{
					Type: amqp.RabbitMQMessage_NEWS_GAME, Game: game,
					NewsGameMessage: &amqp.NewsGameMessage{Version: "1.2.3"},
				}
			},
		},
		{
			name: "guild",
			message: func(game amqp.Game) *amqp.RabbitMQMessage {
				return &amqp.RabbitMQMessage{
					Type: amqp.RabbitMQMessage_NEWS_GUILD, Game: game,
					NewsGuildMessage: &amqp.NewsGuildMessage{
						Id: "guildID", Name: "guild", Event: amqp.NewsGuildMessage_CREATE,
					},
				}
			},
		},
	}

	for _, test := range tests {
		for _, game := range []amqp.Game{amqp.Game_DOFUS_GAME, amqp.Game_DOFUS_TOUCH} {
			t.Run(test.name+"/"+game.String(), func(t *testing.T) {
				t.Parallel()

				sent := make([]sentMessage, 0)
				service := newTestService(&sent, nil)

				service.consume(amqp.Context{CorrelationID: "correlationID"}, test.message(game))

				if len(sent) != 1 {
					t.Fatalf("sent %d message(s), want 1", len(sent))
				}
				if sent[0].game != game {
					t.Errorf("posted by %v's bot, want %v", sent[0].game, game)
				}
				if sent[0].channelID != "reportingChannel" {
					t.Errorf("channelID = %q, want %q", sent[0].channelID, "reportingChannel")
				}
			})
		}
	}
}

// A tweet is routed to the news channel of its own game's account row. An account
// that exists only for another game must not borrow that game's channel.
func TestTwitterNewsUsesItsGamesAccount(t *testing.T) {
	t.Parallel()

	accounts := []entities.TwitterAccount{
		{ID: "account", Game: amqp.Game_DOFUS_GAME, NewsChannelID: "dofusChannel"},
		{ID: "account", Game: amqp.Game_DOFUS_TOUCH, NewsChannelID: "touchChannel"},
		{ID: "dofusOnly", Game: amqp.Game_DOFUS_GAME, NewsChannelID: "dofusChannel"},
	}

	tests := []struct {
		name      string
		accountID string
		game      amqp.Game
		channelID string
	}{
		{name: "dofus", accountID: "account", game: amqp.Game_DOFUS_GAME, channelID: "dofusChannel"},
		{name: "touch", accountID: "account", game: amqp.Game_DOFUS_TOUCH, channelID: "touchChannel"},
		{name: "wrong game", accountID: "dofusOnly", game: amqp.Game_DOFUS_TOUCH, channelID: ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			sent := make([]sentMessage, 0)
			service := newTestService(&sent, accounts)

			service.consume(amqp.Context{CorrelationID: "correlationID"}, &amqp.RabbitMQMessage{
				Type:     amqp.RabbitMQMessage_NEWS_TWITTER,
				Game:     test.game,
				Language: amqp.Language_FR,
				NewsTwitterMessage: &amqp.NewsTwitterMessage{
					TwitterId: test.accountID, Title: "@account", Description: "text",
				},
			})

			if test.channelID == "" {
				if len(sent) != 0 {
					t.Fatalf("sent %d message(s), want 0: no account row for this game", len(sent))
				}
				return
			}

			if len(sent) != 1 {
				t.Fatalf("sent %d message(s), want 1", len(sent))
			}
			if sent[0].channelID != test.channelID {
				t.Errorf("channelID = %q, want %q", sent[0].channelID, test.channelID)
			}
			if sent[0].game != test.game {
				t.Errorf("posted by %v's bot, want %v", sent[0].game, test.game)
			}
		})
	}
}

// An unknown game is dropped rather than posted under another bot's branding.
func TestConsumeDropsUnknownGame(t *testing.T) {
	t.Parallel()

	sent := make([]sentMessage, 0)
	service := newTestService(&sent, nil)

	service.consume(amqp.Context{CorrelationID: "correlationID"}, &amqp.RabbitMQMessage{
		Type:            amqp.RabbitMQMessage_NEWS_GAME,
		NewsGameMessage: &amqp.NewsGameMessage{Version: "1.2.3"},
	})

	if len(sent) != 0 {
		t.Errorf("sent %d message(s), want 0", len(sent))
	}
}

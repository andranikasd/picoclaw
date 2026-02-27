package discord

import (
	"github.com/andranikasd/picoclaw/pkg/bus"
	"github.com/andranikasd/picoclaw/pkg/channels"
	"github.com/andranikasd/picoclaw/pkg/config"
)

func init() {
	channels.RegisterFactory("discord", func(cfg *config.Config, b *bus.MessageBus) (channels.Channel, error) {
		return NewDiscordChannel(cfg.Channels.Discord, b)
	})
}

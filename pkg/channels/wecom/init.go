package wecom

import (
	"github.com/andranikasd/picoclaw/pkg/bus"
	"github.com/andranikasd/picoclaw/pkg/channels"
	"github.com/andranikasd/picoclaw/pkg/config"
)

func init() {
	channels.RegisterFactory("wecom", func(cfg *config.Config, b *bus.MessageBus) (channels.Channel, error) {
		return NewWeComBotChannel(cfg.Channels.WeCom, b)
	})
	channels.RegisterFactory("wecom_app", func(cfg *config.Config, b *bus.MessageBus) (channels.Channel, error) {
		return NewWeComAppChannel(cfg.Channels.WeComApp, b)
	})
}

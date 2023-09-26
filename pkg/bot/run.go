package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (b *bot) Run() {
	updetesCfg := tgbotapi.UpdateConfig{
		Offset:  0,
		Timeout: 10,
	}

	for upd := range b.GetUpdatesChan(updetesCfg) {
		go b.processUpdate(upd)
	}
}

func (b *bot) processUpdate(upd tgbotapi.Update) {
	if upd.MyChatMember != nil {
		status := upd.MyChatMember.NewChatMember.Status
		if status == "left" || status == "kicked" {

		}
	}
}

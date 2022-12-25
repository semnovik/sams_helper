package buttons

import tele "gopkg.in/telebot.v3"

type InlineSelector struct {
	Markup *tele.ReplyMarkup
}

func NewInlineSelector() InlineSelector {
	return InlineSelector{&tele.ReplyMarkup{}}
}

func (i InlineSelector) HolidaysBtn() tele.Btn {
	return i.Markup.Data("🎉Праздник сегодня", "holiday")
}

func (i InlineSelector) CurrencyBtn() tele.Btn {
	return i.Markup.Data("💵Курсы валют", "currency")
}

func (i InlineSelector) SquadCallBtn() tele.Btn {
	return i.Markup.Data("🆘Созвать сквад🆘", "squad call")
}

func (i InlineSelector) LoveBtn() tele.Btn {
	return i.Markup.Data("💌Отправить милость", "honey")
}

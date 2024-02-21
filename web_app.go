package telebot

// WebAppData describes data sent from a Web App to the bot.
type WebAppData struct {
	// Data is the data sent from the Web App. Be aware that a bad client can send arbitrary data in this field.
	Data string `json:"data"`

	// ButtonText is the text of the web_app keyboard button from which the Web App was opened.
	// Be aware that a bad client can send arbitrary data in this field.
	ButtonText string `json:"button_text"`
}

// WebAppInfo describes a Web App.
type WebAppInfo struct {
	// URL is an HTTPS URL of a Web App to be opened with additional data as specified in Initializing Web Apps.
	URL string `json:"url"`
}

type SentWebAppMessage struct {
	InlineMessageID *string `json:"inline_message_id,omitempty"`
}

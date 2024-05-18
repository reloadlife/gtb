package telebot

import "time"

type sendAnimationRequest struct {
	ChatID    any   `json:"chat_id"`
	Animation *File `json:"animation" file:"1"`

	BusinessID          *BusinessID      `json:"business_connection_id,omitempty"`
	ThreadID            *MessageThreadID `json:"message_thread_id,omitempty"`
	Caption             *string          `json:"caption,omitempty"`
	ParseMode           ParseMode        `json:"parse_mode,omitempty"`
	Entities            []Entity         `json:"caption_entities,omitempty"`
	HasSpoiler          *bool            `json:"has_spoiler,omitempty"`
	DisableNotification *bool            `json:"disable_notification,omitempty"`
	Protected           *bool            `json:"protect_content,omitempty"`
	ReplyParameters     *ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup         ReplyMarkup      `json:"reply_markup,omitempty"`

	Duration  time.Duration `json:"duration,omitempty"`
	Width     *int          `json:"width,omitempty"`
	Height    *int          `json:"height,omitempty"`
	Thumbnail *File         `json:"thumbnail,omitempty" file:"1"`
}

package telebot

import (
	"errors"
)

// Verify in MaybeInaccessibleMessage needs a custom implementation
// because it can be either AccessibleMessage or InaccessibleMessage.
func (u *MaybeInaccessibleMessage) Verify() error {
	if u.InaccessibleMessage == nil && u.AccessibleMessage == nil {
		return errors.New("both AccessibleMessage and InaccessibleMessage are nil")
	}
	if u.InaccessibleMessage != nil && u.AccessibleMessage != nil {
		return errors.New("both AccessibleMessage and InaccessibleMessage are not nil")
	}

	if u.AccessibleMessage != nil {
		return verify(u.AccessibleMessage)
	}
	return verify(u.InaccessibleMessage)
}

func (u *Update) Verify() error                       { return verify(u) }
func (r *MessageReaction) Verify() error              { return verify(r) }
func (r *MessageReactionCountUpdated) Verify() error  { return verify(r) }
func (u *AccessibleMessage) Verify() error            { return verify(u) }
func (u *InaccessibleMessage) Verify() error          { return verify(u) }
func (c *Chat) Verify() error                         { return verify(c) }
func (c *ChatPhoto) Verify() error                    { return verify(c) }
func (c *ChatLocation) Verify() error                 { return verify(c) }
func (c *ChatPermissions) Verify() error              { return verify(c) }
func (c *CallbackGame) Verify() error                 { return verify(c) }
func (c *LoginURL) Verify() error                     { return verify(c) }
func (c *ChatMember) Verify() error                   { return verify(c) }
func (u *User) Verify() error                         { return verify(u) }
func (c *ReactionType) Verify() error                 { return verify(c) }
func (c *ReactionCount) Verify() error                { return verify(c) }
func (c *ChatMemberUpdated) Verify() error            { return verify(c) }
func (c *ChatInviteLink) Verify() error               { return verify(c) }
func (r *Rights) Verify() error                       { return verify(r) }
func (m *MessageOrigin) Verify() error                { return verify(m) }
func (u *ExternalReplyInfo) Verify() error            { return verify(u) }
func (m *ForceReplyMarkup) Verify() error             { return verify(m) }
func (m *ReplyKeyboardRemove) Verify() error          { return verify(m) }
func (m *InlineKeyboardMarkup) Verify() error         { return verify(m) }
func (m *ReplyKeyboardMarkup) Verify() error          { return verify(m) }
func (i *KeyboardButton) Verify() error               { return verify(i) }
func (i *InlineKeyboardButton) Verify() error         { return verify(i) }
func (i *SwitchInlineQueryChosenChat) Verify() error  { return verify(i) }
func (i *KeyboardButtonRequestChat) Verify() error    { return verify(i) }
func (i *KeyboardButtonRequestUsers) Verify() error   { return verify(i) }
func (i *KeyboardButtonPollType) Verify() error       { return verify(i) }
func (c *ChatBoostRemoved) Verify() error             { return verify(c) }
func (c *ChatBoostUpdated) Verify() error             { return verify(c) }
func (c *ChatBoostAdded) Verify() error               { return verify(c) }
func (c *ChatJoinRequest) Verify() error              { return verify(c) }
func (c *Poll) Verify() error                         { return verify(c) }
func (c *PollAnswer) Verify() error                   { return verify(c) }
func (c *ShippingQuery) Verify() error                { return verify(c) }
func (c *PreCheckoutQuery) Verify() error             { return verify(c) }
func (c *Callback) Verify() error                     { return verify(c) }
func (c *ChosenInlineResult) Verify() error           { return verify(c) }
func (c *InlineQuery) Verify() error                  { return verify(c) }
func (c *ShippingAddress) Verify() error              { return verify(c) }
func (c *OrderInfo) Verify() error                    { return verify(c) }
func (c *LabeledPrice) Verify() error                 { return verify(c) }
func (c *Invoice) Verify() error                      { return verify(c) }
func (c *SuccessfulPayment) Verify() error            { return verify(c) }
func (c *PassportData) Verify() error                 { return verify(c) }
func (c *EncryptedPassportElement) Verify() error     { return verify(c) }
func (c *EncryptedCredentials) Verify() error         { return verify(c) }
func (c *PassportFile) Verify() error                 { return verify(c) }
func (c *PassportElementError) Verify() error         { return verify(c) }
func (c *InputSticker) Verify() error                 { return verify(c) }
func (c *InputInvoiceMessageContent) Verify() error   { return verify(c) }
func (c *InputLocationMessageContent) Verify() error  { return verify(c) }
func (c *InputVenueMessageContent) Verify() error     { return verify(c) }
func (c *InputContactMessageContent) Verify() error   { return verify(c) }
func (c *InputTextMessageContent) Verify() error      { return verify(c) }
func (c *ChatBoostSource) Verify() error              { return verify(c) }
func (c *ChatBoost) Verify() error                    { return verify(c) }
func (f *File) Verify() error                         { return verify(f) }
func (f *PhotoSize) Verify() error                    { return verify(f) }
func (f *Audio) Verify() error                        { return verify(f) }
func (f *Document) Verify() error                     { return verify(f) }
func (f *Video) Verify() error                        { return verify(f) }
func (f *Animation) Verify() error                    { return verify(f) }
func (f *Voice) Verify() error                        { return verify(f) }
func (f *VideoNote) Verify() error                    { return verify(f) }
func (f *Contact) Verify() error                      { return verify(f) }
func (f *Location) Verify() error                     { return verify(f) }
func (f *Venue) Verify() error                        { return verify(f) }
func (f *PollOption) Verify() error                   { return verify(f) }
func (f *Dice) Verify() error                         { return verify(f) }
func (f *Game) Verify() error                         { return verify(f) }
func (f *Sticker) Verify() error                      { return verify(f) }
func (f *StickerSet) Verify() error                   { return verify(f) }
func (f *MaskPosition) Verify() error                 { return verify(f) }
func (f *ForumTopic) Verify() error                   { return verify(f) }
func (f *ForumTopicClosed) Verify() error             { return verify(f) }
func (f *ForumTopicReopened) Verify() error           { return verify(f) }
func (f *ForumTopicEdited) Verify() error             { return verify(f) }
func (f *ForumTopicCreated) Verify() error            { return verify(f) }
func (f *GameHighScore) Verify() error                { return verify(f) }
func (f *MenuButton) Verify() error                   { return verify(f) }
func (f *TextQuote) Verify() error                    { return verify(f) }
func (f *Entity) Verify() error                       { return verify(f) }
func (f *LinkPreviewOptions) Verify() error           { return verify(f) }
func (f *Story) Verify() error                        { return verify(f) }
func (f *AutoDeleteTimerChanged) Verify() error       { return verify(f) }
func (f *ProximityAlertTriggered) Verify() error      { return verify(f) }
func (f *UsersShared) Verify() error                  { return verify(f) }
func (f *ChatShared) Verify() error                   { return verify(f) }
func (f *WriteAccessAllowed) Verify() error           { return verify(f) }
func (f *GeneralForumTopicHidden) Verify() error      { return verify(f) }
func (f *GeneralForumTopicUnhidden) Verify() error    { return verify(f) }
func (f *GiveawayCreated) Verify() error              { return verify(f) }
func (f *GiveawayCompleted) Verify() error            { return verify(f) }
func (f *GiveawayWinners) Verify() error              { return verify(f) }
func (f *Giveaway) Verify() error                     { return verify(f) }
func (f *VideoChatScheduled) Verify() error           { return verify(f) }
func (f *VideoChatEnded) Verify() error               { return verify(f) }
func (f *VideoChatParticipantsInvited) Verify() error { return verify(f) }
func (f *VideoChatStarted) Verify() error             { return verify(f) }
func (f *WebAppData) Verify() error                   { return verify(f) }
func (f *WebAppInfo) Verify() error                   { return verify(f) }
func (f *BotCommand) Verify() error                   { return verify(f) }
func (f *BotCommandScope) Verify() error              { return verify(f) }
func (f *UserProfilePhotos) Verify() error            { return verify(f) }
func (f *UserChatBoosts) Verify() error               { return verify(f) }
func (f *SentWebAppMessage) Verify() error            { return verify(f) }
func (f *ReplyParameters) Verify() error              { return verify(f) }

package telebot

import (
	"errors"
)

func (u *Update) Verify() error                      { return verify(u) }
func (r *MessageReaction) Verify() error             { return verify(r) }
func (r *MessageReactionCountUpdated) Verify() error { return verify(r) }

func (u *AccessibleMessage) Verify() error   { return verify(u) }
func (u *InaccessibleMessage) Verify() error { return verify(u) }

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

func (c *Chat) Verify() error { return verify(c) }

func (c *ChatPhoto) Verify() error       { return verify(c) }
func (c *ChatLocation) Verify() error    { return verify(c) }
func (c *ChatPermissions) Verify() error { return verify(c) }
func (c *CallbackGame) Verify() error    { return verify(c) }
func (c *LoginURL) Verify() error        { return verify(c) }
func (c *ChatMember) Verify() error      { return verify(c) }
func (u *User) Verify() error            { return verify(u) }
func (c *ReactionType) Verify() error    { return verify(c) }
func (c *ReactionCount) Verify() error   { return verify(c) }

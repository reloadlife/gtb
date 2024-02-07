package telebot

// ChatBoostUpdated represents a boost added to a chat or changed.
type ChatBoostUpdated struct {
	// Chat is the chat which was boosted.
	Chat Chat `json:"chat"`

	// Boost is the information about the chat boost.
	Boost ChatBoost `json:"boost"`
}

// ChatBoostRemoved represents a boost removed from a chat.
type ChatBoostRemoved struct {
	// Chat is the chat which was boosted.
	Chat Chat `json:"chat"`

	// BoostID is the unique identifier of the boost.
	BoostID string `json:"boost_id"`

	// RemoveDate is the point in time (Unix timestamp) when the boost was removed.
	RemoveDate int64 `json:"remove_date"`

	// Source is the source of the removed boost.
	Source ChatBoostSource `json:"source"`
}
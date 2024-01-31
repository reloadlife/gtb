package telebot

import "encoding/json"

func (b *bot) GetMe() (*User, error) {
	raw, err := b.sendMethodRequest(methodGetMe, nil)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Result User
	}

	if err = json.NewDecoder(raw.Body).Decode(&resp); err != nil {
		return nil, err
	}

	return &resp.Result, nil
}
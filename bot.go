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

	if err = json.Unmarshal(raw, &resp); err != nil {
		return nil, err
	}

	return &resp.Result, nil
}

func (b *bot) SetMyCommands(commands []BotCommand, opts ...any) error {
	params := setMyCommands{
		Commands: commands,
	}

	for _, opt := range opts {
		switch o := opt.(type) {
		case *BotCommandScope:
			params.Scope = o
		case string:
			params.Language = o
		}
	}

	_, err := b.sendMethodRequest(methodSetMyCommands, params)
	return err
}

func (b *bot) GetMyCommands(opts ...any) ([]BotCommand, error) {
	params := getMyCommands{}

	for _, opt := range opts {
		switch o := opt.(type) {
		case *BotCommandScope:
			params.Scope = o
		case string:
			params.Language = o
		}
	}

	raw, err := b.sendMethodRequest(methodGetMyCommands, params)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Result []BotCommand
	}

	if err = json.Unmarshal(raw, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}

func (b *bot) DeleteMyCommands(opts ...any) error {
	params := deleteMyCommands{}

	for _, opt := range opts {
		switch o := opt.(type) {
		case *BotCommandScope:
			params.Scope = o
		case string:
			params.Language = o
		}
	}

	_, err := b.sendMethodRequest(methodDeleteMyCommands, params)
	return err
}

func (b *bot) SetMyName(name string, opts ...any) error {
	params := setMyNameRequest{
		Name: name,
	}

	for _, opt := range opts {
		switch o := opt.(type) {
		case string:
			params.Language = o
		}
	}

	_, err := b.sendMethodRequest(methodSetMyName, params)
	return err
}

func (b *bot) GetMyName(opts ...any) (*string, error) {
	params := setMyNameRequest{}

	for _, opt := range opts {
		switch o := opt.(type) {
		case string:
			params.Language = o
		}
	}

	raw, err := b.sendMethodRequest(methodGetMyName, nil)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Result string
	}

	if err = json.Unmarshal(raw, &resp); err != nil {
		return nil, err
	}

	return &resp.Result, nil
}

func (b *bot) SetMyDescription(description string, opts ...any) error {
	params := setMyDescriptionRequest{
		Description: description,
	}

	for _, opt := range opts {
		switch o := opt.(type) {
		case string:
			params.Language = o
		}
	}

	_, err := b.sendMethodRequest(methodSetMyDescription, params)
	return err
}

func (b *bot) GetMyDescription(opts ...any) (*string, error) {
	params := setMyDescriptionRequest{}

	for _, opt := range opts {
		switch o := opt.(type) {
		case string:
			params.Language = o
		}
	}

	raw, err := b.sendMethodRequest(methodGetMyDescription, nil)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Result string
	}

	if err = json.Unmarshal(raw, &resp); err != nil {
		return nil, err
	}

	return &resp.Result, nil
}

func (b *bot) SetMyShortDescription(description string, opts ...any) error {
	params := setMyDescriptionRequest{
		ShortDescription: description,
	}

	for _, opt := range opts {
		switch o := opt.(type) {
		case string:
			params.Language = o
		}
	}

	_, err := b.sendMethodRequest(methodSetMyShortDescription, params)
	return err
}

func (b *bot) GetMyShortDescription(opts ...any) (*string, error) {
	params := setMyDescriptionRequest{}

	for _, opt := range opts {
		switch o := opt.(type) {
		case string:
			params.Language = o
		}
	}

	raw, err := b.sendMethodRequest(methodGetMyShortDescription, nil)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Result string
	}

	if err = json.Unmarshal(raw, &resp); err != nil {
		return nil, err
	}

	return &resp.Result, nil
}

func (b *bot) SetChatMenuButton(opts ...any) error {
	params := setChatMenuButton{}

	for _, opt := range opts {
		switch o := opt.(type) {
		case Recipient:
			params.ChatID = o.Recipient()
		case *MenuButton:
			params.Button = o
		}
	}

	_, err := b.sendMethodRequest(methodSetChatMenuButton, params)
	return err
}

func (b *bot) GetChatMenuButton(opts ...any) (*MenuButton, error) {
	params := setChatMenuButton{}

	for _, opt := range opts {
		switch o := opt.(type) {
		case Recipient:
			params.ChatID = o.Recipient()
		case *MenuButton:
			params.Button = o
		}
	}

	raw, err := b.sendMethodRequest(methodGetChatMenuButton, params)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Result MenuButton
	}

	if err = json.Unmarshal(raw, &resp); err != nil {
		return nil, err
	}

	return &resp.Result, nil
}

func (b *bot) SetMyDefaultAdministratorRights(opts ...any) error {
	params := setMyDefaultAdministratorRights{}

	for _, opt := range opts {
		switch o := opt.(type) {
		case *Rights:
			params.Rights = o
		case bool:
			params.ForChannels = o
		}
	}

	_, err := b.sendMethodRequest(methodSetMyDefaultAdministratorRights, params)
	return err
}

func (b *bot) GetMyDefaultAdministratorRights(opts ...any) (*Rights, error) {
	params := setMyDefaultAdministratorRights{}

	for _, opt := range opts {
		switch o := opt.(type) {
		case bool:
			params.ForChannels = o
		}
	}

	raw, err := b.sendMethodRequest(methodGetMyDefaultAdministratorRights, params)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Result Rights
	}

	if err = json.Unmarshal(raw, &resp); err != nil {
		return nil, err
	}

	return &resp.Result, nil
}

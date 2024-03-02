package config

import (
	"fmt"
	tele "go.mamad.dev/telebot"
	"strings"
)

type Handler interface {
	GetKeyboard(l ...string) tele.ReplyMarkup
	GetText(l ...string) string
}

type Handle interface {
	GetCommand(l ...string) []any
	GetHandler() Handler
}

type handle struct {
	conf    *config
	Command any     `yaml:"cmd" json:"cmd"`
	Handler handler `yaml:"handlerResponse" json:"handler"`
}

type handler struct {
	conf *config

	Keyboard string `yaml:"keyboard" json:"keyboard"`
	Text     string `yaml:"text" json:"text"`
}

func (h *handle) GetCommand(l ...string) []any {
	locale := h.conf.GetDefaultLocale()
	if len(l) > 0 {
		locale = l[0]
	}

	commands := make([]any, 0)
	switch hc := h.Command.(type) {
	case string:
		if strings.HasPrefix(hc, "locale:") {
			lKey := strings.TrimPrefix(hc, "locale:")
			commands = append(commands, h.conf.L(locale, lKey))
		} else {
			commands = append(commands, hc)
		}
	case []string:
		for _, c := range hc {
			if locale != "" {
				if strings.HasPrefix(c, "locale:") {
					lKey := strings.TrimPrefix(c, "locale:")
					commands = append(commands, h.conf.L(locale, lKey))
				} else {
					commands = append(commands, c)
				}
			}
		}
	case []any:
		for _, c := range hc {
			switch c := c.(type) {
			case string:
				if strings.HasPrefix(c, "locale:") {
					lKey := strings.TrimPrefix(c, "locale:")
					commands = append(commands, h.conf.L(locale, lKey))
				} else {
					commands = append(commands, c)
				}
			default:
				panic(fmt.Errorf("`cmd` should be either string or array of strings, but got %T", c))
			}
		}
	default:
		panic(fmt.Errorf("`cmd` should be either string or array of strings, but got %T", hc))
	}
	return commands
}

func (h *handle) GetHandler() Handler {
	g := h.Handler
	g.conf = h.conf
	return &g
}

func (h *handler) GetKeyboard(l ...string) tele.ReplyMarkup {
	locale := h.conf.GetDefaultLocale()
	if len(l) > 0 {
		locale = l[0]
	}

	return h.conf.Keyboard(h.Keyboard, locale)
}

func (h *handler) GetText(l ...string) string {
	locale := h.conf.GetDefaultLocale()
	if len(l) > 0 {
		locale = l[0]
	}

	if locale != "" {
		lKey := strings.TrimPrefix(h.Text, "locale:")
		return h.conf.L(locale, lKey)
	}
	return h.Text
}

func (c *config) GetHandles() []Handle {
	handles := make([]Handle, 0)

	for _, h := range c.conf.Handles {
		ha := h
		ha.conf = c
		handles = append(handles, &ha)
	}

	return handles
}

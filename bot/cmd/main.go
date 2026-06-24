package main

import (
	"context"
	"flowmoney/bot/internal/config"
	"flowmoney/bot/internal/handlers"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"

	tele "gopkg.in/telebot.v3"

	"golang.org/x/net/proxy"
)

func main() {
	cfg := config.MustLoad()

	httpClient, err := newHTTPClient(cfg.ProxyUrl)
	if err != nil {
		log.Fatal(err)
	}

	bot, err := tele.NewBot(tele.Settings{
		Token:     cfg.Token,
		Client:    httpClient,
		ParseMode: tele.ModeHTML,
		Poller:    &tele.LongPoller{Timeout: 60 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Бот запущен: @%s\n", bot.Me.Username)

	handlers.Register(bot, handlers.NewStore(cfg.Bot.ApiUrl))

	bot.Start()
}

// newHTTPClient опционально заворачивает HTTP-клиент бота в прокси —
// например, в локальный SOCKS5/HTTP порт Happ, если api.telegram.org
// недоступен напрямую. По умолчанию (proxyUrl == "") идёт напрямую.
func newHTTPClient(proxyUrl string) (*http.Client, error) {
	if proxyUrl == "" {
		return &http.Client{}, nil
	}

	parsed, err := url.Parse(proxyUrl)
	if err != nil {
		return nil, fmt.Errorf("invalid proxy url: %w", err)
	}

	switch parsed.Scheme {
	case "http", "https":
		return &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(parsed)}}, nil

	case "socks5", "socks5h":
		dialer, err := proxy.FromURL(parsed, proxy.Direct)
		if err != nil {
			return nil, fmt.Errorf("socks5 dialer: %w", err)
		}
		return &http.Client{Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return dialer.Dial(network, addr)
			},
		}}, nil

	default:
		return nil, fmt.Errorf("unsupported proxy scheme %q (use http:// or socks5://)", parsed.Scheme)
	}
}

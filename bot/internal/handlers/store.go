package handlers

import (
	"flowmoney/bot/internal/client"
	"sync"
)

// Account хранит авторизованного клиента flowmoney API для одного чата.
type Account struct {
	Client *client.Client
	UserId int
}

func (a *Account) LoggedIn() bool {
	return a.UserId != 0
}

// Store — по одному клиенту flowmoney API на чат, токен у каждого свой.
type Store struct {
	mu       sync.Mutex
	accounts map[int64]*Account
	apiUrl   string
}

func NewStore(apiUrl string) *Store {
	return &Store{accounts: make(map[int64]*Account), apiUrl: apiUrl}
}

func (s *Store) Get(chatId int64) *Account {
	s.mu.Lock()
	defer s.mu.Unlock()

	acc, ok := s.accounts[chatId]
	if !ok {
		acc = &Account{Client: client.NewClient(s.apiUrl)}
		s.accounts[chatId] = acc
	}
	return acc
}

func (s *Store) Drop(chatId int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.accounts, chatId)
}

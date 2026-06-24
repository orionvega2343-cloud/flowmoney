package handlers

import (
	"flowmoney/bot/internal/client"
	"sync"
)

// Step — что бот ждёт от пользователя следующим текстовым сообщением.
type Step string

const (
	StepNone Step = ""

	StepRegisterName     Step = "register_name"
	StepRegisterEmail    Step = "register_email"
	StepRegisterPassword Step = "register_password"

	StepLoginEmail    Step = "login_email"
	StepLoginPassword Step = "login_password"

	StepBalanceAmount Step = "balance_amount"

	StepCategoryCreateTitle Step = "category_create_title"
	StepCategoryGetId       Step = "category_get_id"
	StepCategoryUpdateId    Step = "category_update_id"
	StepCategoryUpdateTitle Step = "category_update_title"

	StepTransactionAmount     Step = "transaction_amount"
	StepTransactionCategoryId Step = "transaction_category_id"
	StepTransactionGetId      Step = "transaction_get_id"

	StepBudgetCreateCategoryId Step = "budget_create_category_id"
	StepBudgetCreateAmount     Step = "budget_create_amount"
	StepBudgetCreateMonth      Step = "budget_create_month"
	StepBudgetCreateYear       Step = "budget_create_year"
	StepBudgetGetId            Step = "budget_get_id"
	StepBudgetCategoryId       Step = "budget_category_id"
	StepBudgetMonthMonth       Step = "budget_month_month"
	StepBudgetMonthYear        Step = "budget_month_year"
	StepBudgetUpdateId         Step = "budget_update_id"
	StepBudgetUpdateAmount     Step = "budget_update_amount"
	StepBudgetDeleteId         Step = "budget_delete_id"
)

// Session хранит состояние диалога с одним чатом: авторизованного клиента
// flowmoney API, текущий шаг ввода и временный буфер для многошаговых форм.
type Session struct {
	Client *client.Client
	UserId int
	Step   Step
	Data   map[string]string
}

func (s *Session) LoggedIn() bool {
	return s.UserId != 0
}

func (s *Session) Reset() {
	s.Step = StepNone
	s.Data = make(map[string]string)
}

// SessionStore — по одному клиенту flowmoney API на чат, токен у каждого свой.
type SessionStore struct {
	mu       sync.Mutex
	sessions map[int64]*Session
	apiUrl   string
}

func NewSessionStore(apiUrl string) *SessionStore {
	return &SessionStore{sessions: make(map[int64]*Session), apiUrl: apiUrl}
}

func (s *SessionStore) Get(chatId int64) *Session {
	s.mu.Lock()
	defer s.mu.Unlock()

	sess, ok := s.sessions[chatId]
	if !ok {
		sess = &Session{Client: client.NewClient(s.apiUrl), Data: make(map[string]string)}
		s.sessions[chatId] = sess
	}
	return sess
}

func (s *SessionStore) Drop(chatId int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.sessions, chatId)
}

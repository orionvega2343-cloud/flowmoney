 ## Dependenties(deps.go)
 
Общие зависимости обработчиков

    type Deps struct {
	Store *Store 
    }
База запросов, в ней хранятся все запросы клиента

    func (d Deps) account(c tele.Context) *Account {
	acc := d.Store.Get(c.Chat().ID)
	acc.Step = nil
	return acc
    }

Получаем id чата от telegram.org, берем id и узнаем кто написал, в случае отмены или смены действия бот забывает, что было и обрабатывает следующее, например: Нажали "Добавить категорию", не стали ее вводить и тут-же нажали на кнопку профиля, в ином случае бот бы все равно ждал ввода категории

    func (d Deps) requireLogin(c tele.Context, acc *Account) bool {
	if acc.LoggedIn() {
		return true
	}
	_ = c.Send("🔒 Сначала войдите в аккаунт:", authMarkup())
	return false
    }

Бот проверяет зарегистрирован ли аккаунт, в случае если да, пропускаем дальше, иначе перенаправляем на форму входа

    func (d Deps) fail(c tele.Context, action string, err error) error {
	return c.Send("❌ " + action + ": " + err.Error())
    }

Объясняем пользователю, где произошло падение


## Store

    type Account struct {
	Client *client.Client
	UserId int
	Step   *Step
    }

Аккаунт пользователя


    type Store struct {
	mu       sync.Mutex         
	accounts map[int64]*Account 
	apiUrl   string             
    }

База активных пользователей 


    func NewStore(apiUrl string) *Store {
	return &Store{accounts: make(map[int64]*Account), apiUrl: apiUrl}
    }

Коструктор, принимает поля аккаунта и URL-адрес

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

Идентификация аккаунта: избегаем race condition, идентифицируем id чата в случае нового пользователя создаем новый аккаунт с чистым клиентом

    func (s *Store) Drop(chatId int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.accounts, chatId)
    }

Удаляем аккаунт


## Router

    func Register(bot *tele.Bot, store *Store) {
	d := Deps{Store: store}

	bot.Use(ackCallbacks)

	NewUserHandlers(d).Register(bot)
	NewCategoryHandlers(d).Register(bot)
	NewTransactionHandlers(d).Register(bot)
	NewBudgetHandlers(d).Register(bot)

	bot.Handle(tele.OnText, func(c tele.Context) error {
		acc := store.Get(c.Chat().ID)
		if acc.Step != nil {
			return continueDialog(c, acc)
		}
		return c.Send("Не понимаю 🤔 Воспользуйтесь клавиатурой ниже или командой /start.")
	})
    }


Регистрируем обработчики, в случае неверной команды или ввода отвечаем

    func ackCallbacks(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		if c.Callback() != nil {
			_ = c.Respond()
		}
		return next(c)
	}
    }

Убираем "часики" с кнопки


## Keyboards 

    const (
	txtProfile      = "👤 Профиль"
	txtCategories   = "📁 Категории"
	txtTransactions = "💸 Транзакции"
	txtBudget       = "📊 Бюджет"
	txtLogout       = "🚪 Выйти"
    )

Объявляем кнопки

    func mainKeyboard() *tele.ReplyMarkup {
	kb := &tele.ReplyMarkup{ResizeKeyboard: true}
	kb.Reply(
		tele.Row{{Text: txtProfile}, {Text: txtCategories}},
		tele.Row{{Text: txtTransactions}, {Text: txtBudget}},
		tele.Row{{Text: txtLogout}},
	)
	return kb
    }

Reply клавиатура с разделами нашего бота

## Dialog

    type Step struct {
	Prompt string
	Next   func(reply string) StepResult
    }

Шаг, отслеживает на каком этапе находится пользователь

    type StepResult struct {
	Next   *Step
	Text   string
	Markup *tele.ReplyMarkup
	Err    error
    }

Результат шага, показывает, что случилось на текущем шаге

    func ask(prompt string, next func(reply string) StepResult) StepResult {
	return StepResult{Next: &Step{Prompt: prompt, Next: next}}
    }

Конструктор

    func done(text string, markup ...*tele.ReplyMarkup) StepResult {
	r := StepResult{Text: text}
	if len(markup) > 0 {
		r.Markup = markup[0]
	}
	return r
    }

Записываем значение шага, в случае перехода дальше, сбрасываем шаги

        func fail(err error) StepResult {
	return StepResult{Err: err}
    }

Если диалог не найден, выбросим ошибку

    func startDialog(c tele.Context, acc *Account, step *Step) error {
	acc.Step = step
	return c.Send(step.Prompt, &tele.ReplyMarkup{ForceReply: true})
    }

Начало диалога

    func continueDialog(c tele.Context, acc *Account) error {
	step := acc.Step
	acc.Step = nil

	result := step.Next(c.Text())
	switch {
	case result.Err != nil:
		return c.Send("❌ " + result.Err.Error())
	case result.Next != nil:
		acc.Step = result.Next
		return c.Send(result.Next.Prompt, &tele.ReplyMarkup{ForceReply: true})
	case result.Markup != nil:
		return c.Send(result.Text, result.Markup)
	default:
		return c.Send(result.Text)
	}
    }


Продолжение диалога: в случае если ошибка вернем ❌ и ошибку, в случае Next != nil перейдем на следующий этап, в случае кнопки вернем текст и кнопку, иначе просто текст
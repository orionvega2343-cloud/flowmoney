package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// Callback data для всех кнопок бота. Формат "domain:action[:arg]".
const (
	cbRegister = "auth:register"
	cbLogin    = "auth:login"
	cbLogout   = "auth:logout"

	cbMenuMain        = "menu:main"
	cbMenuCategory    = "menu:category"
	cbMenuTransaction = "menu:transaction"
	cbMenuBudget      = "menu:budget"

	cbProfile       = "user:profile"
	cbBalanceTopUp  = "user:balance:topup"
	cbBalanceDeduct = "user:balance:deduct"

	cbCategoryCreate = "category:create"
	cbCategoryList   = "category:list"
	cbCategoryGet    = "category:get"
	cbCategoryUpdate = "category:update"

	cbTransactionCreateIncome  = "transaction:create:income"
	cbTransactionCreateExpense = "transaction:create:expense"
	cbTransactionList          = "transaction:list"
	cbTransactionGet           = "transaction:get"

	cbBudgetCreate     = "budget:create"
	cbBudgetGet        = "budget:get"
	cbBudgetByCategory = "budget:by_category"
	cbBudgetByMonth    = "budget:by_month"
	cbBudgetUpdate     = "budget:update"
	cbBudgetDelete     = "budget:delete"
)

func authMenu() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("📝 Регистрация", cbRegister),
			tgbotapi.NewInlineKeyboardButtonData("🔑 Войти", cbLogin),
		),
	)
}

func mainMenu() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("👤 Профиль", cbProfile),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("📁 Категории", cbMenuCategory),
			tgbotapi.NewInlineKeyboardButtonData("💸 Транзакции", cbMenuTransaction),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("📊 Бюджет", cbMenuBudget),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🚪 Выйти", cbLogout),
		),
	)
}

func profileMenu() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("➕ Пополнить баланс", cbBalanceTopUp),
			tgbotapi.NewInlineKeyboardButtonData("➖ Списать с баланса", cbBalanceDeduct),
		),
		backRow(cbMenuMain),
	)
}

func categoryMenu() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("➕ Создать категорию", cbCategoryCreate),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("📋 Мои категории", cbCategoryList),
			tgbotapi.NewInlineKeyboardButtonData("🔎 Категория по ID", cbCategoryGet),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("✏️ Изменить категорию", cbCategoryUpdate),
		),
		backRow(cbMenuMain),
	)
}

func transactionMenu() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("💵 Новый доход", cbTransactionCreateIncome),
			tgbotapi.NewInlineKeyboardButtonData("💳 Новый расход", cbTransactionCreateExpense),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("📋 Мои транзакции", cbTransactionList),
			tgbotapi.NewInlineKeyboardButtonData("🔎 Транзакция по ID", cbTransactionGet),
		),
		backRow(cbMenuMain),
	)
}

func budgetMenu() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("➕ Создать бюджет", cbBudgetCreate),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔎 Бюджет по ID", cbBudgetGet),
			tgbotapi.NewInlineKeyboardButtonData("🏷 По категории", cbBudgetByCategory),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("📅 За месяц", cbBudgetByMonth),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("✏️ Изменить", cbBudgetUpdate),
			tgbotapi.NewInlineKeyboardButtonData("🗑 Удалить", cbBudgetDelete),
		),
		backRow(cbMenuMain),
	)
}

func backRow(to string) []tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("⬅️ Назад", to),
	)
}

func kbPtr(kb tgbotapi.InlineKeyboardMarkup) *tgbotapi.InlineKeyboardMarkup {
	return &kb
}

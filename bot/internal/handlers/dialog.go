package handlers

import tele "gopkg.in/telebot.v3"

// Step — один шаг диалога: вопрос, который бот задаёт пользователю, и
// функция, разбирающая ответ. Кнопка запускает первый Step, дальше каждый
// текстовый ответ ведёт по цепочке Next, пока шаг не вернёт финальный текст.
type Step struct {
	Prompt string
	Next   func(reply string) StepResult
}

// StepResult — что делать после ответа пользователя: либо продолжить
// диалог следующим Step, либо завершить его текстом (с опциональной
// клавиатурой), либо сообщить об ошибке.
type StepResult struct {
	Next   *Step
	Text   string
	Markup *tele.ReplyMarkup
	Err    error
}

func ask(prompt string, next func(reply string) StepResult) StepResult {
	return StepResult{Next: &Step{Prompt: prompt, Next: next}}
}

func done(text string, markup ...*tele.ReplyMarkup) StepResult {
	r := StepResult{Text: text}
	if len(markup) > 0 {
		r.Markup = markup[0]
	}
	return r
}

func fail(err error) StepResult {
	return StepResult{Err: err}
}

// startDialog запускает цепочку шагов: запоминает её в аккаунте чата и
// присылает первый вопрос с принудительным ответом (force reply).
func startDialog(c tele.Context, acc *Account, step *Step) error {
	acc.Step = step
	return c.Send(step.Prompt, &tele.ReplyMarkup{ForceReply: true})
}

// continueDialog обрабатывает очередное сообщение чата, находящегося в
// диалоге: передаёт его текущему шагу и либо задаёт следующий вопрос,
// либо завершает диалог результатом.
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

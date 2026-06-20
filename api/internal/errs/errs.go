package errs

import "errors"

var ErrBudgetExceeded = errors.New("бюджет превышен")
var ErrInsufficientFunds = errors.New("insufficient funds")

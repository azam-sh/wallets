package e

import "errors"

var (
	ErrAccNotFound          = errors.New("аккаунт не найден")
	ErrIncorrectPhoneNumber = errors.New("неправильный номер телефона")
	ErrInternalServer       = errors.New("внутренняя ошибка сервера")
	ErrInvalidInput         = errors.New("введены неверные данные")
	ErrExceededLimit        = errors.New("превышен лимит")
)

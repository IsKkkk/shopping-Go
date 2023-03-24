package cart

import (
	"errors"
)

var (
	ErrItemAlreadyExistInCart = errors.New("Item 已经存在")
	ErrCountInvalid           = errors.New("数量不能是负值")
)

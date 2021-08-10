package models

import val "github.com/go-ozzo/ozzo-validation"

// очень удобный пакет для валидаций структур
// правда здесь его использовать было бессмысленно.

func (u *User) ValidateCreate() error {
	return val.ValidateStruct(u,
		val.Field(&u.Login, val.Required),
		// val.Field(&u.Login, val.Required, val.Lenght(4, 16)) если делать что то крупное, можно указывать длину,
		// или же проверять регулярными выражениями, например проверка на валидность email
	)
}

func (u *User) Validate() error {
	return val.ValidateStruct(u,
		val.Field(&u.Login, val.Required),
		val.Field(&u.ID, val.Required),
	)
}

func (u *User) ValidateEdit() error {
	return val.ValidateStruct(u,
		val.Field(&u.ID, val.Required),
	)
}

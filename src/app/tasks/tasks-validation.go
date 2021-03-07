package tasks

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

func (task Model) Validation() error {
	return validation.ValidateStruct(&task,
		validation.Field(&task.Title, validation.Required),
		validation.Field(&task.Status, validation.In("done", "pending")),
		validation.Field(&task.Comment),
	)
}


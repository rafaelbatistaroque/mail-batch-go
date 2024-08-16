package validation

import (
	"errors"
	"fmt"
	"strings"
)

type InputValidation struct {
	erros []string
}

func (i *InputValidation) IsInvalid() bool {
	return len(i.erros) > 0
}

func (i *InputValidation) GetErrors() error {
	if len(i.erros) == 0 {
		return nil
	}

	return errors.New(strings.Join(i.erros, " | "))
}

func (i *InputValidation) AppendError(propertyName, erroMessagem string) {
	i.erros = append(i.erros, fmt.Sprintf(erroMessagem, propertyName))
}

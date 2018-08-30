package brazil

import "time"

// CPF
type cpf struct {
	Number string `json:"number"`
}

// Título de eleitor
type tituloEleitoral struct {
	Number string `json:"number"`
}

// PIS
type pis struct {
	Number string `json:"number"`
}

// Date
type brDate struct {
	Date time.Time `json:"date"`
	Err  error     `json:"err"`
}

type documents interface {
	hasExpectedFormat() Validation
	isValid() Validation
}

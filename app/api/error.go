package api

// ErrorResponse представляет структуру для ответа с ошибкой.
type ErrorResponse struct {
	// Сообщение об ошибке.
	Error string `json:"error" example:"Invalid request"`
}

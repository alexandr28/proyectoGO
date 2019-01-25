package models
// message para el cliente de la api
type Message struct {
	Message  string `json:"message"`
	CodState int    `json:"cod_state"`
}

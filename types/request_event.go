package types

type RequestEvent struct {
	UserId    string `json:"userId"`
	Component string `json:"component"`
	Path      string `json:"path"`
	Method    string `json:"method"`
	Payload   string `json:"payload"`
}

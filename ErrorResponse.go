package exactonline

// ErrorResponse stores general ExactOnline API error response
//
type ErrorResponse struct {
	Error struct {
		Code    int `json:"code"`
		Message struct {
			Lang  string `json:"lang"`
			Value string `json:"value"`
		} `json:"message"`
	} `json:"error"`
}

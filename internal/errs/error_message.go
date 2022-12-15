package errs

type ReplyError struct {
	ErrorMessage ErrorMessage `json:"errorMessage"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

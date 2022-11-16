package response

type errorJSON struct {
	Error string `json:"error"`
}

func Error(c Context, code int, message string) {
	c.AbortWithStatusJSON(code, &errorJSON{
		Error: message,
	})
}

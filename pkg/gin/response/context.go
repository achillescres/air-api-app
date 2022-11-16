package response

type Context interface {
	AbortWithStatusJSON(code int, jsonObj any)
}

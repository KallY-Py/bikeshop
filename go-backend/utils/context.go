package utils

type ContextKey string

const (
	UserIDKey   ContextKey = "user_id"
	UsernameKey ContextKey = "username"
	UserRoleKey ContextKey = "user_role"
)

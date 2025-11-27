package auth

import (
	"context"
)

type contextKey string

const userIDKey contextKey = "user_id"

func SetUserID(ctx context.Context, userID int) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

func GetUserID(ctx context.Context) int {
	userID, ok := ctx.Value(userIDKey).(int)
	if !ok {
		return 0
	}

	return userID
}

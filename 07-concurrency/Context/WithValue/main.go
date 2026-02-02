package main

import (
	"context"
	"fmt"
)

type ctxKey string

const userRoleKey ctxKey = "userRole"

func main() {

	ctx := context.WithValue(context.Background(), userRoleKey, "admin") //guest

	deleteUser(ctx)
}

func deleteUser(ctx context.Context) {
	role := ctx.Value(userRoleKey)

	if role == "admin" {
		fmt.Println("User deleted successfully.")
	} else {
		fmt.Println("Permission denied. Admin role required.")
	}
}

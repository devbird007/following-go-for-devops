package main

import (
	"fmt"
	"context"
)

type key int

const claimsKey key = 0

func NewContext(ctx context.Context, claims Claims)

context.Context {
	return context.WithValue(ctx, claimsKey, claims)
}

func ClaimsFromContext(ctx context.Context) (Claims, bool)

{
	// ctx.Value returns nil if ctx has no value for the key;
	// the Claims type assertion returns ok=false for nil.
	claims, ok := ctx.Value(userIPKey).(Claims)
	return claims, ok
}

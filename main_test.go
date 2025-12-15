package main

import (
	"context"
	"testing"
)

func TestLoaderWithTracing(t *testing.T) {
	loader := newLoader()
	ctx := context.Background()

	_, err := loader.Load(ctx, "key")()
	if err != nil {
		t.Error(err)
	}
}

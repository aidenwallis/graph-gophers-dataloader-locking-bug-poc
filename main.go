package main

import (
	"context"
	"log"

	"github.com/graph-gophers/dataloader/v7"
)

func main() {
	log.Println("run go test -timeout 5s to simulate this bug")
}

func newLoader() dataloader.Interface[string, struct{}] {
	return dataloader.NewBatchedLoader(func(ctx context.Context, k []string) []*dataloader.Result[struct{}] {
		out := make([]*dataloader.Result[struct{}], len(k))
		for i := range k {
			out[i] = &dataloader.Result[struct{}]{
				Data: struct{}{},
			}
		}
		return out
	}, dataloader.WithTracer(&Tracer[string, struct{}]{}))
}

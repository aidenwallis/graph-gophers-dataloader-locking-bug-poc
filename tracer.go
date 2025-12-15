package main

import (
	"context"
	"log"
	"time"

	"github.com/graph-gophers/dataloader/v7"
)

type Tracer[K comparable, V any] struct{}

func (Tracer[K, V]) TraceLoad(ctx context.Context, key K) (context.Context, dataloader.TraceLoadFinishFunc[V]) {
	return ctx, func(thunk dataloader.Thunk[V]) {
		if _, err := thunk(); err != nil {
			log.Println("tracer.Load had an error", err)
		}
	}
}

// TraceLoadMany is a noop function
func (Tracer[K, V]) TraceLoadMany(ctx context.Context, keys []K) (context.Context, dataloader.TraceLoadManyFinishFunc[V]) {
	return ctx, func(thunk dataloader.ThunkMany[V]) {
		_, errs := thunk()

		errorCount := 0
		for _, err := range errs {
			if err != nil {
				errorCount++
			}
		}

		if errorCount > 0 {
			log.Println("tracer had", errorCount, "errors")
		}
	}
}

// TraceBatch is a noop function
func (Tracer[K, V]) TraceBatch(ctx context.Context, keys []K) (context.Context, dataloader.TraceBatchFinishFunc[V]) {
	start := time.Now()

	return ctx, func(result []*dataloader.Result[V]) {
		log.Println("batch took", time.Since(start))
	}
}

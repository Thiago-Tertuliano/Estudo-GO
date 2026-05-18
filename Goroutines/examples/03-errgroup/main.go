package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func task(ctx context.Context, id int, fail bool) error {
	select {
	case <-time.After(500 * time.Millisecond):
		if fail {
			return fmt.Errorf("task %d: falhou", id)
		}
		fmt.Printf("task %d: ok\n", id)
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error { return task(ctx, 1, false) })
	g.Go(func() error { return task(ctx, 2, true) }) //tarefa falha, cancela o contexto
	g.Go(func() error { return task(ctx, 3, false) })

	if err := g.Wait(); err != nil {
		fmt.Println("errgroup:", err)
		if errors.Is(err, context.Canceled) {
			fmt.Println("(outras goroutines foram canceladas pelo contexto do grupo)")
		}
	}
}
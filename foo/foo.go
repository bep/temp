package foo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

const failID = -1

func Run() {
	ints := make(chan int)

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	g, ctx := errgroup.WithContext(ctx)

	defer cancel()

	g.Go(func() error {
		defer close(ints)
		for i := 1; i <= 5; i++ {
			select {
			case ints <- i:
				time.Sleep(1 * time.Second)
			case <-ctx.Done():
				fmt.Println(">>> A DONE")
				return ctx.Err()
			}
		}
		return nil
	})

	for i := 0; i < 5; i++ {
		id := i + 1
		g.Go(func() error {
			return doStuff(g, ctx, id, ints)
		})
	}

	err := g.Wait()

	fmt.Println("Err", err)
}

func doStuff(g *errgroup.Group, ctx context.Context, id int, ints <-chan int) error {
	for i := 0; i < 5; i++ {
		id := id + 10
		g.Go(func() error {
			return doSubStuff(ctx, id, ints)
		})
	}

	for {
		select {
		case i, ok := <-ints:
			if !ok {
				return nil
			}
			if failID == -1 {
				time.Sleep(3 * time.Second)
			} else if failID == id && i > 1 {
				return errors.New("failed1")
			}
			fmt.Println(id, " >>I:", i)
		case <-ctx.Done():
			fmt.Println(id, ">>> B DONE")
			return ctx.Err()
		}
	}

	return nil
}

func doSubStuff(ctx context.Context, id int, ints <-chan int) error {
	for {
		select {
		case i, ok := <-ints:
			if !ok {
				return nil
			}
			if failID == id && i > 1 {
				return errors.New("sub failed1")
			}

			fmt.Println(id, " >>SI:", i)
		case <-ctx.Done():
			fmt.Println(id, ">>> SB DONE")
			return ctx.Err()
		}
	}
	return nil
}

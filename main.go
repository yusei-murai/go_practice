package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("miss")
	}
}

func run(ctx context.Context) error {
	s := &http.Server{
		Addr: ":18080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w)
		}),
	}

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("miss")
			return err
		}
		return nil
	})

	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("miss")
	}

	return eg.Wait()
}

package main

import (
	"fmt"
	"github.com/tomasen/realip"
	"golang.org/x/time/rate"
	"net/http"
	"sync"
	"time"
)

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.Header().Set("Connect", "close")
					app.serverError(w, r, fmt.Errorf("%s", err))
				}
			}()
		},
	)
}

// rateLimit limits the rate of requests using the golang.org/x/time/rate
// package. It also handles race conditions.
func (app *application) rateLimit(next http.Handler) http.Handler {
	type client struct {
		limiter  *rate.Limiter
		lastSeen time.Time
	}

	var (
		mu      sync.Mutex
		clients = make(map[string]*client)
	)

	// Creates a go routine that checks when a client was last seen,
	// so that we can delete old clients who haven't been seen in a
	// while.
	go func() {
		for {
			time.Sleep(time.Minute)

			mu.Lock()

			for ip, client := range clients {
				if time.Since(client.lastSeen) > 3*time.Minute {
					delete(clients, ip)
				}
			}

			mu.Unlock()
		}
	}()

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if app.config.limiter.enabled {
				ip := realip.FromRequest(r)

				mu.Lock()

				if _, found := clients[ip]; !found {
					clients[ip] = &client{
						limiter: rate.NewLimiter(
							rate.Limit(app.config.limiter.rps),
							app.config.limiter.burst,
						),
					}
				}

				clients[ip].lastSeen = time.Now()

				if !clients[ip].limiter.Allow() {
					mu.Unlock()
					app.rateLimitExceededResponse(w, r)
					return
				}

				mu.Unlock()
			}

			next.ServeHTTP(w, r)
		},
	)
}

package middleware

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

// This file is responsible for setting a limit to requests by user ip.
// Current limit: 125

// Rate limit tested and working on a local IP address !!!
// Still, tests must be made in a production environment !!!
// Also config WAF/Nginx in production !!!

var (
	store           = memory.NewStore()                             // In-memory for now, since application is small (*subject to changes)
	rate            = limiter.Rate{Period: time.Minute, Limit: 125} // Rate limit of 125 request per minute (*subject to changes)
	limiterInstance = limiter.New(store, rate)
)

func getIP(req *http.Request) string {
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return req.RemoteAddr
	}
	return ip
}

func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		clientIP := getIP(req)
		context, err := limiterInstance.Get(req.Context(), clientIP)
		if err != nil {
			http.Error(w, "Error handling rate limit.", http.StatusInternalServerError)
			return
		}

		// Debugging rate limite (ip making the request, remaining requests and limit)
		// fmt.Printf("Request from %s | Remaining: %d | Limit Reached: %v\n",
		// clientIP, context.Remaining, context.Reached)

		w.Header().Set("X-RateLimit-Limit", fmt.Sprintf("%d", rate.Limit))
		w.Header().Set("X-RateLimit-Remaining", fmt.Sprintf("%d", context.Remaining))

		if context.Reached {
			http.Error(w, "> Error! 429: Too Many Requests", http.StatusTooManyRequests)
			time.Sleep(5 * time.Second)
			return
		}

		next.ServeHTTP(w, req)
	})
}

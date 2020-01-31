package duo

import (
	"context"
	"log"
	"net/http"
	"sync"
	"testing"
)

func startHTTPServer(wg *sync.WaitGroup) *http.Server {
	srv := &http.Server{Addr: "localhost:3000"}

	fs := http.FileServer(http.Dir("testing"))
	http.Handle("/", fs)

	go func() {
		defer wg.Done()
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()
	return srv
}
func TestParser(t *testing.T) {
	httpServerExitDone := &sync.WaitGroup{}

	httpServerExitDone.Add(1)
	srv := startHTTPServer(httpServerExitDone)

	ParseMainPage("http://localhost:3000/index.html")

	if err := srv.Shutdown(context.TODO()); err != nil {
		panic(err)
	}
	httpServerExitDone.Wait()
}

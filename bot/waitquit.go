package bot

// Handle SIGINT and SIGTERM with a graceful shutdown

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	sigs := make(chan os.Signal, 1)

	// `signal.Notify` registers the given channel to
	// receive notifications of the specified signals.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// This goroutine executes a blocking receive for
	// signals. When it gets one it'll print it out
	// and then notify the program that it can finish.
	go func() {
		sig := <-sigs
		shutdownMutex.Lock()
		shuttingDown = true
		shutdownMutex.Unlock()
		Log(Info, fmt.Sprintf("Received signal: %s, shutting down gracefully", sig))
		// Wait for all plugins to stop running
		plugRunningWaitGroup.Wait()
		// Get the dataLock to make sure the brain is in a consistent state
		dataLock.Lock()
		Log(Info, fmt.Sprintf("Exiting on signal: %s", sig))
		// How long does it _actually_ take for the message to go out?
		time.Sleep(time.Second)
		os.Exit(0)
	}()
}
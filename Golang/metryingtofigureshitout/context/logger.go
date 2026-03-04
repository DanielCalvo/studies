package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

// https://pkg.go.dev/context#WithValue
// The provided key must be comparable and should not be of type string or any other built-in type to avoid collisions between packages using context.
type ctxKey struct{}

var loggerKey ctxKey

func main() {
	opts := slog.HandlerOptions{Level: slog.LevelDebug}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &opts))
	ctx := context.WithValue(context.Background(), loggerKey, logger)
	doSomethingWithContext(ctx)

	//lets do it again using the helper functions
	logger2 := slog.New(slog.NewTextHandler(os.Stdout, nil))
	ctx2 := WithLogger(context.Background(), logger2)
	doSomethingWithContext2(ctx2)

}

func doSomethingWithContext(ctx context.Context) {
	fmt.Println("I did something here, but oh noes, an error happened! I want to log this error")
	logger, ok := ctx.Value(loggerKey).(*slog.Logger)
	//context can panic if the key/value is missing, so check if that is not nil
	//but even if its not nil, its also important to check if the type assertion succeeded with ok, maybe there is a value in there but it can be the wrong type!
	if !ok || logger == nil {
		return
	}
	logger.Error("something went wrong!")
}

//lets create some helper functions to make our life a bit easier

// this one returns a context with a logger for us
func WithLogger(ctx context.Context, l *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, l)

}

// this one retrieves the logger with the safety checks we have above
func LoggerFromContext(ctx context.Context) (*slog.Logger, bool) {
	logger, ok := ctx.Value(loggerKey).(*slog.Logger)
	if !ok || logger == nil {
		return nil, false
	}
	return logger, true
}

// so heres another implementation of our function, now using the helper!
func doSomethingWithContext2(ctx context.Context) {
	fmt.Println("I did something here, but oh noes, an error happened! I want to log this error2")
	logger, ok := LoggerFromContext(ctx)
	if !ok {
		return
	}
	logger.Error("something went wrong2!")
}

//https://go.dev/blog/slog

package main

import (
	"log/slog"
	"os"
)

func main() {
	slog.Info("hello, world")
	//Unlike with the log package, we can easily add key-value pairs to our output by writing them after the message:
	slog.Info("hello, world", "user", os.Getenv("USER"))

	//log’s top-level functions use the default logger. We can get this logger explicitly, and call its methods:
	logger := slog.Default()
	logger.Info("hello, world", "user", os.Getenv("USER"))

	//you can create a json handler!
	logger2 := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger2.Info("hello, world", "user", os.Getenv("USER")) //neat, everything is turned into a key value pair!

	//you can also create your handler by implementing the slog.Handler interface!
	//you can pass a context.Context to some log functions so a handler can extract context information like trace IDs.
	//neat: You can call Logger.With to add attributes to a logger that will appear in all of its output, effectively factoring out the common parts of several log statements.

}

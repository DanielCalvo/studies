## Context
Some note about context!
- Context is really useful for inter-process communication and finishing multiple go routines when a certain thing finishes, like closing go routinges inside a webhandler function when a certain time finishes!
- Or cancelling some sort of value generator go routine that is running in the background when you already have enough values!
- The context is not meant to pass parameter functions, but using it for request-scoped data is OK. You can use the context to use the context of the request throughout your entire system.
- Some useful usage guidelines are here: https://pkg.go.dev/context
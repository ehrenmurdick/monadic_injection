# Go error handling with monads!

Tired of writing if err != nil freaking everywhere? Try monads!

Start by looking at [main.go](main.go)

My favorite part of this approach is that you can combine the monads with DI. A server struct could for example have a MaybeItemRepo injected, errors from which to be handled at leisure! (Hence the name of the repo)

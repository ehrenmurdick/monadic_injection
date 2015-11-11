# Go error handling with monads!

Tired of writing if err != nil freaking everywhere? Try monads!

Start by looking at [main.go](main.go). Run it, and try requesting `http://localhost:8080/1`, `http://localhost:8080/10` and `http://localhost:8080/50`.

My favorite part of this approach is that you can combine the monads with DI. A server struct could for example have a ResultItemRepo injected, errors from which to be handled at leisure! (Hence the name of the repo)

I suggest reading [this](http://codon.com/refactoring-ruby-with-monads) if the monads bits don't make sense.
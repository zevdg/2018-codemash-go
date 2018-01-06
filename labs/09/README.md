# Service scoped variables

In this lab, we'll be refactoring our service so that we only initialize the wordbank once on startup instead of on every request.

1. Create a `Service` type with a `*WordBank` field.
1. Create a `NewService` constructor function that takes a filepath string as input and returns a `*Service` and an `error`.
1. Move the initialization of the wordbank from `generateIpsum` into this constructor.
1. Make `generateSentence`, `generateIpsum`, and `ipsumHandler` all methods on this new struct and fix any sytax errors
1. Create an instance of this type named `svc` in your main method right before the call to `http.HandleFunc` and then pass `svc.ipsumHandler` into `http.HandleFunc`.
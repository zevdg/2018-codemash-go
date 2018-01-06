# The net/http package

In this lab, we'll be using the [net/http](https://golang.org/pkg/net/http/) package to convert our cli tool into a very simple webserver.

1. Create a new function with the following signature.

        func ipsumHandler(w http.ResponseWriter, r *http.Request)

1. Move the call to `generateIpsum` and the corresponding error handling from `main` into `ipsumHandler`.
1. In `main`, use [http.HandleFunc](https://golang.org/pkg/net/http/#HandleFunc) to attach our `ipsumHandler` function to the pattern "/ipsum".  
*Note: `http.Handlefunc` is the first function we've seen that takes another function as input.  This is why the step 1 instructions have such a specific function signature.  The signature of `ipsumHandler` must exactly match the signature expected by `http.HandleFunc`.*
1. At the end of your main method, use [http.ListenAndServe](https://golang.org/pkg/net/http/#ListenAndServe) to start your webservice.  
*Hint: Don't forget to handle the error returned by ListenAndServe*
1. At this point, you can run your server and hit the endpoint you've created (e.g. `localhost:8080/ipsum`) in a browser.  When you do, you'll see your lorem ipsum printed to the command line.
1. To get the lorem ipsum to be printed to the webpage itself, use [fmt.Fprint](https://golang.org/pkg/fmt/#Fprint) to print the paragraph to the `http.ResponseWriter`.
1. Finally, we shouldn't only send success responses to the web client.  When errors happen, it is customary to print some debug info locally but **also** return an apportiate http status code to the client.  If `generateIpsum` returns and error, use [w.WriteHeader](https://golang.org/pkg/net/http/#ResponseWriter) to send [http.StatusInternalServerError](https://golang.org/pkg/net/http/#pkg-constants) to the client.
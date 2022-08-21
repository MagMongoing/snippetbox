### Fixed path and subtree patterns
Go's servemux supports two different type of URL patterns:
- fixed paths don't end with a trainling slash
- subtree paths do end with a trainling slash

Our two new patterns -- "/snippet/view" and "/snippet/create" -- are both examples of fixed paths. In Go's servemux, fixed path patterns like these only matched (and corresponding handler called) when the request URL path exactly matches the fixed path.

In contrast, our pattern "/" is an example of a subtree path (because it ends in a trailing slash). Another example would be something like "/static/". Subtree path patterns are matched (and the corresponding handler called) whenever the start of a request URL path matches the subtree path. If it helps your understanding, you can think of subtree paths as acting a bit like they have a wildcard at the end, like "/**" or "/static/**".

This helps explain why the "/" pattern is acting like a catch-all. The pattern essentiallly means match a single slash, followed by anyting (or nothing at all).

### Restricting the root url pattern
So what if you don't want the "/" pattern to act like a catch-all.
For instance, in the application we're building we want home page to be display if --- and only if -- the request ULR path exactly match "/". Otherwise, we want the user to receive a 404 page not found response.

It's not possible to change the behavior of Go's servemux to do this, but you can include a simple check in the home handler which ultimately has the same effect:

### The DefaultSeveMux
If you've been working with Go for a while you might have come across the http.Handle() and http.HandleFunc() functions. These allow you to register routs without declaring a servemux, like this:
```go
http.HandleFunc("/", home)
```
## Http Headers
### Content-Type
The `http.DetectContentType()` function generally works quite well, but a common gotcha for web developers new to go is that it can't distinguish JSON from plan text. So, by default, JSON response will be sent with a Content-Type: text/plain; charset=utf-8 header. You can prevent this from happening by setting the correct header manually like so:

```go
w.Header().Set("Content-Type", "application/json")
w.Write([]byte(`{"name":"Alex"}`))
```
### Manipulating the header map
```go
// Set a new cache-control header. If an existing "Cache-Control" header exists
// it'll be overwritten
w.Header().Set("Cache-Control","public, max-age=31536000")
// In contrast, the Add() method appends a new "Cache-Control" header and can
// be called multiple times.
w.Header().Add("Cache-Control", "public")
w.Header().add("Cache-Control", "max-age=315360000")
// Delete all values for the "Cache-Control" header.
w.Header().Delete("Cache-Control")
// Retrieve the first value for the "Cache-Control" header.
w.Header().Get("Cache-Control")
// Retrieve a slice of all values for the "Cache-Control" header.
w.Header().Values("Cache-Control")
```
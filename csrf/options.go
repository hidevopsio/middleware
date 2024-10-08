package csrf

import "github.com/hidevopsio/iris/context"

// Option describes a functional option for configuring the CSRF handler.
type Option func(*Csrf)

// MaxAge sets the maximum age (in seconds) of a CSRF token's underlying cookie.
// Defaults to 12 hours.
func MaxAge(age int) Option {
	return func(cs *Csrf) {
		cs.opts.MaxAge = age
	}
}

// Domain sets the cookie domain. Defaults to the current domain of the request
// only (recommended).
//
// This should be a hostname and not a URL. If set, the domain is treated as
// being prefixed with a '.' - e.g. "example.com" becomes ".example.com" and
// matches "www.example.com" and "secure.example.com".
func Domain(domain string) Option {
	return func(cs *Csrf) {
		cs.opts.Domain = domain
	}
}

// Path sets the cookie path. Defaults to the path the cookie was issued from
// (recommended).
//
// This instructs clients to only respond with cookie for that path and its
// subpaths - i.e. a cookie issued from "/register" would be included in requests
// to "/register/step2" and "/register/submit".
func Path(p string) Option {
	return func(cs *Csrf) {
		cs.opts.Path = p
	}
}

// Secure sets the 'Secure' flag on the cookie. Defaults to true (is recommended but requires https).
// Set this to 'false' in your development environment otherwise the cookie won't
// be sent over an insecure channel. Setting this via the presence of a 'DEV'
// environmental variable is a good way of making sure this won't make it to a
// production environment.
func Secure(s bool) Option {
	return func(cs *Csrf) {
		cs.opts.Secure = s
	}
}

// HTTPOnly sets the 'HTTPOnly' flag on the cookie. Defaults to true (recommended).
func HTTPOnly(h bool) Option {
	return func(cs *Csrf) {
		// Note that the function and field names match the case of the
		// related http.Cookie field instead of the "correct" HTTPOnly name
		// that golint suggests.
		cs.opts.HTTPOnly = h
	}
}

// ErrorHandler allows you to change the handler called when CSRF request
// processing encounters an invalid token or request. A typical use would be to
// provide a handler that returns a static HTML file with a HTTP 403 status. By
// default a HTTP 403 status and a plain text CSRF failure reason are served.
//
// Note that a custom error handler can also access the csrf.FailureReason(r)
// function to retrieve the CSRF validation reason from the request context.
func ErrorHandler(h context.Handler) Option {
	return func(cs *Csrf) {
		cs.opts.ErrorHandler = h
	}
}

// RequestHeader allows you to change the request header the CSRF middleware
// inspects. The default is X-CSRF-Token.
func RequestHeader(header string) Option {
	return func(cs *Csrf) {
		cs.opts.RequestHeader = header
	}
}

// FieldName allows you to change the name attribute of the hidden <input> field
// inspected by this package. The default is 'gorilla.csrf.Token'.
func FieldName(name string) Option {
	return func(cs *Csrf) {
		cs.opts.FieldName = name
	}
}

// CookieName changes the name of the CSRF cookie issued to clients.
//
// Note that cookie names should not contain whitespace, commas, semicolons,
// backslashes or control characters as per RFC6265.
func CookieName(name string) Option {
	return func(cs *Csrf) {
		cs.opts.CookieName = name
	}
}

// setStore sets the store used by the CSRF middleware.
// Note: this is private (for now) to allow for internal API changes.
func setStore(s store) Option {
	return func(cs *Csrf) {
		cs.st = s
	}
}

// parseOptions parses the supplied options functions and returns a configured
// csrf handler.
func parseOptions(opts ...Option) *Csrf {
	cs := new(Csrf)

	// Default to true. See Secure & HttpOnly function comments for rationale.
	// Set here to allow package users to override the default.
	cs.opts.Secure = true
	cs.opts.HTTPOnly = true

	// Range over each options function and apply it
	// to our csrf type to configure it. Options functions are
	// applied in order, with any conflicting options overriding
	// earlier calls.
	for _, option := range opts {
		option(cs)
	}

	return cs
}

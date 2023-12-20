package constants

const (
	AllowOrigin     = "*, localhost:5080" // more specific "localhost:3000, google.com"
	AllowCredential = "true"
	AllowHeader     = "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, User-Agent, Accept"
	AllowMethods    = "POST, GET, PUT, DELETE, PATCH"
	MaxAge          = "43200" // for 12 hour
)

package char

var (
	CRLF  = []byte{'\r', '\n'}
	CRLF2 = []byte{'\r', '\n', '\r', '\n'}

	// Semis    = []byte{';'}
	// Comma    = []byte{','}

	Spaces   = []byte{' '}
	Hashtags = []byte{'#'}
	Colons   = []byte(":")

	Charset  = []byte("charset")
	Boundary = []byte("boundary")
	Chunked  = []byte("chunked")

	At            = byte('@')
	Colon         = byte(':')
	Comma         = byte(',')
	Equal         = byte('=')
	Hashtag       = byte('#')
	Hyphen        = byte('-')
	Space         = byte(' ')
	Semi          = byte(';')
	QuotationMark = byte('"')
	QuestionMark  = byte('?')

	SemiSpace = []byte("; ")

	ColonSlashSlash          = []byte("://")
	Slash                    = []byte("/")
	SlashSlash               = []byte("//")
	SlashDotSlash            = []byte("/./")
	SlashDotDot              = []byte("/..")
	SlashDotDotSlash         = []byte("/../")
	BackSlashDotDot          = []byte(`\..`)
	BackSlashDotBackSlash    = []byte(`\.\`)
	BackSlashDotDotBackSlash = []byte(`\..\`)
	SlashDotDotBackSlash     = []byte(`/..\`)

	HTTP  = []byte("http")
	HTTPS = []byte("https")

	Date               = []byte("Date")
	TE                 = []byte("TE")
	Trailer            = []byte("Trailer")
	TransferEncoding   = []byte("Transfer-Encoding")
	AcceptRanges       = []byte("Accept-Ranges")
	ContentRange       = []byte("Content-Range")
	IfRange            = []byte("If-Range")
	Range              = []byte("Range")
	Server             = []byte("Server")
	From               = []byte("From")
	Host               = []byte("Host")
	Referer            = []byte("Referer")
	ReferrerPolicy     = []byte("Referrer-Policy")
	UserAgent          = []byte("User-Agent")
	Location           = []byte("Location")
	ContentEncoding    = []byte("Content-Encoding")
	ContentLanguage    = []byte("Content-Language")
	ContentLength      = []byte("Content-Length")
	ContentLocation    = []byte("Content-Location")
	ContentType        = []byte("Content-Type")
	ContentDisposition = []byte("Content-Disposition")
	Connection         = []byte("Connection")
	ProxyConnection    = []byte("Proxy-Connection")
	Accept             = []byte("Accept")
	AcceptCharset      = []byte("Accept-Charset")
	AcceptEncoding     = []byte("Accept-Encoding")
	AcceptLanguage     = []byte("Accept-Language")
	Cookie             = []byte("Cookie")
	Expect             = []byte("Expect")
	MaxForwards        = []byte("Max-Forwards")
	SetCookie          = []byte("Set-Cookie")
	ETag               = []byte("ETag")
	IfMatch            = []byte("If-Match")
	IfModifiedSince    = []byte("If-Modified-Since")
	IfNoneMatch        = []byte("If-None-Match")
	IfUnmodifiedSince  = []byte("If-Unmodified-Since")
	LastModified       = []byte("Last-Modified")
	Vary               = []byte("Vary")

	KeepAlive = []byte("keep-alive")
	Close     = []byte("close")
	Bytes     = []byte("bytes")
	Timeout   = []byte("timeout")
	Max       = []byte("max")

	Expires  = []byte("Expires")
	MaxAge   = []byte("Max-Age")
	Domain   = []byte("Domain")
	Path     = []byte("Path")
	Secure   = []byte("Secure")
	HttpOnly = []byte("HttpOnly")
	SameSite = []byte("SameSite")

	None   = []byte("None")
	Lax    = []byte("Lax")
	Strict = []byte("Strict")
)

// 	n := bytes.LastIndex(b, strSlashDotDot)
//			n := bytes.Index(b, strBackSlashDotBackSlash)

type Server struct {
	Addr string
	Handler Handler
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	TLSConfig *tls.Config
}

type Handler interface {
	ServerHTTP(ResponseWriter, *Request)
}

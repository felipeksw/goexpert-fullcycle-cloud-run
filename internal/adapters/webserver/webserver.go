package webserver

import "net/http"

type WebServer struct {
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	s.Handlers[path] = handler
}

/*
A criação do mux eu estava achando que deveria ser no NewWebServer
para que o novo Mux criado faça parte do "objeto" WebServer
Eu não consegui colocar isso no NewWebServer e não sei se seria qual o melhor lugar
*/
func (s *WebServer) Start() {
	mux := http.NewServeMux()

	for path, handler := range s.Handlers {
		mux.Handle(path, handler)
	}

	err := http.ListenAndServe(s.WebServerPort, mux)
	if err != nil {
		panic(err)
	}
}

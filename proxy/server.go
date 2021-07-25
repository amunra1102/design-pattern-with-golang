package proxy

type Server interface {
	handleRequest(url string, method string) (int, string)
}

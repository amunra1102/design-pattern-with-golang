package proxy

type application struct {}

func (a *application) handleRequest(url, method string) (int, string) {
	if url == "app/status" && method == "GET" {
		return 200, "OK"
	}

	if url == "create/user" && method == "POST" {
		return 200, "user created"
	}

	return 404, "not found"
}

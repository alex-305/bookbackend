package swagger

import "net/http"

func Handle(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./docs/api.json")
}

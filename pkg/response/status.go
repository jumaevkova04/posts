package response

import "net/http"

func Status(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status": "ok",
	}

	if err := JSON(w, http.StatusOK, data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

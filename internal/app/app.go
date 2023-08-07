package app

import (
	"net/http"

	"github.com/Jokaru-py/API-for-arithmetic/internal/utils"
)

var UserAccessAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := make(map[string]interface{})
		accessHeader := r.Header.Get("User-Access") //Получение токена

		if accessHeader == "" || accessHeader != "superuser" { //Токен отсутствует, возвращаем  403 http-код Unauthorized
			response = utils.Message(false, "Missing User-Access")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}

		next.ServeHTTP(w, r) //передать управление следующему обработчику!
	})
}

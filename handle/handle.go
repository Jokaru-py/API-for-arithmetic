package handle

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Jokaru-py/API-for-arithmetic/internal/utils"
	"github.com/Jokaru-py/API-for-arithmetic/models"
)

func Operation(w http.ResponseWriter, r *http.Request) {
	operation := &models.Operation{}
	err := json.NewDecoder(r.Body).Decode(operation)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	fmt.Println(operation)
}

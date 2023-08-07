package handle

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

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

	str := strings.ReplaceAll(operation.Str, " ", "")

	var sum, number int
	var v rune
	strRune := []rune(str)
	for i := len(strRune) - 1; i >= 0; i-- {

		v = strRune[i]

		if v != 43 && v != 45 {
			number, err = strconv.Atoi(string(v))
			if err != nil {
				log.Println(err)
				// Ошибка
			}

			continue
		}

		if v == 43 { // Сложение
			sum += number
		} else if v == 45 { // Вычитание
			sum -= number
		}
	}

	fmt.Printf("sum: %v\n", sum)

	resp := utils.Message(true, "succes")
	resp["data"] = sum

	utils.Respond(w, resp)
}

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"qtim/internal/entities"
	"strings"
)

func main() {
	//Инициализируем мультиплексор
	mux := http.NewServeMux()
	//Прикручиваем к нему ручку
	mux.HandleFunc("/detect", detectHandler)
	//Скармливаем его серверу и слушаем
	log.Fatalln(http.ListenAndServe("localhost:8080", mux))
}

//Внутри логика ручки
func detectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		req := &entities.Request{}
		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			buildResponse(w, http.StatusBadRequest, nil)
			return
		}
		researchedStr := req.Str
		researchedChar := req.Char
		res := calculate(researchedStr, researchedChar)
		result := map[string]int{"count": res}
		response, err := json.Marshal(result)
		if err != nil {
			buildResponse(w, http.StatusInternalServerError, nil)
			return
		}
		buildResponse(w, http.StatusOK, response)
	}
	buildResponse(w, http.StatusBadRequest, nil)
}

//Считающая вхождения char в str функция
func calculate(researchedStr string, researchedChar string) (result int) {
	toLowerStr := strings.ToLower(researchedStr)
	toLowerChar := strings.ToLower(researchedChar)
	result = strings.Count(toLowerStr, toLowerChar)
	return
}

//Шаблон ответа
func buildResponse(w http.ResponseWriter, statusCode int, body []byte) {
	w.Header().Set("Content-Type", "application/json ; charset=utf-8")
	w.WriteHeader(statusCode)
	w.Write(body)
}

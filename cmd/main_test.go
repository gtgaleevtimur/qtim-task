package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"qtim/internal/entities"
	"testing"
)

//Тест логики функции подсчета внутри ручки
//Для ошибок разкоментить 35-40
func TestCalculate(t *testing.T) {
	for _, v := range []struct {
		x   string
		y   string
		exp int
	}{
		{
			x:   "Ехал я под рождество!",
			y:   "о",
			exp: 3,
		},
		{
			x:   "Погонял коня ножом и бичом!",
			y:   "ом",
			exp: 2,
		},
		{
			x:   "Передо мной темный лес стоит",
			y:   "лес",
			exp: 1,
		},
		//{
		//	x:   "Передо мной темный лес стоит",
		//	y:   "Передо мной темный лес стоит",
		//	exp: 0,
		//},
	} {

		res := calculate(v.x, v.y)
		if res != v.exp {
			t.Logf("В <%v> нет %v вхождений <%v>", v.x, v.exp, v.y)
			t.Fail()
		}
	}
}

//Тест логики ручки.
func TestDetectHandler(t *testing.T) {
	request := entities.Request{"Hello world!", "o"}
	bytesRequest, err := json.Marshal(request)
	if err != nil {
		t.Log(err)
	}
	reader := bytes.NewReader(bytesRequest)
	req, err := http.NewRequest("POST", "/detect", reader)
	if err != nil {
		t.Log(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(detectHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"count":2}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

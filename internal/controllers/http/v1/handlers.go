package v1

import (
	"encoding/json"
	"ims/internal/store"
	"net/http"
)

// определение входящих/выходящих запросов
// TODO: пока можно хранить только строки, добавить поддержку других типов данных
type setKeyInput struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Seconds int    `json:"seconds"`
}

type getKeyInput struct {
	Key string `json:"key"`
}

// возвращаем бул-наличие ключа помимо самих значений, так можно иметь одну структуру ответа, не в зависимости от того, есть ключ или нет
type getKeyOutput struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Exists bool   `json:"exists"`
}

type deleteKeyInput struct {
	Key string `json:"key"`
}

// хендл для задания ключа, отвечает только на POST-запрос
func SetKey(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var input setKeyInput
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	store := store.GetStore()
	store.Set(input.Key, input.Value, input.Seconds)
	//store.Set(input.Key, input.Value, time.Minute*5)
	w.WriteHeader(http.StatusOK)
}

// хендл для поиска ключа, отвечает только на GET-запрос
func GetKey(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var input getKeyInput
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	store := store.GetStore()
	value, err := store.Get(input.Key)
	var output getKeyOutput
	// можно было не проверять, значание по умолчанию у bool всегда false
	// НО для ясности действий
	if err != nil {
		output.Exists = false
	} else {
		output.Key = input.Key
		output.Value = value.(string)
		output.Exists = true
	}
	data, err := json.Marshal(&output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// хендл для удаления ключа, отвечает только на DELETE-запрос
func DeleteKey(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var input deleteKeyInput
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	store := store.GetStore()
	store.Delete(input.Key)
	w.WriteHeader(http.StatusOK)
}

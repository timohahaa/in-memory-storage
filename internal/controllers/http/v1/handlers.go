package v1

import (
	"encoding/json"
	"ims/internal/store"
	"net/http"
)

// определение интерфейса хранилища
// это интерфейс моего ИМС, избавляемся от деталей реализации, легче тестить, больше независимости кода
type Store interface {
	// задать значение ключа и продолжительность его жизни
	Set(key string, value any, seconds int)
	// получить ключ, если такого нет - возвращается ошибка
	Get(key string) (any, error)
	Delete(key string)
}

// кастомный handler для запросов
type myHttpHandler struct {
	store Store
}

func NewHttpHandler() *myHttpHandler {
	return &myHttpHandler{
		store: store.GetStore(),
	}
}

// ServeHTTP метод для совместимости с интерфейсом http.Handler
func (h *myHttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/set":
		h.setKey(w, r)
	case "/get":
		h.getKey(w, r)
	case "/delete":
		h.deleteKey(w, r)
	case "/health":
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}(w, r)
	}
}

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
func (h *myHttpHandler) setKey(w http.ResponseWriter, r *http.Request) {
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
	h.store.Set(input.Key, input.Value, input.Seconds)
	//store.Set(input.Key, input.Value, time.Minute*5)
	w.WriteHeader(http.StatusOK)
}

// хендл для поиска ключа, отвечает только на GET-запрос
func (h *myHttpHandler) getKey(w http.ResponseWriter, r *http.Request) {
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
	value, err := h.store.Get(input.Key)
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
func (h *myHttpHandler) deleteKey(w http.ResponseWriter, r *http.Request) {
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
	h.store.Delete(input.Key)
	w.WriteHeader(http.StatusOK)
}

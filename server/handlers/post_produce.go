package handlers

import (
	"encoding/json"
	"gannet/services"
	"gannet/services/produce"
	"io/ioutil"
	"net/http"
)

// PostProduce
// endpoint for the creation of produce,
// expects a body of an array of produce items,
// recieves the produce items, an than validates them, building an array of
// errors as a `[]map[string]string` in order to return a detailed list of whats wrong and why
// if no errors are found, than the data is inserted
func PostProduce(w http.ResponseWriter, r *http.Request) {
	var data []produce.ProduceItem
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = json.Unmarshal(body, &data)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	errors := []map[string]string{}
	for _, item := range data {
		e := make(map[string]string)

		if priceError := produce.ValidateUnitPrice(item.UnitPrice); priceError != nil {
			e["unitPrice"] = priceError.Error()
		}

		if nameError := produce.ValidateName(item.Name); nameError != nil {
			e["name"] = nameError.Error()
		}

		if produceCodeError := produce.ValidateProduceCode(item.ProduceCode); produceCodeError != nil {
			e["produceCode"] = produceCodeError.Error()
		}

		if len(e) != 0 {
			errors = append(errors, e)
		}
	}

	if len(errors) != 0 {
		w.WriteHeader(http.StatusBadRequest)

		data, err := json.Marshal(errors)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write(data)
		return
	}

	status, err := services.Insert(data...)

	if err != nil {
		w.WriteHeader(status)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

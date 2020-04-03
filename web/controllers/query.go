package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"github.com/chainHero/heroes-service/web/model"
)

func writeJsonResponse(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Content-Type", "application/json; text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, OPTIONS")
	// w.Header().Set("Access-Control-Allow-Methods", "Content-Type, Accept, Authorization, X-Requested-With, Origin, Accept")
	// w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	// w.Header().Set("Access-Control-Allow-Headers","Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.Write(bytes)
}

func (app *Application) ReadOrderHandler(w http.ResponseWriter, r *http.Request) {
	vals := mux.Vars(r)
	orderId := vals["orderId"]
	payload, err := app.Fabric.QueryReadOrder(orderId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJsonResponse(w, []byte(payload))
}

// queryOrderDetail
func (app *Application) QueryOrderDetailHandler(w http.ResponseWriter, r *http.Request) {
	vals := mux.Vars(r)
	orderId := vals["orderId"]
	payload, err := app.Fabric.QueryOrderDetail(orderId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJsonResponse(w, []byte(payload))
}

func (app *Application) ReadOrderPositionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId := vars["orderId"]
	payload, err := app.Fabric.QueryReadOrderPosition(orderId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJsonResponse(w, []byte(payload))
}

func (app *Application) GetOrdersByRangeHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	indexRange := new(model.IndexRange)
	err = json.Unmarshal(body, indexRange)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	payload, err := app.Fabric.QueryGetOrdersByRange(indexRange.StartIndex, indexRange.EndIndex)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJsonResponse(w, []byte(payload))
}

func (app *Application) GetHistoryForOrderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId := vars["orderId"]
	payload, err := app.Fabric.QueryGetHistoryForOrder(orderId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJsonResponse(w, []byte(payload))
}

func (app *Application) QueryOrdersByBrokerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	brokerId := vars["brokerId"]
	payload, err := app.Fabric.QueryOrdersByBroker(brokerId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJsonResponse(w, []byte(payload))
}

func (app *Application) QueryOrdersHandler(w http.ResponseWriter, r *http.Request) {
	payload, err := app.Fabric.QueryOrders()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJsonResponse(w, []byte(payload))
}

func (app *Application) QueryOrdersWithPaginationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	borkerId := vars["brokerId"]
	page_string := vars["page"]

	payload, err := app.Fabric.QueryOrdersWithPagination(borkerId, page_string)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJsonResponse(w, []byte(payload))
}

func (app *Application) QueryUsersListHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page_string := vars["page"]

	payload, err := app.Fabric.QueryUsersWithPagination(page_string)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJsonResponse(w, []byte(payload))
}

func (app *Application) QueryUserDetailHandler(w http.ResponseWriter, r *http.Request) {
	vals := mux.Vars(r)
	userId := vals["userId"]
	payload, err := app.Fabric.QueryUserDetail(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJsonResponse(w, []byte(payload))
}


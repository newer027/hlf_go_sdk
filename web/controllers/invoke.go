package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"github.com/chainHero/heroes-service/web/model"
)

func (app *Application) InitOrderHandler(w http.ResponseWriter, r *http.Request) {
	// decoder := json.NewDecoder(r.Body)

	var order model.Order
	// err := decoder.Decode(&order)

	bytesBody, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(bytesBody, &order) 

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	payload, err := app.Fabric.InvokeInitOrder(order.OrderId, order.FromAddress, order.ToAddress, 
		order.Content, order.WeightTon, order.TransFee, order.OrderState, order.GoodsOwnerId, 
		order.BrokerId, order.DriverId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJsonResponse(w, []byte(payload))
}

func (app *Application) ChangeStateOrderHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	changeState := new(model.ChangeState)
	err = json.Unmarshal(body, changeState)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	payload, err := app.Fabric.InvokeChangeStateOrder(changeState.OrderId, changeState.NewState)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJsonResponse(w, []byte(payload))
}

func (app *Application) UpdatePositionOrderHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	updatePosition := new(model.UpdatePosition)
	err = json.Unmarshal(body, updatePosition)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	payload, err := app.Fabric.InvokeUpdatePositionOrder(updatePosition.PositionId, updatePosition.OrderId,
		updatePosition.Sequence, updatePosition.TimePosition, updatePosition.PositionString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJsonResponse(w, []byte(payload))
}

func (app *Application) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId := vars["orderId"]
	payload, err := app.Fabric.InvokeDelete(orderId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJsonResponse(w, []byte(payload))
}

// InitStringHash
// "dataId", "orderId", "dataUrl", "shaResult", "comment"
func (app *Application) InitStringHashHandler(w http.ResponseWriter, r *http.Request) {
	var stringHash model.StringHash
	bytesBody, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(bytesBody, &stringHash) 

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	payload, err := app.Fabric.InvokeInitStringHash(stringHash.DataId, stringHash.OrderId, stringHash.DataUrl, stringHash.ShaResult, stringHash.Comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJsonResponse(w, []byte(payload))
}

// InitFileHash
// "fileId", "orderId", "dataUrl", "shaResult", "comment"
func (app *Application) InitFileHashHandler(w http.ResponseWriter, r *http.Request) {
	var fileHash model.FileHash
	bytesBody, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(bytesBody, &fileHash) 

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	payload, err := app.Fabric.InvokeInitFileHash(fileHash.FileId, fileHash.OrderId, fileHash.DataUrl, fileHash.ShaResult, fileHash.Comment, fileHash.Valid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJsonResponse(w, []byte(payload))
}

func (app *Application) InitUserHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	bytesBody, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(bytesBody, &user) 

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	payload, err := app.Fabric.InvokeInitUser(user.UserId, user.UserName, user.Role, user.Telephone, user.Valid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJsonResponse(w, []byte(payload))
}

func (app *Application) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	bytesBody, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(bytesBody, &user) 

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	payload, err := app.Fabric.InvokeUpdateUser(user.UserId, user.UserName, user.Role, user.Telephone, user.Valid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJsonResponse(w, []byte(payload))
}

func (app *Application) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	payload, err := app.Fabric.InvokeDeleteUser(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJsonResponse(w, []byte(payload))
}

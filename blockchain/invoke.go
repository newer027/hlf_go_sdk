package blockchain

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

// InvokeHello
func (setup *FabricSetup) InvokeInitOrder(orderId string, fromAddress string, 
	toAddress string, content string, weightTon string, transFee string, orderState string,
	goodsOwnerId string, brokerId string, driverId string) (string, error) {
	// Prepare arguments
	var args []string
	args = append(args, "initOrder")
	args = append(args, orderId)
	args = append(args, fromAddress)
	args = append(args, toAddress)
	args = append(args, content)
	args = append(args, weightTon)
	args = append(args, transFee)
	args = append(args, orderState)
	args = append(args, goodsOwnerId)
	args = append(args, brokerId)
	args = append(args, driverId)

	// Add data that will be visible in the proposal, like a description of the invoke request
	transientDataMap := make(map[string][]byte)
	transientDataMap["result"] = []byte("Transient data in hello invoke")

	// Create a request (proposal) and send it
	response, err := setup.client.Execute(channel.Request{ChaincodeID: setup.ChainCodeID, 
		Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2]), []byte(args[3]), []byte(args[4]), []byte(args[5]), []byte(args[6]), []byte(args[7]), []byte(args[8]), []byte(args[9]), []byte(args[10])}, TransientMap: transientDataMap})
	if err != nil {
		return "", fmt.Errorf("failed to create order: %v", err)
	}
	return string(response.TransactionID), nil
}

// InvokeHello
func (setup *FabricSetup) InvokeChangeStateOrder(orderId string, newState string) (string, error) {

	// Prepare arguments
	var args []string
	args = append(args, "changeStateOrder")
	args = append(args, orderId)
	args = append(args, newState)

	// Add data that will be visible in the proposal, like a description of the invoke request
	transientDataMap := make(map[string][]byte)
	transientDataMap["result"] = []byte("Transient data in hello invoke")

	// Create a request (proposal) and send it
	response, err := setup.client.Execute(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], 
		Args: [][]byte{[]byte(args[1]), []byte(args[2])}, TransientMap: transientDataMap})
	if err != nil {
		return "", fmt.Errorf("failed to move funds: %v", err)
	}

	return string(response.TransactionID), nil
}

// InvokeHello
func (setup *FabricSetup) InvokeUpdatePositionOrder(positionId string, orderId string, sequence string, 
	timePosition string, positionString string) (string, error) {

	// Prepare arguments
	var args []string
	args = append(args, "updatePositionOrder")
	args = append(args, orderId)
	args = append(args, positionId)
	args = append(args, sequence)
	args = append(args, timePosition)
	args = append(args, positionString)

	// Add data that will be visible in the proposal, like a description of the invoke request
	transientDataMap := make(map[string][]byte)
	transientDataMap["result"] = []byte("Transient data in hello invoke")

	// Create a request (proposal) and send it
	response, err := setup.client.Execute(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], 
		Args: [][]byte{[]byte(args[1]), []byte(args[2]), []byte(args[3]),[]byte(args[4]), []byte(args[5])}, TransientMap: transientDataMap})
	if err != nil {
		return "", fmt.Errorf("failed to move funds: %v", err)
	}

	return string(response.TransactionID), nil
}

// InvokeHello
func (setup *FabricSetup) InvokeDelete(orderId string) (string, error) {

	// Prepare arguments
	var args []string
	args = append(args, "delete")
	args = append(args, orderId)

	// Add data that will be visible in the proposal, like a description of the invoke request
	transientDataMap := make(map[string][]byte)
	transientDataMap["result"] = []byte("Transient data in hello invoke")

	// Create a request (proposal) and send it
	response, err := setup.client.Execute(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1])}, TransientMap: transientDataMap})
	if err != nil {
		return "", fmt.Errorf("failed to delete order: %v", err)
	}

	return string(response.TransactionID), nil
}

// InvokeHello
func (setup *FabricSetup) InvokeInitStringHash(dataId string, orderId string, dataUrl string, shaResult string, comment string) (string, error) {

	// Prepare arguments
	// "dataId", "orderId", "dataUrl", "shaResult", "comment"
	var args []string
	args = append(args, "initStringHash")
	args = append(args, dataId)
	args = append(args, orderId)
	args = append(args, dataUrl)
	args = append(args, shaResult)
	args = append(args, comment)

	// Add data that will be visible in the proposal, like a description of the invoke request
	transientDataMap := make(map[string][]byte)
	transientDataMap["result"] = []byte("Transient data in hello invoke")

	// Create a request (proposal) and send it
	response, err := setup.client.Execute(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], 
		Args: [][]byte{[]byte(args[1]), []byte(args[2]), []byte(args[3]),[]byte(args[4]), []byte(args[5])}, TransientMap: transientDataMap})
	if err != nil {
		return "", fmt.Errorf("failed to delete order: %v", err)
	}

	return string(response.TransactionID), nil
}

// InvokeHello
func (setup *FabricSetup) InvokeInitFileHash(fileId string, orderId string, dataUrl string, shaResult string, comment string, valid string) (string, error) {

	// Prepare arguments
	var args []string
	args = append(args, "initFileHash")
	args = append(args, fileId)
	args = append(args, orderId)
	args = append(args, dataUrl)
	args = append(args, shaResult)
	args = append(args, comment)
	args = append(args, valid)

	// Add data that will be visible in the proposal, like a description of the invoke request
	transientDataMap := make(map[string][]byte)
	transientDataMap["result"] = []byte("Transient data in hello invoke")

	// Create a request (proposal) and send it
	response, err := setup.client.Execute(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2]), []byte(args[3]),[]byte(args[4]), []byte(args[5]), []byte(args[6])}, TransientMap: transientDataMap})
	if err != nil {
		return "", fmt.Errorf("failed to delete order: %v", err)
	}

	return string(response.TransactionID), nil
}

// InvokeHello
func (setup *FabricSetup) InvokeInitUser(userId string, userName string, role string, telephone string, valid string) (string, error) {

	// Prepare arguments
	var args []string
	args = append(args, "initUser")
	args = append(args, userId)
	args = append(args, userName)
	args = append(args, role)
	args = append(args, telephone)
	args = append(args, valid)

	// Add data that will be visible in the proposal, like a description of the invoke request
	transientDataMap := make(map[string][]byte)
	transientDataMap["result"] = []byte("Transient data in hello invoke")

	// Create a request (proposal) and send it
	response, err := setup.client.Execute(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2]), []byte(args[3]),[]byte(args[4]), []byte(args[5])}, TransientMap: transientDataMap})
	if err != nil {
		return "", fmt.Errorf("failed to init user: %v", err)
	}

	return string(response.TransactionID), nil
}


// InvokeHello
func (setup *FabricSetup) InvokeUpdateUser(userId string, userName string, role string, telephone string, valid string) (string, error) {

	// Prepare arguments
	var args []string
	args = append(args, "updateUser")
	args = append(args, userId)
	args = append(args, userName)
	args = append(args, role)
	args = append(args, telephone)
	args = append(args, valid)

	// Add data that will be visible in the proposal, like a description of the invoke request
	transientDataMap := make(map[string][]byte)
	transientDataMap["result"] = []byte("Transient data in hello invoke")

	// Create a request (proposal) and send it
	response, err := setup.client.Execute(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2]), []byte(args[3]),[]byte(args[4]), []byte(args[5])}, TransientMap: transientDataMap})
	if err != nil {
		return "", fmt.Errorf("failed to delete order: %v", err)
	}

	return string(response.TransactionID), nil
}

// InvokeHello
func (setup *FabricSetup) InvokeDeleteUser(userId string) (string, error) {

	// Prepare arguments
	var args []string
	args = append(args, "deleteUser")
	args = append(args, userId)

	// Add data that will be visible in the proposal, like a description of the invoke request
	transientDataMap := make(map[string][]byte)
	transientDataMap["result"] = []byte("Transient data in hello invoke")

	// Create a request (proposal) and send it
	response, err := setup.client.Execute(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1])}, TransientMap: transientDataMap})
	if err != nil {
		return "", fmt.Errorf("failed to delete order: %v", err)
	}

	return string(response.TransactionID), nil
}


package blockchain

import (
	"fmt"
	"bytes"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func (setup *FabricSetup) QueryReadOrder(orderId string) (string, error) {
	var args []string
	args = append(args, "readOrder")
	args = append(args, orderId)

	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1])}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}
	return string(response.Payload), nil
}

// queryOrderDetail
func (setup *FabricSetup) QueryOrderDetail(orderId string) (string, error) {
	var args []string
	args = append(args, "queryOrderDetail")
	args = append(args, orderId)

	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1])}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}
	return string(response.Payload), nil
}

func (setup *FabricSetup) QueryReadOrderPosition(orderId string) (string, error) {
	var args []string
	args = append(args, "queryAssets")

	var buffer bytes.Buffer
	buffer.WriteString("{\"selector\":{\"orderId\":\"")
	buffer.WriteString(orderId)
	buffer.WriteString("\", \"docType\":\"position\"}}")
	args = append(args, buffer.String())
	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1])}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}
	return string(response.Payload), nil
}

func (setup *FabricSetup) QueryGetOrdersByRange(startIndex string, endIndex string) (string, error) {
	var args []string
	args = append(args, "getOrdersByRange")
	args = append(args, startIndex)
	args = append(args, endIndex)
	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2])}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}
	return string(response.Payload), nil
}

func (setup *FabricSetup) QueryGetHistoryForOrder(orderId string) (string, error) {
	var args []string
	args = append(args, "getHistoryForOrder")
	args = append(args, orderId)
	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1])}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}
	return string(response.Payload), nil
}

func (setup *FabricSetup) QueryOrdersByBroker(brokerId string) (string, error) {
	var args []string
	args = append(args, "queryOrdersByBroker")
	args = append(args, brokerId)
	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1])}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}
	return string(response.Payload), nil
}

func (setup *FabricSetup) QueryOrders() (string, error) {
	var args []string
	args = append(args, "queryAssets")

	var buffer bytes.Buffer
	buffer.WriteString("{\"selector\":{\"docType\":\"order\"}}")
	args = append(args, buffer.String())
	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1])}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}
	return string(response.Payload), nil
}

func (setup *FabricSetup) QueryOrdersWithPagination(brokerId string, page string) (string, error) {
	var args []string
	args = append(args, "queryOrdersWithPagination")

	var buffer bytes.Buffer
	buffer.WriteString("{\"selector\":{\"docType\":\"order\",\"brokerId\":\"")
	buffer.WriteString(brokerId)
	buffer.WriteString("\"}}")
	// value := fmt.Sprintf("{\"selector\":{\"docType\":\"order\",\"brokerId\":\"%s\"}}", brokerId)

	args = append(args, buffer.String())
	args = append(args, page)
	args = append(args, "")

	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]),[]byte(args[2]),[]byte(args[3])}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}
	return string(response.Payload), nil
}

func (setup *FabricSetup) QueryUsersWithPagination(page string) (string, error) {
	var args []string
	args = append(args, "queryOrdersWithPagination")

	var buffer bytes.Buffer
	buffer.WriteString("{\"selector\":{\"docType\":\"user\"}}")
	// value := fmt.Sprintf("{\"selector\":{\"docType\":\"order\",\"brokerId\":\"%s\"}}", brokerId)

	args = append(args, buffer.String())
	args = append(args, page)
	args = append(args, "")

	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]),[]byte(args[2]),[]byte(args[3])}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}
	return string(response.Payload), nil
}

// queryUserDetail
func (setup *FabricSetup) QueryUserDetail(userId string) (string, error) {
	var args []string
	args = append(args, "readUser")
	args = append(args, userId)

	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1])}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}
	return string(response.Payload), nil
}

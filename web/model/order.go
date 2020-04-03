package model

type IndexRange struct {
	StartIndex  	string `json:"startIndex"`
	EndIndex   		string `json:"endIndex"`
}

type ChangeState struct {
	OrderId  		string `json:"orderId"`
	NewState   		string `json:"newState"`
}

type Order struct {
	OrderId 		string `json:"orderId"`
	FromAddress 	string `json:"fromAddress"`
	ToAddress 		string `json:"toAddress"`
	Content 		string `json:"content"`
	WeightTon 		string `json:"weightTon"`
	TransFee 		string `json:"transFee"`
	OrderState 		string `json:"orderState"`
	GoodsOwnerId 	string `json:"goodsOwnerId"`
	BrokerId 		string `json:"brokerId"`
	DriverId 		string `json:"driverId"`
}

type UpdatePosition struct {
	PositionId 		string `json:"positionId"`
	OrderId 		string `json:"orderId"`
	Sequence 		string `json:"sequence"`
	TimePosition 	string `json:"timePosition"`
	PositionString 	string `json:"positionString"`	
}

type StringHash struct {
	DataId 			string `json:"dataId"`
	OrderId 		string `json:"orderId"`
	DataUrl 		string `json:"dataUrl"`
	ShaResult 		string `json:"shaResult"`
	Comment 		string `json:"comment"`	
}

type FileHash struct {
	FileId 			string `json:"fileId"`
	OrderId 		string `json:"orderId"`
	DataUrl 		string `json:"dataUrl"`
	ShaResult 		string `json:"shaResult"`
	Comment 		string `json:"comment"`	
	Valid 			string `json:"valid"`	
}

type User struct {
	ObjectType 		string  `json:"docType"`
	UserId			string	`json:"userId"`
	UserName		string	`json:"userName"`
	Role			string	`json:"role"`
	Telephone		string	`json:"telephone"`
	Valid			string	`json:"valid"`
}

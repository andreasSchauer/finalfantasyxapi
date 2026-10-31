package api

type serviceParams interface {
	AlBhedParams | TurnOrderParams | DropChanceParams | CritChanceParams
}

type serviceResponse interface {
	AlBhedResponse | TurnOrderResponse | DropChanceResponse | CritChanceResponse
}
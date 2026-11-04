package api

type ServiceParams interface {
	AlBhedParams | TurnOrderParams | DropChanceParams | CritChanceParams
}

type ServiceResponse interface {
	AlBhedResponse | TurnOrderResponse | DropChanceResponse | CritChanceResponse
}

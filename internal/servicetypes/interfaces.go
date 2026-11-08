// Package servicetypes ajsidoaüsdji
package servicetypes

type ServiceParams interface {
	AlBhedParams | TurnOrderParams | DropChanceParams | CritChanceParams
}

type ServiceResponse interface {
	AlBhedResponse | TurnOrderResponse | DropChanceResponse | CritChanceResponse
}

package orderstatus

//go:generate stringer -type OrderStatus

// OrderStatus new order status
type OrderStatus int32

const (
	Pending OrderStatus = iota
	Dispatched
	ArrivedTransit
	InTransit
	ArrivedDestination
	Received
	Completed
	Cancelled
)

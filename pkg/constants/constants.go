package constants

type OrderStatusStruct struct {
	Unpaid      string `json:"unpaid"`
	Pending     string `json:"pending"`
	Packing     string `json:"packing"`
	ReadyToShip string `json:"ready_to_ship"`
	Shiped      string `json:"shiped"`
	Completed   string `json:"completed"`
}

const (
	PathFileAddress = "./asset/thai_address.json"
)

var (
	OrderStatus = OrderStatusStruct{
		Unpaid:      "UNPAID",
		Pending:     "PENDING",
		Packing:     "PACKING",
		ReadyToShip: "READYTOSHIP",
		Shiped:      "SHIPED",
		Completed:   "COMPLETED",
	}
)

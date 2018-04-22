package warehouse

//go:generate stringer -type Warehouse

// Warehouse determines where you need to send the order
type Warehouse int32

const (
	JiaXingCN Warehouse = iota
	DongGuanCN
	UnitedStatesUS
	TaiWanCN
	SingaporeSG
	MalayLocalMY
	IncheonKR
)

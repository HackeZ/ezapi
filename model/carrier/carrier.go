package carrier

//go:generate stringer -type Carrier

// Carrier code in CN
type Carrier int32

const (
	SFExpress       Carrier = iota // 顺丰快递
	STOExpress                     // 申通快递
	YTOExpress                     // 圆通快递
	ZTOExpress                     // 中通快递
	BestExpress                    // 汇通快递（百世汇通）
	YundaExpress                   // 韵达快递
	TiantianExpress                // 天天快递
	EMSExpress                     // 邮政EMS
	QFKDExpress                    // 全峰快递
	UCExpress                      // 优速快递
	BESExpress                     // 百世快递
	ANWLExpress                    // 安能物流
	FASTExpress                    // 快捷快递
	SUERExpress                    // 速尔快递
	LBWLExpress                    // 龙邦物流
	GTKDExpress                    // 国通快递
	ZJSExpress                     // 宅急送
	YFKYExpress                    // 亚风快运
	CGWLExpress                    // 晨光物流
	DEWLExpress                    // 大恩物流
	XFWLExpress                    // 信丰快递
	JGWLExpress                    // 京广快递
	ANKDExpress                    // 安能快递
	YTKDExpress                    // 运通中港快递
	CRHExpress                     // 高铁快递
	DepponExpress                  // 德邦物流
	OtherExpress                   // 其他
	KYExpress                      // 跨越物流
)

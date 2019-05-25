package ops

import (
	"gxclient-go/core/types"
	"gxclient-go/core/transaction"
)

func NewWorkerCreateOperation(owner types.GrapheneID, workBeginDate, workEndDate types.Time, dailyPay types.UInt64,
	name, url string, initializer types.WorkerInitializer) *WorkerCreateOperation {
	op := &WorkerCreateOperation{
		Owner:				owner,
		WorkBeginDate:		workBeginDate,
		WorkEndDate:		workEndDate,
		DailyPay:			dailyPay,
		Name:				name,
		URL:				url,
		Initializer:		initializer,
	}

	return op
}

type WorkerCreateOperation struct {
	types.OperationFee
	Owner         types.GrapheneID        `json:"owner"`
	WorkBeginDate types.Time              `json:"work_begin_date"`
	WorkEndDate   types.Time              `json:"work_end_date"`
	DailyPay      types.UInt64            `json:"daily_pay"`
	Name          string                  `json:"name"`
	URL           string                  `json:"url"`
	Initializer   types.WorkerInitializer `json:"initializer"`
}

func (p WorkerCreateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(int8(p.Type()))
	enc.Encode(p.Fee)
	enc.Encode(p.Owner)
	enc.Encode(p.WorkBeginDate)
	enc.Encode(p.WorkEndDate)
	enc.Encode(p.DailyPay)
	enc.Encode(p.Name)
	enc.Encode(p.URL)
	enc.Encode(p.Initializer)
	return enc.Err()
}

func (op *WorkerCreateOperation) Type() types.OpType { return types.WorkerCreateOpType }
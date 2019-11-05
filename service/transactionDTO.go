package services

import (
	transferService "github.com/the-final-codedown/tfc-transfer-validator/proto"
	"time"
)

//TransactionDTO dto served to send the transaction on the kafka
type TransactionDTO struct {
	Id       string    `json:"id"`;
	Source   string    `json:"source"`;
	Receiver string    `json:"receiver"`;
	Amount   int32       `json:"amount"`;
	Date     time.Time `json:"date"`;
}

func (t *TransactionDTO) FromTransfer(transfer * transferService.Transfer) {
	t.Amount = transfer.Amount;
	t.Receiver = transfer.Destination;
	t.Source = transfer.Origin;
}

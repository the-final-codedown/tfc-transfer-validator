package main

import (
	"context"
	capUpdater "github.com/the-final-codedown/tfc-cap-updater/proto"
	transferService "github.com/the-final-codedown/tfc-transfer-validator/proto"
	"github.com/the-final-codedown/tfc-transfer-validator/services"
	"log"
	"time"
)

func (t TransferValidator) Pay(context context.Context, transfer *transferService.Transfer) (*transferService.TransferValidation, error) {
	log.Println("Payment validation")
	paymentCap, err := t.capReader.GetCap(transfer.Origin)

	if err != nil {
		log.Println("Error fetching cap : ", err)
		return &transferService.TransferValidation{
			Transfer:  transfer,
			Validated: false,
			Reason:    "Error fetching cap : account not found",
		}, nil
	}
	log.Println("Payment cap : ", paymentCap)
	if paymentCap < transfer.Amount {
		return &transferService.TransferValidation{
			Transfer:  transfer,
			Validated: false,
			Reason:    "Exceeding Cap",
		}, nil
	}

	transaction := services.TransactionDTO{Date: time.Now()}
	transaction.FromTransfer(transfer)
	err = t.kafkaClient.SendTransaction(&transaction)
	if err != nil {
		log.Println("error sending to main app", err)
	}
	downscale := &capUpdater.CapDownscale{
		AccountID: transfer.Origin,
		Value:     transfer.Amount,
	}
	resp, err := t.capUpdaterClient.DownscaleCap(context, downscale)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(resp.Accepted)
	}

	log.Println("Payment validated")

	return &transferService.TransferValidation{
		Transfer:  transfer,
		Validated: true,
	}, nil
}

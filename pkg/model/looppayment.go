package models

import (
	"context"
	"encoding/json"
	"strconv"
)

// {
//     "event": "TransferProcessed",
//     "transaction": "0x4e5e20a1cfca858b1def7ad70b9a286d046b084b47970b1850063b2ea86e8405",
//     "networkId": 11155111,
//     "networkName": "sepolia",
//     "contractAddress": "0xb436D38bC878E5a202Da9e609a549249D178f7fE",
//     "email": "a@s.com",
//     "company": "DIA Data (085b1ed6-c637-4f99-a034-3ea718bcce34)",
//     "parent": "-",
//     "transferId": "252fdc42-0935-4a89-bda2-7618f6dfcc40",
//     "success": true,
//     "paymentTokenAddress": "0x7b79995e5f793A07Bc00c21412e50Ecae098E7f9",
//     "paymentTokenSymbol": "WETH",
//     "endUser": "0xF231DB04c5d92396235506232Ca5F40fcf8dAfb2",
//     "reason": "",
//     "invoiceId": "DIADA-13",
//     "amountPaid": 0.00029615,
//     "agreementId": "7f75d1f9-0c6c-4b71-8892-0ab2aaab07c1",
//     "refId": "",
//     "batchId": "ded7256e-4722-4706-b768-4da06fd930f8",
//     "usdAmount": "1.00"
// }

type LoopPaymentTransferProcessed struct {
	Event               string  `json:"event"`
	Transaction         string  `json:"transaction"`
	NetworkID           int     `json:"networkId"`
	NetworkName         string  `json:"networkName"`
	ContractAddress     string  `json:"contractAddress"`
	Email               string  `json:"email"`
	Company             string  `json:"company"`
	Parent              string  `json:"parent"`
	TransferID          string  `json:"transferId"`
	Success             bool    `json:"success"`
	PaymentTokenAddress string  `json:"paymentTokenAddress"`
	PaymentTokenSymbol  string  `json:"paymentTokenSymbol"`
	EndUser             string  `json:"endUser"`
	Reason              string  `json:"reason"`
	InvoiceID           string  `json:"invoiceId"`
	AmountPaid          float64 `json:"amountPaid"`
	AgreementID         string  `json:"agreementId"`
	RefID               string  `json:"refId"`
	BatchID             string  `json:"batchId"`
	UsdAmount           string  `json:"usdAmount"`
	RedirectUrl         string  `json:"redirectURL"`
}

func (reldb *RelDB) InsertLoopPaymentTransferProcessed(ctx context.Context, record LoopPaymentTransferProcessed) error {
	query := `
    INSERT INTO loop_payment_transfer_processed (
        event, transaction, network_id, network_name, contract_address, email, company, parent, transfer_id, success, 
        payment_token_address, payment_token_symbol, end_user, reason, invoice_id, amount_paid, agreement_id, ref_id, batch_id, usd_amount
    ) VALUES (
        $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20
    )`

	_, err := reldb.postgresClient.Exec(ctx, query,
		record.Event, record.Transaction, record.NetworkID, record.NetworkName, record.ContractAddress,
		record.Email, record.Company, record.Parent, record.TransferID, record.Success,
		record.PaymentTokenAddress, record.PaymentTokenSymbol, record.EndUser, record.Reason,
		record.InvoiceID, record.AmountPaid, record.AgreementID, record.RefID, record.BatchID, record.UsdAmount)

	return err
}

func (reldb *RelDB) GetLastPaymentByEndUser(endUser string) (LoopPaymentTransferProcessed, error) {
	query := `
    SELECT event, transaction, network_id, network_name, contract_address, email, company, parent, transfer_id, success, 
           payment_token_address, payment_token_symbol, end_user, reason, invoice_id, amount_paid, agreement_id, ref_id, batch_id, usd_amount
    FROM loop_payment_transfer_processed
    WHERE end_user = $1
    ORDER BY transaction DESC
    LIMIT 1`

	var record LoopPaymentTransferProcessed
	err := reldb.postgresClient.QueryRow(context.Background(), query, endUser).Scan(
		&record.Event, &record.Transaction, &record.NetworkID, &record.NetworkName, &record.ContractAddress,
		&record.Email, &record.Company, &record.Parent, &record.TransferID, &record.Success,
		&record.PaymentTokenAddress, &record.PaymentTokenSymbol, &record.EndUser, &record.Reason,
		&record.InvoiceID, &record.AmountPaid, &record.AgreementID, &record.RefID, &record.BatchID, &record.UsdAmount)

	if err != nil {
		return record, err
	}

	return record, nil
}

func (reldb *RelDB) GetLastPaymentByAgreementID(agreementID string) (*LoopPaymentTransferProcessed, error) {
	query := `
    SELECT event, transaction, network_id, network_name, contract_address, email, company, parent, transfer_id, success, 
           payment_token_address, payment_token_symbol, end_user, reason, invoice_id, amount_paid, agreement_id, ref_id, batch_id, usd_amount
    FROM loop_payment_transfer_processed
    WHERE agreement_id = $1
    ORDER BY transaction DESC
    LIMIT 1`

	var record LoopPaymentTransferProcessed
	err := reldb.postgresClient.QueryRow(context.Background(), query, agreementID).Scan(
		&record.Event, &record.Transaction, &record.NetworkID, &record.NetworkName, &record.ContractAddress,
		&record.Email, &record.Company, &record.Parent, &record.TransferID, &record.Success,
		&record.PaymentTokenAddress, &record.PaymentTokenSymbol, &record.EndUser, &record.Reason,
		&record.InvoiceID, &record.AmountPaid, &record.AgreementID, &record.RefID, &record.BatchID, &record.UsdAmount)

	if err != nil {
		return nil, err
	}

	return &record, nil
}

type LoopPaymentTransferCreated struct {
	Event           string `json:"event"`
	Transaction     string `json:"transaction"`
	NetworkID       int    `json:"networkId"`
	NetworkName     string `json:"networkName"`
	ContractAddress string `json:"contractAddress"`
	Email           string `json:"email"`
	Company         string `json:"company"`
	Parent          string `json:"parent"`
	ID              string `json:"id"`
	InvoiceID       string `json:"invoiceId"`
	BillDate        int    `json:"billDate"`
	ToAddress       string `json:"toAddress"`
	FromAddress     string `json:"fromAddress"`
	TokenSymbol     string `json:"tokenSymbol"`
	TokenAddress    string `json:"tokenAddress"`
	PaymentType     string `json:"paymentType"`
	Usd             bool   `json:"usd"`
	Amount          string `json:"amount"`
	Item            string `json:"item"`
	ItemID          int    `json:"itemId"`
	Source          string `json:"source"`
	BatchID         string `json:"batchId"`
	RefID           string `json:"refId"`
	AgreementID     string `json:"agreementId"`
	TransferID      string `json:"transferId"`
}

// {
//     "event": "AgreementSignedUp",
//     "transaction": "-",
//     "networkId": 11155111,
//     "networkName": "sepolia",
//     "contractAddress": "0xb436d38bc878e5a202da9e609a549249d178f7fe",
//     "email": "a@s.com",
//     "company": "DIA Data (085b1ed6-c637-4f99-a034-3ea718bcce34)",
//     "parent": "-",
//     "subscriber": "0xf231db04c5d92396235506232ca5f40fcf8dafb2",
//     "item": "Product 2",
//     "itemId": "5597e580-5026-46b6-a0bf-97ae1a88bd0a",
//     "agreementId": "7f75d1f9-0c6c-4b71-8892-0ab2aaab07c1",
//     "agreementAmount": "1.00",
//     "frequencyNumber": 1,
//     "frequencyUnit": "Day",
//     "addOnAgreements": "",
//     "addOnItems": "",
//     "addOnItemIds": "",
//     "addOnTotalAmount": "0.00",
//     "paymentTokenSymbol": "WETH",
//     "paymentTokenAddress": "0x7b79995e5f793a07bc00c21412e50ecae098e7f9",
//     "eventDate": 1722344804,
//     "refId": "",
//     "metadata": {}
// }

type LoopPaymentAgreementSignedUp struct {
	Event               string `json:"event"`
	Transaction         string `json:"transaction"`
	NetworkID           int    `json:"networkId"`
	NetworkName         string `json:"networkName"`
	ContractAddress     string `json:"contractAddress"`
	Email               string `json:"email"`
	Company             string `json:"company"`
	Parent              string `json:"parent"`
	Subscriber          string `json:"subscriber"`
	Item                string `json:"item"`
	ItemID              string `json:"itemId"`
	AgreementID         string `json:"agreementId"`
	AgreementAmount     string `json:"agreementAmount"`
	FrequencyNumber     int    `json:"frequencyNumber"`
	FrequencyUnit       string `json:"frequencyUnit"`
	AddOnAgreements     string `json:"addOnAgreements"`
	AddOnItems          string `json:"addOnItems"`
	AddOnItemIds        string `json:"addOnItemIds"`
	AddOnTotalAmount    string `json:"addOnTotalAmount"`
	PaymentTokenSymbol  string `json:"paymentTokenSymbol"`
	PaymentTokenAddress string `json:"paymentTokenAddress"`
	EventDate           int    `json:"eventDate"`
	RefID               string `json:"refId"`
	CustomerID          string `json:"customerId"`

	Metadata struct {
	} `json:"metadata"`
}

type LoopPaymentResponse struct {
	Event string `json:"event"`

	Transaction         string `json:"transaction"`
	NetworkID           int    `json:"networkId"`
	NetworkName         string `json:"networkName"`
	ContractAddress     string `json:"contractAddress"`
	Email               string `json:"email"`
	Company             string `json:"company"`
	Parent              string `json:"parent"`
	Subscriber          string `json:"subscriber"`
	Item                string `json:"item"`
	ItemID              string `json:"itemId"`
	AgreementID         string `json:"agreementId"`
	AgreementAmount     string `json:"agreementAmount"`
	FrequencyNumber     int    `json:"frequencyNumber"`
	FrequencyUnit       string `json:"frequencyUnit"`
	AddOnAgreements     string `json:"addOnAgreements"`
	AddOnItems          string `json:"addOnItems"`
	AddOnItemIds        string `json:"addOnItemIds"`
	AddOnTotalAmount    string `json:"addOnTotalAmount"`
	PaymentTokenSymbol  string `json:"paymentTokenSymbol"`
	PaymentTokenAddress string `json:"paymentTokenAddress"`
	EventDate           int    `json:"eventDate"`
	RefID               string `json:"refId"`
	InvoiceID           string `json:"invoiceId"`
	CustomerID          string `json:"customerId"`

	Metadata struct {
	} `json:"metadata"`
}

func (reldb *RelDB) InsertLoopPaymentResponse(ctx context.Context, response LoopPaymentResponse) error {
	query := `
        INSERT INTO loop_payment_responses (
            event, transaction, network_id, network_name, contract_address, email, company,
            parent, subscriber, item, item_id, agreement_id, agreement_amount, frequency_number,
            frequency_unit, add_on_agreements, add_on_items, add_on_item_ids, add_on_total_amount,
            payment_token_symbol, payment_token_address, event_date, ref_id, invoice_id, metadata, customer_id
        ) VALUES (
            $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, 
            $20, $21, $22, $23, $24, $25, $26
        )
    `
	_, err := reldb.postgresClient.Exec(ctx, query,
		response.Event, response.Transaction, response.NetworkID, response.NetworkName,
		response.ContractAddress, response.Email, response.Company, response.Parent,
		response.Subscriber, response.Item, response.ItemID, response.AgreementID,
		response.AgreementAmount, response.FrequencyNumber, response.FrequencyUnit,
		response.AddOnAgreements, response.AddOnItems, response.AddOnItemIds,
		response.AddOnTotalAmount, response.PaymentTokenSymbol, response.PaymentTokenAddress,
		response.EventDate, response.RefID, response.InvoiceID, response.Metadata, response.CustomerID,
	)
	return err
}

func (reldb *RelDB) GetLoopPaymentResponseByAgreementID(ctx context.Context, agreementID string) (*LoopPaymentResponse, error) {
	query := `
        SELECT event, transaction, network_id, network_name, contract_address, email, company,
               parent, subscriber, item, item_id, agreement_id, agreement_amount, frequency_number,
               frequency_unit, add_on_agreements, add_on_items, add_on_item_ids, add_on_total_amount,
               payment_token_symbol, payment_token_address, event_date, ref_id, invoice_id, metadata, customer_id
        FROM loop_payment_responses
        WHERE agreement_id = $1
    `

	row := reldb.postgresClient.QueryRow(ctx, query, agreementID)

	var (
		response     LoopPaymentResponse
		metadataJSON []byte
		custID       int
	)

	err := row.Scan(
		&response.Event, &response.Transaction, &response.NetworkID, &response.NetworkName,
		&response.ContractAddress, &response.Email, &response.Company, &response.Parent,
		&response.Subscriber, &response.Item, &response.ItemID, &response.AgreementID,
		&response.AgreementAmount, &response.FrequencyNumber, &response.FrequencyUnit,
		&response.AddOnAgreements, &response.AddOnItems, &response.AddOnItemIds,
		&response.AddOnTotalAmount, &response.PaymentTokenSymbol, &response.PaymentTokenAddress,
		&response.EventDate, &response.RefID, &response.InvoiceID, &metadataJSON, &custID,
	)
	if err != nil {
		return nil, err
	}

	response.CustomerID = strconv.Itoa(custID)

	if err := json.Unmarshal(metadataJSON, &response.Metadata); err != nil {
		return nil, err
	}

	return &response, nil
}

func (reldb *RelDB) GetLoopPaymentResponseByCustomerID(ctx context.Context, customerID string) (*LoopPaymentResponse, error) {
	query := `
        SELECT event, transaction, network_id, network_name, contract_address, email, company,
               parent, subscriber, item, item_id, agreement_id, agreement_amount, frequency_number,
               frequency_unit, add_on_agreements, add_on_items, add_on_item_ids, add_on_total_amount,
               payment_token_symbol, payment_token_address, event_date, ref_id, invoice_id, metadata, customer_id
        FROM loop_payment_responses
        WHERE customer_id = $1 ORDER BY event_date DESC LIMIT 1
    `

	row := reldb.postgresClient.QueryRow(ctx, query, customerID)

	var (
		response     LoopPaymentResponse
		metadataJSON []byte
		custID       int
	)

	err := row.Scan(
		&response.Event, &response.Transaction, &response.NetworkID, &response.NetworkName,
		&response.ContractAddress, &response.Email, &response.Company, &response.Parent,
		&response.Subscriber, &response.Item, &response.ItemID, &response.AgreementID,
		&response.AgreementAmount, &response.FrequencyNumber, &response.FrequencyUnit,
		&response.AddOnAgreements, &response.AddOnItems, &response.AddOnItemIds,
		&response.AddOnTotalAmount, &response.PaymentTokenSymbol, &response.PaymentTokenAddress,
		&response.EventDate, &response.RefID, &response.InvoiceID, &metadataJSON, &custID,
	)
	if err != nil {
		return nil, err
	}

	response.CustomerID = strconv.Itoa(custID)

	if err := json.Unmarshal(metadataJSON, &response.Metadata); err != nil {
		return nil, err
	}

	return &response, nil
}

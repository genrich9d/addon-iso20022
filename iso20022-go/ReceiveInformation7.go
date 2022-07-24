package iso20022

// Parameters applied to the settlement of a security transfer.
type ReceiveInformation7 struct {

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Indicates whether the settlement amount includes the stamp duty amount.
	StampDuty *StampDutyType2Code `xml:"StmpDty,omitempty"`

	// Deal amount.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *ReceivingPartiesAndAccount8 `xml:"SttlmPtiesDtls"`

	// Charge related to the transfer of a financial instrument.
	ChargeDetails []*Charge20 `xml:"ChrgDtls,omitempty"`

	// Commission related to the transfer of a financial instrument.
	CommissionDetails []*Commission12 `xml:"ComssnDtls,omitempty"`

	// Tax related to the transfer of a financial instrument.
	TaxDetails []*Tax15 `xml:"TaxDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalTransfer *PhysicalTransferType1Code `xml:"PhysTrf,omitempty"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters4 `xml:"PhysTrfDtls,omitempty"`
}

func (r *ReceiveInformation7) SetSettlementAmount(value, currency string) {
	r.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *ReceiveInformation7) SetStampDuty(value string) {
	r.StampDuty = (*StampDutyType2Code)(&value)
}

func (r *ReceiveInformation7) SetNetAmount(value, currency string) {
	r.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *ReceiveInformation7) AddSettlementPartiesDetails() *ReceivingPartiesAndAccount8 {
	r.SettlementPartiesDetails = new(ReceivingPartiesAndAccount8)
	return r.SettlementPartiesDetails
}

func (r *ReceiveInformation7) AddChargeDetails() *Charge20 {
	newValue := new(Charge20)
	r.ChargeDetails = append(r.ChargeDetails, newValue)
	return newValue
}

func (r *ReceiveInformation7) AddCommissionDetails() *Commission12 {
	newValue := new(Commission12)
	r.CommissionDetails = append(r.CommissionDetails, newValue)
	return newValue
}

func (r *ReceiveInformation7) AddTaxDetails() *Tax15 {
	newValue := new(Tax15)
	r.TaxDetails = append(r.TaxDetails, newValue)
	return newValue
}

func (r *ReceiveInformation7) SetPhysicalTransfer(value string) {
	r.PhysicalTransfer = (*PhysicalTransferType1Code)(&value)
}

func (r *ReceiveInformation7) AddPhysicalTransferDetails() *DeliveryParameters4 {
	r.PhysicalTransferDetails = new(DeliveryParameters4)
	return r.PhysicalTransferDetails
}

package iso20022

// Instruction from an investor to sell investment fund units back to the fund.
type RedemptionOrder5 struct {

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of an order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for an order cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Account impacted by an investment fund order.
	InvestmentAccountDetails *InvestmentAccount21 `xml:"InvstmtAcctDtls"`

	// Specifies the category of the investment fund order.
	OrderType []*FundOrderType2 `xml:"OrdrTp,omitempty"`

	// Additional information about the investor.
	BeneficiaryDetails *IndividualPerson12 `xml:"BnfcryDtls,omitempty"`

	// Quantity of investment fund units redeemed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Amount of money used to derive the quantity of investment fund units to be sold, before deduction of charges, commissions, and taxes.
	// [Quantity * Price]
	GrossAmount *ActiveOrHistoricCurrencyAndAmount `xml:"GrssAmt"`

	// Amount of money to be received following redemption after deduction of charges, commissions and taxes and used to derive the quantity of investment fund units to be sold.
	// [(Quantity * Price) - (Charges + Commissions +Taxes)]
	NetAmount *ActiveOrHistoricCurrencyAndAmount `xml:"NetAmt"`

	// Portion of the investor's holdings, in a specific investment fund/ fund class, that is redeemed.
	HoldingsRedemptionRate *PercentageRate `xml:"HldgsRedRate"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Method by which the transaction is settled.
	SettlementMethod *DeliveryReceiptType2Code `xml:"SttlmMtd,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms6 `xml:"FXDtls,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Tax group to which the purchased investment fund units belong. The investor indicates to the intermediary operating pooled nominees, which type of unit is to be sold.
	Group1Or2Units *UKTaxGroupUnitCode `xml:"Grp1Or2Units,omitempty"`

	// Charge for the placement of an order.
	ChargeDetails []*Charge17 `xml:"ChrgDtls,omitempty"`

	// Amount of money due to a party as compensation for a service.
	CommissionDetails []*Commission10 `xml:"ComssnDtls,omitempty"`

	// Tax applicable to an investment fund order.
	TaxDetails []*Tax16 `xml:"TaxDtls,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters3 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to physical delivery of the securities.
	PhysicalDeliveryDetails *DeliveryParameters3 `xml:"PhysDlvryDtls,omitempty"`

	// Payment processes required to transfer cash from the debtor to the creditor.
	CashSettlementDetails *PaymentTransaction21 `xml:"CshSttlmDtls,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Breakdown of the net amount per type of order.
	StaffClientBreakdown []*InvestmentFundsOrderBreakdown1 `xml:"StffClntBrkdwn,omitempty"`

	// Specifies if advice has been received from an independent financial advisor.
	FinancialAdvice *FinancialAdvice1Code `xml:"FinAdvc,omitempty"`

	// Specifies whether the trade is negotiated.
	NegotiatedTrade *NegotiatedTrade1Code `xml:"NgtdTrad,omitempty"`

	// Information about parties related to the transaction.
	RelatedPartyDetails []*Intermediary8 `xml:"RltdPtyDtls,omitempty"`

	// Part of an investor's subscription amount that is held by the fund in order to pay incentive / performance fees at the end of the fiscal year.
	Equalisation *Equalisation1 `xml:"Equlstn,omitempty"`
}

func (r *RedemptionOrder5) SetOrderReference(value string) {
	r.OrderReference = (*Max35Text)(&value)
}

func (r *RedemptionOrder5) SetClientReference(value string) {
	r.ClientReference = (*Max35Text)(&value)
}

func (r *RedemptionOrder5) SetCancellationReference(value string) {
	r.CancellationReference = (*Max35Text)(&value)
}

func (r *RedemptionOrder5) AddInvestmentAccountDetails() *InvestmentAccount21 {
	r.InvestmentAccountDetails = new(InvestmentAccount21)
	return r.InvestmentAccountDetails
}

func (r *RedemptionOrder5) AddOrderType() *FundOrderType2 {
	newValue := new(FundOrderType2)
	r.OrderType = append(r.OrderType, newValue)
	return newValue
}

func (r *RedemptionOrder5) AddBeneficiaryDetails() *IndividualPerson12 {
	r.BeneficiaryDetails = new(IndividualPerson12)
	return r.BeneficiaryDetails
}

func (r *RedemptionOrder5) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	r.UnitsNumber = new(FinancialInstrumentQuantity1)
	return r.UnitsNumber
}

func (r *RedemptionOrder5) SetGrossAmount(value, currency string) {
	r.GrossAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RedemptionOrder5) SetNetAmount(value, currency string) {
	r.NetAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RedemptionOrder5) SetHoldingsRedemptionRate(value string) {
	r.HoldingsRedemptionRate = (*PercentageRate)(&value)
}

func (r *RedemptionOrder5) SetRounding(value string) {
	r.Rounding = (*RoundingDirection2Code)(&value)
}

func (r *RedemptionOrder5) SetSettlementAmount(value, currency string) {
	r.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionOrder5) SetCashSettlementDate(value string) {
	r.CashSettlementDate = (*ISODate)(&value)
}

func (r *RedemptionOrder5) SetSettlementMethod(value string) {
	r.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (r *RedemptionOrder5) AddForeignExchangeDetails() *ForeignExchangeTerms6 {
	r.ForeignExchangeDetails = new(ForeignExchangeTerms6)
	return r.ForeignExchangeDetails
}

func (r *RedemptionOrder5) SetIncomePreference(value string) {
	r.IncomePreference = (*IncomePreference1Code)(&value)
}

func (r *RedemptionOrder5) SetGroup1Or2Units(value string) {
	r.Group1Or2Units = (*UKTaxGroupUnitCode)(&value)
}

func (r *RedemptionOrder5) AddChargeDetails() *Charge17 {
	newValue := new(Charge17)
	r.ChargeDetails = append(r.ChargeDetails, newValue)
	return newValue
}

func (r *RedemptionOrder5) AddCommissionDetails() *Commission10 {
	newValue := new(Commission10)
	r.CommissionDetails = append(r.CommissionDetails, newValue)
	return newValue
}

func (r *RedemptionOrder5) AddTaxDetails() *Tax16 {
	newValue := new(Tax16)
	r.TaxDetails = append(r.TaxDetails, newValue)
	return newValue
}

func (r *RedemptionOrder5) AddSettlementAndCustodyDetails() *FundSettlementParameters3 {
	r.SettlementAndCustodyDetails = new(FundSettlementParameters3)
	return r.SettlementAndCustodyDetails
}

func (r *RedemptionOrder5) SetPhysicalDeliveryIndicator(value string) {
	r.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionOrder5) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	r.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return r.PhysicalDeliveryDetails
}

func (r *RedemptionOrder5) AddCashSettlementDetails() *PaymentTransaction21 {
	r.CashSettlementDetails = new(PaymentTransaction21)
	return r.CashSettlementDetails
}

func (r *RedemptionOrder5) SetNonStandardSettlementInformation(value string) {
	r.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (r *RedemptionOrder5) AddStaffClientBreakdown() *InvestmentFundsOrderBreakdown1 {
	newValue := new(InvestmentFundsOrderBreakdown1)
	r.StaffClientBreakdown = append(r.StaffClientBreakdown, newValue)
	return newValue
}

func (r *RedemptionOrder5) SetFinancialAdvice(value string) {
	r.FinancialAdvice = (*FinancialAdvice1Code)(&value)
}

func (r *RedemptionOrder5) SetNegotiatedTrade(value string) {
	r.NegotiatedTrade = (*NegotiatedTrade1Code)(&value)
}

func (r *RedemptionOrder5) AddRelatedPartyDetails() *Intermediary8 {
	newValue := new(Intermediary8)
	r.RelatedPartyDetails = append(r.RelatedPartyDetails, newValue)
	return newValue
}

func (r *RedemptionOrder5) AddEqualisation() *Equalisation1 {
	r.Equalisation = new(Equalisation1)
	return r.Equalisation
}

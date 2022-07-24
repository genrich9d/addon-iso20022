package iso20022

// Execution of a redemption order.
type RedemptionExecution5 struct {

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of an order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for an order execution, as assigned by a confirming party.
	DealReference *Max35Text `xml:"DealRef"`

	// Specifies the category of the investment fund order.
	OrderType []*FundOrderType3 `xml:"OrdrTp,omitempty"`

	// Additional information about the investor.
	BeneficiaryDetails *IndividualPerson12 `xml:"BnfcryDtls,omitempty"`

	// Number of investment funds units redeemed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Amount of money paid to the investor as a result of the redemption after deduction of charges, commissions and taxes.
	// [(Quantity * Price) - (Charges + Commissions +Taxes)]
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
	InvestmentAccountDetails *InvestmentAccount21 `xml:"InvstmtAcctDtls"`

	// Portion of the investor's holdings, in a specific investment fund/ fund class, that is redeemed.
	HoldingsRedemptionRate *PercentageRate `xml:"HldgsRedRate,omitempty"`

	// Amount of money resulting from the redemption before deduction of  charges, commissions and taxes.
	// [Quantity * Price]
	GrossAmount *ActiveCurrencyAndAmount `xml:"GrssAmt,omitempty"`

	// Date and time at which a price is applied, according to the terms stated in the prospectus.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Method by which the transaction is settled.
	SettlementMethod *DeliveryReceiptType2Code `xml:"SttlmMtd,omitempty"`

	// Price at which the order was executed.
	DealingPriceDetails *UnitPrice10 `xml:"DealgPricDtls"`

	// Other quoted price than the one at which the order was executed.
	InformativePriceDetails []*UnitPrice10 `xml:"InftvPricDtls,omitempty"`

	// Indicates whether the order has been partially executed, ie, the confirmed quantity does not match the ordered quantity for a given financial instrument.
	PartiallyExecutedIndicator *YesNoIndicator `xml:"PrtlyExctdInd"`

	// Specifies that the execution was subject to best execution rules as defined by MiFID.
	BestExecution *BestExecution1Code `xml:"BestExctn,omitempty"`

	// Indicates whether the dividend is included, ie, cum-dividend, in the executed price. When the dividend is not included, the price will be ex-dividend.
	CumDividendIndicator *YesNoIndicator `xml:"CumDvddInd"`

	// Part of the price deemed as accrued income or profit rather than capital. The interim profit amount is used for tax purposes.
	InterimProfitAmount *ProfitAndLoss1Choice `xml:"IntrmPrftAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails []*ForeignExchangeTerms7 `xml:"FXDtls,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Tax group to which the purchased investment fund units belong. The investor indicates to the intermediary operating pooled nominees, which type of unit is to be sold.
	Group1Or2Units *UKTaxGroupUnitCode `xml:"Grp1Or2Units,omitempty"`

	// Amount of money associated with a service.
	ChargeGeneralDetails *TotalCharges3 `xml:"ChrgGnlDtls,omitempty"`

	// Amount of money due to a party as compensation for a service.
	CommissionGeneralDetails *TotalCommissions3 `xml:"ComssnGnlDtls,omitempty"`

	// Breakdown of the net amount per type of order.
	StaffClientBreakdown []*InvestmentFundsOrderBreakdown1 `xml:"StffClntBrkdwn,omitempty"`

	// Tax related to an investment fund order.
	TaxGeneralDetails *TotalTaxes3 `xml:"TaxGnlDtls,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters4 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Parameters of a physical delivery.
	PhysicalDeliveryDetails *DeliveryParameters3 `xml:"PhysDlvryDtls,omitempty"`

	// Payment processes required to transfer cash from the debtor to the creditor.
	CashSettlementDetails *PaymentTransaction22 `xml:"CshSttlmDtls,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Percentage of units partially settled.
	PartialSettlementOfUnits *PercentageRate `xml:"PrtlSttlmOfUnits,omitempty"`

	// Percentage of cash partially settled.
	PartialSettlementOfCash *PercentageRate `xml:"PrtlSttlmOfCsh,omitempty"`

	// Amount retained by the Fund and paid out later at a time decided by the Fund.
	PartialRedemptionWithholdingAmount *CurrencyAndAmount `xml:"PrtlRedWhldgAmt,omitempty"`

	// Specifies if advice has been received from an independent financial advisor.
	FinancialAdvice *FinancialAdvice1Code `xml:"FinAdvc,omitempty"`

	// Specifies whether the trade is negotiated.
	NegotiatedTrade *NegotiatedTrade1Code `xml:"NgtdTrad,omitempty"`

	// Specifies whether the order execution confirmation is late.
	LateReport *LateReport1Code `xml:"LateRpt,omitempty"`

	// Information about parties related to the transaction.
	RelatedPartyDetails []*Intermediary9 `xml:"RltdPtyDtls,omitempty"`

	// Part of an investor's subscription amount that is held by the fund in order to pay incentive / performance fees at the end of the fiscal year.
	Equalisation *Equalisation1 `xml:"Equlstn,omitempty"`
}

func (r *RedemptionExecution5) SetOrderReference(value string) {
	r.OrderReference = (*Max35Text)(&value)
}

func (r *RedemptionExecution5) SetClientReference(value string) {
	r.ClientReference = (*Max35Text)(&value)
}

func (r *RedemptionExecution5) SetDealReference(value string) {
	r.DealReference = (*Max35Text)(&value)
}

func (r *RedemptionExecution5) AddOrderType() *FundOrderType3 {
	newValue := new(FundOrderType3)
	r.OrderType = append(r.OrderType, newValue)
	return newValue
}

func (r *RedemptionExecution5) AddBeneficiaryDetails() *IndividualPerson12 {
	r.BeneficiaryDetails = new(IndividualPerson12)
	return r.BeneficiaryDetails
}

func (r *RedemptionExecution5) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	r.UnitsNumber = new(FinancialInstrumentQuantity1)
	return r.UnitsNumber
}

func (r *RedemptionExecution5) SetRounding(value string) {
	r.Rounding = (*RoundingDirection2Code)(&value)
}

func (r *RedemptionExecution5) SetNetAmount(value, currency string) {
	r.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionExecution5) AddInvestmentAccountDetails() *InvestmentAccount21 {
	r.InvestmentAccountDetails = new(InvestmentAccount21)
	return r.InvestmentAccountDetails
}

func (r *RedemptionExecution5) SetHoldingsRedemptionRate(value string) {
	r.HoldingsRedemptionRate = (*PercentageRate)(&value)
}

func (r *RedemptionExecution5) SetGrossAmount(value, currency string) {
	r.GrossAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionExecution5) AddTradeDateTime() *DateAndDateTimeChoice {
	r.TradeDateTime = new(DateAndDateTimeChoice)
	return r.TradeDateTime
}

func (r *RedemptionExecution5) SetSettlementAmount(value, currency string) {
	r.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionExecution5) SetCashSettlementDate(value string) {
	r.CashSettlementDate = (*ISODate)(&value)
}

func (r *RedemptionExecution5) SetSettlementMethod(value string) {
	r.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (r *RedemptionExecution5) AddDealingPriceDetails() *UnitPrice10 {
	r.DealingPriceDetails = new(UnitPrice10)
	return r.DealingPriceDetails
}

func (r *RedemptionExecution5) AddInformativePriceDetails() *UnitPrice10 {
	newValue := new(UnitPrice10)
	r.InformativePriceDetails = append(r.InformativePriceDetails, newValue)
	return newValue
}

func (r *RedemptionExecution5) SetPartiallyExecutedIndicator(value string) {
	r.PartiallyExecutedIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionExecution5) SetBestExecution(value string) {
	r.BestExecution = (*BestExecution1Code)(&value)
}

func (r *RedemptionExecution5) SetCumDividendIndicator(value string) {
	r.CumDividendIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionExecution5) AddInterimProfitAmount() *ProfitAndLoss1Choice {
	r.InterimProfitAmount = new(ProfitAndLoss1Choice)
	return r.InterimProfitAmount
}

func (r *RedemptionExecution5) AddForeignExchangeDetails() *ForeignExchangeTerms7 {
	newValue := new(ForeignExchangeTerms7)
	r.ForeignExchangeDetails = append(r.ForeignExchangeDetails, newValue)
	return newValue
}

func (r *RedemptionExecution5) SetIncomePreference(value string) {
	r.IncomePreference = (*IncomePreference1Code)(&value)
}

func (r *RedemptionExecution5) SetGroup1Or2Units(value string) {
	r.Group1Or2Units = (*UKTaxGroupUnitCode)(&value)
}

func (r *RedemptionExecution5) AddChargeGeneralDetails() *TotalCharges3 {
	r.ChargeGeneralDetails = new(TotalCharges3)
	return r.ChargeGeneralDetails
}

func (r *RedemptionExecution5) AddCommissionGeneralDetails() *TotalCommissions3 {
	r.CommissionGeneralDetails = new(TotalCommissions3)
	return r.CommissionGeneralDetails
}

func (r *RedemptionExecution5) AddStaffClientBreakdown() *InvestmentFundsOrderBreakdown1 {
	newValue := new(InvestmentFundsOrderBreakdown1)
	r.StaffClientBreakdown = append(r.StaffClientBreakdown, newValue)
	return newValue
}

func (r *RedemptionExecution5) AddTaxGeneralDetails() *TotalTaxes3 {
	r.TaxGeneralDetails = new(TotalTaxes3)
	return r.TaxGeneralDetails
}

func (r *RedemptionExecution5) AddSettlementAndCustodyDetails() *FundSettlementParameters4 {
	r.SettlementAndCustodyDetails = new(FundSettlementParameters4)
	return r.SettlementAndCustodyDetails
}

func (r *RedemptionExecution5) SetPhysicalDeliveryIndicator(value string) {
	r.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionExecution5) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	r.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return r.PhysicalDeliveryDetails
}

func (r *RedemptionExecution5) AddCashSettlementDetails() *PaymentTransaction22 {
	r.CashSettlementDetails = new(PaymentTransaction22)
	return r.CashSettlementDetails
}

func (r *RedemptionExecution5) SetNonStandardSettlementInformation(value string) {
	r.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (r *RedemptionExecution5) SetPartialSettlementOfUnits(value string) {
	r.PartialSettlementOfUnits = (*PercentageRate)(&value)
}

func (r *RedemptionExecution5) SetPartialSettlementOfCash(value string) {
	r.PartialSettlementOfCash = (*PercentageRate)(&value)
}

func (r *RedemptionExecution5) SetPartialRedemptionWithholdingAmount(value, currency string) {
	r.PartialRedemptionWithholdingAmount = NewCurrencyAndAmount(value, currency)
}

func (r *RedemptionExecution5) SetFinancialAdvice(value string) {
	r.FinancialAdvice = (*FinancialAdvice1Code)(&value)
}

func (r *RedemptionExecution5) SetNegotiatedTrade(value string) {
	r.NegotiatedTrade = (*NegotiatedTrade1Code)(&value)
}

func (r *RedemptionExecution5) SetLateReport(value string) {
	r.LateReport = (*LateReport1Code)(&value)
}

func (r *RedemptionExecution5) AddRelatedPartyDetails() *Intermediary9 {
	newValue := new(Intermediary9)
	r.RelatedPartyDetails = append(r.RelatedPartyDetails, newValue)
	return newValue
}

func (r *RedemptionExecution5) AddEqualisation() *Equalisation1 {
	r.Equalisation = new(Equalisation1)
	return r.Equalisation
}

package iso20022

// Execution of a redemption order.
type RedemptionExecution16 struct {

	// Unique and unambiguous identifier for the order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of the order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for the order execution, as assigned by the confirming party.
	DealReference *Max35Text `xml:"DealRef"`

	// Category of the investment fund order.
	OrderType []*FundOrderType4Choice `xml:"OrdrTp,omitempty"`

	// Additional information about the investor.
	BeneficiaryDetails []*IndividualPerson32 `xml:"BnfcryDtls,omitempty"`

	// Number of investment funds units redeemed.
	UnitsNumber *DecimalNumber `xml:"UnitsNb"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Amount of money paid to the investor when redeeming fund units.
	// Net amount = (Quantity * Price) – (Fees + Taxes).
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Account impacted by the investment fund order execution.
	InvestmentAccountDetails *InvestmentAccount58 `xml:"InvstmtAcctDtls"`

	// Portion of the investor's holdings redeemed.
	HoldingsRedemptionRate *PercentageRate `xml:"HldgsRedRate,omitempty"`

	// Amount of money redeemed from the fund.
	// Gross Amount = Quantity * Price.
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
	DealingPriceDetails *UnitPrice22 `xml:"DealgPricDtls"`

	// Other quoted price than the one at which the order was executed.
	InformativePriceDetails []*UnitPrice22 `xml:"InftvPricDtls,omitempty"`

	// Indicates whether the order has been partially executed, that is, the confirmed quantity does not match the ordered quantity for a given financial instrument.
	PartiallyExecutedIndicator *YesNoIndicator `xml:"PrtlyExctdInd"`

	// Specifies that the execution was subject to best execution rules as defined by MiFID.
	BestExecution *BestExecution1Code `xml:"BestExctn,omitempty"`

	// Indicates whether the dividend is included, that is, cum-dividend, in the executed price. When the dividend is not included, the price will be ex-dividend.
	CumDividendIndicator *YesNoIndicator `xml:"CumDvddInd"`

	// Part of the price deemed as accrued income or profit rather than capital. The interim profit amount is used for tax purposes.
	InterimProfitAmount *ProfitAndLoss2Choice `xml:"IntrmPrftAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	// How the exchange rate is expressed determines which currency is the Unit Currency and Quoted Currency. If the amounts concerned are EUR 1000 and USD 1300, the exchange rate may be expressed as per either of the following examples:
	// EXAMPLE 1
	// UnitCurrency  EUR
	// QuotedCurrency  USD
	// ExchangeRate  1.300
	// EXAMPLE 2
	// UnitCurrency  USD
	// QuotedCurrency  EUR
	// ExchangeRate  0.769
	ForeignExchangeDetails []*ForeignExchangeTerms33 `xml:"FXDtls,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Tax group to which the purchased investment fund units belong. The investor indicates to the intermediary operating pooled nominees, which type of unit is to be sold.
	Group1Or2Units *UKTaxGroupUnit1Code `xml:"Grp1Or2Units,omitempty"`

	// Fees (charges/commission) and taxes that are taken into consideration for the transaction, so that the total difference between the net amount and gross amount is known, without taking into account equalisation.
	TransactionOverhead *TotalFeesAndTaxes40 `xml:"TxOvrhd,omitempty"`

	// Additional information about tax that does not have an impact on the transaction overhead.
	InformativeTaxDetails *InformativeTax1 `xml:"InftvTaxDtls,omitempty"`

	// Breakdown of the net amount per type of order.
	StaffClientBreakdown []*InvestmentFundsOrderBreakdown2 `xml:"StffClntBrkdwn,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters11 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to the physical delivery of the securities.
	PhysicalDeliveryDetails *DeliveryParameters3 `xml:"PhysDlvryDtls,omitempty"`

	// Payment process for the transfer of cash from the debtor to the creditor.
	CashSettlementDetails *PaymentTransaction72 `xml:"CshSttlmDtls,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Percentage of units partially settled.
	PartialSettlementOfUnits *PercentageRate `xml:"PrtlSttlmOfUnits,omitempty"`

	// Percentage of cash partially settled.
	PartialSettlementOfCash *PercentageRate `xml:"PrtlSttlmOfCsh,omitempty"`

	// Amount retained by the fund and paid out later at a time decided by the fund.
	PartialRedemptionWithholdingAmount *ActiveCurrencyAndAmount `xml:"PrtlRedWhldgAmt,omitempty"`

	// Specifies if advice has been received from an independent financial advisor.
	FinancialAdvice *FinancialAdvice1Code `xml:"FinAdvc,omitempty"`

	// Specifies whether the trade is negotiated.
	NegotiatedTrade *NegotiatedTrade1Code `xml:"NgtdTrad,omitempty"`

	// Specifies whether the order execution confirmation is late.
	LateReport *LateReport1Code `xml:"LateRpt,omitempty"`

	// Party related to the transaction.
	RelatedPartyDetails []*Intermediary39 `xml:"RltdPtyDtls,omitempty"`

	// Part of an investor's subscription amount that was held by the fund in order to pay incentive/performance fees at the end of the fiscal year, and is returned due to the redemption.
	Equalisation *Equalisation1 `xml:"Equlstn,omitempty"`

	// Assessment of the customer’s behaviour at the time of the account opening application.
	CustomerConductClassification *CustomerConductClassification1Choice `xml:"CstmrCndctClssfctn,omitempty"`

	// Means by which the investor or account owner submits the open account form.
	TransactionChannelType *TransactionChannelType1Choice `xml:"TxChanlTp,omitempty"`

	// Type of signature.
	SignatureType *SignatureType1Choice `xml:"SgntrTp,omitempty"`

	// Information about a non-standard order.
	OrderWaiverDetails *OrderWaiver1 `xml:"OrdrWvrDtls,omitempty"`

	// Information about gating and hold back of redemption proceeds.
	GatingOrHoldBackDetails *HoldBackInformation2 `xml:"GtgOrHldBckDtls,omitempty"`
}

func (r *RedemptionExecution16) SetOrderReference(value string) {
	r.OrderReference = (*Max35Text)(&value)
}

func (r *RedemptionExecution16) SetClientReference(value string) {
	r.ClientReference = (*Max35Text)(&value)
}

func (r *RedemptionExecution16) SetDealReference(value string) {
	r.DealReference = (*Max35Text)(&value)
}

func (r *RedemptionExecution16) AddOrderType() *FundOrderType4Choice {
	newValue := new(FundOrderType4Choice)
	r.OrderType = append(r.OrderType, newValue)
	return newValue
}

func (r *RedemptionExecution16) AddBeneficiaryDetails() *IndividualPerson32 {
	newValue := new(IndividualPerson32)
	r.BeneficiaryDetails = append(r.BeneficiaryDetails, newValue)
	return newValue
}

func (r *RedemptionExecution16) SetUnitsNumber(value string) {
	r.UnitsNumber = (*DecimalNumber)(&value)
}

func (r *RedemptionExecution16) SetRounding(value string) {
	r.Rounding = (*RoundingDirection2Code)(&value)
}

func (r *RedemptionExecution16) SetNetAmount(value, currency string) {
	r.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionExecution16) AddInvestmentAccountDetails() *InvestmentAccount58 {
	r.InvestmentAccountDetails = new(InvestmentAccount58)
	return r.InvestmentAccountDetails
}

func (r *RedemptionExecution16) SetHoldingsRedemptionRate(value string) {
	r.HoldingsRedemptionRate = (*PercentageRate)(&value)
}

func (r *RedemptionExecution16) SetGrossAmount(value, currency string) {
	r.GrossAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionExecution16) AddTradeDateTime() *DateAndDateTimeChoice {
	r.TradeDateTime = new(DateAndDateTimeChoice)
	return r.TradeDateTime
}

func (r *RedemptionExecution16) SetSettlementAmount(value, currency string) {
	r.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionExecution16) SetCashSettlementDate(value string) {
	r.CashSettlementDate = (*ISODate)(&value)
}

func (r *RedemptionExecution16) SetSettlementMethod(value string) {
	r.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (r *RedemptionExecution16) AddDealingPriceDetails() *UnitPrice22 {
	r.DealingPriceDetails = new(UnitPrice22)
	return r.DealingPriceDetails
}

func (r *RedemptionExecution16) AddInformativePriceDetails() *UnitPrice22 {
	newValue := new(UnitPrice22)
	r.InformativePriceDetails = append(r.InformativePriceDetails, newValue)
	return newValue
}

func (r *RedemptionExecution16) SetPartiallyExecutedIndicator(value string) {
	r.PartiallyExecutedIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionExecution16) SetBestExecution(value string) {
	r.BestExecution = (*BestExecution1Code)(&value)
}

func (r *RedemptionExecution16) SetCumDividendIndicator(value string) {
	r.CumDividendIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionExecution16) AddInterimProfitAmount() *ProfitAndLoss2Choice {
	r.InterimProfitAmount = new(ProfitAndLoss2Choice)
	return r.InterimProfitAmount
}

func (r *RedemptionExecution16) AddForeignExchangeDetails() *ForeignExchangeTerms33 {
	newValue := new(ForeignExchangeTerms33)
	r.ForeignExchangeDetails = append(r.ForeignExchangeDetails, newValue)
	return newValue
}

func (r *RedemptionExecution16) SetIncomePreference(value string) {
	r.IncomePreference = (*IncomePreference1Code)(&value)
}

func (r *RedemptionExecution16) SetGroup1Or2Units(value string) {
	r.Group1Or2Units = (*UKTaxGroupUnit1Code)(&value)
}

func (r *RedemptionExecution16) AddTransactionOverhead() *TotalFeesAndTaxes40 {
	r.TransactionOverhead = new(TotalFeesAndTaxes40)
	return r.TransactionOverhead
}

func (r *RedemptionExecution16) AddInformativeTaxDetails() *InformativeTax1 {
	r.InformativeTaxDetails = new(InformativeTax1)
	return r.InformativeTaxDetails
}

func (r *RedemptionExecution16) AddStaffClientBreakdown() *InvestmentFundsOrderBreakdown2 {
	newValue := new(InvestmentFundsOrderBreakdown2)
	r.StaffClientBreakdown = append(r.StaffClientBreakdown, newValue)
	return newValue
}

func (r *RedemptionExecution16) AddSettlementAndCustodyDetails() *FundSettlementParameters11 {
	r.SettlementAndCustodyDetails = new(FundSettlementParameters11)
	return r.SettlementAndCustodyDetails
}

func (r *RedemptionExecution16) SetPhysicalDeliveryIndicator(value string) {
	r.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionExecution16) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	r.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return r.PhysicalDeliveryDetails
}

func (r *RedemptionExecution16) AddCashSettlementDetails() *PaymentTransaction72 {
	r.CashSettlementDetails = new(PaymentTransaction72)
	return r.CashSettlementDetails
}

func (r *RedemptionExecution16) SetNonStandardSettlementInformation(value string) {
	r.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (r *RedemptionExecution16) SetPartialSettlementOfUnits(value string) {
	r.PartialSettlementOfUnits = (*PercentageRate)(&value)
}

func (r *RedemptionExecution16) SetPartialSettlementOfCash(value string) {
	r.PartialSettlementOfCash = (*PercentageRate)(&value)
}

func (r *RedemptionExecution16) SetPartialRedemptionWithholdingAmount(value, currency string) {
	r.PartialRedemptionWithholdingAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionExecution16) SetFinancialAdvice(value string) {
	r.FinancialAdvice = (*FinancialAdvice1Code)(&value)
}

func (r *RedemptionExecution16) SetNegotiatedTrade(value string) {
	r.NegotiatedTrade = (*NegotiatedTrade1Code)(&value)
}

func (r *RedemptionExecution16) SetLateReport(value string) {
	r.LateReport = (*LateReport1Code)(&value)
}

func (r *RedemptionExecution16) AddRelatedPartyDetails() *Intermediary39 {
	newValue := new(Intermediary39)
	r.RelatedPartyDetails = append(r.RelatedPartyDetails, newValue)
	return newValue
}

func (r *RedemptionExecution16) AddEqualisation() *Equalisation1 {
	r.Equalisation = new(Equalisation1)
	return r.Equalisation
}

func (r *RedemptionExecution16) AddCustomerConductClassification() *CustomerConductClassification1Choice {
	r.CustomerConductClassification = new(CustomerConductClassification1Choice)
	return r.CustomerConductClassification
}

func (r *RedemptionExecution16) AddTransactionChannelType() *TransactionChannelType1Choice {
	r.TransactionChannelType = new(TransactionChannelType1Choice)
	return r.TransactionChannelType
}

func (r *RedemptionExecution16) AddSignatureType() *SignatureType1Choice {
	r.SignatureType = new(SignatureType1Choice)
	return r.SignatureType
}

func (r *RedemptionExecution16) AddOrderWaiverDetails() *OrderWaiver1 {
	r.OrderWaiverDetails = new(OrderWaiver1)
	return r.OrderWaiverDetails
}

func (r *RedemptionExecution16) AddGatingOrHoldBackDetails() *HoldBackInformation2 {
	r.GatingOrHoldBackDetails = new(HoldBackInformation2)
	return r.GatingOrHoldBackDetails
}

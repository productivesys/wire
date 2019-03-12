/*
 * WIRE API
 *
 * Moov WIRE () implements an HTTP API for creating, parsing and validating WIRE files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// FEDWireMessage
type FedWireMessage struct {
	SenderSuppliedInfo SenderSuppliedInfo `json:"senderSuppliedInfo"`
	TypeSubType TypeSubType `json:"typeSubType"`
	InputMessageAccountabilityData InputMessageAccountabilityData `json:"inputMessageAccountabilityData"`
	Amount Amount `json:"amount"`
	SenderDepositoryInstitution SenderDepositoryInstitution `json:"senderDepositoryInstitution"`
	ReceiverDepositoryInstitution ReceiverDepositoryInstitution `json:"receiverDepositoryInstitution"`
	BusinessFunctionCode BusinessFunctionCode `json:"businessFunctionCode"`
	// SenderReference
	SenderReference string `json:"senderReference,omitempty"`
	// PreviousMessageIdentifier
	PreviousMessageIdentifier string `json:"previousMessageIdentifier,omitempty"`
	LocalInstrument LocalInstrument `json:"localInstrument,omitempty"`
	Charges Charges `json:"charges,omitempty"`
	InstructedAmount InstructedAmount `json:"instructedAmount,omitempty"`
	// ExchangeRate  Must contain at least one numeric character and only one decimal comma marker (e.g., an exchange rate of 1.2345 should be entered as 1,2345). 
	ExchangeRate string `json:"exchangeRate,omitempty"`
	IntermediaryFinancialInstitution FinancialInstitution `json:"intermediaryFinancialInstitution,omitempty"`
	BeneficiaryFinancialInstitution FinancialInstitution `json:"beneficiaryFinancialInstitution,omitempty"`
	Beneficiary Personal `json:"beneficiary,omitempty"`
	// ReferenceForBeneficiary {4320}
	ReferenceForBeneficiary string `json:"referenceForBeneficiary,omitempty"`
	AccountDebitedDrawdown AccountDebitedDrawdown `json:"accountDebitedDrawdown,omitempty"`
	Originator Personal `json:"originator,omitempty"`
	// OriginatorOptionF {5010}
	OriginatorOptionF map[string]interface{} `json:"originatorOptionF,omitempty"`
	OriginatorFinancialInstitution FinancialInstitution `json:"originatorFinancialInstitution,omitempty"`
	InstructingFinancialInstitution FinancialInstitution `json:"instructingFinancialInstitution,omitempty"`
	AccountCreditedDrawdown AccountCreditedDrawdown `json:"accountCreditedDrawdown,omitempty"`
	OriginatorToBeneficiaryInfo OriginatorToBeneficiaryInfo `json:"originatorToBeneficiaryInfo,omitempty"`
	ReceiverFinancialInstitutionInfo FiToFiInfo `json:"receiverFinancialInstitutionInfo,omitempty"`
	DrawdownDebitAccountAdviceInfo AdviceInfo `json:"drawdownDebitAccountAdviceInfo,omitempty"`
	IntermediaryFinancialInstitutionInfo FiToFiInfo `json:"intermediaryFinancialInstitutionInfo,omitempty"`
	IntermediaryFinacialInstitutionAdviceInfo AdviceInfo `json:"intermediaryFinacialInstitutionAdviceInfo,omitempty"`
	BeneficiaryFinancialInstitutionInfo FiToFiInfo `json:"beneficiaryFinancialInstitutionInfo,omitempty"`
	BeneficiaryFinancialInstitutionAdviceInfo AdviceInfo `json:"beneficiaryFinancialInstitutionAdviceInfo,omitempty"`
	BeneficiaryInfo FiToFiInfo `json:"beneficiaryInfo,omitempty"`
	BeneficiaryAdviceInfo AdviceInfo `json:"beneficiaryAdviceInfo,omitempty"`
	PaymentMethodToBeneficiary PaymentMethodToBeneficiary `json:"paymentMethodToBeneficiary,omitempty"`
	AdditionalFIToFIInfo AdditionalFiToFiInfo `json:"additionalFIToFIInfo,omitempty"`
	CurrencyInstructedAmount CoverPaymentInfo `json:"currencyInstructedAmount,omitempty"`
	OrderingCustomer CoverPaymentInfo `json:"orderingCustomer,omitempty"`
	OrderingInstitution CoverPaymentInfo `json:"orderingInstitution,omitempty"`
	IntermediaryInstitution CoverPaymentInfo `json:"intermediaryInstitution,omitempty"`
	InstitutionAccount CoverPaymentInfo `json:"institutionAccount,omitempty"`
	BeneficiaryCustomer CoverPaymentInfo `json:"beneficiaryCustomer,omitempty"`
	RemittanceInfo CoverPaymentInfo `json:"remittanceInfo,omitempty"`
	SenderToReceiverInfo CoverPaymentInfo `json:"senderToReceiverInfo,omitempty"`
	UnstructuredAddendaInfo UnstructuredAddendaInfo `json:"unstructuredAddendaInfo,omitempty"`
}

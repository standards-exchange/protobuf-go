// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: standards/open/v1/RT0009546.proto

package v1

import (
	date "google.golang.org/genproto/googleapis/type/date"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A proto defined object representing the data parsed from the basic Morningstar holdings template
type RT0009546 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of holdings
	Holdings []*RT0009546_Holding `protobuf:"bytes,1,rep,name=holdings,proto3" json:"holdings,omitempty"`
}

func (x *RT0009546) Reset() {
	*x = RT0009546{}
	if protoimpl.UnsafeEnabled {
		mi := &file_standards_open_v1_RT0009546_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RT0009546) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RT0009546) ProtoMessage() {}

func (x *RT0009546) ProtoReflect() protoreflect.Message {
	mi := &file_standards_open_v1_RT0009546_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RT0009546.ProtoReflect.Descriptor instead.
func (*RT0009546) Descriptor() ([]byte, []int) {
	return file_standards_open_v1_RT0009546_proto_rawDescGZIP(), []int{0}
}

func (x *RT0009546) GetHoldings() []*RT0009546_Holding {
	if x != nil {
		return x.Holdings
	}
	return nil
}

// Defining a single record / row within excel. First holding appears on row 13
type RT0009546_Holding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The month end date on which the portfolio data was extracted. Must be the last day of the month (Column A)
	PortfolioDate *date.Date `protobuf:"bytes,1,opt,name=portfolio_date,json=portfolioDate,proto3" json:"portfolio_date,omitempty"`
	// Unique portfolio identifier. Can be an internal code unique to each fund (Column B)
	PortfolioId string `protobuf:"bytes,2,opt,name=portfolio_id,json=portfolioId,proto3" json:"portfolio_id,omitempty"`
	// Fund name (Column C)
	Fund string `protobuf:"bytes,3,opt,name=fund,proto3" json:"fund,omitempty"`
	// Unique security identifier. Can use any of the following identifiers: ISIN, CUSIP, Sedol, Exchange ticker, Bloomberg ticker, Custom identifier (Column D)
	SecurityId string `protobuf:"bytes,4,opt,name=security_id,json=securityId,proto3" json:"security_id,omitempty"`
	// Security description (Column E)
	SecurityDescription string `protobuf:"bytes,5,opt,name=security_description,json=securityDescription,proto3" json:"security_description,omitempty"`
	// Shares for equity holdings, par value for fixed-income holdings or number of contracts for derivative holdings (Column F)
	SharesParValue float64 `protobuf:"fixed64,6,opt,name=shares_par_value,json=sharesParValue,proto3" json:"shares_par_value,omitempty"`
	// Base market value (Column G)
	BaseMarketValue float64 `protobuf:"fixed64,7,opt,name=base_market_value,json=baseMarketValue,proto3" json:"base_market_value,omitempty"`
	// Coupon rate for fixed income instrument (expressed as a percentage) (Column H)
	CouponRate float64 `protobuf:"fixed64,8,opt,name=coupon_rate,json=couponRate,proto3" json:"coupon_rate,omitempty"`
	// Maturity/expiration date (Column I)
	MaturityDate *date.Date `protobuf:"bytes,9,opt,name=maturity_date,json=maturityDate,proto3" json:"maturity_date,omitempty"`
	// Currency in which all the portfolio holding's market values are reported (ISO 4217 currency code) (Column J)
	PortfolioBaseCurrency string `protobuf:"bytes,10,opt,name=portfolio_base_currency,json=portfolioBaseCurrency,proto3" json:"portfolio_base_currency,omitempty"`
	// Market value expressed in the local currency of the holding (Column K)
	LocalMarketValue float64 `protobuf:"fixed64,11,opt,name=local_market_value,json=localMarketValue,proto3" json:"local_market_value,omitempty"`
	// The currency in which the holding was purchased (ISO 4217 currency code) (Column L)
	LocalCurrency string `protobuf:"bytes,12,opt,name=local_currency,json=localCurrency,proto3" json:"local_currency,omitempty"`
	// The original value of cost of an asset (Column M)
	CostBasis float64 `protobuf:"fixed64,13,opt,name=cost_basis,json=costBasis,proto3" json:"cost_basis,omitempty"`
	// The security's domicile (ISO 3166 country code) (Column N)
	RegionCode string `protobuf:"bytes,14,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
	// Total Net Asset, i.e. fund aum (Column O)
	Tna float64 `protobuf:"fixed64,15,opt,name=tna,proto3" json:"tna,omitempty"`
	// Money not invested in a fund (Column P)
	Cash float64 `protobuf:"fixed64,16,opt,name=cash,proto3" json:"cash,omitempty"`
	// The weight of each position/security held in the portfolio expressed in percentage e.g. 0.56 = 56% (Column Q)
	// Very useful for validation
	Percentage float64 `protobuf:"fixed64,17,opt,name=percentage,proto3" json:"percentage,omitempty"`
	// Coupon or interest accrued but not paid to bondholders (Column R)
	AccruedInterest float64 `protobuf:"fixed64,18,opt,name=accrued_interest,json=accruedInterest,proto3" json:"accrued_interest,omitempty"`
	// Financial classification for each security according to the their attributes (Column S)
	SecurityType string `protobuf:"bytes,19,opt,name=security_type,json=securityType,proto3" json:"security_type,omitempty"`
	// Local asset type code for South Africa (Column T)
	LocalAssetType string `protobuf:"bytes,20,opt,name=local_asset_type,json=localAssetType,proto3" json:"local_asset_type,omitempty"`
	// Credit quality (India) (Column U)
	CreditQuality string `protobuf:"bytes,21,opt,name=credit_quality,json=creditQuality,proto3" json:"credit_quality,omitempty"`
	// Market Identifier Code (MIC): exchange code of which exchange the holdings traded. (ISO 10383) (Column V)
	MarketIdCode string `protobuf:"bytes,22,opt,name=market_id_code,json=marketIdCode,proto3" json:"market_id_code,omitempty"`
}

func (x *RT0009546_Holding) Reset() {
	*x = RT0009546_Holding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_standards_open_v1_RT0009546_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RT0009546_Holding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RT0009546_Holding) ProtoMessage() {}

func (x *RT0009546_Holding) ProtoReflect() protoreflect.Message {
	mi := &file_standards_open_v1_RT0009546_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RT0009546_Holding.ProtoReflect.Descriptor instead.
func (*RT0009546_Holding) Descriptor() ([]byte, []int) {
	return file_standards_open_v1_RT0009546_proto_rawDescGZIP(), []int{0, 0}
}

func (x *RT0009546_Holding) GetPortfolioDate() *date.Date {
	if x != nil {
		return x.PortfolioDate
	}
	return nil
}

func (x *RT0009546_Holding) GetPortfolioId() string {
	if x != nil {
		return x.PortfolioId
	}
	return ""
}

func (x *RT0009546_Holding) GetFund() string {
	if x != nil {
		return x.Fund
	}
	return ""
}

func (x *RT0009546_Holding) GetSecurityId() string {
	if x != nil {
		return x.SecurityId
	}
	return ""
}

func (x *RT0009546_Holding) GetSecurityDescription() string {
	if x != nil {
		return x.SecurityDescription
	}
	return ""
}

func (x *RT0009546_Holding) GetSharesParValue() float64 {
	if x != nil {
		return x.SharesParValue
	}
	return 0
}

func (x *RT0009546_Holding) GetBaseMarketValue() float64 {
	if x != nil {
		return x.BaseMarketValue
	}
	return 0
}

func (x *RT0009546_Holding) GetCouponRate() float64 {
	if x != nil {
		return x.CouponRate
	}
	return 0
}

func (x *RT0009546_Holding) GetMaturityDate() *date.Date {
	if x != nil {
		return x.MaturityDate
	}
	return nil
}

func (x *RT0009546_Holding) GetPortfolioBaseCurrency() string {
	if x != nil {
		return x.PortfolioBaseCurrency
	}
	return ""
}

func (x *RT0009546_Holding) GetLocalMarketValue() float64 {
	if x != nil {
		return x.LocalMarketValue
	}
	return 0
}

func (x *RT0009546_Holding) GetLocalCurrency() string {
	if x != nil {
		return x.LocalCurrency
	}
	return ""
}

func (x *RT0009546_Holding) GetCostBasis() float64 {
	if x != nil {
		return x.CostBasis
	}
	return 0
}

func (x *RT0009546_Holding) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

func (x *RT0009546_Holding) GetTna() float64 {
	if x != nil {
		return x.Tna
	}
	return 0
}

func (x *RT0009546_Holding) GetCash() float64 {
	if x != nil {
		return x.Cash
	}
	return 0
}

func (x *RT0009546_Holding) GetPercentage() float64 {
	if x != nil {
		return x.Percentage
	}
	return 0
}

func (x *RT0009546_Holding) GetAccruedInterest() float64 {
	if x != nil {
		return x.AccruedInterest
	}
	return 0
}

func (x *RT0009546_Holding) GetSecurityType() string {
	if x != nil {
		return x.SecurityType
	}
	return ""
}

func (x *RT0009546_Holding) GetLocalAssetType() string {
	if x != nil {
		return x.LocalAssetType
	}
	return ""
}

func (x *RT0009546_Holding) GetCreditQuality() string {
	if x != nil {
		return x.CreditQuality
	}
	return ""
}

func (x *RT0009546_Holding) GetMarketIdCode() string {
	if x != nil {
		return x.MarketIdCode
	}
	return ""
}

var File_standards_open_v1_RT0009546_proto protoreflect.FileDescriptor

var file_standards_open_v1_RT0009546_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x6f, 0x70, 0x65, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x52, 0x54, 0x30, 0x30, 0x30, 0x39, 0x35, 0x34, 0x36, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7,
	0x07, 0x0a, 0x09, 0x52, 0x54, 0x30, 0x30, 0x30, 0x39, 0x35, 0x34, 0x36, 0x12, 0x40, 0x0a, 0x08,
	0x68, 0x6f, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x54, 0x30, 0x30, 0x30, 0x39, 0x35, 0x34, 0x36, 0x2e, 0x48, 0x6f, 0x6c,
	0x64, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x68, 0x6f, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x1a, 0xd7,
	0x06, 0x0a, 0x07, 0x48, 0x6f, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x38, 0x0a, 0x0e, 0x70, 0x6f,
	0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x0d, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69,
	0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x6f, 0x72, 0x74,
	0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x75, 0x6e, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x75, 0x6e, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x14,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x28, 0x0a, 0x10, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x5f, 0x70, 0x61, 0x72, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x73, 0x50, 0x61, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x62, 0x61, 0x73, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x5f,
	0x72, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x63, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x12, 0x36, 0x0a, 0x0d, 0x6d, 0x61, 0x74, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65,
	0x52, 0x0c, 0x6d, 0x61, 0x74, 0x75, 0x72, 0x69, 0x74, 0x79, 0x44, 0x61, 0x74, 0x65, 0x12, 0x36,
	0x0a, 0x17, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x5f, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x15, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x42, 0x61, 0x73, 0x65, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x2c, 0x0a, 0x12, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x10, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x6f, 0x73, 0x74, 0x5f, 0x62, 0x61, 0x73, 0x69, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x09, 0x63, 0x6f, 0x73, 0x74, 0x42, 0x61, 0x73, 0x69, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74,
	0x6e, 0x61, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x74, 0x6e, 0x61, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x61, 0x73, 0x68, 0x18, 0x10, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x63, 0x61, 0x73,
	0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67,
	0x65, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x63, 0x63, 0x72, 0x75, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x65, 0x73, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x61, 0x63, 0x63,
	0x72, 0x75, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x13, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x28, 0x0a, 0x10, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63,
	0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x51, 0x75, 0x61, 0x6c, 0x69,
	0x74, 0x79, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x49, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x73,
	0x2d, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x73, 0x2f,
	0x6f, 0x70, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_standards_open_v1_RT0009546_proto_rawDescOnce sync.Once
	file_standards_open_v1_RT0009546_proto_rawDescData = file_standards_open_v1_RT0009546_proto_rawDesc
)

func file_standards_open_v1_RT0009546_proto_rawDescGZIP() []byte {
	file_standards_open_v1_RT0009546_proto_rawDescOnce.Do(func() {
		file_standards_open_v1_RT0009546_proto_rawDescData = protoimpl.X.CompressGZIP(file_standards_open_v1_RT0009546_proto_rawDescData)
	})
	return file_standards_open_v1_RT0009546_proto_rawDescData
}

var file_standards_open_v1_RT0009546_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_standards_open_v1_RT0009546_proto_goTypes = []interface{}{
	(*RT0009546)(nil),         // 0: standards.open.v1.RT0009546
	(*RT0009546_Holding)(nil), // 1: standards.open.v1.RT0009546.Holding
	(*date.Date)(nil),         // 2: google.type.Date
}
var file_standards_open_v1_RT0009546_proto_depIdxs = []int32{
	1, // 0: standards.open.v1.RT0009546.holdings:type_name -> standards.open.v1.RT0009546.Holding
	2, // 1: standards.open.v1.RT0009546.Holding.portfolio_date:type_name -> google.type.Date
	2, // 2: standards.open.v1.RT0009546.Holding.maturity_date:type_name -> google.type.Date
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_standards_open_v1_RT0009546_proto_init() }
func file_standards_open_v1_RT0009546_proto_init() {
	if File_standards_open_v1_RT0009546_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_standards_open_v1_RT0009546_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RT0009546); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_standards_open_v1_RT0009546_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RT0009546_Holding); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_standards_open_v1_RT0009546_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_standards_open_v1_RT0009546_proto_goTypes,
		DependencyIndexes: file_standards_open_v1_RT0009546_proto_depIdxs,
		MessageInfos:      file_standards_open_v1_RT0009546_proto_msgTypes,
	}.Build()
	File_standards_open_v1_RT0009546_proto = out.File
	file_standards_open_v1_RT0009546_proto_rawDesc = nil
	file_standards_open_v1_RT0009546_proto_goTypes = nil
	file_standards_open_v1_RT0009546_proto_depIdxs = nil
}

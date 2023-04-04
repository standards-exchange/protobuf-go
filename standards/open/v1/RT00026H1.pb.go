// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: standards/open/v1/RT00026H1.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	date "google.golang.org/genproto/googleapis/type/date"
	_ "google.golang.org/genproto/googleapis/type/money"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Enumerations of asset categories
type RT00026H1_AssetCategory int32

const (
	// Source not specified by the user.
	RT00026H1_ASSET_CATEGORY_UNSPECIFIED RT00026H1_AssetCategory = 0
	// Asset
	RT00026H1_ASSET RT00026H1_AssetCategory = 1
	// Cash
	RT00026H1_CASH RT00026H1_AssetCategory = 2
	// Cash Equivalent
	RT00026H1_CASH_EQUIVALENT RT00026H1_AssetCategory = 3
	// Equity
	RT00026H1_EQUITY RT00026H1_AssetCategory = 4
	// Fixed Income
	RT00026H1_FIXED_INCOME RT00026H1_AssetCategory = 7
	// Forex
	RT00026H1_FOREIGN_EXCHANGE RT00026H1_AssetCategory = 8
	// Future
	RT00026H1_FUTURE RT00026H1_AssetCategory = 9
	// Other Liabilities
	RT00026H1_LIABILITY RT00026H1_AssetCategory = 10
	// Options
	RT00026H1_OPTION RT00026H1_AssetCategory = 11
)

// Enum value maps for RT00026H1_AssetCategory.
var (
	RT00026H1_AssetCategory_name = map[int32]string{
		0:  "ASSET_CATEGORY_UNSPECIFIED",
		1:  "ASSET",
		2:  "CASH",
		3:  "CASH_EQUIVALENT",
		4:  "EQUITY",
		7:  "FIXED_INCOME",
		8:  "FOREIGN_EXCHANGE",
		9:  "FUTURE",
		10: "LIABILITY",
		11: "OPTION",
	}
	RT00026H1_AssetCategory_value = map[string]int32{
		"ASSET_CATEGORY_UNSPECIFIED": 0,
		"ASSET":                      1,
		"CASH":                       2,
		"CASH_EQUIVALENT":            3,
		"EQUITY":                     4,
		"FIXED_INCOME":               7,
		"FOREIGN_EXCHANGE":           8,
		"FUTURE":                     9,
		"LIABILITY":                  10,
		"OPTION":                     11,
	}
)

func (x RT00026H1_AssetCategory) Enum() *RT00026H1_AssetCategory {
	p := new(RT00026H1_AssetCategory)
	*p = x
	return p
}

func (x RT00026H1_AssetCategory) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RT00026H1_AssetCategory) Descriptor() protoreflect.EnumDescriptor {
	return file_standards_open_v1_RT00026H1_proto_enumTypes[0].Descriptor()
}

func (RT00026H1_AssetCategory) Type() protoreflect.EnumType {
	return &file_standards_open_v1_RT00026H1_proto_enumTypes[0]
}

func (x RT00026H1_AssetCategory) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RT00026H1_AssetCategory.Descriptor instead.
func (RT00026H1_AssetCategory) EnumDescriptor() ([]byte, []int) {
	return file_standards_open_v1_RT00026H1_proto_rawDescGZIP(), []int{0, 0}
}

// Fund Holding
type RT00026H1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The effective date of the holding.
	// The holdings are as of market close on the effective date
	EffectiveDate *date.Date `protobuf:"bytes,1,opt,name=effective_date,json=effectiveDate,proto3" json:"effective_date,omitempty"`
	// The legal identifier of the umbrella
	// Should conform to: https://openfunds.org/OFST005010 (Umbrella)
	// See: https://openfunds.org/knowledge/whitepapers/fundnames/ for more information
	UmbrellaFund string `protobuf:"bytes,2,opt,name=umbrella_fund,json=umbrellaFund,proto3" json:"umbrella_fund,omitempty"`
	// Name of the relevant sub-fund
	// Should conform to: https://openfunds.org/OFST010110 (Legal Fund Name Only)
	// Refers to the Sub-Fund in the case of an umbrella structure
	// or Fund in the case of a standalone structure
	// See: https://openfunds.org/knowledge/whitepapers/fundnames/ for more information
	Fund string `protobuf:"bytes,3,opt,name=fund,proto3" json:"fund,omitempty"`
	// The ISO compliant three character code for the base currency of the fund
	FundCurrency string `protobuf:"bytes,4,opt,name=fund_currency,json=fundCurrency,proto3" json:"fund_currency,omitempty"`
	// The internalID used by the system producing the data
	InternalAssetId string `protobuf:"bytes,5,opt,name=internal_asset_id,json=internalAssetId,proto3" json:"internal_asset_id,omitempty"`
	// ISO 6166 code of ISIN when available. If an ISIN is used, it is recommended that the MIC or Bloomberg Exchange Code should additionally be populated. This combination allows for a fund holding to be uniquely identified.
	ExternalAssetIdIsin string `protobuf:"bytes,101,opt,name=external_asset_id_isin,json=externalAssetIdIsin,proto3" json:"external_asset_id_isin,omitempty"`
	// CUSIP - The Committee on Uniform Securities Identification Procedures number assigned by the CUSIP Service Bureau for U.S. and Canadian companies when available. If a CUSIP is used, it is recommended that the MIC or Bloomberg Exchange Code should additionally be populated. This combination allows for a fund holding to be uniquely identified.
	ExternalAssetIdCusip string `protobuf:"bytes,102,opt,name=external_asset_id_cusip,json=externalAssetIdCusip,proto3" json:"external_asset_id_cusip,omitempty"`
	// SEDOL - Stock Exchange Daily Official List for the London Stock Exchange (when available). If a SEDOL is used, it is recommended that the MIC or Bloomberg Exchange Code should additionally be populated. This combination allows for a fund holding to be uniquely identified.
	ExternalAssetIdSedol string `protobuf:"bytes,103,opt,name=external_asset_id_sedol,json=externalAssetIdSedol,proto3" json:"external_asset_id_sedol,omitempty"`
	// FIGI - Financial Instrument Global Identifier (when available)
	ExternalAssetIdFigi string `protobuf:"bytes,104,opt,name=external_asset_id_figi,json=externalAssetIdFigi,proto3" json:"external_asset_id_figi,omitempty"`
	// PermID - Refinitiv Permanent Identifiers (when available)
	ExternalAssetIdPermid string `protobuf:"bytes,105,opt,name=external_asset_id_permid,json=externalAssetIdPermid,proto3" json:"external_asset_id_permid,omitempty"`
	// MIC - Market identifier code as defined by ISO 10383 (available at https://www.iso20022.org/market-identifier-codes) (when available)
	ExternalExchangeIdMic string `protobuf:"bytes,201,opt,name=external_exchange_id_mic,json=externalExchangeIdMic,proto3" json:"external_exchange_id_mic,omitempty"`
	// Bloomberg Exchange Code - Two digit market identifier code as defined by Bloomberg (when available)
	ExternalExchangeIdBbExchangeCode string `protobuf:"bytes,202,opt,name=external_exchange_id_bb_exchange_code,json=externalExchangeIdBbExchangeCode,proto3" json:"external_exchange_id_bb_exchange_code,omitempty"`
	// The asset category of the holding
	AssetCategory RT00026H1_AssetCategory `protobuf:"varint,6,opt,name=asset_category,json=assetCategory,proto3,enum=standards.open.v1.RT00026H1_AssetCategory" json:"asset_category,omitempty"`
	// The ISO compliant three character local currency code of the asset
	LocalCurrencyCode string `protobuf:"bytes,7,opt,name=local_currency_code,json=localCurrencyCode,proto3" json:"local_currency_code,omitempty"`
	// The quantity of the asset held
	Quantity float64 `protobuf:"fixed64,8,opt,name=quantity,proto3" json:"quantity,omitempty"`
	// price in local currency units
	PriceLcu float64 `protobuf:"fixed64,9,opt,name=price_lcu,json=priceLcu,proto3" json:"price_lcu,omitempty"`
	// price in fund currency units
	PriceBase float64 `protobuf:"fixed64,10,opt,name=price_base,json=priceBase,proto3" json:"price_base,omitempty"`
	// the total market value in fund currency units
	MarketValueBase float64 `protobuf:"fixed64,11,opt,name=market_value_base,json=marketValueBase,proto3" json:"market_value_base,omitempty"`
	// the total net asset value of the fund on the effective date in base currency units
	FundNav float64 `protobuf:"fixed64,12,opt,name=fund_nav,json=fundNav,proto3" json:"fund_nav,omitempty"`
	// the percentage weight of total net asset value
	PercentNav float64 `protobuf:"fixed64,13,opt,name=percent_nav,json=percentNav,proto3" json:"percent_nav,omitempty"`
	// The date as of which the holdings record was produced.
	AsOfDate *date.Date `protobuf:"bytes,301,opt,name=as_of_date,json=asOfDate,proto3" json:"as_of_date,omitempty"`
}

func (x *RT00026H1) Reset() {
	*x = RT00026H1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_standards_open_v1_RT00026H1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RT00026H1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RT00026H1) ProtoMessage() {}

func (x *RT00026H1) ProtoReflect() protoreflect.Message {
	mi := &file_standards_open_v1_RT00026H1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RT00026H1.ProtoReflect.Descriptor instead.
func (*RT00026H1) Descriptor() ([]byte, []int) {
	return file_standards_open_v1_RT00026H1_proto_rawDescGZIP(), []int{0}
}

func (x *RT00026H1) GetEffectiveDate() *date.Date {
	if x != nil {
		return x.EffectiveDate
	}
	return nil
}

func (x *RT00026H1) GetUmbrellaFund() string {
	if x != nil {
		return x.UmbrellaFund
	}
	return ""
}

func (x *RT00026H1) GetFund() string {
	if x != nil {
		return x.Fund
	}
	return ""
}

func (x *RT00026H1) GetFundCurrency() string {
	if x != nil {
		return x.FundCurrency
	}
	return ""
}

func (x *RT00026H1) GetInternalAssetId() string {
	if x != nil {
		return x.InternalAssetId
	}
	return ""
}

func (x *RT00026H1) GetExternalAssetIdIsin() string {
	if x != nil {
		return x.ExternalAssetIdIsin
	}
	return ""
}

func (x *RT00026H1) GetExternalAssetIdCusip() string {
	if x != nil {
		return x.ExternalAssetIdCusip
	}
	return ""
}

func (x *RT00026H1) GetExternalAssetIdSedol() string {
	if x != nil {
		return x.ExternalAssetIdSedol
	}
	return ""
}

func (x *RT00026H1) GetExternalAssetIdFigi() string {
	if x != nil {
		return x.ExternalAssetIdFigi
	}
	return ""
}

func (x *RT00026H1) GetExternalAssetIdPermid() string {
	if x != nil {
		return x.ExternalAssetIdPermid
	}
	return ""
}

func (x *RT00026H1) GetExternalExchangeIdMic() string {
	if x != nil {
		return x.ExternalExchangeIdMic
	}
	return ""
}

func (x *RT00026H1) GetExternalExchangeIdBbExchangeCode() string {
	if x != nil {
		return x.ExternalExchangeIdBbExchangeCode
	}
	return ""
}

func (x *RT00026H1) GetAssetCategory() RT00026H1_AssetCategory {
	if x != nil {
		return x.AssetCategory
	}
	return RT00026H1_ASSET_CATEGORY_UNSPECIFIED
}

func (x *RT00026H1) GetLocalCurrencyCode() string {
	if x != nil {
		return x.LocalCurrencyCode
	}
	return ""
}

func (x *RT00026H1) GetQuantity() float64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *RT00026H1) GetPriceLcu() float64 {
	if x != nil {
		return x.PriceLcu
	}
	return 0
}

func (x *RT00026H1) GetPriceBase() float64 {
	if x != nil {
		return x.PriceBase
	}
	return 0
}

func (x *RT00026H1) GetMarketValueBase() float64 {
	if x != nil {
		return x.MarketValueBase
	}
	return 0
}

func (x *RT00026H1) GetFundNav() float64 {
	if x != nil {
		return x.FundNav
	}
	return 0
}

func (x *RT00026H1) GetPercentNav() float64 {
	if x != nil {
		return x.PercentNav
	}
	return 0
}

func (x *RT00026H1) GetAsOfDate() *date.Date {
	if x != nil {
		return x.AsOfDate
	}
	return nil
}

// Batch set of RT00026H1 responses
type RT00026H1Batch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response []*RT00026H1 `protobuf:"bytes,1,rep,name=response,proto3" json:"response,omitempty"`
}

func (x *RT00026H1Batch) Reset() {
	*x = RT00026H1Batch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_standards_open_v1_RT00026H1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RT00026H1Batch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RT00026H1Batch) ProtoMessage() {}

func (x *RT00026H1Batch) ProtoReflect() protoreflect.Message {
	mi := &file_standards_open_v1_RT00026H1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RT00026H1Batch.ProtoReflect.Descriptor instead.
func (*RT00026H1Batch) Descriptor() ([]byte, []int) {
	return file_standards_open_v1_RT00026H1_proto_rawDescGZIP(), []int{1}
}

func (x *RT00026H1Batch) GetResponse() []*RT00026H1 {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_standards_open_v1_RT00026H1_proto protoreflect.FileDescriptor

var file_standards_open_v1_RT00026H1_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x6f, 0x70, 0x65, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x52, 0x54, 0x30, 0x30, 0x30, 0x32, 0x36, 0x48, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x09, 0x0a, 0x09, 0x52, 0x54,
	0x30, 0x30, 0x30, 0x32, 0x36, 0x48, 0x31, 0x12, 0x38, 0x0a, 0x0e, 0x65, 0x66, 0x66, 0x65, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61,
	0x74, 0x65, 0x52, 0x0d, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x6d, 0x62, 0x72, 0x65, 0x6c, 0x6c, 0x61, 0x5f, 0x66, 0x75,
	0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x6d, 0x62, 0x72, 0x65, 0x6c,
	0x6c, 0x61, 0x46, 0x75, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x75, 0x6e, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x75, 0x6e, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x75,
	0x6e, 0x64, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x66, 0x75, 0x6e, 0x64, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12,
	0x2a, 0x0a, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x16, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64,
	0x5f, 0x69, 0x73, 0x69, 0x6e, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x49, 0x73, 0x69, 0x6e,
	0x12, 0x35, 0x0a, 0x17, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x5f, 0x69, 0x64, 0x5f, 0x63, 0x75, 0x73, 0x69, 0x70, 0x18, 0x66, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x14, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x49, 0x64, 0x43, 0x75, 0x73, 0x69, 0x70, 0x12, 0x35, 0x0a, 0x17, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x5f, 0x73, 0x65, 0x64,
	0x6f, 0x6c, 0x18, 0x67, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x53, 0x65, 0x64, 0x6f, 0x6c, 0x12, 0x33,
	0x0a, 0x16, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x5f, 0x69, 0x64, 0x5f, 0x66, 0x69, 0x67, 0x69, 0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x46,
	0x69, 0x67, 0x69, 0x12, 0x37, 0x0a, 0x18, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x64, 0x18,
	0x69, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x18,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x5f, 0x69, 0x64, 0x5f, 0x6d, 0x69, 0x63, 0x18, 0xc9, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x15, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x49, 0x64, 0x4d, 0x69, 0x63, 0x12, 0x50, 0x0a, 0x25, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x5f, 0x62,
	0x62, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0xca, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x20, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x64, 0x42, 0x62, 0x45, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x51, 0x0a, 0x0e, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x2a, 0x2e, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x54, 0x30, 0x30, 0x30, 0x32, 0x36, 0x48, 0x31, 0x2e, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0d, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x2e, 0x0a, 0x13, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x5f, 0x6c, 0x63, 0x75, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x4c, 0x63, 0x75, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x62, 0x61,
	0x73, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x42,
	0x61, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x61, 0x73, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x64, 0x5f, 0x6e, 0x61, 0x76, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x07, 0x66, 0x75, 0x6e, 0x64, 0x4e, 0x61, 0x76, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x76, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x76, 0x12, 0x30, 0x0a, 0x0a, 0x61,
	0x73, 0x5f, 0x6f, 0x66, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0xad, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44,
	0x61, 0x74, 0x65, 0x52, 0x08, 0x61, 0x73, 0x4f, 0x66, 0x44, 0x61, 0x74, 0x65, 0x22, 0xb4, 0x01,
	0x0a, 0x0d, 0x41, 0x73, 0x73, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12,
	0x1e, 0x0a, 0x1a, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52,
	0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x09, 0x0a, 0x05, 0x41, 0x53, 0x53, 0x45, 0x54, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x41,
	0x53, 0x48, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x41, 0x53, 0x48, 0x5f, 0x45, 0x51, 0x55,
	0x49, 0x56, 0x41, 0x4c, 0x45, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x51, 0x55,
	0x49, 0x54, 0x59, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x46, 0x49, 0x58, 0x45, 0x44, 0x5f, 0x49,
	0x4e, 0x43, 0x4f, 0x4d, 0x45, 0x10, 0x07, 0x12, 0x14, 0x0a, 0x10, 0x46, 0x4f, 0x52, 0x45, 0x49,
	0x47, 0x4e, 0x5f, 0x45, 0x58, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x08, 0x12, 0x0a, 0x0a,
	0x06, 0x46, 0x55, 0x54, 0x55, 0x52, 0x45, 0x10, 0x09, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x49, 0x41,
	0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x0a, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x50, 0x54, 0x49,
	0x4f, 0x4e, 0x10, 0x0b, 0x22, 0x4b, 0x0a, 0x0f, 0x52, 0x54, 0x30, 0x30, 0x30, 0x32, 0x36, 0x48,
	0x31, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x12, 0x38, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x74, 0x61, 0x6e,
	0x64, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x54,
	0x30, 0x30, 0x30, 0x32, 0x36, 0x48, 0x31, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x73, 0x2d, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x67, 0x6f, 0x2f, 0x73,
	0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_standards_open_v1_RT00026H1_proto_rawDescOnce sync.Once
	file_standards_open_v1_RT00026H1_proto_rawDescData = file_standards_open_v1_RT00026H1_proto_rawDesc
)

func file_standards_open_v1_RT00026H1_proto_rawDescGZIP() []byte {
	file_standards_open_v1_RT00026H1_proto_rawDescOnce.Do(func() {
		file_standards_open_v1_RT00026H1_proto_rawDescData = protoimpl.X.CompressGZIP(file_standards_open_v1_RT00026H1_proto_rawDescData)
	})
	return file_standards_open_v1_RT00026H1_proto_rawDescData
}

var file_standards_open_v1_RT00026H1_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_standards_open_v1_RT00026H1_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_standards_open_v1_RT00026H1_proto_goTypes = []interface{}{
	(RT00026H1_AssetCategory)(0), // 0: standards.open.v1.RT00026H1.AssetCategory
	(*RT00026H1)(nil),            // 1: standards.open.v1.RT00026H1
	(*RT00026H1Batch)(nil),       // 2: standards.open.v1.RT00026H1_batch
	(*date.Date)(nil),            // 3: google.type.Date
}
var file_standards_open_v1_RT00026H1_proto_depIdxs = []int32{
	3, // 0: standards.open.v1.RT00026H1.effective_date:type_name -> google.type.Date
	0, // 1: standards.open.v1.RT00026H1.asset_category:type_name -> standards.open.v1.RT00026H1.AssetCategory
	3, // 2: standards.open.v1.RT00026H1.as_of_date:type_name -> google.type.Date
	1, // 3: standards.open.v1.RT00026H1_batch.response:type_name -> standards.open.v1.RT00026H1
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_standards_open_v1_RT00026H1_proto_init() }
func file_standards_open_v1_RT00026H1_proto_init() {
	if File_standards_open_v1_RT00026H1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_standards_open_v1_RT00026H1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RT00026H1); i {
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
		file_standards_open_v1_RT00026H1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RT00026H1Batch); i {
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
			RawDescriptor: file_standards_open_v1_RT00026H1_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_standards_open_v1_RT00026H1_proto_goTypes,
		DependencyIndexes: file_standards_open_v1_RT00026H1_proto_depIdxs,
		EnumInfos:         file_standards_open_v1_RT00026H1_proto_enumTypes,
		MessageInfos:      file_standards_open_v1_RT00026H1_proto_msgTypes,
	}.Build()
	File_standards_open_v1_RT00026H1_proto = out.File
	file_standards_open_v1_RT00026H1_proto_rawDesc = nil
	file_standards_open_v1_RT00026H1_proto_goTypes = nil
	file_standards_open_v1_RT00026H1_proto_depIdxs = nil
}

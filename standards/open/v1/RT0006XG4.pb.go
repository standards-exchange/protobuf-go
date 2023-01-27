// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: RT0006XG4.proto

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

// Share Class Unit Price
type RT0006XG4 struct {
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
	// The internal umbrella ID used by the system producing the data
	InternalUmbrellaId string `protobuf:"bytes,3,opt,name=internal_umbrella_id,json=internalUmbrellaId,proto3" json:"internal_umbrella_id,omitempty"`
	// Name of the relevant sub-fund
	// Should conform to: https://openfunds.org/OFST010110 (Legal Fund Name Only)
	// Refers to the Sub-Fund in the case of an umbrella structure
	// or Fund in the case of a standalone structure
	// See: https://openfunds.org/knowledge/whitepapers/fundnames/ for more information
	Fund string `protobuf:"bytes,4,opt,name=fund,proto3" json:"fund,omitempty"`
	// The internal fund ID used by the system producing the data
	InternalFundId string `protobuf:"bytes,5,opt,name=internal_fund_id,json=internalFundId,proto3" json:"internal_fund_id,omitempty"`
	// The ISO compliant three character code for the base currency of the fund
	FundCurrency string `protobuf:"bytes,6,opt,name=fund_currency,json=fundCurrency,proto3" json:"fund_currency,omitempty"`
	// Extension that identifies the share class.
	// Should conform to: https://openfunds.org/OFST020050 (Share Class Extension)
	ShareClassExtension string `protobuf:"bytes,7,opt,name=share_class_extension,json=shareClassExtension,proto3" json:"share_class_extension,omitempty"`
	// The internal share class ID used by the system producing the data
	InternalShareClassId string `protobuf:"bytes,8,opt,name=internal_share_class_id,json=internalShareClassId,proto3" json:"internal_share_class_id,omitempty"`
	// ISO 6166 code of ISIN when available. If an ISIN is used, it is recommended that the MIC or Bloomberg Exchange Code should additionally be populated. This combination allows for a fund holding to be uniquely identified.
	ExternalShareClassIdIsin string `protobuf:"bytes,101,opt,name=external_share_class_id_isin,json=externalShareClassIdIsin,proto3" json:"external_share_class_id_isin,omitempty"`
	// CUSIP - The Committee on Uniform Securities Identification Procedures number assigned by the CUSIP Service Bureau for U.S. and Canadian companies when available. If a CUSIP is used, it is recommended that the MIC or Bloomberg Exchange Code should additionally be populated. This combination allows for a fund holding to be uniquely identified.
	ExternalShareClassIdCusip string `protobuf:"bytes,102,opt,name=external_share_class_id_cusip,json=externalShareClassIdCusip,proto3" json:"external_share_class_id_cusip,omitempty"`
	// SEDOL - Stock Exchange Daily Official List for the London Stock Exchange (when available). If a SEDOL is used, it is recommended that the MIC or Bloomberg Exchange Code should additionally be populated. This combination allows for a fund holding to be uniquely identified.
	ExternalShareClassIdSedol string `protobuf:"bytes,103,opt,name=external_share_class_id_sedol,json=externalShareClassIdSedol,proto3" json:"external_share_class_id_sedol,omitempty"`
	// FIGI - Financial_Instrument Global Identifier (when available)
	ExternalShareClassIdFigi string `protobuf:"bytes,104,opt,name=external_share_class_id_figi,json=externalShareClassIdFigi,proto3" json:"external_share_class_id_figi,omitempty"`
	// PermID - Refinitiv Permanent Identifiers (when available)
	ExternalShareClassIdPermid string `protobuf:"bytes,105,opt,name=external_share_class_id_permid,json=externalShareClassIdPermid,proto3" json:"external_share_class_id_permid,omitempty"`
	// MIC - Market identifier code as defined by ISO 10383 (available at https://www.iso20022.org/market-identifier-codes) (when available)
	ExternalExchangeIdMic string `protobuf:"bytes,201,opt,name=external_exchange_id_mic,json=externalExchangeIdMic,proto3" json:"external_exchange_id_mic,omitempty"`
	// Bloomberg Exchange Code - Two digit market identifier code as defined by Bloomberg (when available)
	ExternalExchangeIdBbExchangeCode string `protobuf:"bytes,202,opt,name=external_exchange_id_bb_exchange_code,json=externalExchangeIdBbExchangeCode,proto3" json:"external_exchange_id_bb_exchange_code,omitempty"`
	// The NAV per share on the effective date
	NavPerShare float64 `protobuf:"fixed64,9,opt,name=nav_per_share,json=navPerShare,proto3" json:"nav_per_share,omitempty"`
	// The NAV per share on the prior effective date
	PriorNavPerShare float64 `protobuf:"fixed64,10,opt,name=prior_nav_per_share,json=priorNavPerShare,proto3" json:"prior_nav_per_share,omitempty"`
	// The total net assets attributed to the share class
	TotalNetAssets float64 `protobuf:"fixed64,11,opt,name=total_net_assets,json=totalNetAssets,proto3" json:"total_net_assets,omitempty"`
	// The total shares outstanding per class
	SharesOutstanding float64 `protobuf:"fixed64,12,opt,name=shares_outstanding,json=sharesOutstanding,proto3" json:"shares_outstanding,omitempty"`
}

func (x *RT0006XG4) Reset() {
	*x = RT0006XG4{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RT0006XG4_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RT0006XG4) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RT0006XG4) ProtoMessage() {}

func (x *RT0006XG4) ProtoReflect() protoreflect.Message {
	mi := &file_RT0006XG4_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RT0006XG4.ProtoReflect.Descriptor instead.
func (*RT0006XG4) Descriptor() ([]byte, []int) {
	return file_RT0006XG4_proto_rawDescGZIP(), []int{0}
}

func (x *RT0006XG4) GetEffectiveDate() *date.Date {
	if x != nil {
		return x.EffectiveDate
	}
	return nil
}

func (x *RT0006XG4) GetUmbrellaFund() string {
	if x != nil {
		return x.UmbrellaFund
	}
	return ""
}

func (x *RT0006XG4) GetInternalUmbrellaId() string {
	if x != nil {
		return x.InternalUmbrellaId
	}
	return ""
}

func (x *RT0006XG4) GetFund() string {
	if x != nil {
		return x.Fund
	}
	return ""
}

func (x *RT0006XG4) GetInternalFundId() string {
	if x != nil {
		return x.InternalFundId
	}
	return ""
}

func (x *RT0006XG4) GetFundCurrency() string {
	if x != nil {
		return x.FundCurrency
	}
	return ""
}

func (x *RT0006XG4) GetShareClassExtension() string {
	if x != nil {
		return x.ShareClassExtension
	}
	return ""
}

func (x *RT0006XG4) GetInternalShareClassId() string {
	if x != nil {
		return x.InternalShareClassId
	}
	return ""
}

func (x *RT0006XG4) GetExternalShareClassIdIsin() string {
	if x != nil {
		return x.ExternalShareClassIdIsin
	}
	return ""
}

func (x *RT0006XG4) GetExternalShareClassIdCusip() string {
	if x != nil {
		return x.ExternalShareClassIdCusip
	}
	return ""
}

func (x *RT0006XG4) GetExternalShareClassIdSedol() string {
	if x != nil {
		return x.ExternalShareClassIdSedol
	}
	return ""
}

func (x *RT0006XG4) GetExternalShareClassIdFigi() string {
	if x != nil {
		return x.ExternalShareClassIdFigi
	}
	return ""
}

func (x *RT0006XG4) GetExternalShareClassIdPermid() string {
	if x != nil {
		return x.ExternalShareClassIdPermid
	}
	return ""
}

func (x *RT0006XG4) GetExternalExchangeIdMic() string {
	if x != nil {
		return x.ExternalExchangeIdMic
	}
	return ""
}

func (x *RT0006XG4) GetExternalExchangeIdBbExchangeCode() string {
	if x != nil {
		return x.ExternalExchangeIdBbExchangeCode
	}
	return ""
}

func (x *RT0006XG4) GetNavPerShare() float64 {
	if x != nil {
		return x.NavPerShare
	}
	return 0
}

func (x *RT0006XG4) GetPriorNavPerShare() float64 {
	if x != nil {
		return x.PriorNavPerShare
	}
	return 0
}

func (x *RT0006XG4) GetTotalNetAssets() float64 {
	if x != nil {
		return x.TotalNetAssets
	}
	return 0
}

func (x *RT0006XG4) GetSharesOutstanding() float64 {
	if x != nil {
		return x.SharesOutstanding
	}
	return 0
}

// Batch set of RT0006XG4 responses
type RT0006XG4Batch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response []*RT0006XG4 `protobuf:"bytes,1,rep,name=response,proto3" json:"response,omitempty"`
}

func (x *RT0006XG4Batch) Reset() {
	*x = RT0006XG4Batch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RT0006XG4_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RT0006XG4Batch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RT0006XG4Batch) ProtoMessage() {}

func (x *RT0006XG4Batch) ProtoReflect() protoreflect.Message {
	mi := &file_RT0006XG4_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RT0006XG4Batch.ProtoReflect.Descriptor instead.
func (*RT0006XG4Batch) Descriptor() ([]byte, []int) {
	return file_RT0006XG4_proto_rawDescGZIP(), []int{1}
}

func (x *RT0006XG4Batch) GetResponse() []*RT0006XG4 {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_RT0006XG4_proto protoreflect.FileDescriptor

var file_RT0006XG4_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x52, 0x54, 0x30, 0x30, 0x30, 0x36, 0x58, 0x47, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x2f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xea, 0x07, 0x0a, 0x09, 0x52, 0x54, 0x30, 0x30,
	0x30, 0x36, 0x58, 0x47, 0x34, 0x12, 0x38, 0x0a, 0x0e, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65,
	0x52, 0x0d, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x75, 0x6d, 0x62, 0x72, 0x65, 0x6c, 0x6c, 0x61, 0x5f, 0x66, 0x75, 0x6e, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x6d, 0x62, 0x72, 0x65, 0x6c, 0x6c, 0x61,
	0x46, 0x75, 0x6e, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x5f, 0x75, 0x6d, 0x62, 0x72, 0x65, 0x6c, 0x6c, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x12, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x55, 0x6d, 0x62, 0x72,
	0x65, 0x6c, 0x6c, 0x61, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x75, 0x6e, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x75, 0x6e, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x66, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x46, 0x75,
	0x6e, 0x64, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x75, 0x6e, 0x64, 0x5f, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x75, 0x6e,
	0x64, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x32, 0x0a, 0x15, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x73, 0x68, 0x61, 0x72, 0x65, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a,
	0x17, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x68, 0x61, 0x72, 0x65, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x1c, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x5f,
	0x69, 0x73, 0x69, 0x6e, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x53, 0x68, 0x61, 0x72, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64,
	0x49, 0x73, 0x69, 0x6e, 0x12, 0x40, 0x0a, 0x1d, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x5f,
	0x63, 0x75, 0x73, 0x69, 0x70, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x19, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x68, 0x61, 0x72, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49,
	0x64, 0x43, 0x75, 0x73, 0x69, 0x70, 0x12, 0x40, 0x0a, 0x1d, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69,
	0x64, 0x5f, 0x73, 0x65, 0x64, 0x6f, 0x6c, 0x18, 0x67, 0x20, 0x01, 0x28, 0x09, 0x52, 0x19, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x68, 0x61, 0x72, 0x65, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x49, 0x64, 0x53, 0x65, 0x64, 0x6f, 0x6c, 0x12, 0x3e, 0x0a, 0x1c, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x5f, 0x69, 0x64, 0x5f, 0x66, 0x69, 0x67, 0x69, 0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x68, 0x61, 0x72, 0x65, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x49, 0x64, 0x46, 0x69, 0x67, 0x69, 0x12, 0x42, 0x0a, 0x1e, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x5f, 0x69, 0x64, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x64, 0x18, 0x69, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x1a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x68, 0x61, 0x72, 0x65, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x18,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x5f, 0x69, 0x64, 0x5f, 0x6d, 0x69, 0x63, 0x18, 0xc9, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x15, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x49, 0x64, 0x4d, 0x69, 0x63, 0x12, 0x50, 0x0a, 0x25, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x5f, 0x62,
	0x62, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0xca, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x20, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x64, 0x42, 0x62, 0x45, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x6e, 0x61, 0x76, 0x5f,
	0x70, 0x65, 0x72, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0b, 0x6e, 0x61, 0x76, 0x50, 0x65, 0x72, 0x53, 0x68, 0x61, 0x72, 0x65, 0x12, 0x2d, 0x0a, 0x13,
	0x70, 0x72, 0x69, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x76, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x70, 0x72, 0x69, 0x6f, 0x72,
	0x4e, 0x61, 0x76, 0x50, 0x65, 0x72, 0x53, 0x68, 0x61, 0x72, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6e, 0x65, 0x74, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4e, 0x65, 0x74, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x5f,
	0x6f, 0x75, 0x74, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x11, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x73, 0x74, 0x61, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x22, 0x4b, 0x0a, 0x0f, 0x52, 0x54, 0x30, 0x30, 0x30, 0x36, 0x58, 0x47,
	0x34, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x12, 0x38, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x74, 0x61, 0x6e,
	0x64, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x54,
	0x30, 0x30, 0x30, 0x36, 0x58, 0x47, 0x34, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x6f, 0x70,
	0x65, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RT0006XG4_proto_rawDescOnce sync.Once
	file_RT0006XG4_proto_rawDescData = file_RT0006XG4_proto_rawDesc
)

func file_RT0006XG4_proto_rawDescGZIP() []byte {
	file_RT0006XG4_proto_rawDescOnce.Do(func() {
		file_RT0006XG4_proto_rawDescData = protoimpl.X.CompressGZIP(file_RT0006XG4_proto_rawDescData)
	})
	return file_RT0006XG4_proto_rawDescData
}

var file_RT0006XG4_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_RT0006XG4_proto_goTypes = []interface{}{
	(*RT0006XG4)(nil),      // 0: standards.open.v1.RT0006XG4
	(*RT0006XG4Batch)(nil), // 1: standards.open.v1.RT0006XG4_batch
	(*date.Date)(nil),      // 2: google.type.Date
}
var file_RT0006XG4_proto_depIdxs = []int32{
	2, // 0: standards.open.v1.RT0006XG4.effective_date:type_name -> google.type.Date
	0, // 1: standards.open.v1.RT0006XG4_batch.response:type_name -> standards.open.v1.RT0006XG4
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_RT0006XG4_proto_init() }
func file_RT0006XG4_proto_init() {
	if File_RT0006XG4_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_RT0006XG4_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RT0006XG4); i {
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
		file_RT0006XG4_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RT0006XG4Batch); i {
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
			RawDescriptor: file_RT0006XG4_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RT0006XG4_proto_goTypes,
		DependencyIndexes: file_RT0006XG4_proto_depIdxs,
		MessageInfos:      file_RT0006XG4_proto_msgTypes,
	}.Build()
	File_RT0006XG4_proto = out.File
	file_RT0006XG4_proto_rawDesc = nil
	file_RT0006XG4_proto_goTypes = nil
	file_RT0006XG4_proto_depIdxs = nil
}

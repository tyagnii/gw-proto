// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.0--rc2
// source: exchanger/v1/exchanger.proto

package exchanger

import (
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

// Запрос для получения курса обмена для конкретной валюты
type CurrencyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromCurrency string `protobuf:"bytes,1,opt,name=from_currency,json=fromCurrency,proto3" json:"from_currency,omitempty"`
	ToCurrency   string `protobuf:"bytes,2,opt,name=to_currency,json=toCurrency,proto3" json:"to_currency,omitempty"`
}

func (x *CurrencyRequest) Reset() {
	*x = CurrencyRequest{}
	mi := &file_exchanger_v1_exchanger_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CurrencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyRequest) ProtoMessage() {}

func (x *CurrencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exchanger_v1_exchanger_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyRequest.ProtoReflect.Descriptor instead.
func (*CurrencyRequest) Descriptor() ([]byte, []int) {
	return file_exchanger_v1_exchanger_proto_rawDescGZIP(), []int{0}
}

func (x *CurrencyRequest) GetFromCurrency() string {
	if x != nil {
		return x.FromCurrency
	}
	return ""
}

func (x *CurrencyRequest) GetToCurrency() string {
	if x != nil {
		return x.ToCurrency
	}
	return ""
}

// Ответ с курсом обмена для конкретной валюты
type ExchangeRateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromCurrency string  `protobuf:"bytes,1,opt,name=from_currency,json=fromCurrency,proto3" json:"from_currency,omitempty"`
	ToCurrency   string  `protobuf:"bytes,2,opt,name=to_currency,json=toCurrency,proto3" json:"to_currency,omitempty"`
	Rate         float64 `protobuf:"fixed64,3,opt,name=rate,proto3" json:"rate,omitempty"`
}

func (x *ExchangeRateResponse) Reset() {
	*x = ExchangeRateResponse{}
	mi := &file_exchanger_v1_exchanger_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExchangeRateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeRateResponse) ProtoMessage() {}

func (x *ExchangeRateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exchanger_v1_exchanger_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeRateResponse.ProtoReflect.Descriptor instead.
func (*ExchangeRateResponse) Descriptor() ([]byte, []int) {
	return file_exchanger_v1_exchanger_proto_rawDescGZIP(), []int{1}
}

func (x *ExchangeRateResponse) GetFromCurrency() string {
	if x != nil {
		return x.FromCurrency
	}
	return ""
}

func (x *ExchangeRateResponse) GetToCurrency() string {
	if x != nil {
		return x.ToCurrency
	}
	return ""
}

func (x *ExchangeRateResponse) GetRate() float64 {
	if x != nil {
		return x.Rate
	}
	return 0
}

// Ответ с курсами обмена всех валют
type ExchangeRatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rates map[string]float64 `protobuf:"bytes,1,rep,name=rates,proto3" json:"rates,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"` // ключ: валюта, значение: курс
}

func (x *ExchangeRatesResponse) Reset() {
	*x = ExchangeRatesResponse{}
	mi := &file_exchanger_v1_exchanger_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExchangeRatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeRatesResponse) ProtoMessage() {}

func (x *ExchangeRatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exchanger_v1_exchanger_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeRatesResponse.ProtoReflect.Descriptor instead.
func (*ExchangeRatesResponse) Descriptor() ([]byte, []int) {
	return file_exchanger_v1_exchanger_proto_rawDescGZIP(), []int{2}
}

func (x *ExchangeRatesResponse) GetRates() map[string]float64 {
	if x != nil {
		return x.Rates
	}
	return nil
}

// Пустое сообщение
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_exchanger_v1_exchanger_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_exchanger_v1_exchanger_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_exchanger_v1_exchanger_proto_rawDescGZIP(), []int{3}
}

var File_exchanger_v1_exchanger_proto protoreflect.FileDescriptor

var file_exchanger_v1_exchanger_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x57, 0x0a, 0x0f, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x66,
	0x72, 0x6f, 0x6d, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x6f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x22, 0x70, 0x0a, 0x14, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x72, 0x6f,
	0x6d, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x6f, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x6f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x72,
	0x61, 0x74, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x15, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a,
	0x05, 0x72, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x61,
	0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x72, 0x61, 0x74, 0x65, 0x73, 0x1a,
	0x38, 0x0a, 0x0a, 0x52, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x32, 0xb0, 0x01, 0x0a, 0x0f, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x73, 0x12, 0x0f, 0x2e, 0x65, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e, 0x65, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x1a,
	0x47, 0x65, 0x74, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x46,
	0x6f, 0x72, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x19, 0x2e, 0x65, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x79, 0x61, 0x67, 0x6e, 0x69, 0x69, 0x2f, 0x67, 0x77, 0x2d, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_exchanger_v1_exchanger_proto_rawDescOnce sync.Once
	file_exchanger_v1_exchanger_proto_rawDescData = file_exchanger_v1_exchanger_proto_rawDesc
)

func file_exchanger_v1_exchanger_proto_rawDescGZIP() []byte {
	file_exchanger_v1_exchanger_proto_rawDescOnce.Do(func() {
		file_exchanger_v1_exchanger_proto_rawDescData = protoimpl.X.CompressGZIP(file_exchanger_v1_exchanger_proto_rawDescData)
	})
	return file_exchanger_v1_exchanger_proto_rawDescData
}

var file_exchanger_v1_exchanger_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_exchanger_v1_exchanger_proto_goTypes = []any{
	(*CurrencyRequest)(nil),       // 0: exchange.CurrencyRequest
	(*ExchangeRateResponse)(nil),  // 1: exchange.ExchangeRateResponse
	(*ExchangeRatesResponse)(nil), // 2: exchange.ExchangeRatesResponse
	(*Empty)(nil),                 // 3: exchange.Empty
	nil,                           // 4: exchange.ExchangeRatesResponse.RatesEntry
}
var file_exchanger_v1_exchanger_proto_depIdxs = []int32{
	4, // 0: exchange.ExchangeRatesResponse.rates:type_name -> exchange.ExchangeRatesResponse.RatesEntry
	3, // 1: exchange.ExchangeService.GetExchangeRates:input_type -> exchange.Empty
	0, // 2: exchange.ExchangeService.GetExchangeRateForCurrency:input_type -> exchange.CurrencyRequest
	2, // 3: exchange.ExchangeService.GetExchangeRates:output_type -> exchange.ExchangeRatesResponse
	1, // 4: exchange.ExchangeService.GetExchangeRateForCurrency:output_type -> exchange.ExchangeRateResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_exchanger_v1_exchanger_proto_init() }
func file_exchanger_v1_exchanger_proto_init() {
	if File_exchanger_v1_exchanger_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_exchanger_v1_exchanger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_exchanger_v1_exchanger_proto_goTypes,
		DependencyIndexes: file_exchanger_v1_exchanger_proto_depIdxs,
		MessageInfos:      file_exchanger_v1_exchanger_proto_msgTypes,
	}.Build()
	File_exchanger_v1_exchanger_proto = out.File
	file_exchanger_v1_exchanger_proto_rawDesc = nil
	file_exchanger_v1_exchanger_proto_goTypes = nil
	file_exchanger_v1_exchanger_proto_depIdxs = nil
}

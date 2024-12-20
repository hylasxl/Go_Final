// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: api/proto/airline.proto

package airlinepb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateArlineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AirlineName string `protobuf:"bytes,1,opt,name=airlineName,proto3" json:"airlineName,omitempty"`
	AirlineCode string `protobuf:"bytes,2,opt,name=airlineCode,proto3" json:"airlineCode,omitempty"`
	Country     string `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *CreateArlineRequest) Reset() {
	*x = CreateArlineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_airline_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateArlineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArlineRequest) ProtoMessage() {}

func (x *CreateArlineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_airline_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArlineRequest.ProtoReflect.Descriptor instead.
func (*CreateArlineRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_airline_proto_rawDescGZIP(), []int{0}
}

func (x *CreateArlineRequest) GetAirlineName() string {
	if x != nil {
		return x.AirlineName
	}
	return ""
}

func (x *CreateArlineRequest) GetAirlineCode() string {
	if x != nil {
		return x.AirlineCode
	}
	return ""
}

func (x *CreateArlineRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type UpdateAirlineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AirlineName string `protobuf:"bytes,1,opt,name=airlineName,proto3" json:"airlineName,omitempty"`
	AirlineCode string `protobuf:"bytes,2,opt,name=airlineCode,proto3" json:"airlineCode,omitempty"`
	Country     string `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	AirlineId   string `protobuf:"bytes,4,opt,name=airlineId,proto3" json:"airlineId,omitempty"`
}

func (x *UpdateAirlineRequest) Reset() {
	*x = UpdateAirlineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_airline_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAirlineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAirlineRequest) ProtoMessage() {}

func (x *UpdateAirlineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_airline_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAirlineRequest.ProtoReflect.Descriptor instead.
func (*UpdateAirlineRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_airline_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateAirlineRequest) GetAirlineName() string {
	if x != nil {
		return x.AirlineName
	}
	return ""
}

func (x *UpdateAirlineRequest) GetAirlineCode() string {
	if x != nil {
		return x.AirlineCode
	}
	return ""
}

func (x *UpdateAirlineRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *UpdateAirlineRequest) GetAirlineId() string {
	if x != nil {
		return x.AirlineId
	}
	return ""
}

type GetListAirlineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetListAirlineRequest) Reset() {
	*x = GetListAirlineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_airline_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListAirlineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListAirlineRequest) ProtoMessage() {}

func (x *GetListAirlineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_airline_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListAirlineRequest.ProtoReflect.Descriptor instead.
func (*GetListAirlineRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_airline_proto_rawDescGZIP(), []int{2}
}

type DeleteAirlineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AirlineId string `protobuf:"bytes,1,opt,name=airlineId,proto3" json:"airlineId,omitempty"`
}

func (x *DeleteAirlineRequest) Reset() {
	*x = DeleteAirlineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_airline_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAirlineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAirlineRequest) ProtoMessage() {}

func (x *DeleteAirlineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_airline_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAirlineRequest.ProtoReflect.Descriptor instead.
func (*DeleteAirlineRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_airline_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteAirlineRequest) GetAirlineId() string {
	if x != nil {
		return x.AirlineId
	}
	return ""
}

type GetAirlineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AirlineId string `protobuf:"bytes,1,opt,name=airlineId,proto3" json:"airlineId,omitempty"`
}

func (x *GetAirlineRequest) Reset() {
	*x = GetAirlineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_airline_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAirlineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAirlineRequest) ProtoMessage() {}

func (x *GetAirlineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_airline_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAirlineRequest.ProtoReflect.Descriptor instead.
func (*GetAirlineRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_airline_proto_rawDescGZIP(), []int{4}
}

func (x *GetAirlineRequest) GetAirlineId() string {
	if x != nil {
		return x.AirlineId
	}
	return ""
}

type GetAirlineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Airline *Airline `protobuf:"bytes,1,opt,name=airline,proto3" json:"airline,omitempty"`
}

func (x *GetAirlineResponse) Reset() {
	*x = GetAirlineResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_airline_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAirlineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAirlineResponse) ProtoMessage() {}

func (x *GetAirlineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_airline_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAirlineResponse.ProtoReflect.Descriptor instead.
func (*GetAirlineResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_airline_proto_rawDescGZIP(), []int{5}
}

func (x *GetAirlineResponse) GetAirline() *Airline {
	if x != nil {
		return x.Airline
	}
	return nil
}

type GetListAirlineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Airlines []*Airline `protobuf:"bytes,1,rep,name=airlines,proto3" json:"airlines,omitempty"`
}

func (x *GetListAirlineResponse) Reset() {
	*x = GetListAirlineResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_airline_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListAirlineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListAirlineResponse) ProtoMessage() {}

func (x *GetListAirlineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_airline_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListAirlineResponse.ProtoReflect.Descriptor instead.
func (*GetListAirlineResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_airline_proto_rawDescGZIP(), []int{6}
}

func (x *GetListAirlineResponse) GetAirlines() []*Airline {
	if x != nil {
		return x.Airlines
	}
	return nil
}

type AirlineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *AirlineResponse) Reset() {
	*x = AirlineResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_airline_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AirlineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AirlineResponse) ProtoMessage() {}

func (x *AirlineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_airline_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AirlineResponse.ProtoReflect.Descriptor instead.
func (*AirlineResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_airline_proto_rawDescGZIP(), []int{7}
}

func (x *AirlineResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type Airline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AirlineName string `protobuf:"bytes,1,opt,name=airlineName,proto3" json:"airlineName,omitempty"`
	AirlineCode string `protobuf:"bytes,2,opt,name=airlineCode,proto3" json:"airlineCode,omitempty"`
	Country     string `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	Id          string `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Airline) Reset() {
	*x = Airline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_airline_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Airline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Airline) ProtoMessage() {}

func (x *Airline) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_airline_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Airline.ProtoReflect.Descriptor instead.
func (*Airline) Descriptor() ([]byte, []int) {
	return file_api_proto_airline_proto_rawDescGZIP(), []int{8}
}

func (x *Airline) GetAirlineName() string {
	if x != nil {
		return x.AirlineName
	}
	return ""
}

func (x *Airline) GetAirlineCode() string {
	if x != nil {
		return x.AirlineCode
	}
	return ""
}

func (x *Airline) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Airline) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AirlineGetFlightRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AirlineId string `protobuf:"bytes,1,opt,name=airlineId,proto3" json:"airlineId,omitempty"`
}

func (x *AirlineGetFlightRequest) Reset() {
	*x = AirlineGetFlightRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_airline_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AirlineGetFlightRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AirlineGetFlightRequest) ProtoMessage() {}

func (x *AirlineGetFlightRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_airline_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AirlineGetFlightRequest.ProtoReflect.Descriptor instead.
func (*AirlineGetFlightRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_airline_proto_rawDescGZIP(), []int{9}
}

func (x *AirlineGetFlightRequest) GetAirlineId() string {
	if x != nil {
		return x.AirlineId
	}
	return ""
}

type AirlineGetFlightResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flights []*FlightAL `protobuf:"bytes,1,rep,name=flights,proto3" json:"flights,omitempty"`
	Error   string      `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *AirlineGetFlightResponse) Reset() {
	*x = AirlineGetFlightResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_airline_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AirlineGetFlightResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AirlineGetFlightResponse) ProtoMessage() {}

func (x *AirlineGetFlightResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_airline_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AirlineGetFlightResponse.ProtoReflect.Descriptor instead.
func (*AirlineGetFlightResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_airline_proto_rawDescGZIP(), []int{10}
}

func (x *AirlineGetFlightResponse) GetFlights() []*FlightAL {
	if x != nil {
		return x.Flights
	}
	return nil
}

func (x *AirlineGetFlightResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type FlightAL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlightNumber         string                 `protobuf:"bytes,1,opt,name=FlightNumber,proto3" json:"FlightNumber,omitempty"`
	DepartureAirportID   uint32                 `protobuf:"varint,2,opt,name=DepartureAirportID,proto3" json:"DepartureAirportID,omitempty"`
	ArrivalAirportID     uint32                 `protobuf:"varint,3,opt,name=ArrivalAirportID,proto3" json:"ArrivalAirportID,omitempty"`
	DepartureTime        *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=DepartureTime,proto3" json:"DepartureTime,omitempty"`
	ArrivalTime          *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=ArrivalTime,proto3" json:"ArrivalTime,omitempty"`
	AvailableSeats       uint32                 `protobuf:"varint,6,opt,name=AvailableSeats,proto3" json:"AvailableSeats,omitempty"`
	Duration             uint32                 `protobuf:"varint,7,opt,name=Duration,proto3" json:"Duration,omitempty"`
	Status               string                 `protobuf:"bytes,8,opt,name=Status,proto3" json:"Status,omitempty"`
	UpdatedDepartureTime *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=UpdatedDepartureTime,proto3" json:"UpdatedDepartureTime,omitempty"`
	UpdatedArrivalTime   *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=UpdatedArrivalTime,proto3" json:"UpdatedArrivalTime,omitempty"`
	Reason               string                 `protobuf:"bytes,11,opt,name=Reason,proto3" json:"Reason,omitempty"`
}

func (x *FlightAL) Reset() {
	*x = FlightAL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_airline_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlightAL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlightAL) ProtoMessage() {}

func (x *FlightAL) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_airline_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlightAL.ProtoReflect.Descriptor instead.
func (*FlightAL) Descriptor() ([]byte, []int) {
	return file_api_proto_airline_proto_rawDescGZIP(), []int{11}
}

func (x *FlightAL) GetFlightNumber() string {
	if x != nil {
		return x.FlightNumber
	}
	return ""
}

func (x *FlightAL) GetDepartureAirportID() uint32 {
	if x != nil {
		return x.DepartureAirportID
	}
	return 0
}

func (x *FlightAL) GetArrivalAirportID() uint32 {
	if x != nil {
		return x.ArrivalAirportID
	}
	return 0
}

func (x *FlightAL) GetDepartureTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DepartureTime
	}
	return nil
}

func (x *FlightAL) GetArrivalTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ArrivalTime
	}
	return nil
}

func (x *FlightAL) GetAvailableSeats() uint32 {
	if x != nil {
		return x.AvailableSeats
	}
	return 0
}

func (x *FlightAL) GetDuration() uint32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *FlightAL) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *FlightAL) GetUpdatedDepartureTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedDepartureTime
	}
	return nil
}

func (x *FlightAL) GetUpdatedArrivalTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedArrivalTime
	}
	return nil
}

func (x *FlightAL) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

var File_api_proto_airline_proto protoreflect.FileDescriptor

var file_api_proto_airline_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x69, 0x72, 0x6c,
	0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x73, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x6c, 0x69, 0x6e, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x69, 0x72, 0x6c, 0x69,
	0x6e, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x69,
	0x72, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x69, 0x72,
	0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x92, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x61, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x61, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x34, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x69, 0x72,
	0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61,
	0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x22, 0x31, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x41, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x22, 0x3d, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x41, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x61, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x69, 0x72, 0x6c, 0x69,
	0x6e, 0x65, 0x52, 0x07, 0x61, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x43, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x61, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41,
	0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x08, 0x61, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x73,
	0x22, 0x27, 0x0a, 0x0f, 0x41, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x77, 0x0a, 0x07, 0x41, 0x69, 0x72,
	0x6c, 0x69, 0x6e, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x69, 0x72, 0x6c, 0x69,
	0x6e, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x69, 0x72, 0x6c, 0x69, 0x6e,
	0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x69, 0x72,
	0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x37, 0x0a, 0x17, 0x41, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x47, 0x65, 0x74,
	0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x61, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x22, 0x5a, 0x0a, 0x18, 0x41,
	0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x66, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x41, 0x4c, 0x52, 0x07, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x9a, 0x04, 0x0a, 0x08, 0x46, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x41, 0x4c, 0x12, 0x22, 0x0a, 0x0c, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x46, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x12, 0x44, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x75, 0x72, 0x65, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x75, 0x72, 0x65, 0x41,
	0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x44, 0x12, 0x2a, 0x0a, 0x10, 0x41, 0x72, 0x72, 0x69,
	0x76, 0x61, 0x6c, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x10, 0x41, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x41, 0x69, 0x72, 0x70, 0x6f,
	0x72, 0x74, 0x49, 0x44, 0x12, 0x40, 0x0a, 0x0d, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x75, 0x72,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x75,
	0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x41, 0x72, 0x72, 0x69, 0x76, 0x61,
	0x6c, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x41, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x53, 0x65, 0x61, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x41, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x61, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x4e, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x44, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x75, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x14, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x75, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x4a, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x72, 0x72, 0x69, 0x76,
	0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x32, 0xbc, 0x03, 0x0a, 0x0e, 0x41, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x69, 0x72, 0x6c, 0x69, 0x6e,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1b, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x69, 0x72, 0x6c, 0x69, 0x6e,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x69, 0x72,
	0x6c, 0x69, 0x6e, 0x65, 0x12, 0x17, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x69, 0x72, 0x6c,
	0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0d, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1a, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x69, 0x72, 0x6c, 0x69, 0x6e,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x41, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x52, 0x0a, 0x11, 0x41, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x73, 0x12, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x69, 0x72, 0x6c,
	0x69, 0x6e, 0x65, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x69, 0x72, 0x6c, 0x69,
	0x6e, 0x65, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x61, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_airline_proto_rawDescOnce sync.Once
	file_api_proto_airline_proto_rawDescData = file_api_proto_airline_proto_rawDesc
)

func file_api_proto_airline_proto_rawDescGZIP() []byte {
	file_api_proto_airline_proto_rawDescOnce.Do(func() {
		file_api_proto_airline_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_airline_proto_rawDescData)
	})
	return file_api_proto_airline_proto_rawDescData
}

var file_api_proto_airline_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_api_proto_airline_proto_goTypes = []any{
	(*CreateArlineRequest)(nil),      // 0: auth.CreateArlineRequest
	(*UpdateAirlineRequest)(nil),     // 1: auth.UpdateAirlineRequest
	(*GetListAirlineRequest)(nil),    // 2: auth.GetListAirlineRequest
	(*DeleteAirlineRequest)(nil),     // 3: auth.DeleteAirlineRequest
	(*GetAirlineRequest)(nil),        // 4: auth.GetAirlineRequest
	(*GetAirlineResponse)(nil),       // 5: auth.GetAirlineResponse
	(*GetListAirlineResponse)(nil),   // 6: auth.GetListAirlineResponse
	(*AirlineResponse)(nil),          // 7: auth.AirlineResponse
	(*Airline)(nil),                  // 8: auth.Airline
	(*AirlineGetFlightRequest)(nil),  // 9: auth.AirlineGetFlightRequest
	(*AirlineGetFlightResponse)(nil), // 10: auth.AirlineGetFlightResponse
	(*FlightAL)(nil),                 // 11: auth.FlightAL
	(*timestamppb.Timestamp)(nil),    // 12: google.protobuf.Timestamp
}
var file_api_proto_airline_proto_depIdxs = []int32{
	8,  // 0: auth.GetAirlineResponse.airline:type_name -> auth.Airline
	8,  // 1: auth.GetListAirlineResponse.airlines:type_name -> auth.Airline
	11, // 2: auth.AirlineGetFlightResponse.flights:type_name -> auth.FlightAL
	12, // 3: auth.FlightAL.DepartureTime:type_name -> google.protobuf.Timestamp
	12, // 4: auth.FlightAL.ArrivalTime:type_name -> google.protobuf.Timestamp
	12, // 5: auth.FlightAL.UpdatedDepartureTime:type_name -> google.protobuf.Timestamp
	12, // 6: auth.FlightAL.UpdatedArrivalTime:type_name -> google.protobuf.Timestamp
	0,  // 7: auth.AirlineService.CreateArline:input_type -> auth.CreateArlineRequest
	2,  // 8: auth.AirlineService.GetListAirline:input_type -> auth.GetListAirlineRequest
	4,  // 9: auth.AirlineService.GetAirline:input_type -> auth.GetAirlineRequest
	1,  // 10: auth.AirlineService.UpdateAirline:input_type -> auth.UpdateAirlineRequest
	3,  // 11: auth.AirlineService.DeleteAirline:input_type -> auth.DeleteAirlineRequest
	9,  // 12: auth.AirlineService.AirlineGetFlights:input_type -> auth.AirlineGetFlightRequest
	7,  // 13: auth.AirlineService.CreateArline:output_type -> auth.AirlineResponse
	6,  // 14: auth.AirlineService.GetListAirline:output_type -> auth.GetListAirlineResponse
	5,  // 15: auth.AirlineService.GetAirline:output_type -> auth.GetAirlineResponse
	7,  // 16: auth.AirlineService.UpdateAirline:output_type -> auth.AirlineResponse
	7,  // 17: auth.AirlineService.DeleteAirline:output_type -> auth.AirlineResponse
	10, // 18: auth.AirlineService.AirlineGetFlights:output_type -> auth.AirlineGetFlightResponse
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_api_proto_airline_proto_init() }
func file_api_proto_airline_proto_init() {
	if File_api_proto_airline_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_airline_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateArlineRequest); i {
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
		file_api_proto_airline_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateAirlineRequest); i {
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
		file_api_proto_airline_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetListAirlineRequest); i {
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
		file_api_proto_airline_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteAirlineRequest); i {
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
		file_api_proto_airline_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetAirlineRequest); i {
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
		file_api_proto_airline_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetAirlineResponse); i {
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
		file_api_proto_airline_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetListAirlineResponse); i {
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
		file_api_proto_airline_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*AirlineResponse); i {
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
		file_api_proto_airline_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*Airline); i {
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
		file_api_proto_airline_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*AirlineGetFlightRequest); i {
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
		file_api_proto_airline_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*AirlineGetFlightResponse); i {
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
		file_api_proto_airline_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*FlightAL); i {
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
			RawDescriptor: file_api_proto_airline_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_airline_proto_goTypes,
		DependencyIndexes: file_api_proto_airline_proto_depIdxs,
		MessageInfos:      file_api_proto_airline_proto_msgTypes,
	}.Build()
	File_api_proto_airline_proto = out.File
	file_api_proto_airline_proto_rawDesc = nil
	file_api_proto_airline_proto_goTypes = nil
	file_api_proto_airline_proto_depIdxs = nil
}

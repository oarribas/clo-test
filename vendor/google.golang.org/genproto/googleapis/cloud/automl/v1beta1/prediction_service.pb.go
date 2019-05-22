// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/prediction_service.proto

package automl // import "google.golang.org/genproto/googleapis/cloud/automl/v1beta1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import longrunning "google.golang.org/genproto/googleapis/longrunning"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request message for [PredictionService.Predict][google.cloud.automl.v1beta1.PredictionService.Predict].
type PredictRequest struct {
	// Name of the model requested to serve the prediction.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required.
	// Payload to perform a prediction on. The payload must match the
	// problem type that the model was trained to solve.
	Payload *ExamplePayload `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	// Additional domain-specific parameters, any string must be up to 25000
	// characters long.
	//
	// *  For Image Classification:
	//
	//    `score_threshold` - (float) A value from 0.0 to 1.0. When the model
	//     makes predictions for an image, it will only produce results that have
	//     at least this confidence score. The default is 0.5.
	//
	//  *  For Image Object Detection:
	//    `score_threshold` - (float) When Model detects objects on the image,
	//        it will only produce bounding boxes which have at least this
	//        confidence score. Value in 0 to 1 range, default is 0.5.
	//    `max_bounding_box_count` - (int64) No more than this number of bounding
	//        boxes will be returned in the response. Default is 100, the
	//        requested value may be limited by server.
	Params               map[string]string `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PredictRequest) Reset()         { *m = PredictRequest{} }
func (m *PredictRequest) String() string { return proto.CompactTextString(m) }
func (*PredictRequest) ProtoMessage()    {}
func (*PredictRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_prediction_service_4cd94add91ccba6b, []int{0}
}
func (m *PredictRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictRequest.Unmarshal(m, b)
}
func (m *PredictRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictRequest.Marshal(b, m, deterministic)
}
func (dst *PredictRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictRequest.Merge(dst, src)
}
func (m *PredictRequest) XXX_Size() int {
	return xxx_messageInfo_PredictRequest.Size(m)
}
func (m *PredictRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PredictRequest proto.InternalMessageInfo

func (m *PredictRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PredictRequest) GetPayload() *ExamplePayload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *PredictRequest) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

// Response message for [PredictionService.Predict][google.cloud.automl.v1beta1.PredictionService.Predict].
type PredictResponse struct {
	// Prediction result.
	// Translation and Text Sentiment will return precisely one payload.
	Payload []*AnnotationPayload `protobuf:"bytes,1,rep,name=payload,proto3" json:"payload,omitempty"`
	// Additional domain-specific prediction response metadata.
	//
	// * For Image Object Detection:
	//  `max_bounding_box_count` - (int64) At most that many bounding boxes per
	//      image could have been returned.
	//
	// * For Text Sentiment:
	//  `sentiment_score` - (float, deprecated) A value between -1 and 1,
	//      -1 maps to least positive sentiment, while 1 maps to the most positive
	//      one and the higher the score, the more positive the sentiment in the
	//      document is. Yet these values are relative to the training data, so
	//      e.g. if all data was positive then -1 will be also positive (though
	//      the least).
	//      The sentiment_score shouldn't be confused with "score" or "magnitude"
	//      from the previous Natural Language Sentiment Analysis API.
	Metadata             map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PredictResponse) Reset()         { *m = PredictResponse{} }
func (m *PredictResponse) String() string { return proto.CompactTextString(m) }
func (*PredictResponse) ProtoMessage()    {}
func (*PredictResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_prediction_service_4cd94add91ccba6b, []int{1}
}
func (m *PredictResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictResponse.Unmarshal(m, b)
}
func (m *PredictResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictResponse.Marshal(b, m, deterministic)
}
func (dst *PredictResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictResponse.Merge(dst, src)
}
func (m *PredictResponse) XXX_Size() int {
	return xxx_messageInfo_PredictResponse.Size(m)
}
func (m *PredictResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PredictResponse proto.InternalMessageInfo

func (m *PredictResponse) GetPayload() []*AnnotationPayload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *PredictResponse) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Request message for [PredictionService.BatchPredict][google.cloud.automl.v1beta1.PredictionService.BatchPredict].
type BatchPredictRequest struct {
	// Name of the model requested to serve the batch prediction.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The input configuration for batch prediction.
	InputConfig *BatchPredictInputConfig `protobuf:"bytes,3,opt,name=input_config,json=inputConfig,proto3" json:"input_config,omitempty"`
	// Required. The Configuration specifying where output predictions should
	// be written.
	OutputConfig *BatchPredictOutputConfig `protobuf:"bytes,4,opt,name=output_config,json=outputConfig,proto3" json:"output_config,omitempty"`
	// Additional domain-specific parameters for the predictions, any string must
	// be up to 25000 characters long.
	//
	// *  For Video Classification :
	//    `score_threshold` - (float) A value from 0.0 to 1.0. When the model
	//        makes predictions for a video, it will only produce results that
	//        have at least this confidence score. The default is 0.5.
	//    `segment_classification` - (boolean) Set to true to request
	//        segment-level classification. AutoML Video Intelligence returns
	//        labels and their confidence scores for the entire segment of the
	//        video that user specified in the request configuration.
	//        The default is "true".
	//    `shot_classification` - (boolean) Set to true to request shot-level
	//        classification. AutoML Video Intelligence determines the boundaries
	//        for each camera shot in the entire segment of the video that user
	//        specified in the request configuration. AutoML Video Intelligence
	//        then returns labels and their confidence scores for each detected
	//        shot, along with the start and end time of the shot.
	//        WARNING: Model evaluation is not done for this classification type,
	//        the quality of it depends on training data, but there are no metrics
	//        provided to describe that quality. The default is "false".
	//    `1s_interval_classification` - (boolean) Set to true to request
	//        classification for a video at one-second intervals. AutoML Video
	//        Intelligence returns labels and their confidence scores for each
	//        second of the entire segment of the video that user specified in the
	//        request configuration.
	//        WARNING: Model evaluation is not done for this classification
	//        type, the quality of it depends on training data, but there are no
	//        metrics provided to describe that quality. The default is
	//        "false".
	Params               map[string]string `protobuf:"bytes,5,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *BatchPredictRequest) Reset()         { *m = BatchPredictRequest{} }
func (m *BatchPredictRequest) String() string { return proto.CompactTextString(m) }
func (*BatchPredictRequest) ProtoMessage()    {}
func (*BatchPredictRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_prediction_service_4cd94add91ccba6b, []int{2}
}
func (m *BatchPredictRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchPredictRequest.Unmarshal(m, b)
}
func (m *BatchPredictRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchPredictRequest.Marshal(b, m, deterministic)
}
func (dst *BatchPredictRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchPredictRequest.Merge(dst, src)
}
func (m *BatchPredictRequest) XXX_Size() int {
	return xxx_messageInfo_BatchPredictRequest.Size(m)
}
func (m *BatchPredictRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchPredictRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BatchPredictRequest proto.InternalMessageInfo

func (m *BatchPredictRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BatchPredictRequest) GetInputConfig() *BatchPredictInputConfig {
	if m != nil {
		return m.InputConfig
	}
	return nil
}

func (m *BatchPredictRequest) GetOutputConfig() *BatchPredictOutputConfig {
	if m != nil {
		return m.OutputConfig
	}
	return nil
}

func (m *BatchPredictRequest) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

// Batch predict result.
type BatchPredictResult struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BatchPredictResult) Reset()         { *m = BatchPredictResult{} }
func (m *BatchPredictResult) String() string { return proto.CompactTextString(m) }
func (*BatchPredictResult) ProtoMessage()    {}
func (*BatchPredictResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_prediction_service_4cd94add91ccba6b, []int{3}
}
func (m *BatchPredictResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchPredictResult.Unmarshal(m, b)
}
func (m *BatchPredictResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchPredictResult.Marshal(b, m, deterministic)
}
func (dst *BatchPredictResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchPredictResult.Merge(dst, src)
}
func (m *BatchPredictResult) XXX_Size() int {
	return xxx_messageInfo_BatchPredictResult.Size(m)
}
func (m *BatchPredictResult) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchPredictResult.DiscardUnknown(m)
}

var xxx_messageInfo_BatchPredictResult proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PredictRequest)(nil), "google.cloud.automl.v1beta1.PredictRequest")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.automl.v1beta1.PredictRequest.ParamsEntry")
	proto.RegisterType((*PredictResponse)(nil), "google.cloud.automl.v1beta1.PredictResponse")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.automl.v1beta1.PredictResponse.MetadataEntry")
	proto.RegisterType((*BatchPredictRequest)(nil), "google.cloud.automl.v1beta1.BatchPredictRequest")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.automl.v1beta1.BatchPredictRequest.ParamsEntry")
	proto.RegisterType((*BatchPredictResult)(nil), "google.cloud.automl.v1beta1.BatchPredictResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PredictionServiceClient is the client API for PredictionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PredictionServiceClient interface {
	// Perform an online prediction. The prediction result will be directly
	// returned in the response.
	// Available for following ML problems, and their expected request payloads:
	// * Image Classification - Image in .JPEG, .GIF or .PNG format, image_bytes
	//                          up to 30MB.
	// * Image Object Detection - Image in .JPEG, .GIF or .PNG format, image_bytes
	//                            up to 30MB.
	// * Text Classification - TextSnippet, content up to 10,000 characters,
	//                         UTF-8 encoded.
	// * Text Extraction - TextSnippet, content up to 30,000 characters,
	//                     UTF-8 NFC encoded. * Translation - TextSnippet, content up to 25,000 characters, UTF-8
	//                 encoded.
	// * Tables - Row, with column values matching the columns of the model,
	//            up to 5MB.
	// * Text Sentiment - TextSnippet, content up 500 characters, UTF-8 encoded.
	Predict(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*PredictResponse, error)
	// Perform a batch prediction. Unlike the online [Predict][google.cloud.automl.v1beta1.PredictionService.Predict], batch
	// prediction result won't be immediately available in the response. Instead,
	// a long running operation object is returned. User can poll the operation
	// result via [GetOperation][google.longrunning.Operations.GetOperation]
	// method. Once the operation is done, [BatchPredictResult][google.cloud.automl.v1beta1.BatchPredictResult] is returned in
	// the [response][google.longrunning.Operation.response] field.
	// Available for following ML problems:
	// * Video Classification
	// * Text Extraction
	// * Tables
	BatchPredict(ctx context.Context, in *BatchPredictRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
}

type predictionServiceClient struct {
	cc *grpc.ClientConn
}

func NewPredictionServiceClient(cc *grpc.ClientConn) PredictionServiceClient {
	return &predictionServiceClient{cc}
}

func (c *predictionServiceClient) Predict(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*PredictResponse, error) {
	out := new(PredictResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.automl.v1beta1.PredictionService/Predict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *predictionServiceClient) BatchPredict(ctx context.Context, in *BatchPredictRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.cloud.automl.v1beta1.PredictionService/BatchPredict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PredictionServiceServer is the server API for PredictionService service.
type PredictionServiceServer interface {
	// Perform an online prediction. The prediction result will be directly
	// returned in the response.
	// Available for following ML problems, and their expected request payloads:
	// * Image Classification - Image in .JPEG, .GIF or .PNG format, image_bytes
	//                          up to 30MB.
	// * Image Object Detection - Image in .JPEG, .GIF or .PNG format, image_bytes
	//                            up to 30MB.
	// * Text Classification - TextSnippet, content up to 10,000 characters,
	//                         UTF-8 encoded.
	// * Text Extraction - TextSnippet, content up to 30,000 characters,
	//                     UTF-8 NFC encoded. * Translation - TextSnippet, content up to 25,000 characters, UTF-8
	//                 encoded.
	// * Tables - Row, with column values matching the columns of the model,
	//            up to 5MB.
	// * Text Sentiment - TextSnippet, content up 500 characters, UTF-8 encoded.
	Predict(context.Context, *PredictRequest) (*PredictResponse, error)
	// Perform a batch prediction. Unlike the online [Predict][google.cloud.automl.v1beta1.PredictionService.Predict], batch
	// prediction result won't be immediately available in the response. Instead,
	// a long running operation object is returned. User can poll the operation
	// result via [GetOperation][google.longrunning.Operations.GetOperation]
	// method. Once the operation is done, [BatchPredictResult][google.cloud.automl.v1beta1.BatchPredictResult] is returned in
	// the [response][google.longrunning.Operation.response] field.
	// Available for following ML problems:
	// * Video Classification
	// * Text Extraction
	// * Tables
	BatchPredict(context.Context, *BatchPredictRequest) (*longrunning.Operation, error)
}

func RegisterPredictionServiceServer(s *grpc.Server, srv PredictionServiceServer) {
	s.RegisterService(&_PredictionService_serviceDesc, srv)
}

func _PredictionService_Predict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).Predict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.automl.v1beta1.PredictionService/Predict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).Predict(ctx, req.(*PredictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PredictionService_BatchPredict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchPredictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).BatchPredict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.automl.v1beta1.PredictionService/BatchPredict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).BatchPredict(ctx, req.(*BatchPredictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PredictionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.automl.v1beta1.PredictionService",
	HandlerType: (*PredictionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Predict",
			Handler:    _PredictionService_Predict_Handler,
		},
		{
			MethodName: "BatchPredict",
			Handler:    _PredictionService_BatchPredict_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/automl/v1beta1/prediction_service.proto",
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/prediction_service.proto", fileDescriptor_prediction_service_4cd94add91ccba6b)
}

var fileDescriptor_prediction_service_4cd94add91ccba6b = []byte{
	// 634 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcd, 0x6e, 0xd4, 0x3c,
	0x14, 0x95, 0x33, 0xfd, 0xf9, 0xea, 0x69, 0x3f, 0xc0, 0x54, 0x68, 0x94, 0x82, 0xa8, 0x02, 0x8b,
	0x6a, 0x5a, 0xc5, 0xb4, 0x14, 0x15, 0xa6, 0x65, 0x31, 0x53, 0x55, 0x85, 0x45, 0xd5, 0x28, 0xa0,
	0x22, 0x55, 0x95, 0x46, 0x6e, 0xc6, 0x84, 0x80, 0x63, 0x87, 0xc4, 0xa9, 0xa8, 0x10, 0x1b, 0x5e,
	0x81, 0x2d, 0x0f, 0xc0, 0xb3, 0xb0, 0xe5, 0x15, 0xd8, 0xb0, 0x82, 0x1d, 0x5b, 0x14, 0xdb, 0x93,
	0x66, 0x54, 0x14, 0xcd, 0x2c, 0xd8, 0xc5, 0xc9, 0x3d, 0xe7, 0x9e, 0x73, 0xaf, 0x73, 0xe0, 0x66,
	0x28, 0x44, 0xc8, 0x28, 0x0e, 0x98, 0xc8, 0x07, 0x98, 0xe4, 0x52, 0xc4, 0x0c, 0x9f, 0xad, 0x9f,
	0x52, 0x49, 0xd6, 0x71, 0x92, 0xd2, 0x41, 0x14, 0xc8, 0x48, 0xf0, 0x7e, 0x46, 0xd3, 0xb3, 0x28,
	0xa0, 0x6e, 0x92, 0x0a, 0x29, 0xd0, 0x92, 0x46, 0xb9, 0x0a, 0xe5, 0x6a, 0x94, 0x6b, 0x50, 0xf6,
	0x4d, 0x43, 0x49, 0x92, 0x08, 0x13, 0xce, 0x85, 0x24, 0x05, 0x43, 0xa6, 0xa1, 0x76, 0x6d, 0xc3,
	0x8b, 0xf2, 0x7e, 0x42, 0xce, 0x99, 0x20, 0x03, 0x83, 0x5a, 0xab, 0x43, 0x0d, 0x88, 0x24, 0xfd,
	0x48, 0xd2, 0x78, 0xd8, 0xe3, 0x6e, 0x5d, 0x75, 0x24, 0xc6, 0xe1, 0x14, 0x09, 0x4d, 0x47, 0x74,
	0xdf, 0x31, 0xd5, 0x4c, 0xf0, 0x30, 0xcd, 0x39, 0x8f, 0x78, 0x78, 0xa9, 0xc8, 0xf9, 0x0d, 0xe0,
	0xff, 0x9e, 0x1e, 0x9a, 0x4f, 0xdf, 0xe6, 0x34, 0x93, 0x08, 0xc1, 0x29, 0x4e, 0x62, 0xda, 0x02,
	0xcb, 0x60, 0x65, 0xce, 0x57, 0xcf, 0x68, 0x0f, 0xce, 0x1a, 0x7b, 0x2d, 0x6b, 0x19, 0xac, 0x34,
	0x37, 0x56, 0xdd, 0x9a, 0x81, 0xba, 0x7b, 0xef, 0x48, 0x9c, 0x30, 0xea, 0x69, 0x88, 0x3f, 0xc4,
	0xa2, 0x43, 0x38, 0x93, 0x90, 0x94, 0xc4, 0x59, 0xab, 0xb1, 0xdc, 0x58, 0x69, 0x6e, 0x6c, 0xd5,
	0xb2, 0x8c, 0xea, 0x72, 0x3d, 0x85, 0xdc, 0xe3, 0x32, 0x3d, 0xf7, 0x0d, 0x8d, 0xfd, 0x08, 0x36,
	0x2b, 0xaf, 0xd1, 0x55, 0xd8, 0x78, 0x43, 0xcf, 0x8d, 0xf2, 0xe2, 0x11, 0x2d, 0xc2, 0xe9, 0x33,
	0xc2, 0x72, 0xaa, 0x64, 0xcf, 0xf9, 0xfa, 0xd0, 0xb1, 0x1e, 0x02, 0xe7, 0x17, 0x80, 0x57, 0xca,
	0x0e, 0x59, 0x22, 0x78, 0x46, 0xd1, 0x93, 0x0b, 0x9b, 0x40, 0x09, 0x74, 0x6b, 0x05, 0x76, 0xcb,
	0xe5, 0x5f, 0x72, 0x7a, 0x04, 0xff, 0x8b, 0xa9, 0x24, 0xc5, 0xa2, 0x5b, 0x96, 0xa2, 0xea, 0x8c,
	0xe7, 0x55, 0x2b, 0x71, 0x0f, 0x0c, 0x58, 0xdb, 0x2d, 0xb9, 0xec, 0x6d, 0xb8, 0x30, 0xf2, 0x69,
	0x22, 0xcb, 0x3f, 0x2d, 0x78, 0xbd, 0x47, 0x64, 0xf0, 0x6a, 0x8c, 0x8d, 0xbf, 0x80, 0xf3, 0x11,
	0x4f, 0x72, 0xd9, 0x0f, 0x04, 0x7f, 0x19, 0x85, 0xad, 0x86, 0x5a, 0xfb, 0x66, 0xad, 0x89, 0x2a,
	0xf7, 0xd3, 0x02, 0xbc, 0xab, 0xb0, 0x7e, 0x33, 0xba, 0x38, 0xa0, 0x63, 0xb8, 0x20, 0x72, 0x59,
	0x61, 0x9e, 0x52, 0xcc, 0x0f, 0xc6, 0x66, 0x3e, 0x54, 0x68, 0x43, 0x3d, 0x2f, 0x2a, 0x27, 0xf4,
	0xbc, 0xbc, 0x5f, 0xd3, 0x6a, 0xe6, 0x3b, 0x63, 0x93, 0xfe, 0xa3, 0x4b, 0xb6, 0x08, 0xd1, 0x68,
	0x97, 0x2c, 0x67, 0x72, 0xe3, 0x87, 0x05, 0xaf, 0x79, 0x65, 0x52, 0x3d, 0xd3, 0x41, 0x85, 0xbe,
	0x00, 0x38, 0x6b, 0xde, 0xa2, 0xd5, 0x09, 0x7e, 0x0c, 0x7b, 0x6d, 0x92, 0x9b, 0xe5, 0xf4, 0x3e,
	0x7e, 0xfb, 0xfe, 0xc9, 0xda, 0x71, 0xb6, 0xca, 0xe4, 0x78, 0x5f, 0x2c, 0xfc, 0x71, 0x92, 0x8a,
	0xd7, 0x34, 0x90, 0x19, 0x6e, 0x63, 0x26, 0x02, 0x1d, 0x12, 0xb8, 0x8d, 0x63, 0x31, 0xa0, 0x2c,
	0xc3, 0xed, 0x0f, 0x1d, 0x93, 0xad, 0x1d, 0xd0, 0x2e, 0xa4, 0xce, 0x57, 0x7d, 0xa1, 0x7b, 0x93,
	0x0e, 0xda, 0xbe, 0x35, 0x44, 0x54, 0xe2, 0xc9, 0x3d, 0x1c, 0xc6, 0x93, 0xb3, 0xaf, 0x54, 0x76,
	0x9d, 0x9d, 0x49, 0x55, 0x9e, 0x56, 0x7a, 0x75, 0x40, 0xbb, 0xf7, 0x19, 0xc0, 0xdb, 0x81, 0x88,
	0xeb, 0xf4, 0xf5, 0x6e, 0x5c, 0x5a, 0x86, 0x57, 0x84, 0xa3, 0x07, 0x8e, 0xbb, 0x06, 0x16, 0x0a,
	0x46, 0x78, 0xe8, 0x8a, 0x34, 0xc4, 0x21, 0xe5, 0x2a, 0x3a, 0xb1, 0xfe, 0x44, 0x92, 0x28, 0xfb,
	0x6b, 0x20, 0x6f, 0xeb, 0xe3, 0x57, 0x6b, 0x69, 0x5f, 0x15, 0x9e, 0xec, 0x16, 0x45, 0x27, 0xdd,
	0x5c, 0x8a, 0x03, 0x76, 0x72, 0xa4, 0x8b, 0x4e, 0x67, 0x14, 0xd7, 0xfd, 0x3f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xda, 0xb7, 0xf9, 0x0c, 0xd6, 0x06, 0x00, 0x00,
}
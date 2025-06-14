// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.0
// source: proto/cart.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int32                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	mi := &file_proto_cart_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_proto_cart_proto_rawDescGZIP(), []int{0}
}

func (x *UserRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CartItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        int32                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProductId     int32                  `protobuf:"varint,3,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity      int32                  `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CartItem) Reset() {
	*x = CartItem{}
	mi := &file_proto_cart_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItem) ProtoMessage() {}

func (x *CartItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItem.ProtoReflect.Descriptor instead.
func (*CartItem) Descriptor() ([]byte, []int) {
	return file_proto_cart_proto_rawDescGZIP(), []int{1}
}

func (x *CartItem) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CartItem) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CartItem) GetProductId() int32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *CartItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type CartItemList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         []*CartItem            `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CartItemList) Reset() {
	*x = CartItemList{}
	mi := &file_proto_cart_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CartItemList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItemList) ProtoMessage() {}

func (x *CartItemList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItemList.ProtoReflect.Descriptor instead.
func (*CartItemList) Descriptor() ([]byte, []int) {
	return file_proto_cart_proto_rawDescGZIP(), []int{2}
}

func (x *CartItemList) GetItems() []*CartItem {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_proto_cart_proto protoreflect.FileDescriptor

const file_proto_cart_proto_rawDesc = "" +
	"\n" +
	"\x10proto/cart.proto\x12\x04cart\"&\n" +
	"\vUserRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\x05R\x06userId\"n\n" +
	"\bCartItem\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\x05R\x06userId\x12\x1d\n" +
	"\n" +
	"product_id\x18\x03 \x01(\x05R\tproductId\x12\x1a\n" +
	"\bquantity\x18\x04 \x01(\x05R\bquantity\"4\n" +
	"\fCartItemList\x12$\n" +
	"\x05items\x18\x01 \x03(\v2\x0e.cart.CartItemR\x05items2D\n" +
	"\vCartService\x125\n" +
	"\fGetCartItems\x12\x11.cart.UserRequest\x1a\x12.cart.CartItemListB\x06Z\x04gen/b\x06proto3"

var (
	file_proto_cart_proto_rawDescOnce sync.Once
	file_proto_cart_proto_rawDescData []byte
)

func file_proto_cart_proto_rawDescGZIP() []byte {
	file_proto_cart_proto_rawDescOnce.Do(func() {
		file_proto_cart_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_cart_proto_rawDesc), len(file_proto_cart_proto_rawDesc)))
	})
	return file_proto_cart_proto_rawDescData
}

var file_proto_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_cart_proto_goTypes = []any{
	(*UserRequest)(nil),  // 0: cart.UserRequest
	(*CartItem)(nil),     // 1: cart.CartItem
	(*CartItemList)(nil), // 2: cart.CartItemList
}
var file_proto_cart_proto_depIdxs = []int32{
	1, // 0: cart.CartItemList.items:type_name -> cart.CartItem
	0, // 1: cart.CartService.GetCartItems:input_type -> cart.UserRequest
	2, // 2: cart.CartService.GetCartItems:output_type -> cart.CartItemList
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_cart_proto_init() }
func file_proto_cart_proto_init() {
	if File_proto_cart_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_cart_proto_rawDesc), len(file_proto_cart_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_cart_proto_goTypes,
		DependencyIndexes: file_proto_cart_proto_depIdxs,
		MessageInfos:      file_proto_cart_proto_msgTypes,
	}.Build()
	File_proto_cart_proto = out.File
	file_proto_cart_proto_goTypes = nil
	file_proto_cart_proto_depIdxs = nil
}

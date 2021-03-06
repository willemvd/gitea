// Code generated by protoc-gen-go.
// source: Authentication.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = math.Inf

type TokenIdentifier_Kind int32

const (
	TokenIdentifier_HBASE_AUTH_TOKEN TokenIdentifier_Kind = 0
)

var TokenIdentifier_Kind_name = map[int32]string{
	0: "HBASE_AUTH_TOKEN",
}
var TokenIdentifier_Kind_value = map[string]int32{
	"HBASE_AUTH_TOKEN": 0,
}

func (x TokenIdentifier_Kind) Enum() *TokenIdentifier_Kind {
	p := new(TokenIdentifier_Kind)
	*p = x
	return p
}
func (x TokenIdentifier_Kind) String() string {
	return proto1.EnumName(TokenIdentifier_Kind_name, int32(x))
}
func (x *TokenIdentifier_Kind) UnmarshalJSON(data []byte) error {
	value, err := proto1.UnmarshalJSONEnum(TokenIdentifier_Kind_value, data, "TokenIdentifier_Kind")
	if err != nil {
		return err
	}
	*x = TokenIdentifier_Kind(value)
	return nil
}

type AuthenticationKey struct {
	Id               *int32 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	ExpirationDate   *int64 `protobuf:"varint,2,req,name=expiration_date" json:"expiration_date,omitempty"`
	Key              []byte `protobuf:"bytes,3,req,name=key" json:"key,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *AuthenticationKey) Reset()         { *m = AuthenticationKey{} }
func (m *AuthenticationKey) String() string { return proto1.CompactTextString(m) }
func (*AuthenticationKey) ProtoMessage()    {}

func (m *AuthenticationKey) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *AuthenticationKey) GetExpirationDate() int64 {
	if m != nil && m.ExpirationDate != nil {
		return *m.ExpirationDate
	}
	return 0
}

func (m *AuthenticationKey) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type TokenIdentifier struct {
	Kind             *TokenIdentifier_Kind `protobuf:"varint,1,req,name=kind,enum=proto.TokenIdentifier_Kind" json:"kind,omitempty"`
	Username         []byte                `protobuf:"bytes,2,req,name=username" json:"username,omitempty"`
	KeyId            *int32                `protobuf:"varint,3,req,name=key_id" json:"key_id,omitempty"`
	IssueDate        *int64                `protobuf:"varint,4,opt,name=issue_date" json:"issue_date,omitempty"`
	ExpirationDate   *int64                `protobuf:"varint,5,opt,name=expiration_date" json:"expiration_date,omitempty"`
	SequenceNumber   *int64                `protobuf:"varint,6,opt,name=sequence_number" json:"sequence_number,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *TokenIdentifier) Reset()         { *m = TokenIdentifier{} }
func (m *TokenIdentifier) String() string { return proto1.CompactTextString(m) }
func (*TokenIdentifier) ProtoMessage()    {}

func (m *TokenIdentifier) GetKind() TokenIdentifier_Kind {
	if m != nil && m.Kind != nil {
		return *m.Kind
	}
	return TokenIdentifier_HBASE_AUTH_TOKEN
}

func (m *TokenIdentifier) GetUsername() []byte {
	if m != nil {
		return m.Username
	}
	return nil
}

func (m *TokenIdentifier) GetKeyId() int32 {
	if m != nil && m.KeyId != nil {
		return *m.KeyId
	}
	return 0
}

func (m *TokenIdentifier) GetIssueDate() int64 {
	if m != nil && m.IssueDate != nil {
		return *m.IssueDate
	}
	return 0
}

func (m *TokenIdentifier) GetExpirationDate() int64 {
	if m != nil && m.ExpirationDate != nil {
		return *m.ExpirationDate
	}
	return 0
}

func (m *TokenIdentifier) GetSequenceNumber() int64 {
	if m != nil && m.SequenceNumber != nil {
		return *m.SequenceNumber
	}
	return 0
}

// Serialization of the org.apache.hadoop.security.token.Token class
// Note that this is a Hadoop class, so fields may change!
type Token struct {
	// the TokenIdentifier in serialized form
	// Note: we can't use the protobuf directly because the Hadoop Token class
	// only stores the serialized bytes
	Identifier       []byte `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
	Password         []byte `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Service          []byte `protobuf:"bytes,3,opt,name=service" json:"service,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto1.CompactTextString(m) }
func (*Token) ProtoMessage()    {}

func (m *Token) GetIdentifier() []byte {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *Token) GetPassword() []byte {
	if m != nil {
		return m.Password
	}
	return nil
}

func (m *Token) GetService() []byte {
	if m != nil {
		return m.Service
	}
	return nil
}

// RPC request & response messages
type GetAuthenticationTokenRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetAuthenticationTokenRequest) Reset()         { *m = GetAuthenticationTokenRequest{} }
func (m *GetAuthenticationTokenRequest) String() string { return proto1.CompactTextString(m) }
func (*GetAuthenticationTokenRequest) ProtoMessage()    {}

type GetAuthenticationTokenResponse struct {
	Token            *Token `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetAuthenticationTokenResponse) Reset()         { *m = GetAuthenticationTokenResponse{} }
func (m *GetAuthenticationTokenResponse) String() string { return proto1.CompactTextString(m) }
func (*GetAuthenticationTokenResponse) ProtoMessage()    {}

func (m *GetAuthenticationTokenResponse) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type WhoAmIRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *WhoAmIRequest) Reset()         { *m = WhoAmIRequest{} }
func (m *WhoAmIRequest) String() string { return proto1.CompactTextString(m) }
func (*WhoAmIRequest) ProtoMessage()    {}

type WhoAmIResponse struct {
	Username         *string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	AuthMethod       *string `protobuf:"bytes,2,opt,name=auth_method" json:"auth_method,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *WhoAmIResponse) Reset()         { *m = WhoAmIResponse{} }
func (m *WhoAmIResponse) String() string { return proto1.CompactTextString(m) }
func (*WhoAmIResponse) ProtoMessage()    {}

func (m *WhoAmIResponse) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func (m *WhoAmIResponse) GetAuthMethod() string {
	if m != nil && m.AuthMethod != nil {
		return *m.AuthMethod
	}
	return ""
}

func init() {
	proto1.RegisterEnum("proto.TokenIdentifier_Kind", TokenIdentifier_Kind_name, TokenIdentifier_Kind_value)
}

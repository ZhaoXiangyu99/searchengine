// Code generated by Kitex v0.3.1. DO NOT EDIT.

package userservice

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
	"google.golang.org/protobuf/proto"
	"searchengine/kitex_gen/userModel"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*userModel.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateUser": kitex.NewMethodInfo(createUserHandler, newCreateUserArgs, newCreateUserResult, false),
		"MGetUser":   kitex.NewMethodInfo(mGetUserHandler, newMGetUserArgs, newMGetUserResult, false),
		"CheckUser":  kitex.NewMethodInfo(checkUserHandler, newCheckUserArgs, newCheckUserResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "user",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.3.1",
		Extra:           extra,
	}
	return svcInfo
}

func createUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(userModel.CreateUserRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(userModel.UserService).CreateUser(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CreateUserArgs:
		success, err := handler.(userModel.UserService).CreateUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateUserResult)
		realResult.Success = success
	}
	return nil
}
func newCreateUserArgs() interface{} {
	return &CreateUserArgs{}
}

func newCreateUserResult() interface{} {
	return &CreateUserResult{}
}

type CreateUserArgs struct {
	Req *userModel.CreateUserRequest
}

func (p *CreateUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CreateUserArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CreateUserArgs) Unmarshal(in []byte) error {
	msg := new(userModel.CreateUserRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateUserArgs_Req_DEFAULT *userModel.CreateUserRequest

func (p *CreateUserArgs) GetReq() *userModel.CreateUserRequest {
	if !p.IsSetReq() {
		return CreateUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type CreateUserResult struct {
	Success *userModel.CreateUserResponse
}

var CreateUserResult_Success_DEFAULT *userModel.CreateUserResponse

func (p *CreateUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CreateUserResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CreateUserResult) Unmarshal(in []byte) error {
	msg := new(userModel.CreateUserResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateUserResult) GetSuccess() *userModel.CreateUserResponse {
	if !p.IsSetSuccess() {
		return CreateUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*userModel.CreateUserResponse)
}

func (p *CreateUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func mGetUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(userModel.MGetUserRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(userModel.UserService).MGetUser(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *MGetUserArgs:
		success, err := handler.(userModel.UserService).MGetUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*MGetUserResult)
		realResult.Success = success
	}
	return nil
}
func newMGetUserArgs() interface{} {
	return &MGetUserArgs{}
}

func newMGetUserResult() interface{} {
	return &MGetUserResult{}
}

type MGetUserArgs struct {
	Req *userModel.MGetUserRequest
}

func (p *MGetUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MGetUserArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *MGetUserArgs) Unmarshal(in []byte) error {
	msg := new(userModel.MGetUserRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var MGetUserArgs_Req_DEFAULT *userModel.MGetUserRequest

func (p *MGetUserArgs) GetReq() *userModel.MGetUserRequest {
	if !p.IsSetReq() {
		return MGetUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MGetUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type MGetUserResult struct {
	Success *userModel.MGetUserResponse
}

var MGetUserResult_Success_DEFAULT *userModel.MGetUserResponse

func (p *MGetUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MGetUserResult")
	}
	return proto.Marshal(p.Success)
}

func (p *MGetUserResult) Unmarshal(in []byte) error {
	msg := new(userModel.MGetUserResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MGetUserResult) GetSuccess() *userModel.MGetUserResponse {
	if !p.IsSetSuccess() {
		return MGetUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MGetUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*userModel.MGetUserResponse)
}

func (p *MGetUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func checkUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(userModel.CheckUserRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(userModel.UserService).CheckUser(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CheckUserArgs:
		success, err := handler.(userModel.UserService).CheckUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CheckUserResult)
		realResult.Success = success
	}
	return nil
}
func newCheckUserArgs() interface{} {
	return &CheckUserArgs{}
}

func newCheckUserResult() interface{} {
	return &CheckUserResult{}
}

type CheckUserArgs struct {
	Req *userModel.CheckUserRequest
}

func (p *CheckUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CheckUserArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CheckUserArgs) Unmarshal(in []byte) error {
	msg := new(userModel.CheckUserRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CheckUserArgs_Req_DEFAULT *userModel.CheckUserRequest

func (p *CheckUserArgs) GetReq() *userModel.CheckUserRequest {
	if !p.IsSetReq() {
		return CheckUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CheckUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type CheckUserResult struct {
	Success *userModel.CheckUserResponse
}

var CheckUserResult_Success_DEFAULT *userModel.CheckUserResponse

func (p *CheckUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CheckUserResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CheckUserResult) Unmarshal(in []byte) error {
	msg := new(userModel.CheckUserResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CheckUserResult) GetSuccess() *userModel.CheckUserResponse {
	if !p.IsSetSuccess() {
		return CheckUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CheckUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*userModel.CheckUserResponse)
}

func (p *CheckUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateUser(ctx context.Context, Req *userModel.CreateUserRequest) (r *userModel.CreateUserResponse, err error) {
	var _args CreateUserArgs
	_args.Req = Req
	var _result CreateUserResult
	if err = p.c.Call(ctx, "CreateUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MGetUser(ctx context.Context, Req *userModel.MGetUserRequest) (r *userModel.MGetUserResponse, err error) {
	var _args MGetUserArgs
	_args.Req = Req
	var _result MGetUserResult
	if err = p.c.Call(ctx, "MGetUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CheckUser(ctx context.Context, Req *userModel.CheckUserRequest) (r *userModel.CheckUserResponse, err error) {
	var _args CheckUserArgs
	_args.Req = Req
	var _result CheckUserResult
	if err = p.c.Call(ctx, "CheckUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

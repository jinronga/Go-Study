// Code generated by Kitex v0.8.0. DO NOT EDIT.

package itemservice

import (
	"context"
	item "example_shop/kitex_gen/example/shop/item"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return itemServiceServiceInfo
}

var itemServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "ItemService"
	handlerType := (*item.ItemService)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetItem": kitex.NewMethodInfo(getItemHandler, newItemServiceGetItemArgs, newItemServiceGetItemResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "item",
		"ServiceFilePath": `idl/item.thrift`,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.8.0",
		Extra:           extra,
	}
	return svcInfo
}

func getItemHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*item.ItemServiceGetItemArgs)
	realResult := result.(*item.ItemServiceGetItemResult)
	success, err := handler.(item.ItemService).GetItem(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newItemServiceGetItemArgs() interface{} {
	return item.NewItemServiceGetItemArgs()
}

func newItemServiceGetItemResult() interface{} {
	return item.NewItemServiceGetItemResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetItem(ctx context.Context, req *item.GetItemReq) (r *item.GetItemResp, err error) {
	var _args item.ItemServiceGetItemArgs
	_args.Req = req
	var _result item.ItemServiceGetItemResult
	if err = p.c.Call(ctx, "GetItem", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

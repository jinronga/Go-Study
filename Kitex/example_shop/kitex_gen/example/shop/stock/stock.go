// Code generated by thriftgo (0.3.12). DO NOT EDIT.

package stock

import (
	"context"
	"example_shop/kitex_gen/example/shop/base"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

type GetItemStockReq struct {
	ItemId int64 `thrift:"item_id,1,required" frugal:"1,required,i64" json:"item_id"`
}

func NewGetItemStockReq() *GetItemStockReq {
	return &GetItemStockReq{}
}

func (p *GetItemStockReq) InitDefault() {
	*p = GetItemStockReq{}
}

func (p *GetItemStockReq) GetItemId() (v int64) {
	return p.ItemId
}
func (p *GetItemStockReq) SetItemId(val int64) {
	p.ItemId = val
}

var fieldIDToName_GetItemStockReq = map[int16]string{
	1: "item_id",
}

func (p *GetItemStockReq) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetItemId bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetItemId = true
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetItemId {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_GetItemStockReq[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_GetItemStockReq[fieldId]))
}

func (p *GetItemStockReq) ReadField1(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.ItemId = _field
	return nil
}

func (p *GetItemStockReq) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("GetItemStockReq"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *GetItemStockReq) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("item_id", thrift.I64, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.ItemId); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *GetItemStockReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetItemStockReq(%+v)", *p)

}

func (p *GetItemStockReq) DeepEqual(ano *GetItemStockReq) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.ItemId) {
		return false
	}
	return true
}

func (p *GetItemStockReq) Field1DeepEqual(src int64) bool {

	if p.ItemId != src {
		return false
	}
	return true
}

type GetItemStockResp struct {
	Stock    int64          `thrift:"stock,1" frugal:"1,default,i64" json:"stock"`
	BaseResp *base.BaseResp `thrift:"BaseResp,255" frugal:"255,default,base.BaseResp" json:"BaseResp"`
}

func NewGetItemStockResp() *GetItemStockResp {
	return &GetItemStockResp{}
}

func (p *GetItemStockResp) InitDefault() {
	*p = GetItemStockResp{}
}

func (p *GetItemStockResp) GetStock() (v int64) {
	return p.Stock
}

var GetItemStockResp_BaseResp_DEFAULT *base.BaseResp

func (p *GetItemStockResp) GetBaseResp() (v *base.BaseResp) {
	if !p.IsSetBaseResp() {
		return GetItemStockResp_BaseResp_DEFAULT
	}
	return p.BaseResp
}
func (p *GetItemStockResp) SetStock(val int64) {
	p.Stock = val
}
func (p *GetItemStockResp) SetBaseResp(val *base.BaseResp) {
	p.BaseResp = val
}

var fieldIDToName_GetItemStockResp = map[int16]string{
	1:   "stock",
	255: "BaseResp",
}

func (p *GetItemStockResp) IsSetBaseResp() bool {
	return p.BaseResp != nil
}

func (p *GetItemStockResp) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 255:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField255(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_GetItemStockResp[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *GetItemStockResp) ReadField1(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Stock = _field
	return nil
}
func (p *GetItemStockResp) ReadField255(iprot thrift.TProtocol) error {
	_field := base.NewBaseResp()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.BaseResp = _field
	return nil
}

func (p *GetItemStockResp) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("GetItemStockResp"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField255(oprot); err != nil {
			fieldId = 255
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *GetItemStockResp) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("stock", thrift.I64, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.Stock); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *GetItemStockResp) writeField255(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("BaseResp", thrift.STRUCT, 255); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.BaseResp.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 255 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 255 end error: ", p), err)
}

func (p *GetItemStockResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetItemStockResp(%+v)", *p)

}

func (p *GetItemStockResp) DeepEqual(ano *GetItemStockResp) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Stock) {
		return false
	}
	if !p.Field255DeepEqual(ano.BaseResp) {
		return false
	}
	return true
}

func (p *GetItemStockResp) Field1DeepEqual(src int64) bool {

	if p.Stock != src {
		return false
	}
	return true
}
func (p *GetItemStockResp) Field255DeepEqual(src *base.BaseResp) bool {

	if !p.BaseResp.DeepEqual(src) {
		return false
	}
	return true
}

type StockService interface {
	GetItemStock(ctx context.Context, req *GetItemStockReq) (r *GetItemStockResp, err error)
}

type StockServiceClient struct {
	c thrift.TClient
}

func NewStockServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *StockServiceClient {
	return &StockServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewStockServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *StockServiceClient {
	return &StockServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewStockServiceClient(c thrift.TClient) *StockServiceClient {
	return &StockServiceClient{
		c: c,
	}
}

func (p *StockServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *StockServiceClient) GetItemStock(ctx context.Context, req *GetItemStockReq) (r *GetItemStockResp, err error) {
	var _args StockServiceGetItemStockArgs
	_args.Req = req
	var _result StockServiceGetItemStockResult
	if err = p.Client_().Call(ctx, "GetItemStock", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type StockServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      StockService
}

func (p *StockServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *StockServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *StockServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewStockServiceProcessor(handler StockService) *StockServiceProcessor {
	self := &StockServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("GetItemStock", &stockServiceProcessorGetItemStock{handler: handler})
	return self
}
func (p *StockServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type stockServiceProcessorGetItemStock struct {
	handler StockService
}

func (p *stockServiceProcessorGetItemStock) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := StockServiceGetItemStockArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("GetItemStock", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := StockServiceGetItemStockResult{}
	var retval *GetItemStockResp
	if retval, err2 = p.handler.GetItemStock(ctx, args.Req); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing GetItemStock: "+err2.Error())
		oprot.WriteMessageBegin("GetItemStock", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("GetItemStock", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type StockServiceGetItemStockArgs struct {
	Req *GetItemStockReq `thrift:"req,1" frugal:"1,default,GetItemStockReq" json:"req"`
}

func NewStockServiceGetItemStockArgs() *StockServiceGetItemStockArgs {
	return &StockServiceGetItemStockArgs{}
}

func (p *StockServiceGetItemStockArgs) InitDefault() {
	*p = StockServiceGetItemStockArgs{}
}

var StockServiceGetItemStockArgs_Req_DEFAULT *GetItemStockReq

func (p *StockServiceGetItemStockArgs) GetReq() (v *GetItemStockReq) {
	if !p.IsSetReq() {
		return StockServiceGetItemStockArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *StockServiceGetItemStockArgs) SetReq(val *GetItemStockReq) {
	p.Req = val
}

var fieldIDToName_StockServiceGetItemStockArgs = map[int16]string{
	1: "req",
}

func (p *StockServiceGetItemStockArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *StockServiceGetItemStockArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_StockServiceGetItemStockArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *StockServiceGetItemStockArgs) ReadField1(iprot thrift.TProtocol) error {
	_field := NewGetItemStockReq()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Req = _field
	return nil
}

func (p *StockServiceGetItemStockArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("GetItemStock_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *StockServiceGetItemStockArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Req.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *StockServiceGetItemStockArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("StockServiceGetItemStockArgs(%+v)", *p)

}

func (p *StockServiceGetItemStockArgs) DeepEqual(ano *StockServiceGetItemStockArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Req) {
		return false
	}
	return true
}

func (p *StockServiceGetItemStockArgs) Field1DeepEqual(src *GetItemStockReq) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

type StockServiceGetItemStockResult struct {
	Success *GetItemStockResp `thrift:"success,0,optional" frugal:"0,optional,GetItemStockResp" json:"success,omitempty"`
}

func NewStockServiceGetItemStockResult() *StockServiceGetItemStockResult {
	return &StockServiceGetItemStockResult{}
}

func (p *StockServiceGetItemStockResult) InitDefault() {
	*p = StockServiceGetItemStockResult{}
}

var StockServiceGetItemStockResult_Success_DEFAULT *GetItemStockResp

func (p *StockServiceGetItemStockResult) GetSuccess() (v *GetItemStockResp) {
	if !p.IsSetSuccess() {
		return StockServiceGetItemStockResult_Success_DEFAULT
	}
	return p.Success
}
func (p *StockServiceGetItemStockResult) SetSuccess(x interface{}) {
	p.Success = x.(*GetItemStockResp)
}

var fieldIDToName_StockServiceGetItemStockResult = map[int16]string{
	0: "success",
}

func (p *StockServiceGetItemStockResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *StockServiceGetItemStockResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_StockServiceGetItemStockResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *StockServiceGetItemStockResult) ReadField0(iprot thrift.TProtocol) error {
	_field := NewGetItemStockResp()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Success = _field
	return nil
}

func (p *StockServiceGetItemStockResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("GetItemStock_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *StockServiceGetItemStockResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *StockServiceGetItemStockResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("StockServiceGetItemStockResult(%+v)", *p)

}

func (p *StockServiceGetItemStockResult) DeepEqual(ano *StockServiceGetItemStockResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *StockServiceGetItemStockResult) Field0DeepEqual(src *GetItemStockResp) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}

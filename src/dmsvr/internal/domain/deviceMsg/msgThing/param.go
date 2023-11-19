package msgThing

import (
	"encoding/json"
	"fmt"
	"github.com/i-Things/things/shared/domain/schema"
	"github.com/i-Things/things/shared/errors"
	"github.com/spf13/cast"
	"math"
)

const (
	//是否校验参数范围
	validateDataRange = true
)

type Param struct {
	Identifier string              `json:"identifier"` //标识符
	Name       string              `json:"name"`       //功能名称
	Desc       string              `json:"gesc"`       //描述
	Mode       schema.PropertyMode `json:"mode"`       //读写乐行:rw(可读可写) r(只读)
	Required   bool                `json:"required"`   //是否必须
	Type       schema.EventType    `json:"type"`       //事件类型: 信息:info  告警alert  故障:fault
	Value      struct {
		Type    schema.DataType   `json:"type"`              //参数类型:bool int string struct float timestamp array enum
		Mapping map[string]string `json:"mapping,omitempty"` //枚举及bool类型:bool enum
		Min     string            `json:"min,omitempty"`     //数值最小值:int string float
		Max     string            `json:"max,omitempty"`     //数值最大值:int string float
		Start   string            `json:"start,omitempty"`   //初始值:int float
		Step    string            `json:"step,omitempty"`    //步长:int float
		Unit    string            `json:"unit,omitempty"`    //单位:int float
		Value   any               `json:"value"`
		/*
			读到的数据  如果是是数组则类型为[]interface{}  如果是结构体类型则为map[id]Param
				interface 为数据内容  					string为结构体的key value 为数据内容
		*/
	} `json:"value"` //数据定义
}

func (tp *Param) SetByDefine(d *schema.Define, val any) (err error) {
	tp.Value.Type = d.Type
	tp.Value.Type = d.Type
	tp.Value.Mapping = make(map[string]string)
	for k, v := range d.Mapping {
		tp.Value.Mapping[k] = v
	}
	tp.Value.Mapping = d.Mapping
	tp.Value.Min = d.Min
	tp.Value.Max = d.Max
	tp.Value.Start = d.Start
	tp.Value.Step = d.Step
	tp.Value.Unit = d.Unit
	tp.Value.Value, err = GetVal(d, val)
	return err
}

func ToVal(tp map[string]Param) (map[string]any, error) {
	ret := make(map[string]any, len(tp))
	var err error
	for k, v := range tp {
		ret[k], err = v.ToVal()
		if err != nil {
			return nil, err
		}
	}
	return ret, nil
}

func (tp *Param) ToVal() (any, error) {
	if tp == nil {
		return nil, errors.Parameter.AddMsgf("Param is nil")
	}
	var err error
	switch tp.Value.Type {
	case schema.DataTypeStruct:
		v, ok := tp.Value.Value.(map[string]Param)
		if ok == false {
			return nil, errors.Parameter.AddMsgf("struct Param is not find")
		}
		val := make(map[string]any, len(v)+1)
		for _, tp := range v {
			val[tp.Identifier], err = tp.ToVal()
			if err != nil {
				return nil, err
			}
		}
		return val, nil
	case schema.DataTypeArray:
		array, ok := tp.Value.Value.([]any)
		if ok == false {
			return nil, errors.Parameter.AddMsgf("array Param is not find")
		}
		val := make([]any, 0, len(array)+1)
		for _, value := range array {
			switch value.(type) {
			case map[string]Param:
				valMap := make(map[string]any, len(array)+1)
				for _, tp := range value.(map[string]Param) {
					valMap[tp.Identifier], err = tp.ToVal()
					if err != nil {
						return nil, err
					}
				}
				val = append(val, valMap)
			default:
				val = append(val, value)
			}
		}
		return val, nil
	default:
		return tp.Value.Value, nil
	}
}

func GetVal(d *schema.Define, val any) (any, error) {
	switch d.Type {
	case schema.DataTypeBool:
		return cast.ToBoolE(val)
	case schema.DataTypeInt:
		if num, err := cast.ToInt64E(val); err != nil {
			return nil, errors.Parameter.AddDetail(val)
		} else {
			if validateDataRange && (num > cast.ToInt64(d.Max) || num < cast.ToInt64(d.Min)) {
				return nil, errors.OutRange.AddDetailf("value %v out of range:[%s,%s]", val, d.Max, d.Min)
			}
			step := cast.ToInt64(d.Step)
			if step != 0 {
				num = num / step * step
			}
			return num, nil
		}
	case schema.DataTypeFloat:
		if num, err := cast.ToFloat64E(val); err != nil {
			return nil, errors.Parameter.AddDetail(val)
		} else {
			if validateDataRange && (num > cast.ToFloat64(d.Max) || num < cast.ToFloat64(d.Min)) {
				return nil, errors.OutRange.AddDetailf(
					"value %v out of range:[%s,%s]", val, d.Max, d.Min)
			}
			step := cast.ToFloat64(d.Step)
			if step != 0 && !math.IsNaN(step) && !math.IsInf(step, 0) {
				num = math.Floor(num/step) * step
			}
			return num, nil
		}
	case schema.DataTypeString:
		if str, ok := val.(string); !ok {
			return nil, errors.Parameter.AddDetail(val)
		} else {
			if validateDataRange && (len(str) > cast.ToInt(d.Max)) {
				return nil, errors.OutRange.AddDetailf("value %v out of range:%s", val, d.Max)
			}
			return str, nil
		}
	case schema.DataTypeEnum: //枚举类型 报文中传递的是数字
		if num, ok := val.(json.Number); !ok {
			return nil, errors.Parameter.AddDetail(val)
		} else {
			ret, err := num.Int64()
			if err != nil {
				return nil, errors.Parameter.AddDetail(val)
			}
			_, ok := d.Mapping[string(num)]
			if !ok {
				return nil, errors.OutRange.AddDetailf("value %v not in enum", val)
			}
			return ret, nil
		}
	case schema.DataTypeTimestamp:
		switch val.(type) {
		case json.Number:
			ret, err := val.(json.Number).Int64()
			if err != nil {
				return nil, errors.Parameter.AddDetail(val)
			}
			return ret, nil
		case string:
			ret, err := cast.ToInt64E(val)
			if err != nil {
				return nil, errors.Parameter.AddDetail(val)
			}
			return ret, nil
		}
	case schema.DataTypeStruct:
		if strut, ok := val.(map[string]any); !ok {
			return nil, errors.Parameter.AddDetail(val)
		} else {
			getParam := make(map[string]Param, len(strut))
			for k, v := range strut {
				sv, ok := d.Spec[k]
				if ok == false {
					continue
				}
				tp := Param{
					Identifier: sv.Identifier,
					Name:       sv.Name,
				}
				err := tp.SetByDefine(&sv.DataType, v)
				if err == nil {
					getParam[k] = tp
				} else if !errors.Cmp(err, errors.NotFind) {
					return nil, errors.Fmt(err).AddDetail(sv.Identifier)
				}
			}
			return getParam, nil
		}
	case schema.DataTypeArray:
		if arr, ok := val.([]any); !ok {
			return nil, errors.Parameter.AddDetail(fmt.Sprint(val))
		} else {
			if len(arr) == 0 {
				return d, errors.NotFind
			}
			getParam := make([]any, 0, len(arr)+1)
			for _, v := range arr {
				param, err := GetVal(d.ArrayInfo, v)
				if err == nil {
					getParam = append(getParam, param)
				} else if !errors.Cmp(err, errors.NotFind) {
					return nil, errors.Fmt(err).AddDetail(fmt.Sprint(v))
				}
			}
			return getParam, nil
		}
	}
	return nil, errors.Parameter.AddDetail("need param")
}

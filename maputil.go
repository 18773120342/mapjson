package mapjson

import "strconv"

type Json struct {
}

func NewJson() IfaceJson {
	return &Json{}
}
func (this *Json) SetJson(data map[string]interface{}, values ...interface{}) {
	if data == nil {
		return
	}
	num := len(values)
	if num < 2 {
		return
	}
	if num > 2 {
		for k := 0; k < num-2; k++ {
			v := data[values[k].(string)]
			switch c := v.(type) {
			case map[string]interface{}:
				data = c
			default:
				f := make(map[string]interface{})
				data[values[k].(string)] = f
				data = f
			}
		}
	}
	data[values[num-2].(string)] = values[num-1]
}

func (this *Json) GetJsonString(data map[string]interface{}, values ...interface{}) string {
	if data == nil {
		return ""
	}
	// brsm x y
	num := len(values)
	if num < 1 {
		return ""
	}
	if num > 1 {
		for k := 0; k < num-1; k++ {
			v := data[values[k].(string)]
			switch c := v.(type) {
			case map[string]interface{}:
				data = c
			default:
				return ""
			}
		}
	}
	return TypeForString(data[values[num-1].(string)])
}

func (this *Json) GetJsonStringArr(data map[string]interface{}, values ...interface{}) []string {
	if data == nil {
		return nil
	}
	// brsm x y
	num := len(values)
	if num < 1 {
		return nil
	}
	if num > 1 {
		for k := 0; k < num-1; k++ {
			v := data[values[k].(string)]
			switch c := v.(type) {
			case map[string]interface{}:
				data = c
			default:
				return nil
			}
		}
	}
	return TypeForStringArr(data[values[num-1].(string)])
}

func (this *Json) GetJsonIntArr(data map[string]interface{}, values ...interface{}) []int {
	if data == nil {
		return nil
	}
	// brsm x y
	num := len(values)
	if num < 1 {
		return nil
	}
	if num > 1 {
		for k := 0; k < num-1; k++ {
			v := data[values[k].(string)]
			switch c := v.(type) {
			case map[string]interface{}:
				data = c
			default:
				return nil
			}
		}
	}
	return TypeForIntArr(data[values[num-1].(string)])
}
func (this *Json) GetJsonObject(data map[string]interface{}, values ...interface{}) map[string]interface{} {
	if data == nil {
		return nil
	}
	// brsm x y
	num := len(values)
	if num < 1 {
		return nil
	}
	if num > 1 {
		for k := 0; k < num-1; k++ {
			v := data[values[k].(string)]
			switch c := v.(type) {
			case map[string]interface{}:
				data = c
			default:
				return nil
			}
		}
	}
	return TypeForMap(data[values[num-1].(string)])
}

func (this *Json) GetJsonInt8(data map[string]interface{}, values ...interface{}) int8 {
	// brsm x y
	if data == nil {
		return 0
	}
	num := len(values)
	if num < 1 {
		return 0
	}
	if num > 1 {
		for k := 0; k < num-1; k++ {
			v := data[values[k].(string)]
			switch c := v.(type) {
			case map[string]interface{}:
				data = c
			default:
				return 0
			}
		}
	}
	return TypeForInt8(data[values[num-1].(string)])
}

func (this *Json) GetJsonInt16(data map[string]interface{}, values ...interface{}) int16 {
	// brsm x y
	if data == nil {
		return 0
	}
	num := len(values)
	if num < 1 {
		return 0
	}
	if num > 1 {
		for k := 0; k < num-1; k++ {
			v := data[values[k].(string)]
			switch c := v.(type) {
			case map[string]interface{}:
				data = c
			default:
				return 0
			}
		}
	}
	return TypeForInt16(data[values[num-1].(string)])
}
func (this *Json) GetJsonFloat64(data map[string]interface{}, values ...interface{}) float64 {
	// brsm x y
	if data == nil {
		return 0
	}
	num := len(values)
	if num < 1 {
		return 0
	}
	if num > 1 {
		for k := 0; k < num-1; k++ {
			v := data[values[k].(string)]
			switch c := v.(type) {
			case map[string]interface{}:
				data = c
			default:
				return 0
			}
		}
	}
	return TypeForFloat64(data[values[num-1].(string)])
}
func (this *Json) GetJsonInt(data map[string]interface{}, values ...interface{}) int {
	// brsm x y
	if data == nil {
		return 0
	}
	num := len(values)
	if num < 1 {
		return 0
	}
	if num > 1 {
		for k := 0; k < num-1; k++ {
			v := data[values[k].(string)]
			switch c := v.(type) {
			case map[string]interface{}:
				data = c
			default:
				return 0
			}
		}
	}
	return TypeForInt(data[values[num-1].(string)])
}
func (this *Json) GetJsonInt32(data map[string]interface{}, values ...interface{}) int32 {
	// brsm x y
	if data == nil {
		return 0
	}
	num := len(values)
	if num < 1 {
		return 0
	}
	if num > 1 {
		for k := 0; k < num-1; k++ {
			v := data[values[k].(string)]
			switch c := v.(type) {
			case map[string]interface{}:
				data = c
			default:
				return 0
			}
		}
	}
	return TypeForInt32(data[values[num-1].(string)])
}

func (this *Json) GetJsonInt64(data map[string]interface{}, values ...interface{}) int64 {
	// brsm x y
	if data == nil {
		return 0
	}
	num := len(values)
	if num < 1 {
		return 0
	}
	if num > 1 {
		for k := 0; k < num-1; k++ {
			v := data[values[k].(string)]
			switch c := v.(type) {
			case map[string]interface{}:
				data = c
			default:
				return 0
			}
		}
	}
	return TypeForInt64(data[values[num-1].(string)])
}

func (this *Json) GetJsonIntUin8(data map[string]interface{}, values ...interface{}) uint8 {
	// brsm x y
	if data == nil {
		return 0
	}
	num := len(values)
	if num < 1 {
		return 0
	}
	if num > 1 {
		for k := 0; k < num-1; k++ {
			v := data[values[k].(string)]
			switch c := v.(type) {
			case map[string]interface{}:
				data = c
			default:
				return 0
			}
		}
	}
	return TypeForUint8(data[values[num-1].(string)])
}

func (this *Json) GetJsonIntUin16(data map[string]interface{}, values ...interface{}) uint16 {
	// brsm x y
	if data == nil {
		return 0
	}
	num := len(values)
	if num < 1 {
		return 0
	}
	if num > 1 {
		for k := 0; k < num-1; k++ {
			v := data[values[k].(string)]
			switch c := v.(type) {
			case map[string]interface{}:
				data = c
			default:
				return 0
			}
		}
	}
	return TypeForUint16(data[values[num-1].(string)])
}
func (this *Json) GetJsonIntUin32(data map[string]interface{}, values ...interface{}) uint32 {
	// brsm x y
	if data == nil {
		return 0
	}
	num := len(values)
	if num < 1 {
		return 0
	}
	if num > 1 {
		for k := 0; k < num-1; k++ {
			v := data[values[k].(string)]
			switch c := v.(type) {
			case map[string]interface{}:
				data = c
			default:
				return 0
			}
		}
	}
	return TypeForUint32(data[values[num-1].(string)])
}
func (this *Json) GetJsonIntUin64(data map[string]interface{}, values ...interface{}) uint64 {
	// brsm x y
	if data == nil {
		return 0
	}
	num := len(values)
	if num < 1 {
		return 0
	}
	if num > 1 {
		for k := 0; k < num-1; k++ {
			v := data[values[k].(string)]
			switch c := v.(type) {
			case map[string]interface{}:
				data = c
			default:
				return 0
			}
		}
	}
	return TypeForUint64(data[values[num-1].(string)])
}
func (this *Json) GetJsonInterFace(data map[string]interface{}, values ...interface{}) interface{} {
	if data == nil {
		return nil
	}
	num := len(values)
	if num < 1 {
		return 0
	}
	if num > 1 {
		for k := 0; k < num-1; k++ {
			v := data[values[k].(string)]
			switch c := v.(type) {
			case map[string]interface{}:
				data = c
			default:
				return nil
			}
		}
	}
	return data[values[num-1].(string)]
}

//返回一个增加后的值
func (this *Json) SetJsonInc(data map[string]interface{}, values ...interface{}) int32 {
	if data == nil {
		return 0
	}
	if len(values) < 2 {
		return 0
	}
	v1 := TypeForInt32(values[len(values)-1])
	v2 := this.GetJsonInt32(data, values[:len(values)-1]...)
	v := values[:len(values)-1]
	v = append(v, v1+v2)
	this.SetJson(data, v...)
	return v1 + v2

}

//返回一个增加后的值
func (this *Json) SetJsonDec(data map[string]interface{}, values ...interface{}) int32 {
	if data == nil {
		return 0
	}
	if len(values) < 2 {
		return 0
	}
	v1 := TypeForInt32(values[len(values)-1])
	v2 := this.GetJsonInt32(data, values[:len(values)-1]...)
	v := values[:len(values)-1]
	v = append(v, v2-v1)
	this.SetJson(data, v...)
	return v2 - v1
}

func (this *Json) DelJson(data map[string]interface{}, values ...interface{}) {
	if data == nil {
		return
	}
	num := len(values)
	if num < 1 {
		return
	}
	if num > 1 {
		for k := 0; k < num-1; k++ {
			v := data[values[k].(string)]
			switch c := v.(type) {
			case map[string]interface{}:
				data = c
			default:
				return
			}
		}
	}
	delete(data, values[num-1].(string))
}
func TypeForInt8(value interface{}) int8 {
	switch s := value.(type) {
	case float64:
		return int8(s)
	case int8:
		return s
	case int16:
		return int8(s)
	case int32:
		return int8(s)
	case int64:
		return int8(s)
	case uint8:
		return int8(s)
	case uint16:
		return int8(s)
	case uint32:
		return int8(s)
	case uint64:
		return int8(s)
	case float32:
		return int8(s)
	case int:
		return int8(s)
	case uint:
		return int8(s)
	}
	return 0
}
func TypeForInt16(value interface{}) int16 {
	switch s := value.(type) {
	case float64:
		return int16(s)
	case int8:
		return int16(s)
	case int16:
		return s
	case int32:
		return int16(s)
	case int64:
		return int16(s)
	case uint8:
		return int16(s)
	case uint16:
		return int16(s)
	case uint32:
		return int16(s)
	case uint64:
		return int16(s)
	case float32:
		return int16(s)
	case int:
		return int16(s)
	case uint:
		return int16(s)
	}
	return 0
}

func TypeForFloat64(value interface{}) float64 {
	switch s := value.(type) {
	case float64:
		return float64(s)
	case int8:
		return float64(s)
	case int16:
		return float64(s)
	case int32:
		return float64(s)
	case int64:
		return float64(s)
	case uint8:
		return float64(s)
	case uint16:
		return float64(s)
	case uint32:
		return float64(s)
	case uint64:
		return float64(s)
	case float32:
		return float64(s)
	case int:
		return float64(s)
	case uint:
		return float64(s)
	}
	return 0
}
func TypeForInt(value interface{}) int {
	switch s := value.(type) {
	case float64:
		return int(s)
	case int8:
		return int(s)
	case int16:
		return int(s)
	case int32:
		return int(s)
	case int64:
		return int(s)
	case uint8:
		return int(s)
	case uint16:
		return int(s)
	case uint32:
		return int(s)
	case uint64:
		return int(s)
	case float32:
		return int(s)
	case int:
		return int(s)
	case uint:
		return int(s)
	}
	return 0
}
func TypeForInt32(value interface{}) int32 {
	switch s := value.(type) {
	case float64:
		return int32(s)
	case int8:
		return int32(s)
	case int16:
		return int32(s)
	case int32:
		return int32(s)
	case int64:
		return int32(s)
	case uint8:
		return int32(s)
	case uint16:
		return int32(s)
	case uint32:
		return int32(s)
	case uint64:
		return int32(s)
	case float32:
		return int32(s)
	case int:
		return int32(s)
	case uint:
		return int32(s)
	}
	return 0
}
func TypeForInt64(value interface{}) int64 {
	switch s := value.(type) {
	case float64:
		return int64(s)
	case int8:
		return int64(s)
	case int16:
		return int64(s)
	case int32:
		return int64(s)
	case int64:
		return int64(s)
	case uint8:
		return int64(s)
	case uint16:
		return int64(s)
	case uint32:
		return int64(s)
	case uint64:
		return int64(s)
	case float32:
		return int64(s)
	case int:
		return int64(s)
	case uint:
		return int64(s)
	}
	return 0
}

func TypeForUint8(value interface{}) uint8 {
	switch s := value.(type) {
	case float64:
		return uint8(s)
	case int8:
		return uint8(s)
	case int16:
		return uint8(s)
	case int32:
		return uint8(s)
	case int64:
		return uint8(s)
	case uint8:
		return uint8(s)
	case uint16:
		return uint8(s)
	case uint32:
		return uint8(s)
	case uint64:
		return uint8(s)
	case float32:
		return uint8(s)
	case int:
		return uint8(s)
	case uint:
		return uint8(s)
	}
	return 0
}
func TypeForUint16(value interface{}) uint16 {
	switch s := value.(type) {
	case float64:
		return uint16(s)
	case int8:
		return uint16(s)
	case int16:
		return uint16(s)
	case int32:
		return uint16(s)
	case int64:
		return uint16(s)
	case uint8:
		return uint16(s)
	case uint16:
		return uint16(s)
	case uint32:
		return uint16(s)
	case uint64:
		return uint16(s)
	case float32:
		return uint16(s)
	case int:
		return uint16(s)
	case uint:
		return uint16(s)
	}
	return 0
}
func TypeForUint32(value interface{}) uint32 {
	switch s := value.(type) {
	case float64:
		return uint32(s)
	case int8:
		return uint32(s)
	case int16:
		return uint32(s)
	case int32:
		return uint32(s)
	case int64:
		return uint32(s)
	case uint8:
		return uint32(s)
	case uint16:
		return uint32(s)
	case uint32:
		return uint32(s)
	case uint64:
		return uint32(s)
	case float32:
		return uint32(s)
	case int:
		return uint32(s)
	case uint:
		return uint32(s)
	}
	return 0
}
func TypeForString(value interface{}) string {
	switch s := value.(type) {
	case string:
		return s
	}
	return ""
}
func TypeForStringArr(value interface{}) []string {
	switch s := value.(type) {
	case []string:
		return s
	case []interface{}:
		varparamSlice := []string{}
		for _, v := range s {
			switch v1 := v.(type) {
			case string:
				varparamSlice = append(varparamSlice, v1)
			case float64:
				strV := strconv.FormatInt(int64(v1), 10)
				varparamSlice = append(varparamSlice, strV)
			case int:
				strV := strconv.FormatInt(int64(v1), 10)
				varparamSlice = append(varparamSlice, strV)
			case int8:
				strV := strconv.FormatInt(int64(v1), 10)
				varparamSlice = append(varparamSlice, strV)
			case int16:
				strV := strconv.FormatInt(int64(v1), 10)
				varparamSlice = append(varparamSlice, strV)
			case int32:
				strV := strconv.FormatInt(int64(v1), 10)
				varparamSlice = append(varparamSlice, strV)
			case int64:
				strV := strconv.FormatInt(int64(v1), 10)
				varparamSlice = append(varparamSlice, strV)
			case uint:
				strV := strconv.FormatInt(int64(v1), 10)
				varparamSlice = append(varparamSlice, strV)

			case uint8:
				strV := strconv.FormatInt(int64(v1), 10)
				varparamSlice = append(varparamSlice, strV)
			case uint16:
				strV := strconv.FormatInt(int64(v1), 10)
				varparamSlice = append(varparamSlice, strV)
			case uint32:
				strV := strconv.FormatInt(int64(v1), 10)
				varparamSlice = append(varparamSlice, strV)
			case uint64:
				strV := strconv.FormatInt(int64(v1), 10)
				varparamSlice = append(varparamSlice, strV)
			}
		}
		return varparamSlice
	}
	return nil
}
func TypeForIntArr(value interface{}) []int {
	switch s := value.(type) {
	case []int:
		return s
	case []interface{}:
		varparamSlice := []int{}
		for _, v := range s {
			switch v1 := v.(type) {
			case string:
				v2, _ := strconv.Atoi(v1)
				varparamSlice = append(varparamSlice, v2)
			case float64:
				varparamSlice = append(varparamSlice, int(v1))
			case int:
				varparamSlice = append(varparamSlice, int(v1))
			case int8:
				varparamSlice = append(varparamSlice, int(v1))
			case int16:
				varparamSlice = append(varparamSlice, int(v1))
			case int32:
				varparamSlice = append(varparamSlice, int(v1))
			case int64:
				varparamSlice = append(varparamSlice, int(v1))
			case uint:
				varparamSlice = append(varparamSlice, int(v1))
			case uint8:
				varparamSlice = append(varparamSlice, int(v1))
			case uint16:
				varparamSlice = append(varparamSlice, int(v1))
			case uint32:
				varparamSlice = append(varparamSlice, int(v1))
			case uint64:
				varparamSlice = append(varparamSlice, int(v1))
			}
		}
		return varparamSlice
	}
	return nil
}
func TypeForMap(value interface{}) map[string]interface{} {
	switch s := value.(type) {
	case map[string]interface{}:
		arr := make(map[string]interface{})
		for k, v := range s {
			arr[k] = v
		}
		return arr
	}
	return nil
}
func TypeForUint64(value interface{}) uint64 {
	switch s := value.(type) {
	case float64:
		return uint64(s)
	case int8:
		return uint64(s)
	case int16:
		return uint64(s)
	case int32:
		return uint64(s)
	case int64:
		return uint64(s)
	case uint8:
		return uint64(s)
	case uint16:
		return uint64(s)
	case uint32:
		return uint64(s)
	case uint64:
		return uint64(s)
	case float32:
		return uint64(s)
	case int:
		return uint64(s)
	case uint:
		return uint64(s)
	}
	return 0
}

package mapjson

type IfaceJson interface {
	SetJson(data map[string]interface{}, values ...interface{})
	GetJsonString(data map[string]interface{}, values ...interface{}) string
	GetJsonObject(data map[string]interface{}, values ...interface{}) map[string]interface{}
	GetJsonInt8(data map[string]interface{}, values ...interface{}) int8
	GetJsonInt16(data map[string]interface{}, values ...interface{}) int16
	GetJsonFloat64(data map[string]interface{}, values ...interface{}) float64
	GetJsonInt32(data map[string]interface{}, values ...interface{}) int32
	GetJsonInt64(data map[string]interface{}, values ...interface{}) int64
	GetJsonIntUin8(data map[string]interface{}, values ...interface{}) uint8
	GetJsonIntUin16(data map[string]interface{}, values ...interface{}) uint16
	GetJsonIntUin32(data map[string]interface{}, values ...interface{}) uint32
	GetJsonIntUin64(data map[string]interface{}, values ...interface{}) uint64
	SetJsonInc(data map[string]interface{}, values ...interface{}) int32
	SetJsonDec(data map[string]interface{}, values ...interface{}) int32
	GetJsonInt(data map[string]interface{}, values ...interface{}) int
	GetJsonInterFace(data map[string]interface{}, values ...interface{}) interface{}
	DelJson(data map[string]interface{}, values ...interface{})
	GetJsonStringArr(data map[string]interface{}, values ...interface{}) []string
}

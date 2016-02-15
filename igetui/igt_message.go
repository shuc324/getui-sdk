package igetui

import "med3-go-worker/lib/getui/igetui/template"

type IGtMessage struct {
	IsOffline         bool
	OfflineExpireTime int32
	Data              template.ITemplate
	PushNetWorkType  byte
	Priority     int32
	
}

type IGtSingleMessage struct {
	IGtMessage
	
}

func NewIGtSingleMessage(isoffline bool, offlineexpiretime int32, templatee template.ITemplate) *IGtSingleMessage {
	return &IGtSingleMessage{
		IGtMessage: IGtMessage{
			IsOffline:         isoffline,
			OfflineExpireTime: offlineexpiretime,
			Data:              templatee,	
		},
	}
}

type IGtListMessage struct {
	IGtMessage
}

func NewIGtListMessage(isoffline bool, offlineexpiretime int32, templatee template.ITemplate) *IGtListMessage {
	return &IGtListMessage{
		IGtMessage: IGtMessage{
			IsOffline:         isoffline,
			OfflineExpireTime: offlineexpiretime,
			Data:              templatee,
		},
	}
}

type IGtAppMessage struct {
	IGtMessage	
	Speed  int32
	AppIdList     []string
	PhoneTypeList []string
	ProvinceList  []string
}

func NewIGtAppMessage(isoffline bool, offlineexpiretime int32, templatee template.ITemplate) *IGtAppMessage {
	return &IGtAppMessage{
		IGtMessage: IGtMessage{
			IsOffline:         isoffline,
			OfflineExpireTime: offlineexpiretime,
			Data:              templatee,		
		},		
	}
}







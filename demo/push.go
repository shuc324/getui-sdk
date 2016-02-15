package demo

import (
	"log"
	"strconv"
	"encoding/json"
	p "getui-sdk/push"
	i "getui-sdk/igetui"
	t "getui-sdk/igetui/template"
	"strings"
	"sync"
	"getui-sdk/util"
)

// 常量
const (
	PUSH_NUM_KEY = "med-msg-svr:push-number:"
	GE_TUI_CONF_PATH string = "../config/getui.yaml"
	IS_OFFLINE_ANDROID, IS_OFFLINE_IOS bool = true, true
	PHONE_TYPE_ANDROID, PHONE_TYPE_IOS string = "ANDROID", "IOS"
	PUSH_NETWORK_TYPE_ANDROID, PUSH_NETWORK_TYPE_IOS byte = 0, 0
	PUSH_MESSAGE_SPEED_ANDROID, PUSH_MESSAGE_SPEED_IOS int32 = 1000, 1000
	PUSH_SINGLE, PUSH_GROUP_LIMIT_MIN, PUSH_GROUP_LIMIT_MAX int = 1, 2, 10
	OFFLINE_EXPIRE_TIME_ANDROID, OFFLINE_EXPIRE_TIME_IOS int32 = 300 * 12, 300 * 12
)

// 个推配置
type GeTuiConf struct {
	Host         string `yaml:"host"`
	AppId        string `yaml:"app_id"`
	AppKey       string `yaml:"app_key"`
	MasterSecret string `yaml:"master_secret"`
}

// push包字体段
type PushField struct {
	Id       int `json:"id"`
	TargetId int `json:"targetId"`
	Extra    string `json:"extra"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	ClientId string `json:"-"`
	PlatFrom string `json:"-"`
}

// push包转json
func (p *PushField) toString() string {
	bytes, _ := json.Marshal(p)
	return string(bytes)
}

// 取得个推配置
var GetGeTuiConf = func() GeTuiConf {
	GeTuiConf := GeTuiConf{}
	err := util.YamlFileDecode(GE_TUI_CONF_PATH, &GeTuiConf)
	if err != nil {
		panic(err)
	}
	return GeTuiConf
}

// 线程锁
var wg sync.WaitGroup

// 梳理Push
var hecklePush = func(pushList []PushField) ([]PushField, []PushField) {
	iosPushList, androidPushList := []PushField{}, []PushField{}
	for _, push := range pushList {
		switch clientType := strings.ToUpper(push.PlatFrom); clientType {
		case PHONE_TYPE_IOS:
			iosPushList = append(iosPushList, push)
		case PHONE_TYPE_ANDROID:
			androidPushList = append(androidPushList, push)
		}
	}
	return iosPushList, androidPushList
}

// ios 广播
var PushBroadCastToIos = func(broadCast []PushField) {
	GeTuiConf := GetGeTuiConf()
	push := p.NewIGeTui(GeTuiConf.Host, GetGeTuiConf().AppKey, GetGeTuiConf().MasterSecret, )
	template := t.NewTransmissionTemplate(GeTuiConf.AppId, GeTuiConf.AppKey, 2, "医联", )
	broadCastMessage := i.NewIGtAppMessage(IS_OFFLINE_IOS, OFFLINE_EXPIRE_TIME_IOS, template, )
	broadCastMessage.PushNetWorkType = PUSH_NETWORK_TYPE_IOS
	template.SetPushInfo("", "1", broadCast[0].Title, "", broadCast[0].toString(), "", "", "", 1)
	broadCastMessage.Speed = PUSH_MESSAGE_SPEED_IOS
	broadCastMessage.AppIdList = append(broadCastMessage.AppIdList, GeTuiConf.AppId)
	broadCastMessage.PhoneTypeList = append(broadCastMessage.PhoneTypeList, PHONE_TYPE_IOS)
	broadCastRet := push.PushMessageToApp(*broadCastMessage)
	if broadCastRet["result"] != nil {
		log.Printf("push broadcast to ios failed: %s", broadCast[0].toString())
	} else {
		log.Printf("push broadcast to ios successfully")
	}
}

// android 广播
var PushBroadCastToAndroid = func(broadCast []PushField) {
	GeTuiConf := GetGeTuiConf()
	push := p.NewIGeTui(GeTuiConf.Host, GeTuiConf.AppKey, GeTuiConf.MasterSecret, )
	template := t.NewTransmissionTemplate(GeTuiConf.AppId, GeTuiConf.AppKey, 2, "医联", )
	broadCastMessage := i.NewIGtAppMessage(IS_OFFLINE_ANDROID, OFFLINE_EXPIRE_TIME_ANDROID, template, )
	broadCastMessage.PushNetWorkType = PUSH_NETWORK_TYPE_ANDROID
	template.SetPushInfo("", "1", broadCast[0].Title, "", broadCast[0].toString(), "", "", "", 1)
	broadCastMessage.Speed = PUSH_MESSAGE_SPEED_ANDROID
	broadCastMessage.AppIdList = append(broadCastMessage.AppIdList, GeTuiConf.AppId)
	broadCastMessage.PhoneTypeList = append(broadCastMessage.PhoneTypeList, PHONE_TYPE_ANDROID)
	broadCastRet := push.PushMessageToApp(*broadCastMessage)
	if broadCastRet["result"] != nil {
		log.Printf("push broadcast to android failed: %s", broadCast[0].toString())
	} else {
		log.Printf("push broadcast to android successfully")
	}
}

// ios 发送push
var PushMessageToIos = func(pushList []PushField) {
	defer wg.Done()
	GeTuiConf := GetGeTuiConf()
	push := p.NewIGeTui(GeTuiConf.Host, GeTuiConf.AppKey, GeTuiConf.MasterSecret, )
	template := t.NewTransmissionTemplate(GeTuiConf.AppId, GeTuiConf.AppKey, 2, "医联", )

	// 单推
	if isSingle := len(pushList); isSingle == PUSH_SINGLE {
		singleMessage := i.NewIGtSingleMessage(IS_OFFLINE_IOS, OFFLINE_EXPIRE_TIME_IOS, template, )
		singleMessage.PushNetWorkType = PUSH_NETWORK_TYPE_IOS
		// 单推取出push条数
		//join := []string{PUSH_NUM_KEY, strconv.Itoa(pushList[0].TargetId)}
		//pushNumStr, _ := Redis.Chat.Get(strings.Join(join, "")).Result()
		//pushNum, _ := strconv.Atoi(pushNumStr)
		//log.Println(strconv.Itoa(pushNum + 1))
		template.SetPushInfo("", strconv.Itoa(10 + 1), pushList[0].Title, "", pushList[0].toString(), "", "", "", 1)
		target := i.NewTarget(GeTuiConf.AppId, "", "", )
		singleRet := push.PushAPNMessageToSingle(*singleMessage, *target, pushList[0].ClientId)
		if singleRet["result"] != "ok" {
			log.Printf("has push failed to ios user: %s -> %s", singleRet["result"], pushList[0].toString())
		} else {
			log.Printf("has push successfully to ios user: %s", pushList[0].toString())
		}
	}

	// 群推
	if isList := len(pushList); isList > PUSH_GROUP_LIMIT_MAX {
		log.Printf("push the number of more than %d", PUSH_GROUP_LIMIT_MAX)
	} else if isList >= PUSH_GROUP_LIMIT_MIN {
		listMessage := i.NewIGtListMessage(IS_OFFLINE_IOS, OFFLINE_EXPIRE_TIME_IOS, template, )
		listMessage.PushNetWorkType = PUSH_NETWORK_TYPE_IOS
		// 群推无条数,badge置为0 添加推送记录
		template.SetPushInfo("", "0", pushList[0].Title, "", pushList[0].toString(), "", "", "", 1)
		// 此处用到起了redis记数
		//for _, push := range pushList {
		//	pushNumKey := []string{PUSH_NUM_KEY, strconv.Itoa(push.TargetId)}
		//	Redis.Chat.Incr(strings.Join(pushNumKey, ""))
		//}
		deviceTokenList := make([]string, len(pushList))
		for _, singlePush := range pushList {
			deviceTokenList = append(deviceTokenList, singlePush.ClientId)
		}
		contentId := push.GetAPNContentId(GeTuiConf.AppId, *listMessage)
		listRet := push.PushAPNMessageToList(GeTuiConf.AppId, contentId.(string), deviceTokenList)
		bytes, _ := json.Marshal(pushList)
		if listRet["result"] != "ok" {
			log.Printf("has push failed to ios users : %s -> %s", listRet["result"], string(bytes))
		} else {
			log.Printf("has push successfully to ios users: %s", string(bytes))
		}
	}
}

// android 发送push
var PushMessageToAndroid = func(pushList []PushField) {
	defer wg.Done()
	GeTuiConf := GetGeTuiConf()
	push := p.NewIGeTui(GeTuiConf.Host, GeTuiConf.AppKey, GeTuiConf.MasterSecret, )
	template := t.NewTransmissionTemplate(GeTuiConf.AppId, GeTuiConf.AppKey, 2, pushList[0].toString(), )

	// 单推
	if isSingle := len(pushList); isSingle == PUSH_SINGLE {
		singleMessage := i.NewIGtSingleMessage(IS_OFFLINE_ANDROID, OFFLINE_EXPIRE_TIME_ANDROID, template, )
		singleMessage.PushNetWorkType = PUSH_NETWORK_TYPE_ANDROID
		template.SetPushInfo("", strconv.Itoa(1), pushList[0].Title, "", pushList[0].toString(), "", "", "", 1)
		target := i.NewTarget(GeTuiConf.AppId, pushList[0].ClientId, "", )
		singleRet := push.PushMessageToSingle(*singleMessage, *target)
		if singleRet["result"] != "ok" {
			log.Printf("push failed to android user: %s -> %s", singleRet["result"], pushList[0].toString())
		} else {
			log.Printf("has push successfully to android user: %s", pushList[0].toString())
		}
	}

	// 群推
	if isList := len(pushList); isList > PUSH_GROUP_LIMIT_MAX {
		log.Printf("push the number of more than %d", PUSH_GROUP_LIMIT_MAX)
	} else if isList >= PUSH_GROUP_LIMIT_MIN {
		listMessage := i.NewIGtListMessage(IS_OFFLINE_ANDROID, OFFLINE_EXPIRE_TIME_ANDROID, template, )
		listMessage.PushNetWorkType = PUSH_NETWORK_TYPE_ANDROID
		template.SetPushInfo("", "0", pushList[0].Title, "", pushList[0].toString(), "", "", "", 1)
		targetList := []i.Target{}
		for _, singlePush := range pushList {
			target := i.NewTarget(GeTuiConf.AppId, singlePush.ClientId, "", )
			targetList = append(targetList, *target)
		}
		contentId := push.GetContentId(*listMessage)
		listRet := push.PushMessageToList(contentId.(string), targetList)
		bytes, _ := json.Marshal(pushList)
		if listRet["result"] != "ok" {
			log.Printf("push failed to android users: %s -> %s", listRet["result"], string(bytes))
		} else {
			log.Printf("has push successfully to android users: %s", string(bytes))
		}
	}
}

var distributePush = func(list []PushField) {
	groupNum := len(list) / PUSH_GROUP_LIMIT_MAX
	phoneType := strings.ToUpper(list[0].PlatFrom)
	for times := 0; times <= groupNum; times++ {
		wg.Add(1)
		sliceMin := times * PUSH_GROUP_LIMIT_MAX
		sliceMax := (times + 1) * PUSH_GROUP_LIMIT_MAX
		if times == groupNum {
			if phoneType == PHONE_TYPE_IOS {
				go PushMessageToIos(list[sliceMin : ])
			} else {
				go PushMessageToAndroid(list[sliceMin : ])
			}
		} else {
			if phoneType == PHONE_TYPE_IOS {
				go PushMessageToIos(list[sliceMin : sliceMax])
			} else {
				go PushMessageToAndroid(list[sliceMin : sliceMax])
			}
		}
	}
	wg.Wait()
}

// 发送push/广播
var SendPush = func(pushList []PushField, isBroadCast bool) {
	iosPushList, androidPushList := hecklePush(pushList)
	// push
	if isBroadCast == false {
		if ios := len(iosPushList); ios != 0 {
			distributePush(iosPushList)
		}
		if android := len(androidPushList); android != 0 {
			distributePush(androidPushList)
		}
	}
	// 广播
	if isBroadCast == true {
		// todo
	}
}

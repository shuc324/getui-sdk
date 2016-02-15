package template

import "getui-sdk/protobuf"
import proto "github.com/golang/protobuf/proto"

type LinkTemplate struct {
	//AppId               string
	//AppKey              string
	BaseTemplate
	Text                string
	Title               string
	Logo                string
	Url                 string
	IsRing              bool
	IsVibrate           bool
	IsClearable         bool
	PushType            string
	TransmissionType    int32
	TransmissionContent string
}


func NewLinkTemplate(appid string, appkey string, transmissiontype int32,
	transmissionconntent string, titile string, text string, logo string, url string, isring bool,
	isvibrate bool, isclearable bool) *LinkTemplate {
	return &LinkTemplate{
		BaseTemplate:BaseTemplate{
			AppId: appid,
		    AppKey: appkey,
			},
		TransmissionType:    transmissiontype,
		TransmissionContent: transmissionconntent,
		Title:               titile,
		Text:                text,
		Logo:                logo,
		Url:                 url,
		IsRing:              isring,
		IsVibrate:           isvibrate,
		IsClearable:         isclearable,
		PushType:            "NotifyMsg",
	}
}


func (t *LinkTemplate) GetTransmissionContent() string {
	return t.TransmissionContent
}

func (t *LinkTemplate) GetPushType() string {
	return t.PushType
}



func (t *LinkTemplate) GetTransparent() *protobuf.Transparent {
	transparent := &protobuf.Transparent{
		Id:          proto.String(""),
		Action:      proto.String("pushmessage"),
		TaskId:      proto.String(""),
		AppKey:      proto.String(t.AppKey),
		AppId:       proto.String(t.AppId),
		MessageId:   proto.String(""),
		PushInfo:    t.GetPushInfo(),
		ActionChain: t.GetActionChains(),
		Condition :  t.GetDurCondition(),
	}
	return transparent
}

func (t *LinkTemplate) GetActionChains() []*protobuf.ActionChain {

	//set actionChain
	actionChain1 := &protobuf.ActionChain{
		ActionId: proto.Int32(1),
		Type:     protobuf.ActionChain_Goto.Enum(),
		Next:     proto.Int32(10000),
	}

	//start up app
	actionChain2 := &protobuf.ActionChain{
		ActionId:  proto.Int32(10000),
		Type:      protobuf.ActionChain_notification.Enum(),
		Title:     proto.String(t.Title),
		Text:      proto.String(t.Text),
		Logo:      proto.String(t.Logo),
		Ring:      proto.Bool(t.IsRing),
		Clearable: proto.Bool(t.IsClearable),
		Buzz:      proto.Bool(t.IsVibrate),
		Next:      proto.Int32(10010),
	}

	//goto
	actionChain3 := &protobuf.ActionChain{
		ActionId: proto.Int32(10010),
		Type:     protobuf.ActionChain_Goto.Enum(),
		Next:     proto.Int32(10030),
	}

	//start web
	actionChain4 := &protobuf.ActionChain{
		ActionId: proto.Int32(10030),
		Type:     protobuf.ActionChain_startweb.Enum(),
		Url:      proto.String(t.Url),
		Next:     proto.Int32(100),
	}
	//end
	actionChain5 := &protobuf.ActionChain{
		ActionId: proto.Int32(100),
		Type:     protobuf.ActionChain_eoa.Enum(),
	}

	actionChains := []*protobuf.ActionChain{actionChain1, actionChain2, actionChain3, actionChain4, actionChain5}

	return actionChains
}

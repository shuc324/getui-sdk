package template

import "getui-sdk/protobuf"
import proto "github.com/golang/protobuf/proto"

type NotyPopLoadTemplate struct {
	//AppId               string
	//AppKey              string
	BaseTemplate
	PushType            string
	TransmissionType    int32
	TransmissionContent string	
	  NotyIcon string
      LogoUrl string
      NotyTitle string
      NotyContent string
      IsRing bool //true
      IsVibrate bool //true
      IsClearable bool //true
      PopTitle string
      PopContent string
      PopImage string
      PopButton1 string
      PopButton2 string
      LoadIcon string
      LoadTitle string
      LoadUrl string
      IsAutoInstall bool //False
      IsActive bool //False
      AndroidMark string
      SymbianMark string
      IosMark  string
      
}

func NewNotyPopLoadTemplate (appid string, appkey string, transmissiontype int32,
	transmissionconntent string, notyicon string, notytitle string, notycontent string, loadicon string, loadurl string, loadtitle string,isring bool,
	isvibrate bool, isclearable bool,poptitle string,popcontent string,popimage string,popbutton1 string,popbutton2 string,
	isautoinstall bool,isactive bool,andriodmark string,sybianmark string,iosmark string) *NotyPopLoadTemplate {
		
	return &NotyPopLoadTemplate{
		BaseTemplate:BaseTemplate{
			AppId: appid,
		    AppKey: appkey,
			},
		TransmissionType:    transmissiontype,
		TransmissionContent: transmissionconntent,
		NotyIcon : notyicon,
      NotyTitle :notytitle,
      NotyContent :notycontent,
	  PopTitle : poptitle,
      PopContent :popcontent,
      PopImage :popimage,
      PopButton1 :popbutton1,
      PopButton2 :popbutton2,
      LoadIcon :loadicon,
      LoadTitle :loadtitle,
      LoadUrl :loadurl,
      IsAutoInstall :isautoinstall, //False
      IsActive :isactive, //False
      AndroidMark :andriodmark,
      SymbianMark :sybianmark,
      IosMark  : iosmark,
		IsRing:              isring,
		IsVibrate:           isvibrate,
		IsClearable:         isclearable,
		PushType:            "NotyPopLoadMsg",
	}
}


func (t *NotyPopLoadTemplate) GetTransmissionContent() string {
	return t.TransmissionContent
}

func (t *NotyPopLoadTemplate) GetPushType() string {
	return t.PushType
}

func (t *NotyPopLoadTemplate) GetTransparent() *protobuf.Transparent {
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




func (t *NotyPopLoadTemplate) GetActionChains() []*protobuf.ActionChain {

	//set actionChain
	actionChain1 := &protobuf.ActionChain{
		ActionId: proto.Int32(1),
		Type:     protobuf.ActionChain_Goto.Enum(),
		Next:     proto.Int32(10000),
	}

	//notification
	actionChain2 := &protobuf.ActionChain{
		ActionId:  proto.Int32(10000),
		Type:      protobuf.ActionChain_notification.Enum(),
		Title:     proto.String(t.NotyTitle),
		Text:      proto.String(t.NotyContent),
		Logo:      proto.String(t.NotyIcon),
		LogoURL:   proto.String(t.LogoUrl),
		Ring:      proto.Bool(t.IsRing),
		Clearable: proto.Bool(t.IsClearable),
		Buzz:      proto.Bool(t.IsVibrate),
		Next:      proto.Int32(10010),
	}

	//goto
	actionChain3 := &protobuf.ActionChain{
		ActionId: proto.Int32(10010),
		Type:     protobuf.ActionChain_Goto.Enum(),
		Next:     proto.Int32(10020),
	}

	//start web
	button1 := &protobuf.Button{
		Text : proto.String(t.PopButton1),
		Next : proto.Int32(10040),	
		}
	button2 := &protobuf.Button{
		Text : proto.String(t.PopButton2),
		Next : proto.Int32(100),
			}
	actionChain4 := &protobuf.ActionChain{
		
		
		//proto.		
		ActionId: proto.Int32(10020),
		Buttons : []*protobuf.Button{button1,button2},
		Type:     protobuf.ActionChain_popup.Enum(),
		Title: proto.String(t.PopTitle),
		Text: proto.String(t.PopContent),
		Img:proto.String(t.PopImage),
		Next:     proto.Int32(6),
	}
	
	AppStartUp := &protobuf.AppStartUp{
		Android: proto.String(t.AndroidMark),
		Symbia: proto.String(t.SymbianMark),
		Ios: proto.String(t.IosMark),
	}
	
	
	//end
	actionChain5 := &protobuf.ActionChain{
		ActionId: proto.Int32(10040),
		Type:     protobuf.ActionChain_appdownload.Enum(),
		Name:    proto.String(t.LoadTitle),
		Url:proto.String(t.LoadUrl),
		Logo:proto.String(t.LoadIcon),
		AutoInstall:proto.Bool(t.IsAutoInstall),
		Autostart:proto.Bool(t.IsActive),
		Appstartupid: AppStartUp,
		Next:proto.Int32(6),
	}
	
	actionChain6 := &protobuf.ActionChain{
		ActionId: proto.Int32(100),
		Type:     protobuf.ActionChain_eoa.Enum(),
	}

	actionChains := []*protobuf.ActionChain{actionChain1, actionChain2, actionChain3, actionChain4, actionChain5,actionChain6 }

	return actionChains
}









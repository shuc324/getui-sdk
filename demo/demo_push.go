package main

import (
	"fmt"
	i "med3-go-worker/lib/getui/igetui"
	t "med3-go-worker/lib/getui/igetui/template"
	p "med3-go-worker/lib/getui/push"
	"time"
)

var CID1 = "4d35ffb9b84a20a490a4cbdfa5f4935f"
var CID2 = "a4d0d140e97872f495728dc5e10888e1"

var APPKEY = "tpDVam96sY8pxhwBupJ462"
var APPID = "aK6jeksP5C7CsjSSEqLAA3"
var MASTERSECRET = "TBokfpttQJ6aHIhBE9y867"
var CID = "8ba9a76c746f34c2ec00b93db30b9d21"
var HOST = "http://192.168.10.61:8006/apiex.htm"

//var HOST = "http://sdk.open.api.igexin.com/apiex.htm"
//var MASTERSECRET = "6bLzdFZWFeAlPrjNms4Aq7"
/*
var APPKEY = "vtDm307Hbk6HN3MG6tN1a6"
var APPID = "DzZ5576WbA6IxM0ytqcZR"
var MASTERSECRET = "rIJR2yCSbg9XlcK60iyq32"
var CID = "424bc8af133b06b0f06364f06c0fda2a"
var HOST = "http://192.168.10.61:80/apiex.htm"
*/
var DEVICETOKEN = "3337de7aa297065657c087a041d28b3c90c9ed51bdc37c58e8d13ced523f5f5f"

var PushAPN = func() {
	push := p.NewIGeTui(
		HOST,         //Host:
		APPKEY,       //AppKey:
		MASTERSECRET, //MasterSecret:
	)
	template := t.NewTransmissionTemplate(
		APPID,
		APPKEY,
		2, //TransmissionType
		"这是一条toApp测试消息", //TransmissionContent
	)
	//template.setPushInfo(actionLocKey, badge, message, sound, payload, locKey, locArgs, launchImage)
	template.SetPushInfo("你好", "4", "message", "test1.wav", "payloads", "hh", "hh", "gg", 2)

	/* ex. PushAPNMessageToSingle*/
	message := i.NewIGtSingleMessage(
		true,     //IsOffline
		300*12,   //OfflineExpireTime
		template, //Data
	)
	message.PushNetWorkType = 0
	target := i.NewTarget(
		APPID, //appId
		CID,   //clientId
		"",    //alias
	)
	ret := push.PushAPNMessageToSingle(*message, *target, DEVICETOKEN)
	fmt.Println(ret)

	/* ex. pushAPNMessageToList

		//template.SetPushInfo("你好","4","message","test1.wav","payloadl","hh","hh","gg",1)
		lmessage := i.NewIGtListMessage(
			true,         //IsOffline:
			72*3600*1000, //OfflineExpireTime:
			template,     //Data
		)
	    lmessage.PushNetWorkType = 0 //not care network
	    contentId := push.GetAPNContentId(APPID,*lmessage)
		deviceTokenList:=[]string{DEVICETOKEN}

		lret := push.PushAPNMessageToList(APPID, contentId.(string),deviceTokenList)
		fmt.Println(lret)
	*/
}

var PushMessageToSingle = func() {
	push := p.NewIGeTui(
		HOST,         //Host:
		APPKEY,       //AppKey:
		MASTERSECRET, //MasterSecret:
	)
	// push.Fast()

	template := t.NewNotificationTemplate(
		APPID,  //appId
		APPKEY, //appKey
		2,      //TransmissionType
		"这是一条toSingleNotifi测试消息", //TransmissionContent
		"igetui",                         //Title
		"click to download",              //Text
		"http://www.igetui.com/logo.png", //Logo
		true, //IsRing
		true, //IsVibrate
		true, //IsClearable
	)

	/*

		template := t.NewLinkTemplate(
			APPID,  //AppId:
			APPKEY, //AppKey:
			2,      //TransmissionType
			"这是一条tosinglelink测试消息", //TransmissionContent
			"igetui",                         //Title
			"click to download",              //Text
			"http://www.igetui.com/logo.png", //Logo
			"http://baidu.com",               //Url
			true,                             //IsRing
			true,                             //IsVibrate
			true,                             //IsClearable
		)

	*/
	/*
		  template := t.NewTransmissionTemplate(
			APPID,
			APPKEY,
			2, //TransmissionType
			"这是一条tosingleTransmission测试消息", //TransmissionContent
		)

	*/
	//template.setPushInfo(actionLocKey, badge, message, sound, payload, locKey, locArgs, launchImage)
	//template.SetPushInfo("你好","4","message","test1.wav","c","singleapple","hh2","gg",1)
	//template.SetDuration("2015-03-31 19:41:12","2015-03-31 19:42:12")
	message := i.NewIGtSingleMessage(
		true,     //IsOffline
		300*12,   //OfflineExpireTime
		template, //Data
	)
	message.PushNetWorkType = 0
	target := i.NewTarget(
		APPID, //appId
		CID,   //clientId
		"",    //alias
	)
	/*
		for i:=0;i<10;i++ {

			ret := push.PushMessageToSingle(*message, *target)
		    fmt.Println(ret,push.Host)
		    time.Sleep(1000*time.Millisecond)
		}*/
	ret := push.PushMessageToSingle(*message, *target)

	fmt.Println(ret)
}

var PushMessageToList = func() {
	push := p.NewIGeTui(
		HOST,         //Host:
		APPKEY,       //AppKey:
		MASTERSECRET, //MasterSecret:
	)
	template := t.NewNotificationTemplate(
		APPID,  //appId
		APPKEY, //appKey
		2,      //TransmissionType
		"这是一条toSingleNotifi测试消息", //TransmissionContent
		"igetui",                         //Title
		"click to download",              //Text
		"http://www.igetui.com/logo.png", //Logo
		true, //IsRing
		true, //IsVibrate
		true, //IsClearable
	)
	/*
		template := t.NewLinkTemplate(
			APPID,  //AppId:
			APPKEY, //AppKey:
			2,      //TransmissionType
			"这是一条tolistlink测试消息", //TransmissionContent
			"igetui",                         //Title
			"click to download",              //Text
			"http://www.igetui.com/logo.png", //Logo
			"http://baidu.com",               //Url
			true,                             //IsRing
			true,                             //IsVibrate
			true,                             //IsClearable
		)
	*/
	/*
	    template := t.NewTransmissionTemplate(
			APPID,
			APPKEY,
			2, //TransmissionType
			"这是一条toApp测试消息", //TransmissionContent

		)
	*/
	template.SetPushInfo("你好", "4", "message", "test1.wav", "c", "listapple", "hh", "gg", 1)
	message := i.NewIGtListMessage(
		true,         //IsOffline:
		72*3600*1000, //OfflineExpireTime:
		template,     //Data
	)
	message.PushNetWorkType = 0

	target1 := i.NewTarget(
		APPID, //AppId:
		CID,   //clientId
		"",    //alias
	)

	target2 := i.NewTarget(
		APPID,
		CID,
		"", //alias
	)
	targets := []i.Target{*target1, *target2}
	contentId := push.GetContentId(*message)
	ret := push.PushMessageToList(contentId.(string), targets)
	fmt.Println(ret)
}

var PushMessageToApp = func() {
	push := p.NewIGeTui(
		HOST,         //Host:
		APPKEY,       //AppKey:
		MASTERSECRET, //MasterSecret:
	)
	template := t.NewNotificationTemplate(
		APPID,  //appId
		APPKEY, //appKey
		2,      //TransmissionType
		"这是一条toSingle测试消息", //TransmissionContent
		"igetui",                         //Title
		"click to download",              //Text
		"http://www.igetui.com/logo.png", //Logo
		true, //IsRing
		true, //IsVibrate
		true, //IsClearable
	)
	//template.SetPushInfo("你好","4","message","test1.wav","c","hh","hh","gg",0)
	message := i.NewIGtAppMessage(
		true,         //IsOffline:
		72*3600*1000, //OfflineExpireTime:
		template,     //Data
	)
	message.PushNetWorkType = 0
	message.AppIdList = append(message.AppIdList, APPID)
	message.PhoneTypeList = append(message.PhoneTypeList, "ANDROID", "IOS")
	//message.ProvinceList = append(message.ProvinceList, "浙江")
	message.Speed = 1000
	ret := push.PushMessageToApp(*message)
	fmt.Printf("%v\n", ret)
	fmt.Printf("%v\n", push.GetClientIdStatus(APPID, CID))
}

var SetClientTagtest = func() {
	push := p.NewIGeTui(
		HOST,         //Host:
		APPKEY,       //AppKey:
		MASTERSECRET, //MasterSecret:
	)
	target := i.NewTarget(
		APPID, //appId
		CID,   //clientId
		"",    //alias
	)
	a := []string{"sdsd2"}
	ret1 := push.SetClientTag(a, *target)
	fmt.Println(ret1)
}

var GetUserTagstest = func() {
	push := p.NewIGeTui(
		HOST,         //Host:
		APPKEY,       //AppKey:
		MASTERSECRET, //MasterSecret:
	)
	target := i.NewTarget(
		APPID, //appId
		CID,   //clientId
		"",    //alias
	)
	ret2 := push.GetUserTags(*target)
	fmt.Println(ret2)
}

var BindAliasTest = func() {
	push := p.NewIGeTui(
		HOST,         //Host:
		APPKEY,       //AppKey:
		MASTERSECRET, //MasterSecret:
	)

	ret3 := push.BindAlias("hhh", APPID, CID1)
	//ret := push.QueryAlias(APPID,CID1)
	ret2 := push.UnBindAlias(APPID, "hhh", CID1)
	//ret4 := push.QueryAlias(APPID,CID1)
	//ret3 := push.QueryClientId("hhh",APPID)
	fmt.Println(ret2, ret3)

}
var BindAliasBatchtest = func() {
	push := p.NewIGeTui(
		HOST,         //Host:
		APPKEY,       //AppKey:
		MASTERSECRET, //MasterSecret:
	)
	target1 := i.NewTarget(
		APPID, //appId
		CID,   //clientId
		"abcd",
	)
	target2 := i.NewTarget(
		APPID, //appId
		CID2,  //clientId
		"abcd",
	)
	targets := []i.Target{*target1, *target2}
	ret := push.BindAliasBatch(APPID, targets)
	fmt.Println(ret, push.QueryAlias(APPID, CID), push.QueryAlias(APPID, CID2))
	ret1 := push.UnBindAliasAll(APPID, "abcd")
	fmt.Println(ret1, push.QueryAlias(APPID, CID), push.QueryAlias(APPID, CID2))
}
var Fasttest = func() {

	push := p.NewIGeTui(
		"",           //Host:
		APPKEY,       //AppKey:
		MASTERSECRET, //MasterSecret:
	)
	push.Fast()
	for i := 0; i < 100; i++ {
		fmt.Println(push.Host)
		time.Sleep(600 * 1000 * time.Millisecond)
	}

}

func main() {
	Fasttest()

	//BindAliasTest()
	//BindAliasBatchtest()
	//SetClientTagtest()
	//GetUserTagstest()

	//PushAPN()
	PushMessageToSingle()
	//PushMessageToList()
	//PushMessageToApp()

}

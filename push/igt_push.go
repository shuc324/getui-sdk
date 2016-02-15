package push

import (
	proto "github.com/golang/protobuf/proto"
	//"container/list"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"getui-sdk/igetui"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type IGeTui struct {
	Host         string
	AppKey       string
	MasterSecret string
	flag bool
}

func NewIGeTui(host, appkey, mastersecret string) *IGeTui {
	return &IGeTui{
		Host:         host,
		AppKey:       appkey,
		MasterSecret: mastersecret,
	}
}

func (iGeTui *IGeTui) Fast(){
		if iGeTui.Host ==""{
		ln := iGeTui.GetConnOSServerHostList()
		if len(ln)==1{
			iGeTui.Host=ln[0]
		}
		if len(ln)>1{
			iGeTui.gethost(ln)
			for ;iGeTui.Host=="";{
	   time.Sleep(time.Millisecond)				
			}
		   go iGeTui.timerhost(ln)
		}	 
	}
}

func (iGeTui *IGeTui) test(url string ) {
    _,err := http.Get(url)
    if err !=nil{
    	return 
    }else if iGeTui.flag == false{
    	iGeTui.Host=url
		iGeTui.flag = true
		}
}

func (iGeTui *IGeTui) gethost(ln []string){
	iGeTui.flag =  false
	for i:=0;i <len(ln);i++ {
		go iGeTui.test(ln[i])
	}
}
 
func (iGeTui *IGeTui) timerhost(ln []string){	
	for {	
		time.Sleep(600*1000*time.Millisecond)
		go iGeTui.gethost(ln)
	}
} 

func (iGeTui *IGeTui) GetConnOSServerHostList() []string {
	l := iGeTui.GetConfigOsServerHostList()
	if l == nil || len(l) == 0 {
		ln := []string{"http://sdk.open.api.igexin.com/serviceex",
			"http://sdk.open.api.gepush.com/serviceex",
			"http://sdk.open.api.getui.net/serviceex",
		}
		return ln
	}
	return l
}

func (iGeTui *IGeTui) GetConfigOsServerHostList() []string {
	url := "http://sdk.open.apilist.igexin.com/os_list"
	response, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	l := strings.Split(string(body), "\r\n")
	var ll []string
	for i := 0; i < len(l); i++ {
		if strings.HasPrefix(l[i], "http") {
			ll = append(ll, l[i])
		}
	}
	return l
}


func (iGeTui *IGeTui) connect() bool {
	sign := iGeTui.GetSign(iGeTui.AppKey, iGeTui.GetCurrentTime(), iGeTui.MasterSecret)
	params := map[string]interface{}{}
	params["action"] = "connect"
	params["appkey"] = iGeTui.AppKey
	params["timeStamp"] = iGeTui.GetCurrentTime()
	params["sign"] = sign

	rep := iGeTui.HttpPost(params)
	fmt.Println("rep")
	fmt.Println(rep)
	if "success" == rep["result"] {
		return true
	} else {
		fmt.Println("connect failed")
		panic("")
	}
	return false
}

func (iGeTui *IGeTui) PushAPNMessageToSingle(message igetui.IGtSingleMessage, tartget igetui.Target, deviceToken string) map[string]interface{} {
	params := map[string]interface{}{}
	params["action"] = "apnPushToSingleAction"
	params["appkey"] = iGeTui.AppKey
	params["appId"] = tartget.AppId
	params["DT"] = deviceToken
	byteArray, _ := proto.Marshal(message.Data.GetPushInfo())
	params["PI"] = base64.StdEncoding.EncodeToString(byteArray)
	//fmt.Println(params)
	return iGeTui.HttpPostJson(params)

}
func (iGeTui *IGeTui) PushAPNMessageToList(appId string, contentId string, deviceTokenList []string) map[string]interface{} {
	params := map[string]interface{}{}
	params["action"] = "apnPushToListAction"
	params["appkey"] = iGeTui.AppKey
	params["contentId"] = contentId
	params["DTL"] = deviceTokenList
	params["appId"] = appId
	//fmt.Println(params)
	return iGeTui.HttpPostJson(params)
}

func (iGeTui *IGeTui) PushMessageToSingle(message igetui.IGtSingleMessage, tartget igetui.Target) map[string]interface{} {
	params := map[string]interface{}{}
	params["action"] = "pushMessageToSingleAction"
	params["appkey"] = iGeTui.AppKey
	transparent := message.Data.GetTransparent()
	// fmt.Println(transparent)
	byteArray, _ := proto.Marshal(transparent)
	params["clientData"] = base64.StdEncoding.EncodeToString(byteArray)
	params["transmissionContent"] = message.Data.GetTransmissionContent()
	params["isOffline"] = message.IsOffline
	params["offlineExpireTime"] = message.OfflineExpireTime
	params["appId"] = tartget.AppId
	params["clientId"] = tartget.ClientId
	params["type"] = 2
	params["pushType"] = message.Data.GetPushType()
	//增加pushNetWorkType参数(0:不限;1:wifi;)
	params["pushNetWorkType"] = message.PushNetWorkType
	return iGeTui.HttpPostJson(params)

}

func (iGeTui *IGeTui) PushMessageToAppTg(message igetui.IGtAppMessage,taskGroupName []string) map[string]interface{} {
	params := map[string]interface{}{}
	contentId := iGeTui.GetContentIdTg(message,taskGroupName)
	params["action"] = "pushMessageToAppAction"
	params["appkey"] = iGeTui.AppKey
	params["contentId"] = contentId
    params["type"] = 2
	return iGeTui.HttpPostJson(params)
}

func (iGeTui *IGeTui) GetContentIdTg(message igetui.IGtAppMessage, taskGroupName []string) interface{} {
		params := map[string]interface{}{}
        if taskGroupName != nil {
            if len(taskGroupName) > 40{
            	panic("TaskGroupName is OverLimit 40")
            }else{
            	params["taskGroupName"] = taskGroupName
            }                         
		}       
	params["action"] = "getContentIdAction"
	params["appkey"] = iGeTui.AppKey
	transparent := message.Data.GetTransparent()
	byteArray, _ := proto.Marshal(transparent)
	params["clientData"] = base64.StdEncoding.EncodeToString(byteArray)
	params["transmissionContent"] = message.Data.GetTransmissionContent()
	params["isOffline"] = message.IsOffline
	params["offlineExpireTime"] = message.OfflineExpireTime
	params["pushType"] = message.Data.GetPushType()
	params["pushNetWorkType"] = message.PushNetWorkType
	params["speed"] = message.Speed
	ret := iGeTui.HttpPostJson(params)
	if ret["result"] == "ok" {
		return ret["contentId"]
	} else {
		return " "
	}

}

func (iGeTui *IGeTui) PushMessageToApp(message igetui.IGtAppMessage) map[string]interface{} {
	params := map[string]interface{}{}
	params["action"] = "pushMessageToAppAction"
	params["appkey"] = iGeTui.AppKey
	transparent := message.Data.GetTransparent()
	//fmt.Println(transparent)
	byteArray, _ := proto.Marshal(transparent)
	params["clientData"] = base64.StdEncoding.EncodeToString(byteArray)
	params["transmissionContent"] = message.Data.GetTransmissionContent()
	params["isOffline"] = message.IsOffline
	params["offlineExpireTime"] = message.OfflineExpireTime
	params["appIdList"] = message.AppIdList
	params["phoneTypeList"] = message.PhoneTypeList
	params["provinceList"] = message.ProvinceList
	params["type"] = 2
	params["pushType"] = message.Data.GetPushType()
	params["pushNetWorkType"] = message.PushNetWorkType
	params["speed"] = message.Speed
	return iGeTui.HttpPostJson(params)
}

func (iGeTui *IGeTui) PushMessageToList(contentId string, targets []igetui.Target) map[string]interface{} {
	params := map[string]interface{}{}
	params["action"] = "pushMessageToListAction"
	params["appkey"] = iGeTui.AppKey
	params["contentId"] = contentId

	targetList := []interface{}{}
	for _, target := range targets {
		appId := target.AppId
		clientId := target.ClientId
		targetTmp := map[string]string{"appId": appId, "clientId": clientId}
		targetList = append(targetList, targetTmp)
	}

	params["targetList"] = targetList
	params["type"] = 2

	return iGeTui.HttpPostJson(params)
}

func (iGeTui *IGeTui) GetContentId(message igetui.IGtListMessage) interface{} {
	
	
	params := map[string]interface{}{}
	params["action"] = "getContentIdAction"
	params["appkey"] = iGeTui.AppKey
	transparent := message.Data.GetTransparent()
	byteArray, _ := proto.Marshal(transparent)
	params["clientData"] = base64.StdEncoding.EncodeToString(byteArray)
	params["transmissionContent"] = message.Data.GetTransmissionContent()
	params["isOffline"] = message.IsOffline
	params["offlineExpireTime"] = message.OfflineExpireTime
	params["pushType"] = message.Data.GetPushType()
	params["pushNetWorkType"] = message.PushNetWorkType
	//params["speed"] = message.Speed
	ret := iGeTui.HttpPostJson(params)

	if ret["result"] == "ok" {
		return ret["contentId"]
	} else {
		return " "
	}
}




func (iGeTui *IGeTui) GetAPNContentId(appId string, message igetui.IGtListMessage) interface{} {
	params := map[string]interface{}{}
	params["action"] = "apnGetContentIdAction"
	params["appkey"] = iGeTui.AppKey
	params["appId"] = appId
	byteArray, _ := proto.Marshal(message.Data.GetPushInfo())
	params["PI"] = base64.StdEncoding.EncodeToString(byteArray)

	ret := iGeTui.HttpPostJson(params)

	if ret["result"] == "ok" {
		return ret["contentId"]
	} else {
		return " "
	}
}

func (iGeTui *IGeTui) cancelContentId(contentId string) bool {
	params := map[string]interface{}{}
	params["action"] = "cancelContentIdAction"
	params["contentId"] = contentId

	ret := iGeTui.HttpPostJson(params)

	if ret["result"] == "ok" {
		return true
	} else {
		return false
	}
}



//hxh
func (iGeTui *IGeTui) Stop(contentId string) bool {
	params := map[string]interface{}{}
	params["action"] = "stopTaskAction"
	params["appkey"] = iGeTui.AppKey
	params["contentId"] = contentId
	ret := iGeTui.HttpPostJson(params)
	if ret["result"] == "ok" {
		return true
	} else {
		return false
	}
}

func (iGeTui *IGeTui) GetClientIdStatus(appId string, clientId string) map[string]interface{} {
	params := map[string]interface{}{}
	params["action"] = "getClientIdStatusAction"
	params["appkey"] = iGeTui.AppKey
	params["appId"] = appId
	params["clientId"] = clientId
	return iGeTui.HttpPostJson(params)

}

func (iGeTui *IGeTui) GetPushResult(taskId string) map[string]interface{} {
	params := map[string]interface{}{}
	params["action"] = "getPushMsgResult"
	params["appkey"] = iGeTui.AppKey
	params["taskId"] = taskId
	return iGeTui.HttpPostJson(params)

}

func (iGeTui *IGeTui) SetClientTag(tags []string, tartget igetui.Target) map[string]interface{} {
	params := map[string]interface{}{}
	params["action"] = "setTagAction"
	params["appkey"] = iGeTui.AppKey
	params["appId"] = tartget.AppId
	params["clientId"] = tartget.ClientId
	params["tagList"] = tags
	return iGeTui.HttpPostJson(params)
}

func (iGeTui *IGeTui) GetUserTags(tartget igetui.Target) map[string]interface{} {
	params := map[string]interface{}{}
	params["action"] = "getUserTags"
	params["appkey"] = iGeTui.AppKey
	params["appId"] = tartget.AppId
	params["clientId"] = tartget.ClientId
	return iGeTui.HttpPostJson(params)
}

func (iGeTui *IGeTui) BindAlias(alias string, appId string, clientId string) map[string]interface{} {
	params := map[string]interface{}{}
	params["action"] = "alias_bind"
	params["appkey"] = iGeTui.AppKey
	params["appid"] = appId
	params["cid"] = clientId
	params["alias"] = alias
	return iGeTui.HttpPostJson(params)
}

func (iGeTui *IGeTui) UnBindAlias(appId string, alias string, clientId string) map[string]interface{} {
	params := map[string]interface{}{}
	params["action"] = "alias_unbind"
	params["appkey"] = iGeTui.AppKey
	params["appid"] = appId
	params["cid"] = clientId
	params["alias"] = alias
	return iGeTui.HttpPostJson(params)
}

func (iGeTui *IGeTui) UnBindAliasAll(appId string, alias string) map[string]interface{} {
	params := map[string]interface{}{}
	params["action"] = "alias_unbind"
	params["appkey"] = iGeTui.AppKey
	params["appid"] = appId
	params["alias"] = alias
	return iGeTui.HttpPostJson(params)
}

func (iGeTui *IGeTui) BindAliasBatch(appId string, targets []igetui.Target) map[string]interface{} {
	targetList := []interface{}{}
	params := map[string]interface{}{}
	for _, target := range targets {
		alias := target.Alias
		clientId := target.ClientId
		targetTmp := map[string]string{"cid": clientId, "alias": alias}
		targetList = append(targetList, targetTmp)
	}
	params["aliaslist"] = targetList
	params["action"] = "alias_bind_list"
	params["appkey"] = iGeTui.AppKey
	params["appid"] = appId
	return iGeTui.HttpPostJson(params)
}

func (iGeTui *IGeTui) QueryClientId(alias string, appId string) map[string]interface{} {
	params := map[string]interface{}{}
	params["action"] = "alias_query"
	params["appkey"] = iGeTui.AppKey
	params["appid"] = appId
	params["alias"] = alias
	return iGeTui.HttpPostJson(params)

}

func (iGeTui *IGeTui) QueryAlias(appId string, clientId string) map[string]interface{} {
	params := map[string]interface{}{}
	params["action"] = "alias_query"
	params["appkey"] = iGeTui.AppKey
	params["appid"] = appId
	params["cid"] = clientId
	return iGeTui.HttpPostJson(params)
}

//hxh

func (iGeTui *IGeTui) GetSign(appKey string, timeStamp int64, masterSecret string) string {
	rawValue := appKey + strconv.FormatInt(timeStamp, 10) + masterSecret
	h := md5.New()
	io.WriteString(h, rawValue)
	return hex.EncodeToString(h.Sum(nil))
}

func (iGeTui *IGeTui) GetCurrentTime() int64 {
	t := time.Now().Unix() * 1000
	return t
}

func (iGeTui *IGeTui) HttpPostJson(params map[string]interface{}) map[string]interface{} {
	ret := iGeTui.HttpPost(params)
	if ret["result"] == "sign_error" {
		iGeTui.connect()
		ret = iGeTui.HttpPost(params)
	}

	return ret
}

func (iGeTui *IGeTui) HttpPost(params map[string]interface{}) map[string]interface{} {
	data, _ := json.Marshal(params)
	//fmt.Printf("%s\n", data)
	tryTime := 1
tryAgain:
	res, err := http.Post(iGeTui.Host, "application/json", strings.NewReader(string(data)))
	if err != nil {
		fmt.Println("第"+strconv.Itoa(tryTime)+"次", "请求失败")
		tryTime += 1
		if tryTime < 4 {
			goto tryAgain
		}
		return map[string]interface{}{"result": "post error"}
	}
	body, _ := ioutil.ReadAll(res.Body)
	var ret map[string]interface{}
	json.Unmarshal(body, &ret)
	return ret
}

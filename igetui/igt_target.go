package igetui

type Target struct {
	AppId    string
	ClientId string
	Alias string
}

func NewTarget(appid, clientid ,alias string) *Target {
	return &Target{
		AppId:    appid,
		ClientId: clientid,
		Alias :alias,
	}
}

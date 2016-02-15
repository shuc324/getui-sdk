package demo

import (
	"testing"
	"fmt"
)

func Test_Send_Push_Run(t *testing.T) {
	pushList := make([]PushField, 2)
	for i, _ := range pushList {
		pushList[i].Id = i
		pushList[i].Extra = "/message"
		// pushList[i].ClientId = "84b8083c48f59101f504a41b1382bbf8"
		// pushList[i].PlatFrom = "android"
		pushList[i].ClientId = "79ba230fef04508e527d8c69f93ce5f56b80bfe1b1e6b46efb879f5a58d74717"
		pushList[i].PlatFrom = "ios"
		pushList[i].Title = fmt.Sprintf("rewrite send push by golang")
		pushList[i].Content = fmt.Sprintf("golan send push test")
		pushList[i].TargetId = 30874543
	}
	// 不要将false改为true 改为true会全员发广播
	SendPush(pushList, false)
}
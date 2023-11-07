package main

import (
	"fmt"
	behaviorauditsdk "github.com/Yostardev/behavior-audit-sdk"
	"time"
)

func init() {
	behaviorauditsdk.InitProducer(
		"", // ak
		"", // sk
		"ali-cn-sh-tc-ops-test-all",
		"cn-shanghai.log.aliyuncs.com",
		"behavior-audit", // platform name
		"dev",            // platform env
	)
}

func main() {
	ba := behaviorauditsdk.NewBehaviorAudit()

	// do something
	time.Sleep(time.Second * 2)

	err := ba.
		SetBehavior("test send msg to sls", "context 1", "context 2").
		SetUser("test user", "test username").
		SetResult(behaviorauditsdk.Success, "test error info", "test other info").
		SetData("test original data", "test new data").
		// 自定义行为开始结束时间，不指定 开始时间为sdk.NewBehaviorAudit()时间，结束时间为Send()时间
		//SetTime(time.Now().Add(-10*time.Second), time.Now()).
		Send()
	if err != nil {
		fmt.Println(err)
	}
	behaviorauditsdk.Close()
}

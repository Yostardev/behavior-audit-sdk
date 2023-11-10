package behavior_audit_sdk

import (
	"github.com/aliyun/aliyun-log-go-sdk/producer"
	"strings"
	"time"
)

type BehaviorStatus string

var (
	Success BehaviorStatus = "success"
	Failed  BehaviorStatus = "failed"
)

type BehaviorAudit struct {
	behavior struct {
		category string
		context  string
	}
	time struct {
		start string
		end   string
	}
	user struct {
		name     string
		username string
	}
	result struct {
		status    BehaviorStatus
		errorInfo string
		otherInfo string
	}
	data struct {
		original string
		new      string
	}
}

func NewBehaviorAudit() *BehaviorAudit {
	cstZone := time.FixedZone("GMT", 8*3600)
	return &BehaviorAudit{
		time: struct {
			start string
			end   string
		}{

			start: time.Now().In(cstZone).Format("2006-01-02 15:04:05"),
		},
	}
}

func (b *BehaviorAudit) SetBehavior(category string, contexts ...string) *BehaviorAudit {
	b.behavior.category = category
	b.behavior.context = strings.Join(contexts, "\n")
	return b
}

func (b *BehaviorAudit) SetTime(start, end time.Time) *BehaviorAudit {
	b.time.start = start.Format("2006-01-02 15:04:05")
	b.time.end = end.Format("2006-01-02 15:04:05")
	return b
}

func (b *BehaviorAudit) SetUser(name, username string) *BehaviorAudit {
	b.user.name = name
	b.user.username = username
	return b
}

func (b *BehaviorAudit) SetResult(status BehaviorStatus, errorInfo, otherInfo string) *BehaviorAudit {
	b.result.status = status
	b.result.errorInfo = errorInfo
	b.result.otherInfo = otherInfo
	return b
}

func (b *BehaviorAudit) SetData(original, new string) *BehaviorAudit {
	b.data.original = original
	b.data.new = new
	return b
}

func (b *BehaviorAudit) Send() error {
	if b.time.end == "" {
		cstZone := time.FixedZone("GMT", 8*3600)
		b.time.end = time.Now().In(cstZone).Format("2006-01-02 15:04:05")
	}

	log := producer.GenerateLog(
		uint32(time.Now().Unix()),
		map[string]string{
			"behavior.type":     b.behavior.category,
			"behavior.context":  b.behavior.context,
			"time.start":        b.time.start,
			"time.end":          b.time.end,
			"user.name":         b.user.name,
			"user.username":     b.user.username,
			"result.status":     string(b.result.status),
			"result.error_info": b.result.errorInfo,
			"result.other_info": b.result.otherInfo,
			"data.original":     b.data.original,
			"data.new":          b.data.new,
		},
	)

	err := producerInstance.SendLog(slsConfig.project, slsConfig.logStore, "", "", log)
	if err != nil {
		return err
	}
	return nil
}

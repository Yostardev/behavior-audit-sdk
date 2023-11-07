package behavior_audit_sdk

import (
	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/aliyun/aliyun-log-go-sdk/producer"
)

var producerInstance *producer.Producer

var slsConfig struct {
	project  string
	logStore string
}

func InitProducer(accessKeyID, accessKeySecret, project, endpoint, platformName, platformEnv string) {
	slsConfig.project = project
	slsConfig.logStore = platformName + "-" + platformEnv
	producerConfig := producer.GetDefaultProducerConfig()
	producerConfig.Endpoint = endpoint
	producerConfig.CredentialsProvider = sls.NewStaticCredentialsProvider(accessKeyID, accessKeySecret, "")
	producerConfig.GeneratePackId = true
	producerInstance = producer.InitProducer(producerConfig)
	producerInstance.Start()
}

func Close() {
	producerInstance.SafeClose()
}

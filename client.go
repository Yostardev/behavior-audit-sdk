package behavior_audit_sdk

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	sls20201230 "github.com/alibabacloud-go/sls-20201230/v5/client"
	"github.com/alibabacloud-go/tea/tea"
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
	client, err := CreateSlsClient(accessKeyID, accessKeySecret, endpoint)
	if err != nil {
		panic(err)
	}
	err = CreateLogStore(client)
	if err != nil {
		panic(err)
	}

	err = CreateIndex(client)
	if err != nil {
		panic(err)
	}
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

func CreateSlsClient(accessKeyID, accessKeySecret, endpoint string) (*sls20201230.Client, error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(accessKeyID),
		AccessKeySecret: tea.String(accessKeySecret),
	}
	config.Endpoint = tea.String(endpoint)
	return sls20201230.NewClient(config)
}

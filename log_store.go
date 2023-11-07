package behavior_audit_sdk

import (
	sls20201230 "github.com/alibabacloud-go/sls-20201230/v5/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

func CreateLogStore(client *sls20201230.Client) error {
	createLogStoreRequest := &sls20201230.CreateLogStoreRequest{
		LogstoreName: tea.String(slsConfig.logStore),
		ShardCount:   tea.Int32(2),
		Ttl:          tea.Int32(180),
	}
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		_, err := client.CreateLogStoreWithOptions(tea.String(slsConfig.project), createLogStoreRequest, headers, runtime)
		if err != nil {
			return err
		}

		return nil
	}()

	if tryErr != nil {
		var err = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			err = _t
		} else {
			err.Message = tea.String(tryErr.Error())
		}
		// 如有需要，请打印 error
		_, _err := util.AssertAsString(err.Message)
		if _err != nil {
			return _err
		}
	}
	return nil
}

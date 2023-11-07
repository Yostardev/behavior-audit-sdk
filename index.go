package behavior_audit_sdk

import (
	sls20201230 "github.com/alibabacloud-go/sls-20201230/v5/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

func CreateIndex(client *sls20201230.Client) error {
	line := &sls20201230.CreateIndexRequestLine{
		Chn:           tea.Bool(true),
		CaseSensitive: tea.Bool(true),
		Token:         tea.StringSlice([]string{",", " ", "'", "\"", ";", "=", "(", ")", "[", "]", "{", "}", "?", "@", "<", ">", "/", ":", "\n", "\t", "\r"}),
	}
	keys := map[string]*sls20201230.KeysValue{
		"behavior.type": &sls20201230.KeysValue{
			CaseSensitive: tea.Bool(false),
			Chn:           tea.Bool(true),
			Type:          tea.String("text"),
			Token:         tea.StringSlice([]string{",", " ", "'", "\"", ";", "=", "(", ")", "[", "]", "{", "}", "?", "@", "<", ">", "/", ":", "\n", "\t", "\r"}),
		},
		"behavior.context": &sls20201230.KeysValue{
			CaseSensitive: tea.Bool(false),
			Chn:           tea.Bool(true),
			Type:          tea.String("long"),
		},
		"data.new": &sls20201230.KeysValue{
			CaseSensitive: tea.Bool(false),
			Chn:           tea.Bool(true),
			Type:          tea.String("long"),
		},
		"data.original": &sls20201230.KeysValue{
			CaseSensitive: tea.Bool(false),
			Chn:           tea.Bool(true),
			Type:          tea.String("long"),
		},
		"result.status": &sls20201230.KeysValue{
			CaseSensitive: tea.Bool(false),
			Chn:           tea.Bool(false),
			Type:          tea.String("text"),
			Token:         tea.StringSlice([]string{",", " ", "'", "\"", ";", "=", "(", ")", "[", "]", "{", "}", "?", "@", "<", ">", "/", ":", "\n", "\t", "\r"}),
		},
		"result.error_info": &sls20201230.KeysValue{
			CaseSensitive: tea.Bool(false),
			Chn:           tea.Bool(true),
			Type:          tea.String("long"),
		},
		"result.other_info": &sls20201230.KeysValue{
			CaseSensitive: tea.Bool(false),
			Chn:           tea.Bool(true),
			Type:          tea.String("long"),
		},
		"time.end": &sls20201230.KeysValue{
			CaseSensitive: tea.Bool(false),
			Chn:           tea.Bool(false),
			Type:          tea.String("text"),
			Token:         tea.StringSlice([]string{",", " ", "'", "\"", ";", "=", "(", ")", "[", "]", "{", "}", "?", "@", "<", ">", "/", ":", "\n", "\t", "\r"}),
		},
		"time.start": &sls20201230.KeysValue{
			CaseSensitive: tea.Bool(false),
			Chn:           tea.Bool(false),
			Type:          tea.String("text"),
			Token:         tea.StringSlice([]string{",", " ", "'", "\"", ";", "=", "(", ")", "[", "]", "{", "}", "?", "@", "<", ">", "/", ":", "\n", "\t", "\r"}),
		},
		"user.name": &sls20201230.KeysValue{
			CaseSensitive: tea.Bool(false),
			Chn:           tea.Bool(true),
			Type:          tea.String("text"),
			Token:         tea.StringSlice([]string{",", " ", "'", "\"", ";", "=", "(", ")", "[", "]", "{", "}", "?", "@", "<", ">", "/", ":", "\n", "\t", "\r"}),
		},
		"user.username": &sls20201230.KeysValue{
			CaseSensitive: tea.Bool(false),
			Chn:           tea.Bool(false),
			Type:          tea.String("text"),
			Token:         tea.StringSlice([]string{",", " ", "'", "\"", ";", "=", "(", ")", "[", "]", "{", "}", "?", "@", "<", ">", "/", ":", "\n", "\t", "\r"}),
		},
	}
	createIndexRequest := &sls20201230.CreateIndexRequest{
		Keys: keys,
		Line: line,
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
		_, _err := client.CreateIndexWithOptions(tea.String(slsConfig.project), tea.String(slsConfig.logStore), createIndexRequest, headers, runtime)
		if _err != nil {
			return _err
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

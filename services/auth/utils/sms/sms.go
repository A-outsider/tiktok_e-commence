package sms

import (
	"encoding/json"
	"errors"
	"fmt"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"gomall/services/auth/initialize"
	"strings"
)

func SendCaptcha(phone string, code string) error {
	ma := make(map[string]string, 1)
	ma["code"] = code

	marshal, err := json.Marshal(ma)
	if err != nil {
		return err
	}

	data := string(marshal)

	return Send("阿里云短信测试", "SMS_154950909", phone, data)
}

func Send(signName string, templateCode string, phone string, data string) error {
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String(signName),
		TemplateCode:  tea.String(templateCode),
		PhoneNumbers:  tea.String(phone),
		TemplateParam: &data,
	}
	var err error
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()

		// 复制代码运行请自行打印 API 的返回值
		_, err = initialize.GetSms().SendSmsWithOptions(sendSmsRequest, &util.RuntimeOptions{})
		if err != nil {
			return err
		}

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		var _t *tea.SDKError
		if errors.As(tryErr, &_t) {
			error = _t
		}

		// 诊断地址
		var data interface{}
		d := json.NewDecoder(strings.NewReader(tea.StringValue(error.Data)))
		d.Decode(&data)
		if m, ok := data.(map[string]interface{}); ok {
			recommend, _ := m["Recommend"]
			fmt.Println(recommend)
		}
		_, err = util.AssertAsString(error.Message)
		if err != nil {
			return err
		}
	}
	return err
}

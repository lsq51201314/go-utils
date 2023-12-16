package utils

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

// 短信实例
type YunSMS struct {
	client *dysmsapi.Client
}

// 新建实例
func NewYunSMS(accessKeyId, accessKeySecret string, smsEndpoint ...string) (sms YunSMS, err error) {
	endpoint:="dysmsapi.aliyuncs.com"
	if len(smsEndpoint) > 0 {
		endpoint = smsEndpoint[0]
	}
	config := &openapi.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
		Endpoint:        tea.String(endpoint),
	}
	sms.client, err = dysmsapi.NewClient(config)
	return
}

// 发送短信
func (a *YunSMS) Send(signname, template, phoneNumber, jsonStr string) (err error) {
	sendSmsRequest := &dysmsapi.SendSmsRequest{
		PhoneNumbers:  tea.String(phoneNumber),
		SignName:      tea.String(signname),
		TemplateCode:  tea.String(template),
		TemplateParam: tea.String(jsonStr),
	}
	_, err = a.client.SendSms(sendSmsRequest)
	return
}

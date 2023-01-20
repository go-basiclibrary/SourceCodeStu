package main

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	dysmsapi "github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

func main() {
	config := sdk.NewConfig()

	credential := credentials.NewAccessKeyCredential("LTAI5tEEpaHuxKeLtDZsA8fe", "nMqiVh0ysdfRNq8eWUHlKxB7N3JC1Z")
	/* use STS Token
	credential := credentials.NewStsTokenCredential("<your-access-key-id>", "<your-access-key-secret>", "<your-sts-token>")
	*/
	client, err := dysmsapi.NewClientWithOptions("cn-hangzhou", config, credential)
	if err != nil {
		panic(err)
	}

	request := dysmsapi.CreateSendSmsRequest()

	request.Scheme = "https"

	request.SignName = "阿里云短信测试"
	request.TemplateCode = "SMS_154950909"
	request.PhoneNumbers = "15093710052"
	// 您正在使用阿里云短信测试服务，体验验证码是：${code}，如非本人操作，请忽略本短信！
	request.TemplateParam = "{\"code\":\"1234\"}"

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}

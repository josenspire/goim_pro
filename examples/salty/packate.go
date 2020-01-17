package main

import (
	"context"
	protos "goim-pro/api/protos/salty"
	"goim-pro/pkg/utils"
	"log"
)

func obtainSMSCode(t protos.SMSServiceClient) {
	smsReq := protos.SMSReq{
		CodeType: protos.SMSReq_REGISTER,
		TargetAccount: &protos.SMSReq_Telephone{
			Telephone: "13631210000",
		},
	}
	anyData, _ := utils.MarshalMessageToAny(&smsReq)
	gprcReq := &protos.GrpcReq{
		Data: anyData,
	}
	// 调用 gRPC 接口
	tr, err := t.ObtainSMSCode(context.Background(), gprcReq)
	//tr, err := t.Register(context.Background(), gprcReq)
	if err != nil {
		log.Fatalf("could not greet: %v", err.Error())
	}
	printResp(tr)
}

func register(t protos.UserServiceClient) {
	registerReq := &protos.RegisterReq{
		RegisterType:     0,
		Password:         "1234567890",
		VerificationCode: "123456",
		UserProfile: &protos.UserProfile{
			Telephone:   "13631210001",
			Email:       "123@qq.com",
			Username:    "JAMES001",
			Nickname:    "JAMES001",
			Avatar:      "https://www.baidu.com/avatar/header1.png",
			Description: "Never settle",
			Sex:         0,
			Birthday:    utils.MakeTimestamp(),
			Location:    "CHINA-ZHA",
		},
	}
	anyData, _ := utils.MarshalMessageToAny(registerReq)
	grpcReq := &protos.GrpcReq{
		DeviceID: "",
		Version:  "",
		Language: 0,
		Os:       0,
		Token:    "",
		Data:     anyData,
	}

	tr, err := t.Register(context.Background(), grpcReq)
	if err != nil {
		logger.Errorf("register error: %s", err.Error())
	} else {
		printResp(tr)
	}
}

func printResp(resp *protos.GrpcResp) {
	logger.Infof("[code]: %d", resp.GetCode())
	logger.Infof("[data]: %s", resp.GetData().GetValue())
	logger.Infof("[message]: %s", resp.GetMessage())
}

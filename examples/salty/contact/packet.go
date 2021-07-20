package contact

import (
	"context"
	"github.com/golang/protobuf/proto"
	protos "goim-pro/api/protos/salty"
	"goim-pro/pkg/logs"
	"goim-pro/pkg/utils"
)

var (
	UserId = "01EMK05C80JKJZF9MXGT8XS2KW" // 13631210008
	logger = logs.GetLogger("INFO")
)

func RequestContact(t protos.ContactServiceClient) {
	reqContactReq := &protos.RequestContactReq{
		UserId: "01EMK05C80JKJZF9MXGT8XS2KW", // 13631210001
		// UserId: "01E4QYJBERVD8E5N9SXAEGXMB8", // 13631210002
		Reason: "我是德玛西亚人，来交个朋友！",
	}
	anyData, _ := utils.MarshalMessageToAny(reqContactReq)
	grpcReq := &protos.GrpcReq{
		DeviceId: "LOCAL_DEV",
		Version:  "",
		Language: 0,
		Os:       0,
		Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJNREZGVFVzd01WWlpPRGd5TUVNNU5VMUlNMXBIVGpOS1dGRXNURTlEUVV4ZlJFVlciLCJleHAiOjE2MTI3MDg3NTAsImlhdCI6MTYxMjQ0OTU1MCwiaXNzIjoic2FsdHlfaW0ifQ.n8w-rS3iAL2ieQb35_t8n9-i72GlLCsSrEfOBtEO0e8",
		Data:     anyData,
	}

	tr, err := t.RequestContact(context.Background(), grpcReq)
	if err != nil {
		logger.Errorf("login error: %s", err.Error())
	} else {
		var resp = &protos.RequestContactResp{}
		printResp(tr, resp)
	}
}

func AcceptContact(t protos.ContactServiceClient) {
	acpContactReq := &protos.AcceptContactReq{
		UserId: "01EMK01VY8820C95MH3ZGN3JXQ",
	}
	anyData, _ := utils.MarshalMessageToAny(acpContactReq)
	grpcReq := &protos.GrpcReq{
		DeviceId: "LOCAL_DEV",
		Version:  "",
		Language: 0,
		Os:       0,
		Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJNREZGVFVzd05VTTRNRXBMU2xwR09VMVlSMVE0V0ZNeVMxY3NURTlEUVV4ZlJFVlciLCJleHAiOjE2MTI3MDg3NTAsImlhdCI6MTYxMjQ0OTU1MCwiaXNzIjoic2FsdHlfaW0ifQ.5kizBTqxf9lydrxSA0B5p2ysD_WO9zX9ayVmMzqM6F4",
		Data:     anyData,
	}

	tr, err := t.AcceptContact(context.Background(), grpcReq)
	if err != nil {
		logger.Errorf("accept contact error: %s", err.Error())
	} else {
		var resp = &protos.AcceptContactResp{}
		printResp(tr, resp)
	}
}

func RefusedContact(t protos.ContactServiceClient) {
	refusedContactReq := &protos.RefusedContactReq{
		UserId: "01EN0AHEK8ZVH55V0YYSXZ7CK9",
		Reason: "我喜欢小姐姐",
	}
	anyData, _ := utils.MarshalMessageToAny(refusedContactReq)
	grpcReq := &protos.GrpcReq{
		DeviceId: "LOCAL_DEV",
		Version:  "",
		Language: 0,
		Os:       0,
		Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJNREZGVGpCQlNFTk5VazFOTTFJeFJFZEVNemhMVmpkWFdUZz0iLCJleHAiOjE2MDM1NDExNjAsImlhdCI6MTYwMzI4MTk2MCwiaXNzIjoic2FsdHlfaW0ifQ.OnXfVm9VNh4jouw4TAqi_acSW39uw5ajP_MEf8ztEsI",
		Data:     anyData,
	}

	tr, err := t.RefusedContact(context.Background(), grpcReq)
	if err != nil {
		logger.Errorf("refused contact request error: %s", err.Error())
	} else {
		var resp = &protos.RefusedContactResp{}
		printResp(tr, resp)
	}
}

func UpdateRemarkInfo(t protos.ContactServiceClient) {
	updateRemarkInfoReq := &protos.UpdateRemarkInfoReq{
		UserId: UserId,
		RemarkInfo: &protos.ContactRemark{
			RemarkName:  "喜洋洋",
			Description: "He's a crazy boy.",
			Telephones:  []string{""},
			Tags:        []string{"Friend", "Boy"},
		},
	}
	anyData, _ := utils.MarshalMessageToAny(updateRemarkInfoReq)
	grpcReq := &protos.GrpcReq{
		DeviceId: "LOCAL_DEV",
		Version:  "",
		Language: 0,
		Os:       0,
		Token:    "1234567890",
		Data:     anyData,
	}

	tr, err := t.UpdateRemarkInfo(context.Background(), grpcReq)
	if err != nil {
		logger.Errorf("update remark info error: %s", err.Error())
	} else {
		var resp = &protos.UpdateRemarkInfoResp{}
		printResp(tr, resp)
	}
}

func GetContactList(t protos.ContactServiceClient) {
	grpcReq := &protos.GrpcReq{
		DeviceId: "LOCAL_DEV",
		Version:  "",
		Language: 0,
		Os:       0,
		Token:    "1234567890",
		Data:     nil,
	}

	tr, err := t.GetContactList(context.Background(), grpcReq)
	if err != nil {
		logger.Errorf("get contacts info error: %s", err.Error())
	} else {
		var resp = &protos.GetContactListResp{}
		printResp(tr, resp)
	}
}

func GetContactOperationList(t protos.ContactServiceClient) {
	acpContactReq := &protos.GetContactOperationListReq{
		StartDateTime: 1602605419,
		EndDateTime:   1602605419,
	}
	anyData, _ := utils.MarshalMessageToAny(acpContactReq)
	grpcReq := &protos.GrpcReq{
		DeviceId: "f9b5f73048f3867e",
		Version:  "",
		Language: 0,
		Os:       0,
		Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJNREZHTkVkVE5GSkZNRWM1VGpNMlVsazJWRlExUkV0Qk1sa3NaamxpTldZM016QTBPR1l6T0RZM1pRPT0iLCJleHAiOjE2MjYwMDc3NzcsImlhdCI6MTYyNTc0ODU3NywiaXNzIjoic2FsdHlfaW0ifQ.nA6_Rm8y6KOTc6gTuLtYB6jku5f5XoHlP-8zAxhWmaM",
		Data:     anyData,
	}

	tr, err := t.GetContactOperationList(context.Background(), grpcReq)
	if err != nil {
		logger.Errorf("get notification messages info error: %s", err.Error())
	} else {
		var resp = &protos.GetContactOperationListResp{}
		printResp(tr, resp)
	}
}

func printResp(resp *protos.GrpcResp, pb proto.Message) {
	if err := utils.UnMarshalAnyToMessage(resp.GetData(), pb); err != nil {
		logger.Error("unmarshal failed: %v", err)
	}
	logger.Infof("[code]: %d", resp.GetCode())
	logger.Infof("[data]: %v", pb.String())
	logger.Infof("[message]: %s", resp.GetMessage())
}

package main

import (
	"context"
	"fmt"
	"os"

	client "github.com/larksuite/oapi-sdk-go"
	"github.com/larksuite/oapi-sdk-go/core"
	larkcontact "github.com/larksuite/oapi-sdk-go/service/contact/v3"
)

func main() {
	var appID, appSecret = os.Getenv("APP_ID"), os.Getenv("APP_SECRET")
	var feishu_client = client.NewClient(appID, appSecret, client.WithLogLevel(core.LogLevelDebug), client.WithLogReqRespInfoAtDebugLevel(false))

	user := larkcontact.NewUserBuilder().Build()
	resp, err := feishu_client.Contact.User.Patch(context.Background(),
		larkcontact.NewPatchUserReqBuilder().
			UserId("ou_155184d1e73cbfb8973e5a9e698e74f2").
			UserIdType(larkcontact.UserIdTypeOpenId).
			User(user).
			Build())

	if err != nil {
		fmt.Println(err)
		return
	}

	if resp.Success() {
		fmt.Println(resp.Data.User)
	} else {
		fmt.Println(resp.Msg, resp.Code, resp.RequestId())
	}

}

package code_runner_bot

import (
	"fmt"
	"os"

	"github.com/tahadostifam/RunCode-TBot/configs"

	"github.com/helloyi/go-sshclient"
)

var sshClient *sshclient.Client

func SetSSHClient() {
	addr := fmt.Sprintf("%v:%v", configs.AllConfigs.SSH_HOST, configs.AllConfigs.SSH_PORT)
	client, err := sshclient.DialWithPasswd(addr, configs.AllConfigs.SSH_USER, configs.AllConfigs.SSH_PASS)
	if err != nil {
		fmt.Println("An error occurred on connecting to ssh server")
		fmt.Println("Error:")
		fmt.Println(err)
		os.Exit(1)
	} else {
		sshClient = client
	}
}

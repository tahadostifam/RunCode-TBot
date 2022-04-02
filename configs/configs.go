package configs

type Configs struct {
	BOT_TOKEN         string
	SERVER_SSH_CLIENT struct {
		host     string
		port     string
		username string
		password string
	}
}

func ReadAndSetConfigs() {

}

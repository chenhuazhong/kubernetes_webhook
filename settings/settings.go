package settings

import (
	"kubernetes_webhook/utils"
)

const HOST_IP = "192.168.179.128"
const HOST_PORT = 8088

var (
	IT_DEVOPS_ENV  = utils.GetDefaultEnv("IT_DEVOPS_ENV", "dev")
	CERT_FILE_PATH = utils.GetDefaultEnv("CERT_FILE_PATH", "certs/server-cert.pem")
	KEY_FILE_PATH = utils.GetDefaultEnv("KEY_FILE_PATH", "certs/server-key.pem")

	)


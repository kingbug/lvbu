package init

import (
	"lvbu/utils"
)

func InitWebsocket() {
	go utils.ProDetection()
}

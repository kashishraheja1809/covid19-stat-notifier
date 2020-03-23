package notifications

import (
	"github.com/0xAX/notificator"
)

func SendNotification(
	title string,
	text string,
	severity string,
) {
	var notify *notificator.Notificator
	notify = notificator.New(notificator.Options{
		DefaultIcon: "../covid19.png",
		AppName:     "Covid19 Updates",
	})

	var imagePath, urgency string
	if severity == "CRITICAL" {
		imagePath = "../critical.png"
		urgency = notificator.UR_CRITICAL
	} else {
		imagePath = "../alert.png"
		urgency = notificator.UR_NORMAL
	}

	notify.Push(title, text, imagePath, urgency)
}

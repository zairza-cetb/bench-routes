package notify

import (
	"log"
	"os"
	"strings"
	"testing"
)

func TestSendNotification(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	substr := "/bench-routes/"
	c := strings.Index(dir, substr)
	dir = dir[:c+len(substr)] + "assets/icon.png"
	var n = NotificationData{
		title:       "hi",
		description: "hello",
		icon:        dir,
		urgency:     "critical",
		time:        5000,
	}
	n.SendNotification1()
	n.SendNotification2()
	n.SendNotification3()
	n.SendNotification4()
	n.SendNotification5()
	n.SendNotification6()
}

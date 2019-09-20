package ping

import (
	scrap "github.com/zairza-cetb/bench-routes/src/lib/filters/scraps"
	"log"
	"testing"
)

var (
	urls = []string{
		"google.co.in",
		"facebook.com",
	}
)

func TestHandlerPing(t *testing.T) {
	for _, inst := range urls {
		a := HandlePing(&inst, 20)
		if *a == (scrap.TypePingScrap{}) {
			t.Errorf("invalid response from HandlePing")
		} else {
			log.Println(*a)
		}
	}
}

package providers

import (
	"fmt"
	"log"
	"net/url"
	"os/exec"
	"runtime"

	"github.com/tzneal/ham-go/callsigns"
)

type qrz struct {
}

func init() {
	callsigns.RegisterLookup("qrz", NewQRZ)
}
func NewQRZ(cfg callsigns.LookupConfig) callsigns.Lookup {
	return &qrz{}
}

func (h *qrz) Lookup(call string) (*callsigns.Response, error) {
	var err error
	_, realCall, _ := callsigns.Parse(call)
	url := "https://www.qrz.com/db/" + url.QueryEscape(realCall)
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
	return nil, nil
}

func (h *qrz) RequiresNetwork() bool {
	return true
}

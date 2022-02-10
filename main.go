package main

import (
	"github.com/uniwizardcom/golang-cliarguments/cliarguments"
)

func main() {
	cli := cliarguments.LineService{}
	cli.SetItemSupported(cliarguments.LineServiceItem{
		Name:  "limit",
		Level: 2,
	})
	cli.SetItemSupported(cliarguments.LineServiceItem{
		Name:  "action",
		Level: 2,
	})
	cli.ServiceCmdNew()
}

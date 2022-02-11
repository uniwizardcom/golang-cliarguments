package main

import (
	"fmt"
	"github.com/uniwizardcom/golang-cliarguments/cliarguments"
	"strconv"
)

var cli *cliarguments.LineService

func main() {
	prepareArgs()

	idv := cli.GetArgValue("insert-data", 0, "abc")
	rlv := cli.GetArgValue("records-limit", 1, "777")
	rov := cli.GetArgValue("records-offset", 1, "888")

	rl, _ := strconv.ParseUint(rlv, 10, 64)
	ro, _ := strconv.ParseUint(rov, 10, 64)
	fmt.Printf("idv: [%s]\nrecords-limit=%d\nrecords-offset=%d\n\n", idv, rl, ro)
}

func prepareArgs() {
	cli = &cliarguments.LineService{
		Description: "Proof of concept for RediSearch functionality",
		Using:       "Programu uzywa się standardowo ;)",
	}
	cli.SetItemSupported(cliarguments.LineServiceItem{
		Name: "insert-data",
		Desc: "Akcja dodająca-resetująca zestaw danych (łącznie z indexami) w Redis na podstawie danych z PostgreSQL",
	})
	cli.SetItemSupported(cliarguments.LineServiceItem{
		Name:  "make-test",
		Level: 0,
		Desc:  "Akcja odpalająca pakiet testów przeszukiwania",
	})
	cli.SetItemSupported(cliarguments.LineServiceItem{
		Name:      "records-limit",
		Level:     1,
		Desc:      "Ilość przerzucanych rekordów",
		IsRequire: true,
	})
	cli.SetItemSupported(cliarguments.LineServiceItem{
		Name:      "records-offset",
		Level:     1,
		Desc:      "Offset wybierania rekordów",
		IsRequire: true,
	})
	cli.ServiceCmdNew()
}

package main

import (
	"fmt"
	"github.com/uniwizardcom/golang-cliarguments/cliarguments"
	"strconv"
)

func main() {
	cli := cliarguments.LineService{
		Description: "Proof of concept for RediSearch functionality",
		Using:       "Programu uzywa się standardowo ;)",
	}
	cli.SetItemSupported(cliarguments.LineServiceItem{
		Name:      "insert-data",
		Desc:      "Akcja dodająca-resetująca zestaw danych (łącznie z indexami) w Redis na podstawie danych z PostgreSQL",
		IsRequire: true,
		// DefaultValue: "insert",
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

	idv := cli.GetArgValue("insert-data", 0, "")
	rlv := cli.GetArgValue("records-limit", 2, "")
	rov := cli.GetArgValue("records-offset", 2, "")

	if len(idv) > 0 {
		rl, _ := strconv.ParseInt(rlv, 10, 32)
		ro, _ := strconv.ParseInt(rov, 10, 32)
		fmt.Printf("records-limit=%d; records-offset=%d", rl, ro)
	}
}

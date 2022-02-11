package cliarguments

import (
	"fmt"
)

const (
	ALL_ITEMS = -1
)

func addTabs(cont string, contLenMax int) string {
	spaceLen := contLenMax - len(cont)

	for i := 0; i < spaceLen; i++ {
		cont += " "
	}

	return cont
}

func showInfoAs(title string, items []LineServiceItem) {
	fmt.Printf("%s:\n", title)

	var mess string
	for _, item := range items {
		mess = fmt.Sprintf("- [%s] on level [%d]", item.Name, item.Level)
		if item.IsRequire {
			mess += " (Required!)"
		}
		if len(item.DefaultValue) > 0 {
			mess += fmt.Sprintf(" with value [%s]", item.DefaultValue)
		}
		fmt.Print(addTabs(mess, 50))

		if len(item.Desc) > 0 {
			fmt.Printf(": %s", item.Desc)
		}

		fmt.Println()
	}
}

func showInfoAsBetween(title string, items []LineServiceItem, itemsOpposite []LineServiceItem) {
	fmt.Printf("%s:\n", title)

	for _, item := range items {
		fmt.Printf("- [%s] on level [%d]", item.Name, item.Level)

		if opposite := getItemByNameLevel(item.Name, ALL_ITEMS, itemsOpposite); opposite != nil {
			fmt.Printf(" <=> when supported is: [%s] on level [%d]", opposite.Name, opposite.Level)
		}

		fmt.Println()
	}

	fmt.Println("")
}

func (cli *LineService) ShowHelp() {
	if len(cli.Description) > 0 {
		fmt.Println(cli.Description + "\n")
	}

	if len(cli.Using) > 0 {
		fmt.Printf("Using:\n%s\n\n", cli.Using)
	}

	showInfoAs("Arguments supported", cli.itemsSupported)
	fmt.Println()
}

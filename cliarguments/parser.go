package cliarguments

import (
	"fmt"
	"os"
)

type LineServiceItem struct {
	IsRequire    bool
	Name         string
	Desc         string
	DefaultValue string
	Level        int
	Related      []string
}

type LineService struct {
	args             []string
	items            []LineServiceItem
	itemsSupported   []LineServiceItem
	itemsUnsupported []LineServiceItem
}

func difference(slice1 []LineServiceItem, slice2 []LineServiceItem) []LineServiceItem {
	var diff []LineServiceItem

	for _, s1 := range slice1 {
		found := false
		for _, s2 := range slice2 {
			if s1.Name == s2.Name && s1.Level == s2.Level {
				found = true
				break
			}
		}

		if !found {
			diff = append(diff, s1)
		}
	}

	return diff
}

func parseArg(arg string) (key string, level int, value string) {
	levIsFull := false
	kvIsFull := false

	for index := 0; index < len(arg); index++ {
		if !levIsFull && (arg[index] == '-') {
			level++
		} else if !kvIsFull && (arg[index] != '=') {
			levIsFull = true
			key = key + string(arg[index])
		} else if !kvIsFull && (arg[index] == '=') {
			kvIsFull = true
		} else {
			value = value + string(arg[index])
		}
	}

	return key, level, value
}

func (cli *LineService) importItemsFromArgs() {
	cli.args = os.Args

	for i, val := range cli.args {
		if i > 0 {
			key, level, value := parseArg(val)

			cli.items = append(cli.items, LineServiceItem{
				Name:         key,
				DefaultValue: value,
				Level:        level,
			})
		}
	}
}

func (cli *LineService) SetItemSupported(item LineServiceItem) {
	cli.itemsSupported = append(cli.itemsSupported, item)
}

func (cli *LineService) CheckForSupporting() {
	cli.itemsUnsupported = difference(cli.items, cli.itemsSupported)

	if len(cli.itemsUnsupported) > 0 {
		fmt.Printf("There is unsapported arguments:\n")

		for _, item := range cli.itemsUnsupported {
			fmt.Printf("- [%s] on level [%d] with value [%s]\n", item.Name, item.Level, item.DefaultValue)

		}
	}
}

func (cli *LineService) ServiceCmdNew() {
	cli.importItemsFromArgs()
	cli.CheckForSupporting()
}

// GetArg
// Getting item struct of argument by:
// - name
// - level, if greater or equal then 0 or all level if -1
func (cli *LineService) GetArg(name string, level int) {

}

// Copyright (c) 2022 - Uniwizard Wojciech Niewiadomski <wojtek@uniwizard.com>
// https://uniwizard.com
//
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package cliarguments

import (
	"fmt"
	"os"
)

type LineServiceItemLink struct {
	Name  string
	Level int
}

type LineServiceItem struct {
	IsRequire    bool
	Name         string
	Desc         string
	DefaultValue string
	Level        int
	Related      []LineServiceItemLink
	Action       func(*LineService)
}

type LineService struct {
	Description      string
	Using            string
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

func (cli *LineService) checkForSupporting() {
	cli.itemsUnsupported = difference(cli.items, cli.itemsSupported)

	if len(cli.itemsUnsupported) > 0 {
		cli.ShowHelp()
		showInfoAsBetween("Arguments not supported", cli.itemsUnsupported, cli.itemsSupported)
		exitProgram(1)
	}

	// TODO: it's should be implemented
	/* for _, item := range cli.items {
		if item.Action != nil {
			item.Action(cli)
		}
	} */
}

func (cli *LineService) checkForRequire() {
	requiredCandidates := difference(cli.itemsSupported, cli.items)
	required := make([]LineServiceItem, 0)

	for _, item := range requiredCandidates {
		fmt.Println()
		if item.IsRequire {
			required = append(required, item)
		}
	}

	if len(required) > 0 {
		cli.ShowHelp()
		showInfoAs("Arguments is require", required)
		exitProgram(1)
	}
}

// checkForRelatedSupporting
// TODO: it's should be implemented
func (cli *LineService) checkForRelatedSupporting(item LineServiceItem) {

}

func (cli *LineService) ServiceCmdNew() {
	cli.importItemsFromArgs()

	cli.SetItemSupported(LineServiceItem{
		Name:  "h",
		Level: 1,
		Desc:  "Show this help and stop program",
	})
	cli.checkForSupporting()
	cli.checkForRequire()
}

func getItemsByName(name string, coll []LineServiceItem) []*LineServiceItem {
	res := make([]*LineServiceItem, 0)

	for _, item := range coll {
		if item.Name == name {
			res = append(res, &item)
		}
	}

	return res
}

func getItemByNameLevel(name string, level int, coll []LineServiceItem) *LineServiceItem {
	for _, item := range coll {
		if item.Name == name && (level == ALL_ITEMS || item.Level == level) {
			return &item
		}
	}

	return nil
}

func exitProgram(code int) {
	os.Exit(code)
}

// GetArg
// Getting item struct of argument by:
// - name
// - level, if greater or equal then 0 or all level if -1
func (cli *LineService) GetArg(name string, level int) *LineServiceItem {
	return getItemByNameLevel(name, level, cli.items)
}

// GetArgValue
// Getting item struct of argument by:
// - name
// - level, if greater or equal then 0 or all level if -1
func (cli *LineService) GetArgValue(name string, level int, defaultValue string) string {
	lsi := cli.GetArg(name, level)

	if lsi != nil {
		return lsi.DefaultValue
	}

	return defaultValue
}

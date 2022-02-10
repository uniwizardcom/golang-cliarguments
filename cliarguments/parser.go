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

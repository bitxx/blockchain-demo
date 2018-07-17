package service

import "fmt"

func RandDelegates(delegates []string) []string{
	var randList []string
	randList = delegates[1:]
	randList = append(randList,delegates[0])
	fmt.Printf("%v\n",randList)
	return randList
}



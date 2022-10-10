package main

import (
	"strconv"
	"strings"
)

func isIsomorphic(s string, t string) bool {
	return getStruct(s) == getStruct(t)
}

func getStruct(s string) string {
	sIds := map[string]int{}
	var indexs []string
	for i := 0; i < len(s); i++ {
		c := string(s[i])
		index, exist := sIds[c]
		if !exist {
			index = len(sIds)
			sIds[c] = index
		}
		indexs = append(indexs, strconv.Itoa(index))
	}
	return strings.Join(indexs, "|")
}

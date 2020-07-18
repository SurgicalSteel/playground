package main

import (
	"fmt"
)

func main() {
	urlObject := Initialize("http://www.just.testing.com/bla/bla/bla?user=SurgicalSteel&id=666&name=John%20Petrucci")
	printDomainName(urlObject)
	printPath(urlObject)
	printQueryMap(urlObject)
}

func printDomainName(u URLParser) {
	domainName, err := u.GetHostName()
	if err != nil {
		fmt.Println("[GetDomainName]", err.Error())
		return
	}
	fmt.Println("[GetDomainName] Domain Name :", domainName)
}

func printPath(u URLParser) {
	path, err := u.GetPath()
	if err != nil {
		fmt.Println("[GetPath]", err.Error())
		return
	}
	fmt.Println("[GetPath] Path :", path)
}

func printQueryMap(u URLParser) {
	queryMap, err := u.GetQueryMap()
	if err != nil {
		fmt.Println("[GetQueryMap]", err.Error())
		return
	}
	fmt.Println("[GetQueryMap] Query Map :", queryMap)
}

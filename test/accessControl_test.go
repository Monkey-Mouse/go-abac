package test

import (
	"encoding/json"
	"fmt"
	"github.com/Monkey-Mouse/go-abac"
	"log"
	"os"
)

func ExampleAccessControl_SetGrant() {
	var ac abac.AccessControl
	ac.Grants = make(abac.GrantsType)

	var info abac.IAccessInfo
	info = abac.IAccessInfo{
		Subject:  "sub",
		Action:   "act",
		Resource: "res",
		Rules:    []string{"1", "2", "3"},
	}
	res, err := json.Marshal(ac)
	if err != nil {
		log.Fatal(err)
	} else {
		os.Stdout.Write(res)
	}
	//encoding.TextMarshaler()

	ac.SetGrant(info)

	res, err = json.Marshal(ac)
	if err != nil {
		log.Fatal(err)
	} else {
		os.Stdout.Write(res)

	}
	fmt.Println(ac)

	//output:
	//{"Grants":{}}{"Grants":{"sub":[{"res":[{"act":[["1","2","3"]]}]}]}}{ map[sub:[map[res:[map[act:[[1 2 3]]]]]]]}
}
func ExampleAccessControl_Grant() {
	var ac abac.AccessControl
	ac.Grant(abac.GrantsType{})
	fmt.Println(ac)
	grants := ac.Grants
	grants["test"] = []abac.ResourceGrantsType{{}}
	fmt.Println(ac)

	//output:
	//{ map[]}
	//{ map[test:[map[]]]}
}

func ExampleAccessControl_Can() {
	var ac abac.AccessControl
	ac.Role("test")
	fmt.Println(ac)
	ac.Role("test").Role("another")
	fmt.Println(ac)

	//Output:
	//{test map[]}
	//{another map[]}

}

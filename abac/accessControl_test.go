package abac

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var fooAC AccessControl
var subject1 = SubjectType("account")

func init() {
	fooAC.Grant(GrantsType{"account": ResourceGrantsType{"book": ActionGrantsType{ActionCreate: RulesType{"has user role"}}},
		"role": ResourceGrantsType{"project": ActionGrantsType{ActionCreate: RulesType{"has primer user role"}}},
	})
}

func ExampleAccessControl_SetGrant() {
	var ac AccessControl
	ac.Grants = make(GrantsType)

	var info IAccessInfo
	info = IAccessInfo{
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
	//{"Grants":{}}{"Grants":{"sub":{"res":{"act":["1","2","3"]}}}}{ map[sub:map[res:map[act:[1 2 3]]]]}
}
func ExampleAccessControl_Grant() {
	var ac AccessControl
	ac.Grant(GrantsType{})
	fmt.Println(ac)
	grants := ac.Grants
	grants["test"] = ResourceGrantsType{}
	fmt.Println(ac)

	//output:
	//{ map[]}
	//{ map[test:map[]]}
}

func ExampleAccessControl_GetGrants() {
	fmt.Println(fooAC.GetGrants())

	//output:
	//map[account:map[book:map[create:[has user role]]] role:map[project:map[create:[has primer user role]]]]
}

func ExampleGrantsType_GetSubject() {
	fmt.Println(fooAC.GetGrants().GetSubject(subject1))

	//output:
	//map[book:map[create:[has user role]]]
}

func ExampleResourceGrantsType_GetResource() {
	fmt.Println(fooAC.GetGrants().GetSubject(subject1).GetResource("book"))

	//output:
	//map[create:[has user role]]
}

func ExampleActionGrantsType_GetAction() {
	fmt.Println(fooAC.GetGrants().GetSubject(subject1).GetResource("book").GetAction(ActionDelete))
	fmt.Println(fooAC.GetGrants().GetSubject(subject1).GetResource("book").GetAction(ActionCreate))

	//output:
	//[]
	//[has user role]
}

func ExampleAccessControl_Can() {
	var ac AccessControl
	ac.Role("test")
	fmt.Println(ac)
	ac.Role("test").Role("another")
	fmt.Println(ac)

	//Output:
	//{test map[]}
	//{another map[]}

}

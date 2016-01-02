package main

import (
	singlyLinkedList "github.com/dustinspecker/go-singly-linked-list"
	"bufio"
	"os"
	"fmt"
	"strings"
	"errors"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var list1, list2 *singlyLinkedList.SinglyLinkedList
	var err error
	//Get List1
	if list1,err = getList(reader,list1);(err != nil){
		fmt.Println(err)
		os.Exit(1)
	}
	//Get List2
	if list2,err = getList(reader,list2);(err != nil) {
		fmt.Println(err)
		os.Exit(1)
	}
//	fmt.Printf("%d\t%d\n",list1.Length(),list2.Length())
	if val,err := findFirstCommon(list1,list2);(err != nil){
		fmt.Println(err)
	}else{
		switch v := val.(type) {
		case string:
			fmt.Println(v)
		case int32, int64:
			fmt.Println(v)
		default:
			fmt.Println("unknown")
		}
	}
}

/**
Find the first common element in list1
 */
func findFirstCommon(list1 *singlyLinkedList.SinglyLinkedList,list2 *singlyLinkedList.SinglyLinkedList)(interface{},error){
//	var p singlyLinkedList.Node
	p := list1.Front()
	q := list2.Front()
	for p!= nil{
		temp := q;
		for temp!=nil{
			fmt.Printf("P Val %s Q Val %s\n",p.Value,temp.Value)
			if(temp.Value == p.Value){
				return p.Value,nil
			}
			temp = temp.Next()
		}
		p = p.Next()
	}
	return nil,errors.New("Not found")
}

func getList(reader *bufio.Reader, list *singlyLinkedList.SinglyLinkedList) (*singlyLinkedList.SinglyLinkedList, error) {
	if listAsStr, err := readList(reader); (err == nil) {
		fmt.Printf("%q\n", listAsStr)
		list = singlyLinkedList.New()
		for _, element := range listAsStr {
			n1 := &singlyLinkedList.Node{Value: element}
			list.Append(n1)
		}
		return list, nil
	} else {
		return nil, err
	}
}

func readList(reader *bufio.Reader) ([]string, error) {
	fmt.Print("Enter the elements of List(Ex: 4 2 6 1 5): ")
	if text, err := reader.ReadString('\n'); (err == nil) {
		fmt.Println(text)
		//Remove white space chars
		text = strings.TrimSpace(text)
		//Split List elements
		elements := strings.Split(text, " ")
		return elements, nil;
	}else {
		fmt.Println(err)
		return nil, err;
	}
}
package list

import(
	"testing"
)

func Test_Add(t *testing.T){
	var lst LinkedList;
	element := Element { 1, nil}
	if !lst.Add(element){
		t.Error("LinkedList Add Element is Failing")
	}else{
		t.Log("LinkedList Add method works")
	}
}

func Test_Size(t *testing.T){
	var lst LinkedList;
	
	lst.Add(Element { 1, nil})
	lst.Add(Element{ 2, nil})	
	if lst.Size() != 2 {
		t.Error("LinkedList Size is Failing")
	}else{
		t.Log("LinkedList Size method works")
	}
}

func Test_Contains(t *testing.T){
	var lst LinkedList;
	
	lst.Add(Element { 1, nil})
	lst.Add(Element{ 2, nil})	
	if lst.Contains(Element{ 2, nil}) {
		t.Error("LinkedList Contains is Failing")
	}else{
		t.Log("LinkedList Contains method works")
	}
}

func Test_ContainsAll(t *testing.T){
	var lst,evens LinkedList;
	
	for i := 1; i<=(10); i++ {	
		lst.Add(Element{ i, nil})
		if i%2 ==0 {
			evens.Add(Element{ i, nil})	
		}else{
						
		}
	}

	if lst.ContainsAll(&evens)==false {
		t.Error("LinkedList ContainsAll is Failing")
	}else{
		t.Log("LinkedList ContainsAll method works")
	}
}

func Test_LastIndexOf(t *testing.T){
	var lst,evens LinkedList;	
	for i := 1; i<=(10); i++ {	
		lst.Add(Element{ i, nil})
		if i%2 ==0 {
			evens.Add(Element{ i, nil})	
		}
	}	

	if lst.LastIndexOf(*evens.Get(0))!= 1 {
		t.Error("LinkedList LastIndexOf is Failing")
	}else{
		t.Log("LinkedList LastIndexOf method works")
	}
}

func Test_Remove(t *testing.T){
	var lst LinkedList
	for i := 1; i<=(10); i++ {	
		lst.Add(Element{ i, nil})
	}
	node := lst.Get(0)
	lst.Remove(0)
	if lst.Get(0).Val == node.Val{
		t.Error("Linkedlist Remove is Failing")
	}else{
		t.Log("Linkedlist Remove works")
	}
}

func Test_AddAt(t *testing.T){
	var lst LinkedList;
	
	for i := 1; i<=(10); i++ {	
		lst.Add(Element{ i, nil})
	}
	node := lst.Get(1)
	element := Element{ 5, nil}
	lst.AddAt(element,1)
	t.Log(lst.Get(3).Val)
	if lst.Get(1).Val == node.Val || lst.Get(1).Val !=element.Val {
		t.Error("LinkedList AddAt is Failing")
	}else{
		t.Log("LinkedList AddAt method works")
	}
}
func Test_RemoveElement(t *testing.T){
	var lst LinkedList;
	
	for i := 1; i<=(10); i++ {	
		lst.Add(Element{ i, nil})
	}
	node := lst.Get(1)	
	lst.RemoveElement(*node)
	if lst.Get(1).Val == node.Val {
		t.Error("LinkedList RemoveElement is Failing")
	}else{
		t.Log("LinkedList RemoveElement method works")
	}
}

func Test_RemoveAll(t *testing.T){
	var lst,evens LinkedList;	
	for i := 1; i<=(10); i++ {	
		lst.Add(Element{ i, nil})
		if i%2 ==0 {
			evens.Add(Element{ i, nil})	
		}
	}	

	if lst.RemoveAll(&evens)!=true {
		t.Error("LinkedList RemoveAll is Failing")
	}else{
		t.Log("LinkedList RemoveAll method works")
	}
}
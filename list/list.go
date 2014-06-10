package list

import (
    "reflect" 
    "fmt" 
    "unsafe"      
)


type List interface{
      Add(element Element ) bool
      AddAt(element Element , index int) bool
      AddAll(element interface{} , index int) bool
      Clear()
      Contains(element Element) bool
      ContainsAll(elements List) bool
      Equals(element interface{}) bool
      Get(index int) *Element
      IndexOf(element interface{}) int
      IsEmpty() bool
      LastIndexOf(element Element) int
      Remove(index int) *Element
      RemoveElement(element Element) bool
      RemoveAll(elements List) bool
      RetainAll(elements interface{}) bool
      Set(index int, element interface{}) interface{}
      Size() int
      SubList(fromIndex int, toIndex int ) *List
      SetType(listType reflect.Type)
}


type LinkedList struct{
     head *Element
     tail *Element
     Type reflect.Type
     size int
}

//Definition of a Element
type Element struct{
   Val interface{}
   Next *Element
}

func (n *LinkedList) Add(element Element ) bool{
    var isAdded bool
    isAdded = false
    if n !=nil && element.Val !=nil{
        if n.head ==nil{
            n.head = &element
            n.tail = &element
            n.size++
            isAdded = true
        }else if n.head !=nil && n.tail !=nil {
            n.tail.Next = &element
            n.tail = &element
            n.size++
        }
    }
    return isAdded

}

func (n *LinkedList) AddAt(element Element , index int) bool {
  var isAddedAt bool  = false
  if !n.IsEmpty() && index <=(n.Size()-1){
	  
    if index == 0{
      element.Next = n.head
      n.head = &element
    }else{
       node:=n.head
      counter:=0

      for node.Next != nil {      
        if counter+1 == index {
          break
        }
        node = node.Next      
        counter++
      }
      if node.Next!=nil {  
        fmt.Println(element.Next)   
        element.Next = node.Next
        node.Next = &element
        isAddedAt = true  
      }
    }   	  
  }
  return isAddedAt
}
func (n *LinkedList) AddAll(element interface{} , index int) bool { return false}
func (n *LinkedList) Clear(){
	node := n.head
	for node.Next!=nil{
		node.Val = nil;
		valType := unsafe.Sizeof(node.Val)
		fmt.Println(valType)
		node = node.Next
	}
}
func (n *LinkedList) Contains(element Element) bool{ 
	isPresent := false
	if !n.IsEmpty() {
		node := n.head
		for node.Next != nil {
			if node.Val !=nil && node.Val == element.Val {
				isPresent = true
			}
			node = node.Next
		}
	}
	return isPresent

}
func (n *LinkedList) ContainsAll(elements List) bool{
	var containsAll bool = true
	containsMap := make(map[*Element]bool)
	if elements!=nil && n!=nil {
		head := elements.Get(0)
		for head!=nil && head.Next !=nil {			
			containsAll = n.Contains(*head)
			fmt.Println(containsAll)
			containsMap[head] = containsAll
			if containsAll==false {
				break;
			}
			head = head.Next
		}
	}
	return containsAll
}
func (n *LinkedList) Equals(element interface{}) bool{return false}
func (n *LinkedList) Get(index int) *Element {
    var elementAt *Element = nil
    if !n.IsEmpty() && index <= (n.Size()-1)  {
       node := n.head;
       counter := 0
       for (node.Next!=nil && counter < index){
             node = node.Next
             counter++
        }
        elementAt = node
    }
    return elementAt
}
func (n *LinkedList) IndexOf(element interface{}) int{return 0}
func (n *LinkedList) IsEmpty() bool{
	var isEmpty bool = false
	if n==nil {
		isEmpty = true
	}
	return isEmpty
}
func (n *LinkedList) LastIndexOf(element Element) int{ 
  var position int = -1;
  if !n.IsEmpty() {
    node := n.head
    counter := 0
    for node.Next!= nil {
      if node.Val == element.Val {
        position = counter
      }
      counter++
      node = node.Next
    }
  }
  return position
}
func (n *LinkedList) Remove(index int) *Element{
    var deletedNode *Element
   if !n.IsEmpty() && index <= (n.Size()-1) {
      node := n.head
      counter := 0
      if index==0{
          if n.Size() > 1 {
            deletedNode = n.head
            n.head =  n.head.Next
            n.size--
          }else{
            deletedNode = n.head
            n.head = nil
            n.size = 0  
          }
          
        }else{
           for node.Next!= nil {
              if index == (counter+1){
                deletedNode = node.Next
                node.Next = node.Next.Next
                break;
              }
              counter++
              node = node.Next
          }
        }     
   }
   return deletedNode
}
func (n *LinkedList) RemoveElement(element Element) bool{
  var isRemoved bool = false
  if !n.IsEmpty(){
    node := n.head
    for node.Next != nil{
      if node.Next.Val == element.Val {
        node.Next = node.Next.Next
        isRemoved = true
      }
      node = node.Next
    }
  }
  return isRemoved
}
func (n *LinkedList) RemoveAll(elements List) bool{
  var areAllDeleted bool = false
   
   if !n.IsEmpty() && !elements.IsEmpty() {
    node:=elements.Get(0)
    for node.Next != nil {
      areAllDeleted = n.RemoveElement(*node)
      if !areAllDeleted {
        break
      }
      node= node.Next
    }
   }

  return areAllDeleted
}
func (n *LinkedList) RetainAll(elements interface{}) bool{return false}
func (n *LinkedList) Set(index int, element interface{}) interface{}{ return nil}
func (n *LinkedList) Size() int{ return n.size }
func (n *LinkedList) SubList(fromIndex int, toIndex int ) *List {
    var lstPtr List;
    lstPtr = n
    return &lstPtr
}
func (n *LinkedList) SetType(listType reflect.Type) {}



package main

var stack1 [] int
var stack2 [] int

func Push(node int) {
    stack1 = append(stack1,node)
}

func Pop() int{
    if(len(stack2)==0){
        if(len(stack1)==0){
            value:= -1
            return value
        }
        for len(stack1) > 0{
            index := len(stack1)-1
            value := stack1[index]
            stack2 = append(stack2, value)
            stack1 = stack1[:index]
        }
    }
     
    index := len(stack2)-1
    value := stack2[index]
    stack2 = stack2[:index]
    return value
}
package greetings

import "fmt"
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

//FunctionName(name paramType) returnType{
	//:= declare and initialize variable in one line
	//equal to var message string
	//message = fmt.Sprintf('HI %v, blahblah', name)
//}
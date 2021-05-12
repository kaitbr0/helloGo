package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)
func Hello(name string) (string, error) {
	if name == ""{
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

//FunctionName(name paramType) returnType{
	//:= declare and initialize variable in one line
	//equal to var message string
	//message = fmt.Sprintf('HI %v, blahblah', name)
//}
func Hellos(names []string) (map[string]string, error){
	messages := make(map[string]string)
//loop through names and map a custom message to each name
	for _, name := range names{
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func init(){
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
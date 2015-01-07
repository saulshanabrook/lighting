package mock

import "fmt"

type Mock struct{}

func (m Mock) Send(d Data) (err error) {
	_, err = fmt.Println(d)
	return
}

// eazyE is a simple sugar for error checking. It is a pattern I use often enough,
// so I decided to make a quick library out of it.
package eazyE

import (
	"fmt"
	"log"
)

func New(s string, err error) (fErr error) {
	if err != nil {
		fErr = fmt.Errorf(s, err)
		return
	}
	return
}

func Newl(s string, err error) {
	if err != nil {
		log.Println(fmt.Errorf(s, err))
		return
	}
}

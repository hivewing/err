package err

import (
	"fmt"
)

func Check(e error) {
	if e != nil {
		fmt.Println("error:", e)
		panic(e)
	}
}

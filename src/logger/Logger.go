package logger

import (
	"fmt"
)

type Logger struct {
	On bool
}

func (l Logger) Log(args ...interface{}) {
	if l.On {
		fmt.Println(args);
	}
}

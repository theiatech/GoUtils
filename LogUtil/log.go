package LogUtil

import (
	"fmt"
	"github.com/theiatech/GoUtils/TimeUtil"
	"runtime"
)

func Log(str string)  {
	_,file,line,_ := runtime.Caller(1)
	fmt.Printf("\r\u001B[K%s [%s:%d] %s",TimeUtil.Now(),file,line,str)
}

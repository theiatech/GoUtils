package LogUtil

import (
	"fmt"
	"github.com/theiatech/GoUtils/FontColor"
	"github.com/theiatech/GoUtils/TimeUtil"
	"runtime"
)

func Log(str string)  {
	fmt.Printf("\r\u001B[K%s %s",FontColor.Green(TimeUtil.Now()),str)
}

func LogDebug(str string)  {
	_,file,line,_ := runtime.Caller(1)
	Log(fmt.Sprintf("%s:%d %s",file,line,str))
}


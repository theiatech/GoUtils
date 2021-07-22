package LogUtil

import (
	"fmt"
	"github.com/theiatech/GoUtils/FontColor"
	"github.com/theiatech/GoUtils/TimeUtil"
	"runtime"
)

func Log(str string)  {
	fmt.Printf("\r\u001B[K%s",str)
}

func Logf(format string, a ...interface{}) {
	Log(fmt.Sprintf(format,a...))
}

func LogInfo(str string) {
	Log(fmt.Sprintf("%s %s\n",FontColor.Green(TimeUtil.Now()),str))
}
func LogInfof(format string, a ...interface{}) {
	LogInfo(fmt.Sprintf(format,a...))
}

func LogDebug(str string)  {
	_,file,line,_ := runtime.Caller(1)
	Log(fmt.Sprintf("%s %s:%d %s\n",FontColor.Green(TimeUtil.Now()),file,line,str))
}
func LogDebugf(format string, a ...interface{}) {
	LogDebug(fmt.Sprintf(format,a...))
}

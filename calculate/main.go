package main

import (
	"go.uber.org/zap"
)

func main() {
	L, _ := zap.NewDevelopment()
	//L := zzap.IintZapLogger()
	//IintZapLogger()
	//var a float64 = 1.0
	//var b float64 = 1.0
	//var sum float64 = a + b
	//fmt.Println(sum)
	L.Sugar().Debug("sguasgfasgfa")
}

/*
func IintZapLogger() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

}
*/

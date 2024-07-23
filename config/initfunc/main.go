package main

import "fmt"

type Config struct {
	ConfigFileName string
	ConfigType     string
	ConfigPath     string
	Verbose        bool
	VersionShow    bool
}

var GlobalConfig = Config{
	ConfigFileName: "config.yaml",
	ConfigType:     "yaml",
	ConfigPath:     "./config",
	Verbose:        false,
	VersionShow:    false,
}

func main() {
	fmt.Println(GlobalConfig)
	fmt.Printf("Config Path: %s\n", GlobalConfig.ConfigPath)

}

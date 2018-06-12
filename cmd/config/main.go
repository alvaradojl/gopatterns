package main

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("hello there")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("relayservices")
	viper.BindEnv("color")

	err := viper.ReadInConfig()

	if err != nil { // Handle errors reading the config file
		log.Fatal(err)
	}

	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	host := viper.Get("datastore.metric.host")

	fmt.Printf("\nthe config setting for datastore.metric.host is: %s\n", host)

	isEnabled := viper.GetBool("datastore.warehouse.enabled")

	//var envVariable string
	envVariable := viper.Get("color")

	fmt.Printf("\nEnvironment Variable RELAYSERVICES_COLOR is: %s\n", envVariable)

	if isEnabled {
		fmt.Println("The enabled setting is on")
	} else {
		fmt.Println("The enabled setting is off")
	}

}

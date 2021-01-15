package main

import (
	"fmt"
	"livefun/global"
	"livefun/router"
	"net/http"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {
	v := viper.New()
	v.SetConfigFile("app.yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
		fmt.Println(err)
	}
	router := router.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", global.GVA_CONFIG.App.Addr),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}

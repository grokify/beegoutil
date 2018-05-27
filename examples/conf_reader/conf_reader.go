package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/gotilla/os/osutil"
	ou "github.com/grokify/oauth2more"
	ms "github.com/grokify/oauth2more/multiservice"
	"github.com/joho/godotenv"
)

var ConfPath = "github.com/grokify/beego-oauth2-demo/conf/app.conf"

func main() {
	confPath := filepath.Join(os.Getenv("GOPATH"), "src", ConfPath)

	err := godotenv.Load(confPath)
	if err != nil {
		panic(err)
	}

	fmt.Println(os.Getenv("oauth2configgoogle"))

	ac, err := ou.NewAppCredentialsWrapperFromBytes([]byte(os.Getenv("oauth2configgoogle")))
	if err != nil {
		panic(err)
	}
	cfg, err := ac.Config()
	if err != nil {
		panic(err)
	}
	fmtutil.PrintJSON(cfg)

	env := osutil.Env()
	fmtutil.PrintJSON(env)

	cfgs, err := ms.EnvOAuth2ConfigMap(env, "oauth2config")
	if err != nil {
		panic(err)
	}
	fmtutil.PrintJSON(cfgs)
	fmt.Println("DONE")
}

package configx

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func Init() {

	fmt.Printf("[CONFIGX]  START INIT CONFIGS\n")
	fmt.Printf("[CONFIGX]  SET CONFIGS PATH\n")
	viper.SetConfigName("go-backend") // ชื่อ config file
	viper.AddConfigPath("./conf")     // ระบุ path ของ config file
	viper.AutomaticEnv()              // อ่าน value จาก ENV variable

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// อ่าน config
	fmt.Printf("[CONFIGX]  READ CONFIGS FILE\n")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}

	fmt.Printf("[CONFIGX]  INIT CONFIGS COMPLETED\n")

}

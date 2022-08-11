package main

// "github.com/spf13/viper"

type ServerConfig struct {
	Server Server `json:"server"`
}
type Server struct {
	URL      string `json:"url"`
	Redirect string `json:"redirect"`
	API      string `json:"api"`
}

// func main() {
// 	viper.SetConfigType("yaml")
// 	viper.SetConfigName("config")
// 	viper.AddConfigPath(".")
// 	// viper
// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(viper.GetString("server.url"))
// 	fmt.Println(viper.GetString("server.redirect"))
// 	fmt.Println(viper.AllKeys())
// 	c := &ServerConfig{}
// 	if err = viper.Unmarshal(&c); err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(c)
// }

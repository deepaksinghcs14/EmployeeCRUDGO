package server

import "EmployeeCRUD/config"

func Init() {
	viper := config.GetConfig()
	r := NewRouter()
	err := r.Run(viper.GetString("server.port"))
	if err != nil {
		print(err)
		return
	}
}

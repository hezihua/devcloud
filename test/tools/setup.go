package tools

// 开发环境 单元测试的Setup
func DevelopmentSet() {
	// 测试测试用例的配置
	// err := conf.LoadConfigFromEnv()
	// if err != nil {
	// 	panic(err)
	// }

	// 初始化测试类的注册
	if err := apps.InitApps(); err != nil {
		panic(err)
	}
}

package common

func StartUp() {
	// Initialize AppConfig variable
	InitConfig()
	// Start a MySQL session
	CreateDbSession()
}

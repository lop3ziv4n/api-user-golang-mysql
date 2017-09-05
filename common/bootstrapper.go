package common

func StartUp() {
	// Initialize AppConfig variable
	initConfig()
	// Start a MySQL session
	createDbSession()
}

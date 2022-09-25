package main

import "github.com/WilliamsStudentsOnline/onboarding-backend/internal/router"

// @title           WSO Backend Onboarding Project
// @version         1.0
// @description     This is a sample project for onboarding WSO members.

// @contact.name   WSO Devs
// @contact.url    https://wso.williams.edu/
// @contact.email  wso-dev@wso.williams.edu

// @host      localhost:8080

// @securityDefinitions.basic BasicAuth

func main() {
	r := router.SetupRouter()

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

package main

import "github.com/LegationPro/go-logger/logger"

func main() {
	logger.NewLogger("logs")

	logger.GetLogger().Clean()
}

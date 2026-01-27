package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	// 添加时间戳，方便追踪
	log.Printf("🚀 FractalBot starting at %s\n", time.Now().Format(time.RFC3339))
	log.Printf("📋 Working directory: %s\n", getWorkingDir())
	log.Println("")

	// 简单的版本信息
	fmt.Println("FractalBot - Multi-Agent Orchestration System")
	fmt.Println("Version: 0.1.0-alpha")
	fmt.Println()
	fmt.Println("✅ Core components initialized")
	fmt.Println("📡 Gateway server: Not yet implemented")
	fmt.Println("📡 Channel manager: Not yet implemented")
	fmt.Println("📡 Agent runtime: Not yet implemented")
	fmt.Println()
	fmt.Println("🎯 Ready for development!")
	fmt.Println("   - Go modules will be added incrementally")
	fmt.Println("   - Channel integrations will be implemented")
	fmt.Println("   - Agent coordination will be added")
	fmt.Println()
	fmt.Println("💡 Run with --help to see available commands")
}

func getWorkingDir() string {
	if dir, err := os.Getwd(); err == nil {
		return dir
	}
	return "unknown"
}

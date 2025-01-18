package main

import (
	"fmt"
	"github.com/zcalusic/sysinfo"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func main() {
	var si sysinfo.SysInfo
	si.GetSysInfo()
	println(Green + "OS: " + Reset + runtime.GOOS)
	println(Green + "Arch: " + Reset + runtime.GOARCH)
	println(Green + "CPU: " + Reset + si.CPU.Model)
	println(Green + "OS name: " + Reset + si.OS.Name)
	println(Green + "OS version: " + Reset + si.OS.Version)
	println(Green + "Product: " + Reset + si.Product.Name + " (" + si.Product.Vendor + ")")
	getCurrentTerminal()
	getStorage()
}

func getStorage() {
	fmt.Println(Green + "Storage Info:" + Reset)

	fmt.Printf("%-30s %-20s\n", Yellow+"Model", "Size (GB)"+Reset)
	fmt.Println("---------------------------------------------")

	out, err := exec.Command("lsblk", "-o", "NAME,SIZE,MODEL", "-d", "-n").Output()
	if err != nil {
		fmt.Println("Error executing lsblk:", err)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) >= 3 {
			model := fields[2]
			size := fields[1]
			fmt.Printf("%-30s %-20s\n", model, size)
		}
	}
}

func getCurrentTerminal() {
	shell := os.Getenv("SHELL")
	if shell == "" {
		ppid := os.Getppid()
		cmd := exec.Command("ps", "-p", fmt.Sprint(ppid), "-o", "comm=")
		output, err := cmd.Output()
		if err != nil {
			fmt.Println("Error detecting the current shell:", err)
			return
		}
		shell = strings.TrimSpace(string(output))
	}

	println(Green + "Terminal: " + Reset + shell)
}

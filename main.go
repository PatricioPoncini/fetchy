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
var Green = "\033[32m"
var Yellow = "\033[33m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"

func main() {
	printHeader()

	var si sysinfo.SysInfo
	si.GetSysInfo()

	fmt.Println(Cyan + "System Information:" + Reset)
	println(Green + "OS name:     " + Reset + si.OS.Name)
	println(Green + "OS version:  " + Reset + si.OS.Version)
	println(Green + "Arch:        " + Reset + runtime.GOARCH)
	println(Green + "CPU:         " + Reset + si.CPU.Model)
	println(Green + "Product:     " + Reset + si.Product.Name + " (" + si.Product.Vendor + ")")

	fmt.Println()
	getCurrentTerminal()
	getKernel()
	getMemory()
	getGPU()
	getStorage()

	fmt.Println(Cyan + "---------------------------------------------" + Reset)
}

func printHeader() {
	fmt.Println(Magenta + `
██████╗ ███████╗ ████████╗ ███████╗ ██╗  ██╗ ██╗  ██╗
██╔═══╗ ██╔════╝ ╚══██╔══╝ ██╔════╝ ██║  ██║ ██║  ██║
██████║ ██████╗     ██║    ██║      ███████║ ███████║
██╔═══╝ ██╔══╝      ██║    ██║      ██╔══██║    ██║
██║     ███████╗    ██║    ███████╗ ██║  ██║     ██║
╚═╝     ╚══════╝    ╚═╝    ╚══════╝ ╚═╝  ╚═╝      ██║
` + Reset)
	fmt.Println(Yellow + "             A Lightweight System Info Tool\n" + Reset)
}

func getStorage() {
	fmt.Println(Cyan + "\nStorage Info:" + Reset)

	fmt.Printf("%-30s %-20s\n", Yellow+"Model", "Size (GB)"+Reset)
	fmt.Println(Cyan + "---------------------------------------------" + Reset)

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

	println(Green + "Terminal:    " + Reset + shell)
}

func getKernel() {
	out, err := exec.Command("uname", "-r").Output()
	if err != nil {
		fmt.Println("Error getting kernel version:", err)
		return
	}
	println(Green + "Kernel:      " + Reset + strings.TrimSpace(string(out)))
}

func getMemory() {
	out, err := exec.Command("free", "-h").Output()
	if err != nil {
		fmt.Println("Error getting memory info:", err)
		return
	}

	lines := strings.Split(string(out), "\n")
	if len(lines) > 1 {
		mem := strings.Fields(lines[1])
		if len(mem) >= 3 {
			println(Green + "RAM:         " + Reset + mem[1] + " total, " + mem[2] + " used")
		}
	}
}

func getGPU() {
	out, err := exec.Command("lspci").Output()
	if err != nil {
		fmt.Println("Error getting GPU info:", err)
		return
	}

	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if strings.Contains(line, "VGA") || strings.Contains(line, "3D") {
			println(Green + "GPU:         " + Reset + strings.TrimSpace(line))
			return
		}
	}
	println(Green + "GPU:         " + Reset + "No GPU detected")
}

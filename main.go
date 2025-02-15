package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
)

var Reset = "\033[0m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"

// TODO: Errors with Red println()

func main() {
	printHeader()

	fmt.Println(Cyan + "System Information:" + Reset)
	getUser()
	getHostname()
	getOSName()
	getOSVersion()
	println(Green + "Arch:        " + Reset + runtime.GOARCH)
	getCPU()
	getProduct()

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
    ______     __       __         
   / ____/__  / /______/ /_  __  __
  / /_  / _ \/ __/ ___/ __ \/ / / /
 / __/ /  __/ /_/ /__/ / / / /_/ / 
/_/    \___/\__/\___/_/ /_/\__, /  
                          /____/   
` + Reset)
	fmt.Println(Yellow + "A Lightweight System Info Tool\n" + Reset)
}

func getOSName() {
	out, err := exec.Command("lsb_release", "-d").Output()
	if err != nil {
		fmt.Println("Error executing lsb_release:", err)
		return
	}

	output := string(out)
	description := strings.TrimSpace(strings.TrimPrefix(output, "Description:"))
	println(Green + "OS name:     " + Reset + description)
}

func getCPU() {
	out, err := exec.Command("lscpu").Output()
	if err != nil {
		fmt.Println("Error executing lscpu:", err)
		return
	}

	output := string(out)

	for _, line := range strings.Split(output, "\n") {
		if strings.HasPrefix(line, "Model name:") {
			cpuModel := strings.TrimSpace(strings.Split(line, ":")[1])
			println(Green + "CPU:         " + Reset + cpuModel)
			return
		}
	}

	fmt.Println("CPU Model not found")
}

func getOSVersion() {
	out, err := exec.Command("lsb_release", "-r").Output()
	if err != nil {
		fmt.Println("Error executing lsb_release:", err)
		return
	}

	output := string(out)
	version := strings.TrimSpace(strings.TrimPrefix(output, "Release:"))
	println(Green + "OS version:  " + Reset + version)
}

func getProduct() {
	productPath := "/sys/class/dmi/id/product_name"
	vendorPath := "/sys/class/dmi/id/board_vendor"

	product, err := os.ReadFile(productPath)
	if err != nil {
		fmt.Printf("Error reading product information: %v\n", err)
		return
	}

	vendor, err := os.ReadFile(vendorPath)
	if err != nil {
		fmt.Printf("Error reading vendor information: %v\n", err)
		return
	}

	productName := strings.TrimSpace(string(product))
	vendorName := strings.TrimSpace(string(vendor))

	if productName != "" && vendorName != "" {
		println(Green + "Product:     " + Reset + productName + " (" + vendorName + ")")
	} else {
		fmt.Println("Product information not found")
	}
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

func getUser() {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Printf("Unable to get current user: %s\n", err)
	}

	println(Green + "User:        " + Reset + currentUser.Username)
}

func getHostname() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Unable to get hostname: %s\n", err)
	}
	println(Green + "Hostname:    " + Reset + hostname)
}

# Fetchy 🐀

**Fetchy** is a lightweight system information tool written in Go. It provides detailed information about your system's hardware and software in a clean and colorful terminal output. Perfect for tech enthusiasts who want to quickly fetch system details.

## Features 🛠️

- **System Information**: Operating System, version, architecture, and product details.
- **Hardware Details**: CPU, GPU, memory, and storage information.
- **Terminal and Kernel**: Displays the current terminal and kernel version.
- **Simple and Lightweight**: Minimal dependencies, fast, and easy to use.

## Demo ▶️
```shell
    ______     __       __                                                                                                                                                                                                          
   / ____/__  / /______/ /_  __  __                                                                                                                                                                                                 
  / /_  / _ \/ __/ ___/ __ \/ / / /                                                                                                                                                                                                 
 / __/ /  __/ /_/ /__/ / / / /_/ /                                                                                                                                                                                                  
/_/    \___/\__/\___/_/ /_/\__, /                                                                                                                                                                                                   
                          /____/                                                                                                                                                                                                    
                                                                                                                                                                                                                                    
A Lightweight System Info Tool
                                                                                                                                                                                                                                    
System Information:
User:        gawain
Hostname:    kitt
OS name:     Ubuntu 22.04.5 LTS
OS version:  22.04
Arch:        amd64
CPU:         12th Gen Intel(R) Core(TM) i5-12400F
Product:     System Product Name (ASUSTeK COMPUTER INC.)

Terminal:    /usr/bin/zsh
Kernel:      6.8.0-48-generic
RAM:         6.72GB / 15.43GB [█████████████░░░░░░░░░░░░░░░░░] 44%
GPU:         01:00.0 VGA compatible controller: NVIDIA Corporation GM206 [GeForce GTX 950] (rev a1)

Storage Info:                                                                                                                                                                                                                       
Model                     Size (GB)       
---------------------------------------------
KINGSTON                       894,3G              
WD                             931,5G              
---------------------------------------------
```

## Requirements ⚙️
- **Go**: Version 1.19 or higher

## Installation 💻
1. Clone the repository:
```shell
git clone https://github.com/PatricioPoncini/fetchy.git
```
2. Navigate to the project directory:
```shell
cd fetchy
```
3. Install dependencies:
```shell
go mod tidy
```
4. Build and run:
```shell
make run
```
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
OS name:     Debian GNU/Linux 12 (bookworm)
OS version:  12
Arch:        amd64
CPU:         Intel(R) Core(TM) i5-10210U CPU @ 1.60GHz
Product:     20U4S4T000 (LENOVO)

Terminal:    /bin/zsh
Kernel:      6.1.0-28-amd64
RAM:         15Gi total, 5.5Gi used
GPU:         00:02.0 VGA compatible controller: Intel Corporation CometLake-U GT2 [UHD Graphics] (rev 02)

Storage Info:
Model                     Size (GB)       
---------------------------------------------
KINGSTON                       465.8G              
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
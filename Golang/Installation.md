Here is the complete content for `Installation.md` covering installation and setup of Go on Windows and Ubuntu, along with VSCode integration:



# 🚀 Go Installation and Setup Guide (Windows & Ubuntu)  

This guide will walk you through installing Golang and setting up VS Code on Windows and Ubuntu.  


## 🪟 1. Installing Go on Windows  

### 📥 Step 1: Download Go  
- Visit the official Go download page: [https://go.dev/dl/](https://go.dev/dl/)  
- Download the Windows installer (`.msi` file).  

### 💻 Step 2: Install Go  
1. Run the downloaded `.msi` file.  
2. Follow the installation prompts and complete the setup.  
3. By default, Go will install in:  
   ```text
   C:\Program Files\Go\
   ```
4. The installer will automatically add Go to your system PATH.  

### 🧪 Step 3: Verify Installation  
1. Open Command Prompt or PowerShell.  
2. Run the following commands:  
   ```bash
   go version
   go env
   ```
   You should see the Go version and environment details.  



## 🐧 2. Installing Go on Ubuntu (Linux)  

### 📥 Step 1: Update System  
```bash
sudo apt update && sudo apt upgrade -y
```

### 💻 Step 2: Install Go  
1. Download the latest Go binary (replace version as needed):  
   ```bash
   wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz
   ```
2. Extract and install Go:  
   ```bash
   sudo rm -rf /usr/local/go
   sudo tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz
   ```
3. Add Go to system PATH by editing your shell profile:  
   ```bash
   echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
   source ~/.bashrc
   ```

### 🧪 Step 3: Verify Installation  
```bash
go version
go env
```
You should see the installed Go version and environment details.  



## 🛠️ 3. Setting Up VS Code for Go (Windows & Ubuntu)  

### 📥 Step 1: Install Visual Studio Code  
- Download and install VS Code from [https://code.visualstudio.com/](https://code.visualstudio.com/)  

### 🧩 Step 2: Install Go Extension for VS Code  
1. Open VS Code.  
2. Go to Extensions (`Ctrl+Shift+X`).  
3. Search for `Go` and install the extension by the Go Team at Google.  

### 📝 Step 3: Configure VS Code for Go  
1. Open Settings (`Ctrl + ,`).  
2. Search for `go.gopath` and set your workspace path if needed.  
3. Create a Go workspace folder and open it in VS Code.  

### 🚀 Step 4: Test Your Go Installation  
1. Create a file named `main.go`:  
   ```go
   package main

   import "fmt"

   func main() {
       fmt.Println("Hello, Go!")
   }
   ```
2. Open the terminal in VS Code and run:  
   ```bash
   go run main.go
   ```
   You should see:  
   ```text
   Hello, Go!
   ```


## 🧹 Optional: Set Go Modules (Recommended)  
Enable Go modules for dependency management:  
```bash
go env -w GO111MODULE=on
```



### ✅ Congratulations! You have successfully installed and set up Golang with VS Code on Windows and Ubuntu. Happy Coding! 🚀💙  


# Hello World Go

[![ci pipeline](https://github.com/Obaid-Ghafoori/hello-world-go/actions/workflows/quality-checks.yml/badge.svg)](https://github.com/Obaid-Ghafoori/hello-world-go/actions/workflows/ci-pipeline.yml)
[![codecov](https://codecov.io/gh/Obaid-Ghafoori/hello-world-go/branch/main/graph/badge.svg)](https://codecov.io/gh/Obaid-Ghafoori/hello-world-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/Obaid-Ghafoori/hello-world-go)](https://goreportcard.com/report/github.com/Obaid-Ghafoori/hello-world-go)


## Overview
This project is designed to help automate branch protection rules or any other functionality specified in the code. Follow the steps below to get started and run the program on your local system.
A simple Hello World program written in Go.



## Prerequisites
Ensure you have the following installed on your system:
1. **Git**: [Download Git](https://git-scm.com/)
2. **Go** (Golang): [Install Go](https://golang.org/doc/install)
3. A terminal or command-line interface.

## Getting Started

### Step 1: Clone the Repository
Clone this repository to your local machine using the following command:
```bash
git clone https://github.com/Obaid-Ghafoori/hello-world-go
```



### Step 2: Navigate to the Project Directory
Change to the project directory where the repository was cloned:
```bash
cd hello-world-go
```

### Step 3: Make the Script Executable
Before running the hook installation script, ensure it has executable permissions:
```bash
chmod +x install-hooks.sh
```

### Step 4: Install Git Hooks
Run the `install-hooks.sh` script to set up necessary Git hooks for the repository:
```bash
./install-hooks.sh
```
This step installs the required hooks to manage Git-related actions like branch protection.

### Step 5: Run the Program
Finally, run the program using the `go run` command:
```bash
go run main.go
```

## Additional Information
- **Scripts**:
  - `install-hooks.sh`: Installs Git hooks that help manage specific repository actions.
- **Main Program**:
  - `main.go`: The entry point of the program written in Go.
  
Feel free to explore the codebase and modify it as needed for your use case.



# Commit Duration Checker

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.19+-00ADD8.svg?style=flat&logo=go" alt="Go Version">
  <img src="https://img.shields.io/badge/Git-Required-F05032.svg?style=flat&logo=git" alt="Git Required">
  <img src="https://img.shields.io/badge/Platform-Linux-FCC624.svg?style=flat&logo=linux" alt="Platform">
  <img src="https://img.shields.io/badge/License-MIT-green.svg?style=flat" alt="License">
  <img src="https://img.shields.io/badge/Version-1.0.0-blue.svg?style=flat" alt="Version">
  <img src="https://img.shields.io/badge/Status-Stable-brightgreen.svg?style=flat" alt="Status">
</p>

This project is designed to check the duration of each commit in a Git repository. It helps you track how long each commit took to be completed, providing valuable insights into your project's development process.

## 📋 Requirements

To run this project, you need the following:

- Go 1.19 or higher  
- Git installed on your system  
- Supported platforms: Linux

### 📦 Install Required Packages

Before running the project, you need to install the required Go packages.

Run the following command to install the dependencies:

```bash
go mod tidy
```

## 🚀 Installation and Setup

Follow these steps to set up the project:

**Clone the project:**
```bash
git clone https://github.com/sajad-dev/commit-duration-checker.git
```

**Navigate to the project directory:**
```bash
cd commit-duration-checker
```

**Install dependencies:**
```bash
go mod tidy
```

**Run the project:**
```bash
go run main.go
```

## 📝 Available Commands

Here are the available commands for this project:

**Run Git commands to manage your repository:**
```bash
duration git
```

**Clear logs and activity:**
```bash
duration clear 
```

**Display the help message with available commands:**
```bash
duration help
```

**Run without displaying the panel:**
```bash
duration --no-panel # Or -np
```

## 🧑‍💻 Author

**Mohammad Sajad Poorajam** 👨‍💻🚀


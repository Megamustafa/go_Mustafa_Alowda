#!/bin/bash

# Function to initialize a simple project
init_simple_project() {
  mkdir -p "$1"
  cd "$1" || exit
  go mod init "$1"
  echo 'package main

import "fmt"

func main() {
  fmt.Println("Hello, World!")
}' > main.go
  echo "Simple project initialized in $1"
}

# Function to initialize an API project
init_api_project() {
  mkdir -p "$1"
  cd "$1" || exit
  go mod init "$1"
  mkdir -p routes models repositories services configs controllers
  echo 'package main

import "fmt"

func main() {
  fmt.Println("API Server started")
}' > main.go
  echo "API project initialized in $1"
}

# Main script
read -rp "Enter the project name: " project_name

# Validate project name (no spaces allowed)
if [[ "$project_name" == *" "* ]]; then
  echo "Project name cannot contain spaces."
  exit 1
fi

read -rp "Enter the project type (simple/api): " project_type

if [[ "$project_type" == "simple" ]]; then
  init_simple_project "$project_name"
elif [[ "$project_type" == "api" ]]; then
  init_api_project "$project_name"
else
  echo "Invalid project type. Please enter 'simple' or 'api'."
  exit 1
fi

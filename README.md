# README

## Introduction

This project aims to create a backend server for aggregating data from RSS feeds. The RSS protocol facilitates the distribution of podcasts and blogs, making it easier to access and consume content from various sources.

## Installation

1. Ensure you have Go installed on your system.
   - This line states the prerequisite for the project, which is having Go installed on your system.
2. Clone this repository to your local machine.
   - Clone the repository to your local machine using Git.
3. Navigate to the project directory.
   - Open a terminal or command prompt and change the directory to the project directory.
4. Install the required dependencies using `go mod tidy`.
   - Use the `go mod tidy` command to install the required dependencies for the project.

## Usage

1. Set up environment variables by creating a `.env` file in the project directory.
   - Create a file named `.env` in the project directory to set up environment variables.
2. Define the `PORT` variable in the `.env` file with the desired port number.
   - Set the `PORT` variable in the `.env` file to specify the server port.
3. Run the application using the command `go run main.go`.
   - Execute the application by running the `main.go` file using the `go run` command.

## Example

```bash
$ go run main.go

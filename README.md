# SMTP Scanner

A combo Formatter is a tool for resolving and validating SMTP credentials for a list of email/password combinations. It can check the validity of SMTP servers for popular email providers and perform DNS resolving for other domains.

## Features

- Scan and validate SMTP credentials for email/password combinations

- Check SMTP server details for popular email providers

- Perform DNS resolving for other domains & ports

- Multithreaded execution for faster scanning

- Save valid SMTP credentials to an output file

## Requirements

- Go 1.16 or higher

## Usage

1. Clone the repository:

```shell

git clone https://github.com/Ill-u/COMBOFormat.git

cd COMBOFormat
go build COMBOFormat.go

./COMBOFormat -input input.txt -output output.txt -threads 10



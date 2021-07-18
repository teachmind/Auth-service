# Auth-service

[![Actions Status](https://github.com/teachmind/Auth-service/workflows/build/badge.svg)](https://github.com/teachmind/Auth-service/actions)
[![codecov](https://codecov.io/gh/teachmind/Auth-service/branch/master/graph/badge.svg?token=HivKkjhfjl)](https://codecov.io/gh/teachmind/Auth-service)
[![Go Report Card](https://goreportcard.com/badge/github.com/teachmind/Auth-service)](https://goreportcard.com/report/github.com/teachmind/Auth-service)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/934b654ea9eb4f72b98138b21b5aea94)](https://www.codacy.com/gh/teachmind/Auth-service/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=teachmind/Auth-service&amp;utm_campaign=Badge_Grade)
[![](https://godoc.org/github.com/teachmind/Auth-service?status.svg)](https://godoc.org/github.com/teachmind/Auth-service)

## Features 
-   Server Preparation for Running the project on localhost
-   Database Migration
-   Registration
-   Login
-   Token Validation

## Feature Details
### Registration
-   User Signup
-   Encode and Decoding the HTTP credentials
-   Validating all Credentials 
-   Phone Number Validation using RegExp
### Login
-   User Login
-   Provide JWT token
-   User login validation
    
## Project Structure
    .
    |-- cmd                 # Contains the commands for the project
    |-- images              # Contains all image file
    |-- internal            # Configuration files and Constants
    |-- migration           # Contains migration files
    |-- .env.example        # example/structure of .env file
    |-- Dockerfile          # Used to build docker image.
    |-- go.mode             # Define's the module's import path used for root directory
    |-- go.sum              # Contains the expected cryptographic checksums of the content of specific module versions
    |-- Makefile            # Makefile to run commands after docker up
    |-- readme.md           # Explains project installation and other informations

## Tools and Technology
-   Golang
-   PostgreSQL

## Installation
-   **Step-1:** Copy/rename `.env.example` file as `.env`. Change the `APP_PORT`, `DB_PORT`, `DB_NAME`,`DB_HOST`, `DB_USER`, `DB_PASSWORD` value as per your DB and Project setup.
-   **Step-2:** Run migration command `go run main.go migrate` for Database migration
-   **Step-3:** To start server run `go run main.go server`
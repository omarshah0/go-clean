# Go Clean

## Overview

This Go project leverages the Gin web framework to create a scalable and efficient web application. It includes features such as logging, tracing, and authentication middleware to enhance security and traceability. The authentication middleware is designed with three distinct user types: Admin, Customer, and Driver.

## Features
- **Postgres Database:** Utilizes the Postgres Database.

- **Gin Router:** Utilizes the Gin web framework for handling HTTP requests and routing.

- **Logging:** Implements logging functionality to track and record important events in the application.

- **Tracing:** Incorporates tracing to monitor and analyze the flow of requests within the application, aiding in performance optimization.

- **Authentication Middleware:**
  - *Admin:* Admin users have privileged access to administrative functionalities.
  - *Customer:* Customer users are regular users with specific access rights.
  - *Driver:* Driver users have access to features related to transportation services.

## Pre-Requisites
**Make sure postgress database is up and running.**

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/omarshah0/go-clean
   ```

2. Navigate to the project directory:

   ```bash
   cd go-clean
   ```

3. Install dependencies:

   ```bash
   go mod tidy
   ```

## Configuration

Ensure that you configure the application settings, such as database connection details, logging parameters, and authentication configurations. You can find the configuration file at `config/config.yaml` or `config/config.toml`, depending on your preference.

## Usage

1. Build the application:

   ```bash
   make build
   ```

2. Run the application:

   ```bash
   make run
   ```

3. Access the application at [http://localhost:8000](http://localhost:8000).

## Authentication

The application uses the following authentication middleware:

- **Admin:**

- **Customer:**

- **Driver:**

You can add more as per your requirements.

Ensure to authenticate using the appropriate credentials based on the user type you wish to access.

## Contributing

We welcome contributions to enhance the functionality and features of this project. If you find any issues or have suggestions, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).

---

Feel free to customize this README file according to your project's specifics and add any additional information or instructions as needed.
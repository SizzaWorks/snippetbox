# Snippetbox

![Project Status: Active](https://img.shields.io/badge/Status-Active-green)
[![Go Version](https://img.shields.io/badge/Go-2.23+-00ADD8.svg)](https://golang.org/doc/go1.18)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

A modern, secure web application for creating and sharing text snippets. Built with Go.

## 📋 Table of Contents

- [Features](#features)
- [Requirements](#requirements)
- [Quick Start](#quick-start)
- [Development Setup](#development-setup)
  - [Database Setup](#database-setup)
  - [Running Locally](#running-locally)
  - [Running with Docker](#running-with-docker)
- [Project Structure](#project-structure)
- [Troubleshooting](#troubleshooting)
- [Contributing](#contributing)
- [License](#license)

## ✨ Features

- Create, view, and share text snippets
- User authentication and secure sessions
- Modern, responsive UI
- RESTful API

## 📋 Requirements

- Go 2.23+
- PostgreSQL
- Docker (optional, for containerized deployment)

## 🚀 Quick Start

```bash
# Clone the repository
git clone https://github.com/SizzaWorks/snippetbox.git --depth=1
cd snippetbox

# Set up the database
sudo -u postgres psql -f ./internal/database/setup.sql

# Build and run
go build ./cmd/web/
./web

# Access the application
# The server will start at http://localhost:4000
```

## 🔧 Development Setup

### Database Setup

The application uses PostgreSQL. Make sure you have it installed:

```bash
# Install PostgreSQL if you don't have it already
sudo apt install postgresql postgresql-contrib
```

Create the database and user (default credentials: username `web`, password `pass`):

```bash
# Create database, user, and tables
sudo -u postgres psql -f ./internal/database/setup.sql

# Verify setup
psql -h localhost -U web -d snippetbox -c "\dp snippets"
psql -h localhost -U web -d snippetbox -c "\dt"
```

### Running Locally

The project uses Go modules for dependency management. No need to manually install dependencies.

```bash
# Build the application
go build ./cmd/web/

# Run the server
./web

# By default, the server runs on http://localhost:4000
```

### Running with Docker

The application can also be run using Docker:

```bash
# Start the application with Docker Compose
sudo docker compose up

# Or run in detached mode
sudo docker compose up -d
```

## 📁 Project Structure

```
snippetbox/
├── cmd/                  # Application-specific code
│   └── web/              # Web application entry point
├── internal/             # Non-application specific code
│   ├── database/         # Database models and migrations
│   ├── models/           # Data models
│   ├── validator/        # Input validation
│   └── ...
├── ui/                   # User interface assets
│   ├── html/             # HTML templates
│   ├── static/           # Static files (CSS, JS, images)
│   └── ...
├── docker-compose.yml    # Docker Compose configuration
├── Dockerfile            # Docker configuration
└── README.md             # This file
```

## 🔄 Troubleshooting

### Reset Database

If you need to start over with a clean database:

```bash
# Drop the database and user
sudo -u postgres dropdb snippetbox
sudo -u postgres dropuser web

# If using Docker, remove volumes
sudo docker compose down -v

# Then set up the database again
sudo -u postgres psql -f ./internal/database/setup.sql
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

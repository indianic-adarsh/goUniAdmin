# Go Uni Admin - API

## Prerequisites

| Module    | Version |
| --------- | ------- |
| Go   | 1.24.0 |

## Getting Started

### Local Development Setup


```bash
# Clone the repository
git clone http://git.indianic.com/goUniAdmin.git

# Create a new branch
git checkout -b feature/your-feature-name

# Set up environment variables
cp .env.sample .env
# Edit .env file with your configuration
vi .env

# Install dependencies
go mod tidy

# Start the server
go run ./cmd/api  
```

### Generate Swagger JSON
```bash
go run generate-swagger.go
```

### Docker Deployment

The project includes Docker support:

- `Dockerfile` - Production build
- `Dockerfile-Base-Image` - Base image configuration
- Kubernetes configurations available for dev/qa/uat environments

### CI/CD Pipeline

Automated deployment is configured using GitLab CI/CD pipeline. See `.gitlab-ci.yml` for pipeline configuration.

## Folder structure

```
my-go-project/
├── cmd/
│   └── api/                    # Application entry point
│       └── main.go            # Main application file
│
├── internal/                  # Private application code
│   ├── locales/              # Internationalization files
│   │   ├── en/              # English language files
│   │   │   └── messages.go  # English translations
│   │   ├── de/              # German language files
│   │   │   └── messages.go  # German translations
│   │   └── es/              # Spanish language files
│   │       └── messages.go  # Spanish translations
│   │
│   ├── modules/              # Feature modules
│   │   ├── mastermanagement/ # Master Management module
│   │   │   ├── schema.go    # Data models/structs
│   │   │   ├── service.go   # Business logic
│   │   │   ├── handler.go   # HTTP handlers
│   │   │   ├── routes.go    # Route definitions
│   │   │   └── validator.go # Validation logic
│   │   ├── admin/           # Admin module
│   │   │   ├── schema.go
│   │   │   ├── service.go
│   │   │   ├── handler.go
│   │   │   ├── routes.go
│   │   │   └── validator.go
│   │   ├── emailtemplate/   # Email Template module
│   │   │   ├── schema.go
│   │   │   ├── service.go
│   │   │   ├── handler.go
│   │   │   ├── routes.go
│   │   │   └── validator.go
│   │   ├── settings/        # Settings module
│   │   │   ├── schema.go
│   │   │   ├── service.go
│   │   │   ├── handler.go
│   │   │   ├── routes.go
│   │   │   └── validator.go
│   │   ├── roles/           # Roles module
│   │   │   ├── schema.go
│   │   │   ├── service.go
│   │   │   ├── handler.go
│   │   │   ├── routes.go
│   │   │   └── validator.go
│   │   └── staticpagemanagement/ # Static Page Management module
│   │       ├── schema.go
│   │       ├── service.go
│   │       ├── handler.go
│   │       ├── routes.go
│   │       └── validator.go
│   │
│   ├── services/            # Common services
│   │   ├── middleware/     # Middleware services
│   │   │   └── auth.go     # Example middleware (e.g., authentication)
│   │   ├── validators/     # Common validators
│   │   │   └── common.go   # Shared validation logic
│   │   ├── seed.go         # Seed data logic
│   │   ├── common.go       # Common utility functions
│   │   └── email.go        # Email-related services
│   │
│   └── config/             # Configuration loading
│       └── config.go       # Config structs and loading logic
│
├── migrations/             # Database migrations
│   └── 001_init.go         # Example migration file
│
├── public/                 # Static files
│   └── assets/             # Static assets (e.g., CSS, JS, images)
│
├── scripts/                # Utility scripts
│   └── seed.sh             # Example script for seeding data
│
├── go.mod                  # Go module definition
├── go.sum                  # Dependency checksums
├── README.md               # Project documentation
├── .gitignore              # Git ignore file
├── .gitlab-ci.yml          # GitLab CI/CD configuration
├── .env.sample             # Sample environment variables
├── .env.qa                 # QA environment variables
├── .env.uat                # UAT environment variables
├── .env                    # Local development environment variables
├── sonar-project.properties # SonarQube project properties
└── Dockerfile              # Docker configuration (optional)
```


---

For additional support or questions, please contact the development team.
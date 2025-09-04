# HealthSecure - HIPAA-Compliant Medical Data Access Control System

A production-ready HIPAA-compliant medical data access control system with role-based permissions, comprehensive audit logging, emergency access capabilities, and OAuth2.0 + JWT authentication.

## 🏥 Features

- **Role-Based Access Control**: Doctor, Nurse, and Admin roles with differentiated permissions
- **HIPAA Compliance**: Full audit logging and data protection measures
- **Emergency Access**: Secure break-glass access with justification requirements
- **OAuth2.0 + JWT**: Modern authentication with refresh token rotation
- **Real-time Monitoring**: Comprehensive audit trails and security event logging
- **Production Ready**: Docker deployment with monitoring and observability

## 🚀 Quick Start

### Prerequisites

- Go 1.21+
- Node.js 18+
- MySQL 8.0+
- Redis 6.0+
- Docker & Docker Compose

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd healthsecure
```

2. Copy environment configuration:
```bash
cp configs/.env.example configs/.env
# Edit configs/.env with your configuration
```

3. Start the services:
```bash
docker-compose up -d
```

4. Run database migrations:
```bash
cd backend
go run cmd/migrate/main.go
```

5. Start the development servers:
```bash
# Backend
cd backend
go run cmd/server/main.go

# Frontend (in separate terminal)
cd frontend
npm install
npm start
```

## 📋 API Documentation

### Authentication Endpoints
- `POST /api/auth/login` - User authentication
- `POST /api/auth/refresh` - Token refresh
- `POST /api/auth/logout` - Token invalidation
- `GET /api/auth/me` - Current user info

### Patient Endpoints
- `GET /api/patients` - List patients (role-based filtering)
- `GET /api/patients/:id` - Get patient details
- `POST /api/patients` - Create patient (admin/doctor only)
- `PUT /api/patients/:id` - Update patient

### Medical Records
- `GET /api/patients/:id/records` - Get medical records
- `POST /api/patients/:id/records` - Create medical record
- `PUT /api/records/:id` - Update medical record

### Emergency Access
- `POST /api/emergency/request` - Request emergency access
- `POST /api/emergency/activate/:token` - Activate emergency token
- `GET /api/emergency/active` - List active emergency sessions

### Audit Logs
- `GET /api/audit` - View audit logs (admin only)
- `GET /api/audit/user/:id` - User-specific audit trail

## 🔐 Security Features

### Authentication & Authorization
- JWT tokens with 15-minute expiration
- Refresh token rotation
- Role-based access control (RBAC)
- OAuth2.0 integration support

### Data Protection
- bcrypt password hashing (cost 12+)
- TLS encryption in transit
- Database encryption at rest
- Sensitive data filtering by role

### Audit & Compliance
- All data access logged
- HIPAA-compliant audit trails
- Emergency access monitoring
- Security event alerting

## 🏗️ Architecture

```
healthsecure/
├── backend/           # Go API server
├── frontend/          # React web application
├── database/          # SQL schemas and migrations
├── docker/           # Container configuration
├── tests/            # Test suites
├── configs/          # Environment configuration
└── docs/            # Documentation
```

### Technology Stack

**Backend:**
- Go 1.21+ with Gin framework
- GORM for database operations
- JWT for authentication
- Redis for session management
- MySQL for data persistence

**Frontend:**
- React 18 with hooks
- Material-UI for components
- Axios for API communication
- React Router for navigation

**Infrastructure:**
- Docker for containerization
- MySQL 8.0 for database
- Redis 6.0 for caching
- Nginx for reverse proxy

## 🧪 Testing

Run all tests:
```bash
# Backend tests
cd backend
go test ./...

# Frontend tests
cd frontend
npm test

# Security tests
npm run test:security

# E2E tests
npm run test:e2e
```

## 📊 Monitoring

The system includes comprehensive monitoring:
- Application metrics with Prometheus
- Structured logging with audit trails
- Health check endpoints
- Security event alerting

## 🔧 Configuration

Key environment variables:

```bash
# Database
DB_HOST=localhost
DB_PORT=3306
DB_NAME=healthsecure
DB_USER=healthsecure_user
DB_PASSWORD=secure_password
DB_TLS_MODE=required

# JWT
JWT_SECRET=your-secure-jwt-secret
JWT_EXPIRES=15m
REFRESH_TOKEN_EXPIRES=7d

# Security
BCRYPT_COST=12
RATE_LIMIT_REQUESTS=100
RATE_LIMIT_WINDOW=1h

# Emergency Access
EMERGENCY_ACCESS_DURATION=1h
EMERGENCY_NOTIFICATION_EMAIL=security@yourorg.com
```

## 🏥 HIPAA Compliance

This system implements HIPAA technical safeguards:
- ✅ Access control with unique user identification
- ✅ Audit controls with comprehensive logging
- ✅ Integrity controls for data protection
- ✅ Person or entity authentication
- ✅ Transmission security with encryption

## 📝 User Roles

### Doctor
- Full patient data access including SSN
- Create and modify medical records
- View all diagnostic information
- Emergency access capabilities

### Nurse
- Limited patient data access (no SSN)
- View basic medical information
- Update care notes
- Emergency access for critical situations

### Admin
- User management and system configuration
- Audit log access and monitoring
- No direct patient data access
- Emergency access oversight

## 🚨 Emergency Access

The system provides secure break-glass access:
1. User requests emergency access with justification
2. System generates time-limited access token (1 hour default)
3. All emergency access is logged and audited
4. Administrators are notified immediately
5. Access automatically expires

## 📚 Documentation

- [API Documentation](docs/API.md)
- [Security Implementation](docs/SECURITY.md)
- [Deployment Guide](docs/DEPLOYMENT.md)
- [HIPAA Compliance](docs/HIPAA.md)

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🆘 Support

For support and questions:
- Create an issue in this repository
- Contact the development team
- Review the documentation in the `docs/` directory

---

**⚠️ Security Notice**: This system handles sensitive medical data. Ensure all security configurations are properly set before deploying to production.
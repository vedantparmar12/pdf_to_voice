name: "HealthSecure - HIPAA-Compliant Medical Data Access Control System"
description: |

## Purpose
Build a production-ready HIPAA-compliant medical data access control system with role-based permissions, comprehensive audit logging, emergency access capabilities, and OAuth2.0 + JWT authentication. Demonstrates full-stack development with Go backend, React frontend, and secure multi-role data access patterns.

## Core Principles
1. **Security First**: HIPAA compliance and data protection at every layer
2. **Role-Based Access**: Strict role segregation (Doctor, Nurse, Admin)
3. **Audit Everything**: Comprehensive logging for compliance and monitoring
4. **Emergency Preparedness**: Secure break-glass access with justification
5. **Production Ready**: Docker deployment with monitoring and observability

---

## Goal
Create a complete medical data management system where healthcare professionals can securely access patient information according to their role, with full audit trails and emergency access capabilities.

## Why
- **Business Value**: Streamlines healthcare data access while maintaining HIPAA compliance
- **Security**: Implements industry-standard security patterns for sensitive medical data
- **Problems Solved**: Balances data accessibility with privacy requirements and regulatory compliance

## What
A full-stack web application featuring:
- Role-based authentication (Doctor/Nurse/Admin access levels)
- Secure patient data management with differential access controls
- Real-time audit logging for all data operations
- Emergency access workflow with justification requirements
- OAuth2.0 + JWT authentication with SSO capabilities

### Success Criteria
- [ ] Role-based access controls function correctly for all user types
- [ ] All data access is logged and auditable
- [ ] Emergency access workflow provides secure break-glass capabilities
- [ ] HIPAA compliance requirements are met
- [ ] System handles production load with proper monitoring

## All Needed Context

### Documentation & References
```yaml
# MUST READ - Include these in your context window
- url: https://gin-gonic.com/docs/
  why: Go Gin framework patterns for REST API development
  
- url: https://gorm.io/docs/
  why: Go ORM patterns for database operations and relationships
  
- url: https://reactjs.org/docs/getting-started.html
  why: React patterns for component architecture and state management
  
- url: https://jwt.io/introduction/
  why: JWT token structure and security best practices
  
- url: https://www.hhs.gov/hipaa/for-professionals/security/laws-regulations/index.html
  why: HIPAA security requirements and compliance guidelines
  
- file: backend/internal/models/
  why: Database model patterns and relationships
  
- file: frontend/src/contexts/AuthContext.jsx
  why: Authentication state management patterns
```

### Current Project Structure
```bash
healthsecure/
├── backend/                 # Go backend
├── frontend/               # React frontend  
├── database/              # SQL schemas and migrations
├── docker/                # Container configuration
├── docs/                  # Documentation
├── scripts/              # Build and deployment
└── README.md
```

### Target Architecture
```bash
healthsecure/
├── backend/
│   ├── cmd/
│   │   ├── server/main.go           # Application entry point
│   │   └── migrate/main.go          # Database migration runner
│   ├── internal/
│   │   ├── auth/                    # JWT + OAuth2 handlers
│   │   │   ├── jwt.go              # Token generation/validation
│   │   │   ├── oauth.go            # OAuth2 flow implementation
│   │   │   └── middleware.go        # Auth middleware
│   │   ├── models/                  # Database models
│   │   │   ├── user.go             # User entity with roles
│   │   │   ├── patient.go          # Patient data model
│   │   │   ├── medical_record.go   # Medical records with relations
│   │   │   └── audit_log.go        # Audit trail model
│   │   ├── handlers/                # HTTP route handlers
│   │   │   ├── auth.go             # Authentication endpoints
│   │   │   ├── patients.go         # Patient CRUD with role checks
│   │   │   ├── medical_records.go  # Medical data access
│   │   │   ├── audit.go            # Audit log viewing
│   │   │   └── emergency.go        # Emergency access workflow
│   │   ├── services/                # Business logic layer
│   │   │   ├── user_service.go     # User management logic
│   │   │   ├── patient_service.go  # Patient data business rules
│   │   │   ├── audit_service.go    # Audit logging service
│   │   │   └── emergency_service.go # Emergency access logic
│   │   ├── middleware/              # HTTP middleware
│   │   │   ├── auth.go             # Authentication check
│   │   │   ├── rbac.go             # Role-based access control
│   │   │   ├── audit.go            # Automatic audit logging
│   │   │   └── rate_limit.go       # API rate limiting
│   │   └── database/
│   │       ├── connection.go       # DB connection management
│   │       └── migrations.go       # Schema migration logic
│   ├── configs/
│   │   └── config.go               # Environment configuration
│   ├── go.mod
│   └── go.sum
├── frontend/
│   ├── src/
│   │   ├── components/
│   │   │   ├── Layout/             # App layout components
│   │   │   ├── Auth/               # Authentication components
│   │   │   ├── Patient/            # Patient management UI
│   │   │   ├── Medical/            # Medical records UI
│   │   │   ├── Audit/              # Audit log viewing
│   │   │   └── Emergency/          # Emergency access UI
│   │   ├── pages/                  # Route components
│   │   ├── services/               # API communication
│   │   ├── contexts/               # React context providers
│   │   ├── hooks/                  # Custom React hooks
│   │   └── utils/                  # Helper functions
│   ├── package.json
│   └── package-lock.json
├── database/
│   ├── schema.sql                  # Database schema definition
│   ├── seed_data.sql               # Sample data for testing
│   └── migrations/                 # Database migration files
├── docker/
│   ├── docker-compose.yml          # Multi-container setup
│   ├── Dockerfile.backend          # Go backend container
│   └── Dockerfile.frontend         # React frontend container
├── tests/
│   ├── backend/                    # Go unit and integration tests
│   └── frontend/                   # React component and e2e tests
├── configs/
│   ├── .env.example                # Environment variables template
│   └── nginx.conf                  # Reverse proxy configuration
├── scripts/
│   ├── setup.sh                    # Initial project setup
│   ├── test.sh                     # Run all tests
│   └── deploy.sh                   # Deployment automation
└── docs/
    ├── API.md                      # API documentation
    ├── SECURITY.md                 # Security implementation details
    └── DEPLOYMENT.md               # Deployment guide
```

### Known Gotchas & Critical Requirements
```go
// CRITICAL: HIPAA requires audit logs for ALL data access
// CRITICAL: Emergency access must be time-limited and justified
// CRITICAL: Passwords must be hashed with bcrypt (min cost 12)
// CRITICAL: JWT tokens should expire within 15 minutes, use refresh tokens
// CRITICAL: All database queries must use parameterized statements
// CRITICAL: Never log sensitive patient data (SSN, diagnosis details)
// CRITICAL: MySQL connections must use TLS in production
// CRITICAL: Redis sessions must have expiration times
// CRITICAL: All HTTP responses must include security headers
// CRITICAL: Role checks must happen at both middleware AND service layers
```

## Implementation Blueprint

### Data Models and Core Structures

```go
// models/user.go - Authentication and role management
type User struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Email     string    `json:"email" gorm:"unique;not null;index"`
    Password  string    `json:"-" gorm:"not null"`
    Role      UserRole  `json:"role" gorm:"not null;type:enum('doctor','nurse','admin')"`
    Name      string    `json:"name" gorm:"not null"`
    Active    bool      `json:"active" gorm:"default:true"`
    LastLogin time.Time `json:"last_login"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

// models/patient.go - Patient data with privacy controls
type Patient struct {
    ID            uint                `json:"id" gorm:"primaryKey"`
    FirstName     string             `json:"first_name" gorm:"not null"`
    LastName      string             `json:"last_name" gorm:"not null"`
    DateOfBirth   time.Time          `json:"date_of_birth"`
    SSN           string             `json:"ssn,omitempty" gorm:"unique;column:ssn"` // Only for doctors
    Phone         string             `json:"phone"`
    Address       string             `json:"address"`
    EmergencyContact string          `json:"emergency_contact"`
    MedicalRecords []MedicalRecord   `json:"medical_records,omitempty"`
    CreatedAt     time.Time          `json:"created_at"`
    UpdatedAt     time.Time          `json:"updated_at"`
}

// models/medical_record.go - Medical data with access controls
type MedicalRecord struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    PatientID   uint      `json:"patient_id" gorm:"not null;index"`
    DoctorID    uint      `json:"doctor_id" gorm:"not null;index"`
    Diagnosis   string    `json:"diagnosis" gorm:"type:text"`
    Treatment   string    `json:"treatment" gorm:"type:text"`
    Notes       string    `json:"notes" gorm:"type:text"`
    Medications string    `json:"medications" gorm:"type:text"`
    Severity    string    `json:"severity" gorm:"type:enum('low','medium','high','critical')"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
    
    Patient Patient `json:"patient,omitempty" gorm:"foreignKey:PatientID"`
    Doctor  User    `json:"doctor,omitempty" gorm:"foreignKey:DoctorID"`
}

// models/audit_log.go - Comprehensive audit trail
type AuditLog struct {
    ID           uint      `json:"id" gorm:"primaryKey"`
    UserID       uint      `json:"user_id" gorm:"not null;index"`
    PatientID    *uint     `json:"patient_id,omitempty" gorm:"index"`
    RecordID     *uint     `json:"record_id,omitempty" gorm:"index"`
    Action       string    `json:"action" gorm:"not null;index"`
    Resource     string    `json:"resource" gorm:"not null"`
    IPAddress    string    `json:"ip_address" gorm:"not null"`
    UserAgent    string    `json:"user_agent" gorm:"type:text"`
    EmergencyUse bool      `json:"emergency_use" gorm:"default:false;index"`
    Reason       string    `json:"reason,omitempty" gorm:"type:text"`
    Success      bool      `json:"success" gorm:"default:true;index"`
    ErrorMessage string    `json:"error_message,omitempty"`
    Timestamp    time.Time `json:"timestamp" gorm:"autoCreateTime;index"`
    
    User    User     `json:"user,omitempty" gorm:"foreignKey:UserID"`
    Patient *Patient `json:"patient,omitempty" gorm:"foreignKey:PatientID"`
}

// models/emergency_access.go - Break-glass access tracking
type EmergencyAccess struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    UserID      uint      `json:"user_id" gorm:"not null;index"`
    PatientID   uint      `json:"patient_id" gorm:"not null;index"`
    Reason      string    `json:"reason" gorm:"not null;type:text"`
    AccessToken string    `json:"access_token" gorm:"unique;not null"`
    ExpiresAt   time.Time `json:"expires_at" gorm:"not null;index"`
    UsedAt      *time.Time `json:"used_at,omitempty"`
    RevokedAt   *time.Time `json:"revoked_at,omitempty"`
    CreatedAt   time.Time `json:"created_at"`
    
    User    User    `json:"user,omitempty" gorm:"foreignKey:UserID"`
    Patient Patient `json:"patient,omitempty" gorm:"foreignKey:PatientID"`
}
```

### Task Breakdown and Implementation Order

```yaml
Phase 1: Foundation Setup (Week 1)
Task 1.1: Environment and Dependencies
  CREATE configs/config.go:
    - PATTERN: Use viper for configuration management
    - Load environment variables with validation
    - Database connection strings with TLS options
    - JWT secret management
    - Redis connection configuration

  CREATE database/connection.go:
    - PATTERN: GORM connection with proper error handling
    - Connection pooling configuration
    - MySQL TLS enforcement for production
    - Migration auto-runner

  CREATE docker-compose.yml:
    - Multi-service setup (MySQL, Redis, Backend, Frontend)
    - Health checks for all services
    - Volume persistence for database
    - Network isolation

Task 1.2: Database Models and Migrations
  CREATE internal/models/ (all model files):
    - PATTERN: GORM struct tags with proper indexing
    - JSON serialization controls for sensitive data
    - Foreign key relationships with cascade rules
    - Validation tags for data integrity

  CREATE database/migrations/:
    - Progressive schema changes
    - Index creation for performance
    - Enum constraints for roles and statuses
    - Sample data seeding

Phase 2: Authentication and Authorization (Week 2)
Task 2.1: JWT and OAuth2 Implementation
  CREATE internal/auth/jwt.go:
    - PATTERN: JWT with refresh token rotation
    - Short-lived access tokens (15 minutes)
    - Proper claims structure with role information
    - Token blacklisting capabilities

  CREATE internal/auth/oauth.go:
    - OAuth2 authorization code flow
    - State parameter validation
    - PKCE implementation for security
    - Provider integration (Google, Azure AD)

Task 2.2: Middleware Implementation
  CREATE internal/middleware/auth.go:
    - JWT token validation
    - User context injection
    - Error handling and proper HTTP status codes

  CREATE internal/middleware/rbac.go:
    - Role-based access control checks
    - Resource-specific permissions
    - Emergency access bypass logic

  CREATE internal/middleware/audit.go:
    - Automatic audit log generation
    - Request/response logging (excluding sensitive data)
    - Performance metrics collection

Phase 3: Core Business Logic (Week 3)
Task 3.1: Service Layer Implementation
  CREATE internal/services/user_service.go:
    - User registration with role assignment
    - Password hashing with bcrypt (cost 12+)
    - User activation/deactivation
    - Role management

  CREATE internal/services/patient_service.go:
    - PATTERN: Role-based data filtering
    - Doctor: Full access including SSN
    - Nurse: Limited access, no SSN or sensitive diagnosis
    - Admin: No patient data access
    - Search and pagination with security controls

  CREATE internal/services/medical_record_service.go:
    - Medical record CRUD with doctor verification
    - Access control based on treating physician
    - Data sensitivity filtering by role

Task 3.2: Emergency Access Workflow
  CREATE internal/services/emergency_service.go:
    - Break-glass access token generation
    - Time-limited access (configurable, default 1 hour)
    - Justification requirement and validation
    - Automatic audit trail enhancement
    - Admin notification system

Phase 4: API Handlers and Routes (Week 4)
Task 4.1: HTTP Handlers Implementation
  CREATE internal/handlers/auth.go:
    - POST /api/auth/login - User authentication
    - POST /api/auth/refresh - Token refresh
    - POST /api/auth/logout - Token invalidation
    - GET /api/auth/me - Current user info

  CREATE internal/handlers/patients.go:
    - GET /api/patients - List with role-based filtering
    - GET /api/patients/:id - Detail view with access control
    - POST /api/patients - Create (admin/doctor only)
    - PUT /api/patients/:id - Update with audit trail

  CREATE internal/handlers/emergency.go:
    - POST /api/emergency/request - Request emergency access
    - POST /api/emergency/activate/:token - Activate emergency token
    - GET /api/emergency/active - List active emergency sessions

Phase 5: Frontend Implementation (Week 5-6)
Task 5.1: Authentication and Layout
  CREATE src/contexts/AuthContext.jsx:
    - PATTERN: Context with JWT token management
    - Automatic token refresh logic
    - Role-based component rendering
    - Logout functionality with cleanup

  CREATE src/components/Layout/:
    - Responsive layout with role-based navigation
    - Header with user info and logout
    - Sidebar with permissions-based menu items

Task 5.2: Patient Management Interface
  CREATE src/components/Patient/PatientList.jsx:
    - PATTERN: Role-based data display
    - Search and filtering capabilities
    - Pagination with virtual scrolling
    - Quick actions based on user role

  CREATE src/components/Patient/PatientDetail.jsx:
    - Tabbed interface for patient information
    - Medical records timeline
    - Conditional rendering based on access level
    - Edit capabilities for authorized users

Task 5.3: Emergency Access Interface
  CREATE src/components/Emergency/EmergencyAccess.jsx:
    - Emergency access request form
    - Justification text area with validation
    - Confirmation dialog with warnings
    - Real-time status updates

Phase 6: Testing and Security (Week 7)
Task 6.1: Backend Testing
  CREATE tests/backend/:
    - Unit tests for all service methods
    - Integration tests for API endpoints
    - Security tests for authentication bypass attempts
    - Performance tests for database queries

Task 6.2: Frontend Testing
  CREATE tests/frontend/:
    - Component unit tests with role mocking
    - Integration tests for user workflows
    - E2E tests for critical paths
    - Accessibility testing compliance

Phase 7: Production Deployment (Week 8)
Task 7.1: Container Configuration
  CREATE docker/Dockerfile.backend:
    - Multi-stage build for Go application
    - Security hardening (non-root user)
    - Health check endpoints

Task 7.2: Monitoring and Observability
  CREATE monitoring configuration:
    - Application metrics with Prometheus
    - Log aggregation with structured logging
    - Health checks and uptime monitoring
    - Security event alerting
```

### Critical Security Implementation Patterns

```go
// Authentication middleware with proper error handling
func JWTMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := extractToken(c.GetHeader("Authorization"))
        if token == "" {
            c.JSON(401, gin.H{"error": "Authorization token required"})
            c.Abort()
            return
        }
        
        claims, err := ValidateToken(token)
        if err != nil {
            // CRITICAL: Don't leak token validation details
            c.JSON(401, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }
        
        // PATTERN: Inject user context for downstream handlers
        c.Set("user_id", claims.UserID)
        c.Set("user_role", claims.Role)
        c.Next()
    }
}

// Role-based access control with resource-level permissions
func RequireRole(roles ...string) gin.HandlerFunc {
    return func(c *gin.Context) {
        userRole := c.GetString("user_role")
        
        for _, role := range roles {
            if userRole == role {
                c.Next()
                return
            }
        }
        
        // CRITICAL: Log unauthorized access attempts
        auditService.LogUnauthorizedAccess(
            c.GetUint("user_id"),
            c.Request.URL.Path,
            c.ClientIP(),
        )
        
        c.JSON(403, gin.H{"error": "Insufficient permissions"})
        c.Abort()
    }
}

// Patient data filtering based on user role
func (s *PatientService) GetPatient(userID uint, userRole string, patientID uint) (*Patient, error) {
    patient := &Patient{}
    query := s.db.Where("id = ?", patientID)
    
    // CRITICAL: Apply role-based field selection
    switch userRole {
    case "doctor":
        // Full access including sensitive data
        query = query.Preload("MedicalRecords")
    case "nurse":
        // Limited access, exclude SSN and sensitive diagnosis
        query = query.Select("id, first_name, last_name, date_of_birth, phone, address")
    case "admin":
        // CRITICAL: Admins should not access patient data
        return nil, errors.New("administrators cannot access patient data")
    }
    
    if err := query.First(patient).Error; err != nil {
        return nil, err
    }
    
    // CRITICAL: Audit all patient data access
    s.auditService.LogPatientAccess(userID, patientID, "VIEW", false)
    
    return patient, nil
}

// Emergency access with time limits and enhanced logging
func (s *EmergencyService) RequestEmergencyAccess(userID, patientID uint, reason string) (*EmergencyAccess, error) {
    // CRITICAL: Generate secure random token
    token := generateSecureToken(32)
    
    emergency := &EmergencyAccess{
        UserID:      userID,
        PatientID:   patientID,
        Reason:      reason,
        AccessToken: token,
        ExpiresAt:   time.Now().Add(1 * time.Hour), // CONFIGURABLE
    }
    
    if err := s.db.Create(emergency).Error; err != nil {
        return nil, err
    }
    
    // CRITICAL: Enhanced audit logging for emergency access
    s.auditService.LogEmergencyRequest(userID, patientID, reason)
    
    // CRITICAL: Notify administrators immediately
    s.notificationService.NotifyEmergencyAccess(emergency)
    
    return emergency, nil
}
```

### Integration Points and Environment Configuration

```yaml
ENVIRONMENT VARIABLES:
  # Database Configuration
  DB_HOST: "localhost"
  DB_PORT: "3306"
  DB_NAME: "healthsecure"
  DB_USER: "healthsecure_user"
  DB_PASSWORD: "secure_password_here"
  DB_TLS_MODE: "required"  # CRITICAL for production
  
  # Redis Configuration
  REDIS_HOST: "localhost"
  REDIS_PORT: "6379"
  REDIS_PASSWORD: "redis_password"
  REDIS_DB: "0"
  
  # JWT Configuration
  JWT_SECRET: "super-secure-jwt-secret-minimum-32-chars"
  JWT_EXPIRES: "15m"  # Short-lived tokens
  REFRESH_TOKEN_EXPIRES: "7d"
  
  # OAuth2 Configuration
  OAUTH_CLIENT_ID: "your-oauth-client-id"
  OAUTH_CLIENT_SECRET: "your-oauth-client-secret"
  OAUTH_REDIRECT_URL: "https://yourdomain.com/auth/callback"
  
  # Security Configuration
  BCRYPT_COST: "12"  # Minimum recommended
  RATE_LIMIT_REQUESTS: "100"
  RATE_LIMIT_WINDOW: "1h"
  
  # Emergency Access Configuration
  EMERGENCY_ACCESS_DURATION: "1h"
  EMERGENCY_NOTIFICATION_EMAIL: "security@yourorg.com"
  
  # Application Configuration
  SERVER_PORT: "8080"
  ENVIRONMENT: "production"
  LOG_LEVEL: "info"
  ENABLE_CORS: "true"
  CORS_ORIGINS: "https://yourdomain.com"

DEPENDENCY REQUIREMENTS:
  Backend (Go):
    - github.com/gin-gonic/gin v1.9.1
    - gorm.io/gorm v1.25.4
    - gorm.io/driver/mysql v1.5.1
    - github.com/golang-jwt/jwt/v5 v5.0.0
    - github.com/go-redis/redis/v8 v8.11.5
    - golang.org/x/crypto v0.12.0
    - github.com/google/uuid v1.3.0
    - github.com/spf13/viper v1.16.0
    
  Frontend (React):
    - react v18.2.0
    - react-router-dom v6.15.0
    - @mui/material v5.14.5
    - axios v1.5.0
    - react-hook-form v7.45.4
    - @hookform/resolvers v3.3.1
    - yup v1.3.2
```

## Validation and Testing Strategy

### Level 1: Code Quality and Security
```bash
# Backend validation
go vet ./...                    # Go static analysis
golangci-lint run              # Comprehensive linting
gosec ./...                    # Security vulnerability scanning

# Frontend validation  
npm run lint                   # ESLint and Prettier
npm audit                      # Dependency vulnerability check
npm run test:security          # Security-focused tests
```

### Level 2: Unit and Integration Testing
```go
// Example test for role-based access
func TestPatientService_GetPatient_RoleBasedAccess(t *testing.T) {
    tests := []struct {
        name     string
        userRole string
        wantSSN  bool
        wantErr  bool
    }{
        {"Doctor gets full access", "doctor", true, false},
        {"Nurse gets limited access", "nurse", false, false},
        {"Admin gets no access", "admin", false, true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            patient, err := service.GetPatient(1, tt.userRole, 1)
            
            if tt.wantErr {
                assert.Error(t, err)
                return
            }
            
            assert.NoError(t, err)
            if tt.wantSSN {
                assert.NotEmpty(t, patient.SSN)
            } else {
                assert.Empty(t, patient.SSN)
            }
        })
    }
}

// Example security test
func TestAuth_PreventTokenReuse(t *testing.T) {
    // Test that logout properly blacklists tokens
    token := generateTestToken()
    
    // First request should succeed
    resp1 := makeAuthenticatedRequest(token)
    assert.Equal(t, 200, resp1.StatusCode)
    
    // Logout
    logoutRequest(token)
    
    // Second request should fail
    resp2 := makeAuthenticatedRequest(token)
    assert.Equal(t, 401, resp2.StatusCode)
}
```

### Level 3: End-to-End Security Validation
```bash
# Security testing checklist
npm run test:e2e:security       # Automated security scenarios
npm run test:e2e:roles          # Role-based access scenarios
npm run test:e2e:emergency      # Emergency access workflow

# Manual security validation
curl -X GET /api/patients/1     # Should fail without auth
curl -X GET /api/patients/1 -H "Authorization: Bearer expired_token"  # Should fail
curl -X GET /api/admin/users -H "Authorization: Bearer nurse_token"   # Should fail

# Performance testing
ab -n 1000 -c 10 http://localhost:8080/api/patients  # Load testing
```

### Level 4: HIPAA Compliance Validation
```yaml
COMPLIANCE CHECKLIST:
  Technical Safeguards:
    - [ ] Data encryption in transit (HTTPS/TLS)
    - [ ] Data encryption at rest (database)
    - [ ] Access controls implemented and tested
    - [ ] Audit logs capture all required events
    - [ ] User authentication and authorization working
    - [ ] Emergency access procedures documented and tested
    
  Administrative Safeguards:
    - [ ] Security policies documented
    - [ ] User training materials created
    - [ ] Incident response procedures defined
    - [ ] Regular security assessments planned
    
  Physical Safeguards:
    - [ ] Server access controls documented
    - [ ] Data backup procedures implemented
    - [ ] Disaster recovery plan created
```

## Final Production Readiness Checklist
- [ ] All tests pass: `make test`
- [ ] Security scans pass: `make security-scan`
- [ ] Performance benchmarks meet requirements
- [ ] Database migrations tested on production data copy
- [ ] Monitoring and alerting configured
- [ ] SSL certificates installed and tested
- [ ] Backup and recovery procedures tested
- [ ] HIPAA compliance documentation complete
- [ ] User training materials finalized
- [ ] Incident response procedures documented

---

## Anti-Patterns to Avoid
- ❌ Never log sensitive patient data (SSN, diagnosis details)
- ❌ Don't use admin accounts for regular operations
- ❌ Avoid long-lived JWT tokens (max 15 minutes)
- ❌ Don't skip role checks in service layer (defense in depth)
- ❌ Never commit secrets or credentials to version control
- ❌ Don't allow patient data access for admin users
- ❌ Avoid using HTTP in production (HTTPS only)
- ❌ Don't skip audit logging for any data access
- ❌ Never trust client-side role validation alone

## Confidence Score: 9.5/10

Very high confidence due to:
- Clear security patterns and HIPAA requirements
- Well-established Go + React technology stack
- Comprehensive testing and validation strategy
- Detailed implementation roadmap with security focus
- Real-world applicability for healthcare systems

Minor uncertainty around specific OAuth2 provider integrations, but patterns are well-documented.
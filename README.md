# 🚀 Plantilla de API Go con Clean Architecture

Plantilla de API REST para Go, construida con principios de Clean Architecture. Incluye un sistema de autenticación JWT completo (con refresh tokens) y un módulo de ejemplo (tasks) listo para ser clonado y adaptado a tus necesidades.

## 🚀 Características

- ✅ Autenticación con JWT (Access + Refresh tokens)
- 🔐 Gestión de sesiones activas
- 🏗️ Clean Architecture (Dominio, Aplicación, Infraestructura)
- 🗄️ Base de Datos Dual (PostgreSQL o SQLite) con GORM
- 📝 Módulo de Ejemplo (CRUD de Tareas) para que veas cómo estructurar los tuyos
- ✔️ Validación de datos con go-playground/validator
- 🧩 Inyección de Dependencias (DI) simple y manual
- 🛣️ Router ligero con chi

## 📁 Estructura del Proyecto
La estructura está diseñada para separar responsabilidades y escalar

```
go-easy-list/
├── config/                  # Configuración (Variables de entorno, BBDD)
│   ├── config.go
│   └── database.go
├── internal/
│   ├── auth/                # Módulo de Autenticación (¡Listo para usar!)
│   │   ├── application/
│   │   ├── domain/
│   │   └── infrastructure/
│   ├── tasks/               # Módulo de Ejemplo (renombrar o eliminar)
│   │   ├── application/
│   │   ├── domain/
│   │   └── infrastructure/
│   └── shared/              # Código compartido (Middleware, Handlers, DI)
│       ├── context/
│       ├── http/
│       ├── infrastructure/
│       └── validation/
├── .env.example             # Plantilla de variables de entorno
├── go.mod
├── go.sum
└── main.go                 # Punto de partida

```

## 🛠️ Tecnologías

- **Go 1.23+**
- **PostgreSQL** (Recomendado) o **SQLite**
- **Chi** - Router HTTP
- **GORM** - ORM
- **JWT** - Autenticación
- **Validator** - Validación de datos

## ⚙️ Instalación

### 1. Clonar el repositorio
```bash
git clone https://github.com/DevJorgeRafael/go-task-easy-list.git
cd go-task-easy-list
```

### 2. Instalar dependencias
```bash
go mod download
```

### 3. Configurar variables de entorno

Copia `.env.example` y configura tus variables:
```bash
cp .env.example .env
```
```env
# Server
PORT=8080

# --- Base de Datos (Elegir una) ---

# Opción 1: PostgreSQL (Recomendado)
# Descomentar y ajustar la URL de conexión
DATABASE_URL="postgres://postgres:<password>@localhost:5432/<my_db>?sslmode=disable"

# Opción 2: SQLite
# Descomentar para usar un archivo local
# DB_PATH=./app.db

# JWT (Cambiar por valores seguros)
JWT_SECRET=super-secret-key
JWT_ACCESS_EXPIRATION=1h
JWT_REFRESH_EXPIRATION=7d
```


### 4. Iniciar el servidor
```bash
go run main.go
```

El servidor estará disponible en `http://localhost:8080`

## 📡 API Endpoints

### 🔐 Autenticación (`/api/auth`)

#### Rutas Públicas

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| POST | `/api/auth/register` | Registrar nuevo usuario |
| POST | `/api/auth/login` | Iniciar sesión |
| POST | `/api/auth/refresh` | Renovar access token |

#### Rutas Protegidas (requieren JWT)

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| POST | `/api/auth/logout` | Cerrar sesiones |
| GET | `/api/auth/sessions` | Listar sesiones activas |

### ✅ Tareas (`/api/tasks`)

Todas las rutas requieren autenticación (Header: `Authorization: Bearer <token>`)

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| POST | `/api/tasks` | Crear tarea |
| GET | `/api/tasks` | Listar todas las tareas del usuario |
| GET | `/api/tasks/{id}` | Obtener tarea por ID |
| PUT | `/api/tasks/{id}` | Actualizar tarea |
| DELETE | `/api/tasks/{id}` | Eliminar tarea |


## 🔒 Seguridad

- Contraseñas hasheadas con bcrypt
- JWT con expiración configurable
- Refresh tokens para renovación segura
- Validación de sesiones activas
- Middleware de autenticación en todas las rutas protegidas


## 👤 Autor

Jorge Rafael Rosero - Plantilla Base
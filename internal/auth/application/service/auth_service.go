package service

import (
	"errors"
	"go-task-easy-list/internal/auth/domain/model"
	"go-task-easy-list/internal/auth/domain/repository"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Errores del dominio
var (
	ErrInvalidEmail = errors.New("email inválido")
	ErrEmailExists = errors.New("el email ya está registrado")
	ErrInvalidPassword = errors.New("la contraseña debe tener al menos 8 caracteres")
	ErrUserNotFound = errors.New("usuario no encontrado")
	ErrInvalidCredentials = errors.New("credenciales inválidas")
)

type AuthService struct {
	userRepo   repository.UserRepository
	jwtSecret  string
}

func NewAuthService(userRepo repository.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

// Register- Registrar un usuario
func (s *AuthService) Register(email, password, name string) (*model.User, error) {
	// 1. Validar email (formato, no duplicado)
	if !isValidEmail(email) {
		return nil, ErrInvalidEmail
	}

	existing, _ := s.userRepo.FindByEmail(email)
	if existing != nil {
		return nil, ErrEmailExists
	}

	// 2. Validar y Hash password con bcrypt
	if len(password) < 8 {
		return nil, ErrInvalidPassword
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 3. Crear user
	user := &model.User{
		ID: uuid.New().String(),
		Email: email,
		Password: string(hashedPassword),
		Name: name,
		IsActive: true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// 4. Guardar en DB con userRepo.Create()
	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	// 5. Retornar usuario (SIN password)
	user.Password = ""
	return user, nil
}

// Login - Iniciar sesión
func (s *AuthService) Login(email, password string) (accessToken, refreshToken string, err error) {
	// 1. Buscar user por email con userRepo.FindByEmail()
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", "", err
	}

	// 2. Verificar password con bcrypt.CompareHashAndPassword()
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", "", ErrInvalidCredentials
	}

	// 3. Generar JWT accessToken y refreshToken
	accessToken, err = s.generateAccessToken(user.ID, user.Email)
	if err != nil {
		return "", "", err
	}

	refreshToken = uuid.New().String()

	// 4. Guardar refreshToken en tabla sessions
	return accessToken, refreshToken, nil
}

func (s *AuthService) generateAccessToken(userID, email string) (string ,error) {
	claims := jwt.MapClaims{
		"userId": userID,
		"email": email,
		"exp": time.Now().Add(1 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
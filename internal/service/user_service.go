package service

import (
	"fmt"
	"log"
	"strings"
	"test/app/internal/models"

	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		DB: db,
	}
}

func (u *UserService) GetUsers(users *[]models.User) error {

	if err := u.DB.Order("created_at desc").Find(&users).Error; err != nil {
		log.Println("Ошибка при получении пользователей: ", err)
		return fmt.Errorf("не удалось получить пользователей: %w", err)
	}
	log.Printf("Успешно получено %d пользователей", len(*users))
	return nil
}

func (u *UserService) CreateUser(user *models.User) error {
	if user.Name == "" {
		log.Println("Ошибка: имя пользователя не может быть пустым")
		return fmt.Errorf("Имя пользователя обязательно")
	}

	if user.Age <= 0 {
		log.Println("Ошибка: возраст должен быть больше 0")
		return fmt.Errorf("Указан некорректный возраст")
	}

	if user.Password == "" {
		log.Println("Поле пароля должно быть заполнено")
		return fmt.Errorf("Пароль обязателен к заполнению")
	}

	if user.Email == "" || !strings.Contains(user.Email, "@") {
		log.Println("Введите верный адрес электронной потчы")
		return fmt.Errorf("Некорректный адрес почты")
	}

	if len(user.Profession) <= 5 {
		log.Println("Ошибка: название профессии слишком короткое")
		return fmt.Errorf("название профессии должно содержать больше 10 символов")
	}

	newUser := u.DB.Create(user)
	if newUser.Error != nil {
		log.Printf("Ошибка при создании пользователя: %v", newUser.Error)
		return fmt.Errorf("не удалось создать пользователя: %w", newUser.Error)
	}
	log.Printf("Пользователь %s успешно создан с ID: %d", user.Name, user.ID)
	return nil
}

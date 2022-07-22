package ports

import (
	"github.com/decadevs/lunch-api/internal/core/models"
	"time"
)

type UserService interface {
	FindFoodBenefactorByFullName(fullname string) (*models.FoodBeneficiary, error)
	FindFoodBenefactorByEmail(email string) (*models.FoodBeneficiary, error)
	FindFoodBenefactorByLocation(location string) (*models.FoodBeneficiary, error)
	CreateFoodBenefactor(user *models.FoodBeneficiary) (*models.FoodBeneficiary, error)
	FindKitchenStaffByFullName(fullname string) (*models.KitchenStaff, error)
	FindKitchenStaffByEmail(email string) (*models.KitchenStaff, error)
	FindKitchenStaffByLocation(location string) (*models.KitchenStaff, error)
	CreateKitchenStaff(user *models.KitchenStaff) (*models.KitchenStaff, error)
	FindAdminByEmail(email string) (*models.Admin, error)
	TokenInBlacklist(token *string) bool
	AddTokenToBlacklist(email string, token string) error
	FindUserById(id string) (*models.FoodBeneficiary, error)
	UserResetPassword(id, newPassword string) (*models.FoodBeneficiary, error)
	KitchenStaffResetPassword(id, newPassword string) (*models.KitchenStaff, error)
	AdminResetPassword(id, newPassword string) (*models.Admin, error)
	CreateFoodTimetable(food models.Food) error
	CreateAdmin(user *models.Admin) (*models.Admin, error)
	FindBrunchByDate(year int, month time.Month, day int) (*models.Food, error)
	FindDinnerByDate(year int, month time.Month, day int) (*models.Food, error)
}

// MailerService interface to implement mailing service
type MailerService interface {
	SendMail(subject, body, to, Private, Domain string) error
	GenerateNonAuthToken(UserEmail string, secret string) (*string, error)
	DecodeToken(token, secret string) (string, error)
}

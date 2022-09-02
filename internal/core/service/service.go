package service

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/decadevs/lunch-api/internal/ports"
	"mime/multipart"
)

type userService struct {
	userRepository ports.UserRepository
}

type mailerService struct {
	mailerRepository ports.MailerRepository
}

type AWSService struct {
	AWSRepository ports.AWSRepository
}

func NewUserService(userRepository ports.UserRepository) ports.UserService {
	return &userService{
		userRepository: userRepository,
	}
}
func NewMailerService(mailerRepository ports.MailerRepository) ports.MailerService {
	return &mailerService{
		mailerRepository: mailerRepository,
	}
}

func NewAWSServices(AWSRepository ports.AWSRepository) ports.AWSService {
	return &AWSService{
		AWSRepository: AWSRepository,
	}
}

func (a *AWSService) UploadFileToS3(h *session.Session, file multipart.File, fileName string, size int64) (string, error) {
	return a.AWSRepository.UploadFileToS3(h, file, fileName, size)
}

func (m *mailerService) SendMail(subject, body, to, Private, Domain string) error {
	return m.mailerRepository.SendMail(subject, body, to, Private, Domain)
}
func (m *mailerService) GenerateNonAuthToken(UserEmail string, secret string) (*string, error) {
	return m.mailerRepository.GenerateNonAuthToken(UserEmail, secret)
}
func (m *mailerService) DecodeToken(token, secret string) (string, error) {
	return m.mailerRepository.DecodeToken(token, secret)
}

func (u *userService) FindUserById(id string) (*models.FoodBeneficiary, error) {
	return u.userRepository.FindUserById(id)
}
func (u *userService) UserResetPassword(id, newPassword string) (*models.FoodBeneficiary, error) {
	return u.userRepository.UserResetPassword(id, newPassword)
}

func (u *userService) AdminResetPassword(id, newPassword string) (*models.Admin, error) {
	return u.userRepository.AdminResetPassword(id, newPassword)
}

func (u *userService) KitchenStaffResetPassword(id, newPassword string) (*models.KitchenStaff, error) {
	return u.userRepository.KitchenStaffResetPassword(id, newPassword)
}

func (u *userService) FindActiveUsersByMonth() (interface{}, error) {
	return u.userRepository.FindActiveUsersByMonth()
}

func (u *userService) FindNumbersOfScannedUsers(date string) (int64, error) {
	return u.userRepository.FindNumbersOfScannedUsers(date)
}

func (u *userService) FindFoodBenefactorByFullName(fullname string) (*models.FoodBeneficiary, error) {
	return u.userRepository.FindFoodBenefactorByFullName(fullname)
}

func (u *userService) FindFoodBenefactorByEmail(email string) (*models.FoodBeneficiary, error) {
	return u.userRepository.FindFoodBenefactorByEmail(email)
}

func (u *userService) FindFoodBenefactorByLocation(location string) (*models.FoodBeneficiary, error) {
	return u.userRepository.FindFoodBenefactorByLocation(location)
}

func (u *userService) CreateFoodBenefactor(user *models.FoodBeneficiary) (*models.FoodBeneficiary, error) {
	return u.userRepository.CreateFoodBenefactor(user)
}

func (u *userService) FindKitchenStaffByFullName(fullname string) (*models.KitchenStaff, error) {
	return u.userRepository.FindKitchenStaffByFullName(fullname)
}

func (u *userService) FindKitchenStaffByEmail(email string) (*models.KitchenStaff, error) {
	return u.userRepository.FindKitchenStaffByEmail(email)
}

func (u *userService) FindKitchenStaffByLocation(location string) (*models.KitchenStaff, error) {
	return u.userRepository.FindKitchenStaffByLocation(location)
}

func (u *userService) CreateKitchenStaff(user *models.KitchenStaff) (*models.KitchenStaff, error) {
	return u.userRepository.CreateKitchenStaff(user)
}

func (u *userService) FindAdminByEmail(email string) (*models.Admin, error) {
	return u.userRepository.FindAdminByEmail(email)
}

func (u *userService) TokenInBlacklist(token *string) bool {
	return u.userRepository.TokenInBlacklist(token)
}

func (u *userService) AddTokenToBlacklist(email string, token string) error {
	return u.userRepository.AddTokenToBlacklist(email, token)
}

func (u *userService) CreateFoodTimetable(food models.Food) error {
	return u.userRepository.CreateFoodTimetable(food)
}

func (u *userService) CreateAdmin(user *models.Admin) (*models.Admin, error) {
	return u.userRepository.CreateAdmin(user)
}

func (u *userService) FindBrunchByDate(year int, month int, day int) ([]models.Food, error) {
	return u.userRepository.FindBrunchByDate(year, month, day)
}

func (u *userService) FindDinnerByDate(year int, month int, day int) ([]models.Food, error) {
	return u.userRepository.FindDinnerByDate(year, month, day)
}
func (u *userService) FoodBeneficiaryEmailVerification(id string) (*models.FoodBeneficiary, error) {
	return u.userRepository.FoodBeneficiaryEmailVerification(id)
}

func (u *userService) KitchenStaffEmailVerification(id string) (*models.KitchenStaff, error) {
	return u.userRepository.KitchenStaffEmailVerification(id)
}

func (u *userService) AdminEmailVerification(id string) (*models.Admin, error) {
	return u.userRepository.AdminEmailVerification(id)
}
func (u *userService) FindFoodBenefactorMealRecord(email, date string) (*models.MealRecords, error) {
	return u.userRepository.FindFoodBenefactorMealRecord(email, date)
}
func (u *userService) CreateFoodBenefactorBrunchMealRecord(user *models.FoodBeneficiary) error {
	return u.userRepository.CreateFoodBenefactorBrunchMealRecord(user)
}
func (u *userService) CreateFoodBenefactorDinnerMealRecord(user *models.FoodBeneficiary) error {
	return u.userRepository.CreateFoodBenefactorDinnerMealRecord(user)
}
func (u *userService) UpdateFoodBenefactorBrunchMealRecord(email string) error {
	return u.userRepository.UpdateFoodBenefactorBrunchMealRecord(email)
}
func (u *userService) UpdateFoodBenefactorDinnerMealRecord(email string) error {
	return u.userRepository.UpdateFoodBenefactorDinnerMealRecord(email)
}

func (u *userService) FindAllFoodBeneficiary(pagination *models.Pagination) ([]models.UserDetails, error) {
	return u.userRepository.FindAllFoodBeneficiary(pagination)
}

func (u *userService) GetFoodByID(id string) (*models.Food, error) {
	return u.userRepository.GetFoodByID(id)
}

func (u *userService) UpdateStatus(food []models.Food, status string) error {
	return u.userRepository.UpdateStatus(food, status)
}

func (u *userService) SearchFoodBeneficiary(text string, pagination *models.Pagination) ([]models.UserDetails, error) {
	return u.userRepository.SearchFoodBeneficiary(text, pagination)
}

func (u *userService) GetTotalUsers() (int, error) {
	return u.userRepository.GetTotalUsers()
}

func (u *userService) UpdateMeal(id string, food models.Food) error {
	return u.userRepository.UpdateMeal(id, food)
}
func (u *userService) DeleteMeal(id string) error {
	return u.userRepository.DeleteMeal(id)
}

func (u *userService) FindAllFoodByDate(year int, month int, day int) ([]models.Food, error) {
	return u.userRepository.FindAllFoodByDate(year, month, day)
}

func (u *userService) CreateNotification(notification models.Notification) error {
	return u.userRepository.CreateNotification(notification)
}

func (u *userService) FindNotificationByDate(year int, month int, day int) ([]models.Notification, error) {
	return u.userRepository.FindNotificationByDate(year, month, day)
}

func (u *userService) GetFoodBenefactorById(id string) (*models.FoodBeneficiary, error) {
	return u.userRepository.GetFoodBenefactorById(id)
}

func (u *userService) AdminBlockFoodBeneficiary(userID string) error {
	return u.userRepository.AdminBlockFoodBeneficiary(userID)
}

func (u *userService) AdminRemoveFoodBeneficiary(userID string) error {
	return u.userRepository.AdminRemoveFoodBeneficiary(userID)
}

func (u *userService) FindFoodBenefactorQRCodeMealRecord(mealId, userId string) (*models.QRCodeMealRecords, error) {
	return u.userRepository.FindFoodBenefactorQRCodeMealRecord(mealId, userId)
}
func (u *userService) CreateFoodBenefactorQRMealRecord(mealRecord *models.QRCodeMealRecords) error {
	return u.userRepository.CreateFoodBenefactorQRMealRecord(mealRecord)
}
func (u *userService) NumberOfBlockedBeneficiary() (int64, error) {
	return u.userRepository.NumberOfBlockedBeneficiary()
}

func (u *userService) GetBlockedBeneficiary() ([]models.FoodBeneficiary, error) {
	return u.userRepository.GetBlockedBeneficiary()
}

func (u *userService) UpdateFoodStatusById(id string, status string) error {
	return u.userRepository.UpdateFoodStatusById(id, status)
}

func (u *userService) GetAllFoodBeneficiaries() ([]models.FoodBeneficiary, error) {
	return u.userRepository.GetAllFoodBeneficiaries()
}

func (u *userService) FindFoodByDate(year int, month int, day int) ([]models.Food, error) {
	return u.userRepository.FindFoodByDate(year, month, day)
}

func (u *userService) FindListOfScannedUsers(date string, pagination *models.Pagination) ([]models.UserDetails, error) {
	return u.userRepository.FindListOfScannedUsers(date, pagination)
}

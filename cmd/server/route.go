package server

import (
	"github.com/decadevs/lunch-api/internal/adapters/api"
	"github.com/decadevs/lunch-api/internal/core/middleware"
	"github.com/decadevs/lunch-api/internal/ports"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

func SetupRouter(handler *api.HTTPHandler, userService ports.UserService) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: 12 * time.Hour,
	}))

	r := router.Group("/api/v1")
	{

		r.GET("/ping", handler.PingHandler)
		r.POST("/user/beneficiarysignup", handler.FoodBeneficiarySignUp)
		r.PATCH("/user/beneficiaryverifyemail", handler.BeneficiaryVerifyEmail)
		r.POST("/user/kitchenstaffsignup", handler.KitchenStaffSignUp)
		r.PATCH("/user/kitchenstaffverifyemail", handler.KitchenStaffVerifyEmail)
		r.POST("/user/adminsignup", handler.AdminSignUp)
		r.PATCH("/user/adminverifyemail", handler.AdminVerifyEmail)
		r.POST("/user/kitchenstafflogin", handler.LoginKitchenStaffHandler)
		r.POST("/user/benefactorlogin", handler.LoginFoodBenefactorHandler)
		r.POST("/user/adminlogin", handler.LoginAdminHandler)
		r.POST("/user/beneficiaryforgotpassword", handler.FoodBeneficiaryForgotPassword)
		r.POST("/user/getallbeneficiaries", handler.HandleAdminGetAllBeneficiaries)
		r.PATCH("/user/beneficiaryresetpassword/:token", handler.FoodBeneficiaryResetPassword)
		r.POST("/user/kitchenstaffforgotpassword", handler.KitchenStaffForgotPassword)
		r.PATCH("/user/kitchenstaffresetpassword/:token", handler.KitchenStaffResetPassword)
		r.POST("/user/adminforgotpassword", handler.AdminForgotPassword)
		r.PATCH("/user/adminresetpassword/:token", handler.AdminResetPassword)
		r.GET("/user/notifications", handler.GetNotification)
		r.GET("/user/brunch", handler.GetBrunchHandle)
		r.GET("/user/dinner", handler.GetDinnerHandle)
		r.GET("/user/allfood", handler.GetAllFoodHandler)
	}

	// authorizeKitchenStaff authorizes all authorized kitchen staff handler
	authorizeKitchenStaff := r.Group("/staff")
	authorizeKitchenStaff.Use(middleware.AuthorizeKitchenStaff(userService.FindKitchenStaffByEmail, userService.TokenInBlacklist))
	{
		authorizeKitchenStaff.POST("/kitchenstafflogout", handler.KitchenStaffLogout)
		authorizeKitchenStaff.PUT("/changebrunchstatus", handler.UpdateBrunchFoodStatus)
		authorizeKitchenStaff.PUT("/changedinnerstatus", handler.UpdateDinnerFoodStatus)
		authorizeKitchenStaff.GET("/getusers", handler.GetFoodBeneficiaries)
		authorizeKitchenStaff.GET("/searchbeneficiary/:text", handler.SearchFoodBeneficiaries)
		authorizeKitchenStaff.POST("/createtimetable", handler.CreateFoodTimetableHandle)
		authorizeKitchenStaff.GET("/gettotalusers", handler.GetTotalNumberOfUsers)
		authorizeKitchenStaff.GET("/getbrunchtimetable", handler.GetBrunchTimetable)
		authorizeKitchenStaff.GET("/getdinnertimetable", handler.GetDinnerTimetable)
		authorizeKitchenStaff.GET("/generateqrcode", handler.GetMeal)
	}

	// authorizeBenefactor authorizes all authorized benefactor handler
	authorizeBenefactor := r.Group("/benefactor")
	authorizeBenefactor.Use(middleware.AuthorizeFoodBenefactor(userService.FindFoodBenefactorByEmail, userService.TokenInBlacklist))
	{
		authorizeBenefactor.POST("/beneficiarylogout", handler.FoodBeneficiaryLogout)
		authorizeBenefactor.GET("/qrbrunch", handler.BeneficiaryQRBrunch)
		authorizeBenefactor.GET("/qrdinner", handler.BeneficiaryQRDinner)
		authorizeBenefactor.POST("/qrmealrecords", handler.MealRecord)
	}

	// authorizeAdmin authorizes all authorized admin handler
	authorizeAdmin := r.Group("/admin")
	authorizeAdmin.Use(middleware.AuthorizeAdmin(userService.FindAdminByEmail, userService.TokenInBlacklist))
	{
		authorizeAdmin.POST("/adminlogout", handler.AdminLogout)
		authorizeAdmin.POST("/createtimetable", handler.CreateFoodTimetableHandle)
		authorizeAdmin.DELETE("/deletemeal/:id", handler.DeleteMeal)
		authorizeAdmin.PUT("/updatemeal/:id", handler.UpdateMeal)
		authorizeAdmin.GET("/getfoodbeneficiaryprofile/:id", handler.GetFoodBeneficiaryProfile)
		authorizeAdmin.PUT("/blockfoodbeneficiary/:id", handler.BlockFoodBeneficiary)
		authorizeAdmin.DELETE("/removefoodbeneficiary/:id", handler.RemoveFoodBeneficiary)
		authorizeAdmin.GET("/numberblocked", handler.GetNumberOfBlockedUsers)
		authorizeAdmin.GET("/blockedusers", handler.GetBlockedUsers)
		authorizeAdmin.GET("/getTimetable", handler.HandleAdminGetMealTimetable)
		authorizeAdmin.GET("/getAllBeneficiaries", handler.HandleAdminGetAllBeneficiaries)
		authorizeAdmin.GET("/searchBeneficiaries/:text", handler.HandleAdminSearchFoodBeneficiaries)
		authorizeAdmin.GET("/getTotalNumberOfUsers", handler.HandleAdminGetTotalNumberOfUsers)
		authorizeAdmin.GET("/getTotalNumberOfScannedUsers", handler.HandleGetNumberOfScannedUsers)
		authorizeAdmin.GET("/getGraphData", handler.HandleGetGraphData)
		authorizeAdmin.GET("/getScannedUsersByDate", handler.HandleGetListOfScannedUsersByDate)

	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

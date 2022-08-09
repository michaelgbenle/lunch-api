package api

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func (u *HTTPHandler) UpdateMeal(c *gin.Context) {
	user, err := u.GetAdminFromContext(c)
	if err != nil {
		helpers.JSON(c, "not authorized", http.StatusUnauthorized, nil, []string{"not authorized"})
		return
	}

	id := c.Param("id")

	var food models.Food

	if err = c.BindJSON(&food); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error binding in updateFood handler"})
		return
	}

	err = u.UserService.UpdateMeal(id, food)
	if err != nil {
		helpers.JSON(c, "internal server error", http.StatusInternalServerError, nil, []string{"internal server error"})
		return
	}

	year, month, day := time.Now().Date()

	notification := models.Notification{
		Message: user.FullName + " updated timetable",
		Year:    year,
		Month:   int(month),
		Day:     day,
	}

	err = u.UserService.CreateNotification(notification)
	if err != nil {
		helpers.JSON(c, "Error in getting Notification", http.StatusInternalServerError, nil, []string{"Error in getting Notification"})
		return
	}

	helpers.JSON(c, "Successful updated", http.StatusOK, "Successfully updated", nil)

}

package api

import (
	"fmt"
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (u *HTTPHandler) SearchFoodBeneficiaries(c *gin.Context) {
	_, err := u.GetKitchenStaffFromContext(c)
	if err != nil {
		helpers.JSON(c, "An internal server error", 500, nil, []string{"internal server error"})
		return
	}
	query := c.Param("text")

	beneficiaries, err := u.UserService.SearchFoodBeneficiary(query)
	if err != nil {
		helpers.JSON(c, "An internal server error", 500, nil, []string{"internal server error"})
		return
	}
	fmt.Println(beneficiaries)
	if len(beneficiaries) == 0 {
		helpers.JSON(c, "Record Not Found", 404, nil, []string{"Record Not Found"})
		return
	}
	helpers.JSON(c, "information gotten", http.StatusOK, beneficiaries, nil)
}

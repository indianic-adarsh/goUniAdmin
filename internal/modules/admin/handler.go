package admin

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// AdminHandler handles HTTP requests for admin CRUD
type AdminHandler struct {
	service *AdminService
}

// NewAdminHandler creates a new handler with the service
func NewAdminHandler(service *AdminService) *AdminHandler {
	return &AdminHandler{service: service}
}

// AdminCreateRequest defines the request body for creating an admin
type AdminCreateRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"userName,omitempty"`
	Mobile    string `json:"mobile,omitempty"`
	EmailID   string `json:"emailId"`
	Password  string `json:"password"`
}

// CreateAdmin godoc
// @Summary Create a new admin
// @Description Creates a new admin with hashed password
// @Tags admins
// @Accept json
// @Produce json
// @Param admin body AdminCreateRequest true "Admin data"
// @Success 201 {object} Admin
// @Failure 400 {object} map[string]string "error: Invalid request body or validation error"
// @Failure 500 {object} map[string]string "error: Failed to hash password"
// @Router /admins [post]
func (h *AdminHandler) CreateAdmin(c *gin.Context) {
	var admin Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid request body"})
		return
	}

	// Hash the password using bcrypt before creating the admin
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to hash password"})
		return
	}
	// Replace the plain password with the hashed version
	admin.Password = string(hashedPassword)

	// Create the admin with the hashed password
	created, err := h.service.Create(admin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	// Optionally, remove the password from the response for security
	created.Password = ""
	c.JSON(http.StatusCreated, created)
}

// GetAdmin godoc
// @Summary Get an admin by ID
// @Description Retrieves an admin by their UUID
// @Tags admins
// @Produce json
// @Param id path string true "Admin ID"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} Admin
// @Failure 400 {object} map[string]string "error: Invalid ID"
// @Failure 401 {object} map[string]string "error: Unauthorized"
// @Failure 404 {object} map[string]string "error: Admin not found"
// @Router /admins/{id} [get]
func (h *AdminHandler) GetAdmin(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid ID"})
		return
	}

	admin, err := h.service.Read(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
		return
	}

	// Remove password from response
	admin.Password = ""
	c.JSON(http.StatusOK, gin.H{"success": true, "data": admin, "message": "Get details successfully."})
}

// UpdateAdmin godoc
// @Summary Update an admin
// @Description Updates an existing admin by ID
// @Tags admins
// @Accept json
// @Produce json
// @Param id path string true "Admin ID"
// @Param admin body Admin true "Updated admin data"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} Admin
// @Failure 400 {object} map[string]string "error: Invalid ID or request body"
// @Failure 401 {object} map[string]string "error: Unauthorized"
// @Failure 404 {object} map[string]string "error: Admin not found"
// @Router /admins/{id} [put]
func (h *AdminHandler) UpdateAdmin(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid ID"})
		return
	}

	var admin Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid request body"})
		return
	}

	updated, err := h.service.Update(id, admin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	// Remove password from response
	updated.Password = ""
	c.JSON(http.StatusOK, gin.H{"success": false, "message": "Updated successfully", "data": updated})
}

// DeleteAdmin godoc
// @Summary Delete an admin
// @Description Soft deletes an admin by ID
// @Tags admins
// @Produce json
// @Param id path string true "Admin ID"
// @Param Authorization header string true "Bearer token"
// @Success 204
// @Failure 400 {object} map[string]string "error: Invalid ID"
// @Failure 401 {object} map[string]string "error: Unauthorized"
// @Failure 404 {object} map[string]string "error: Admin not found"
// @Router /admins/{id} [delete]
func (h *AdminHandler) DeleteAdmin(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid ID"})
		return
	}

	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"success": true, "message": "Admin deleted successfully."})
}

// ListAdmins godoc
// @Summary List all admins
// @Description Retrieves a list of all non-deleted admins
// @Tags admins
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Success 200 {array} Admin
// @Failure 500 {object} map[string]string "error: Internal server error"
// @Router /admins [get]
func (h *AdminHandler) ListAdmins(c *gin.Context) {
	// Parse query parameters for pagination
	page, err := strconv.Atoi(c.DefaultQuery("page", "1")) // Default to 1 if "page" is not provided
	if err != nil || page < 1 {
		page = 1 // Ensure page is at least 1 if invalid or non-positive
	}

	page_size, err := strconv.Atoi(c.DefaultQuery("page_size", "10")) // Default to 10 if "page_size" is not provided
	if err != nil || page_size < 1 {
		page_size = 10 // Ensure page_size is at least 1 if invalid or non-positive
	}

	// Calculate the offset and limit for pagination
	offset := (page - 1) * page_size
	limit := page_size

	admins, count, err := h.service.List(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	// Remove passwords from all admins in the list
	for i := range admins {
		admins[i].Password = ""
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Get details successfully.", "data": gin.H{"list": admins, "page": page, "page_size": page_size, "total_count": count}})
}

// AdminLogin godoc
// @Summary Admin login
// @Description Authenticates an admin and returns a JWT token
// @Tags admins
// @Accept json
// @Produce json
// @Param login body object{emailId=string,password=string} true "Login credentials (only emailId and password)"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string "error: Invalid request body"
// @Failure 401 {object} map[string]string "error: Invalid email or password"
// @Failure 500 {object} map[string]string "error: Failed to generate token"
// @Router /admins/login [post]
func (h *AdminHandler) AdminLogin(c *gin.Context) {
	var admin Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid request body"})
		return
	}

	dbAdmin, err := h.service.ReadByEmail(admin.EmailID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": "Invalid email or password"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbAdmin.Password), []byte(admin.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "Invalid email or password"})
		return
	}

	// Remove password from response
	dbAdmin.Password = ""
	// Generate JWT token
	adminId := dbAdmin.ID
	fmt.Println("Generating JWT token ", adminId)

	token, err := h.service.GenerateToken(adminId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to generate token"})
		return
	}
	fmt.Println("line 160")

	// Return the admin data along with the generated token
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Logged in successfully",
		"data": gin.H{
			"admin": dbAdmin,
			"token": token,
		},
	})

}

// GetProfile godoc
// @Summary Get admin profile
// @Description Retrieves the profile of the authenticated admin
// @Tags admins
// @Produce json
// @Security BearerAuth
// @Success 200 {object} Admin
// @Failure 401 {object} map[string]string "error: Unauthorized"
// @Failure 404 {object} map[string]string "error: Admin not found"
// @Router /admins/profile [get]
func (h *AdminHandler) GetProfile(c *gin.Context) {
	adminIDStr, exists := c.Get("adminID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "Admin ID not found in context"})
		return
	}

	adminID, err := uuid.Parse(adminIDStr.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Invalid admin ID in token"})
		return
	}

	admin, err := h.service.Read(adminID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
		return
	}

	admin.Password = ""
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Get details Successfully.", "data": admin})
}

package controllers

import (
	"github.com/gin-gonic/gin"
)

//AuthController interface is a contract what this controller can do
type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

// type authController struct {
// 	authService service.AuthService
// 	jwtService  service.JWTService
// }

//NewAuthController creates a new instance of AuthController
// func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
// 	return &authController{
// 		authService: authService,
// 		jwtService:  jwtService,
// 	}
// }

// func (c *authController) Login(ctx *gin.Context) {
// 	var loginDTO dto.LoginDTO
// 	errDTO := ctx.ShouldBind(&loginDTO)
// 	if errDTO != nil {
// 		response := response.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
// 		return
// 	}
// 	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
// 	if v, ok := authResult.(entity.User); ok {
// 		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
// 		v.Token = generatedToken
// 		response := response.BuildResponse(true, "OK!", v)
// 		ctx.JSON(http.StatusOK, response)
// 		return
// 	}
// 	response := response.BuildErrorResponse("Please check again your credential", "Invalid Credential", response.EmptyObj{})
// 	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
// }

// func (c *authController) Register(ctx *gin.Context) {
// 	var registerDTO dto.RegisterDTO
// 	errDTO := ctx.ShouldBind(&registerDTO)
// 	if errDTO != nil {
// 		response := response.BuildErrorResponse("Failed to process request", errDTO.Error(), response.EmptyObj{})
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
// 		response := response.BuildErrorResponse("Failed to process request", "Duplicate email", response.EmptyObj{})
// 		ctx.JSON(http.StatusConflict, response)
// 	} else {
// 		createdUser := c.authService.CreateUser(registerDTO)
// 		token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
// 		createdUser.Token = token
// 		response := response.BuildResponse(true, "OK!", createdUser)
// 		ctx.JSON(http.StatusCreated, response)
// 	}
// }

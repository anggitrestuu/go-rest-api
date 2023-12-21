package v1

import (
	"fmt"
	"net/http"

	V1Domains "github.com/anggitrestuu/go-rest-api/internal/business/domains/v1"
	"github.com/anggitrestuu/go-rest-api/internal/constants"
	"github.com/anggitrestuu/go-rest-api/internal/datasources/caches"
	"github.com/anggitrestuu/go-rest-api/internal/http/datatransfers/requests"
	"github.com/anggitrestuu/go-rest-api/internal/http/datatransfers/responses"
	"github.com/anggitrestuu/go-rest-api/pkg/jwt"
	"github.com/anggitrestuu/go-rest-api/pkg/validators"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase        V1Domains.UserUsecase
	redisCache     caches.RedisCache
	ristrettoCache caches.RistrettoCache
}

func NewUserHandler(usecase V1Domains.UserUsecase, redisCache caches.RedisCache, ristrettoCache caches.RistrettoCache) UserHandler {
	return UserHandler{
		usecase:        usecase,
		redisCache:     redisCache,
		ristrettoCache: ristrettoCache,
	}
}

// @Summary User registration
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param user body requests.UserRequest true "Register User"
// @Success 200 {object} map[string]interface{} "registration user success"
// @Router /api/v1/auth/regis [post]
func (userH UserHandler) Regis(ctx *gin.Context) {
	var UserRegisRequest requests.UserRequest
	if err := ctx.ShouldBindJSON(&UserRegisRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := validators.ValidatePayloads(UserRegisRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	userDomain := UserRegisRequest.ToV1Domain()
	userDomainn, statusCode, err := userH.usecase.Store(ctx.Request.Context(), userDomain)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "registration user success", map[string]interface{}{
		"user": responses.FromV1Domain(userDomainn),
	})
}

// @Summary User login
// @Description Logs in a user
// @Tags auth
// @Accept json
// @Produce json
// @Param user body requests.UserLoginRequest true "Login User"
// @Success 200 {object} map[string]interface{} "login success"
// @Router /api/v1/auth/login [post]
func (userH UserHandler) Login(ctx *gin.Context) {
	var UserLoginRequest requests.UserLoginRequest
	if err := ctx.ShouldBindJSON(&UserLoginRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := validators.ValidatePayloads(UserLoginRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	userDomain, statusCode, err := userH.usecase.Login(ctx.Request.Context(), UserLoginRequest.ToV1Domain())
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "login success", responses.FromV1Domain(userDomain))
}

// @Summary Send OTP
// @Description Send an OTP to user's email
// @Tags auth
// @Accept json
// @Produce json
// @Param email body requests.UserSendOTPRequest true "Send OTP"
// @Success 200 {object} map[string]interface{} "otp code has been send"
// @Router /api/v1/auth/send-otp [post]
func (userH UserHandler) SendOTP(ctx *gin.Context) {
	var userOTP requests.UserSendOTPRequest

	if err := ctx.ShouldBindJSON(&userOTP); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := validators.ValidatePayloads(userOTP); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	otpCode, statusCode, err := userH.usecase.SendOTP(ctx.Request.Context(), userOTP.Email)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	otpKey := fmt.Sprintf("user_otp:%s", userOTP.Email)
	go userH.redisCache.Set(otpKey, otpCode)

	NewSuccessResponse(ctx, statusCode, fmt.Sprintf("otp code has been send to %s", userOTP.Email), nil)
}

// @Summary Verify OTP
// @Description Verify OTP for a user
// @Tags auth
// @Accept json
// @Produce json
// @Param otp body requests.UserVerifOTPRequest true "Verify OTP"
// @Success 200 {object} map[string]interface{} "otp verification success"
// @Router /api/v1/auth/verif-otp [post]
func (userH UserHandler) VerifOTP(ctx *gin.Context) {
	var userOTP requests.UserVerifOTPRequest

	if err := ctx.ShouldBindJSON(&userOTP); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := validators.ValidatePayloads(userOTP); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	otpKey := fmt.Sprintf("user_otp:%s", userOTP.Email)
	otpRedis, err := userH.redisCache.Get(otpKey)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	statusCode, err := userH.usecase.VerifOTP(ctx.Request.Context(), userOTP.Email, userOTP.Code, otpRedis)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	statusCode, err = userH.usecase.ActivateUser(ctx.Request.Context(), userOTP.Email)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	go userH.redisCache.Del(otpKey)
	go userH.ristrettoCache.Del("users")

	NewSuccessResponse(ctx, statusCode, "otp verification success", nil)
}

// @Summary Get User Data
// @Description Get data of authenticated user
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "user data fetched successfully"
// @Router /api/v1/users/me [get]
func (c UserHandler) GetUserData(ctx *gin.Context) {
	// get authenticated user from context
	userClaims := ctx.MustGet(constants.CtxAuthenticatedUserKey).(jwt.JwtCustomClaim)
	if val := c.ristrettoCache.Get(fmt.Sprintf("user/%s", userClaims.Email)); val != nil {
		NewSuccessResponse(ctx, http.StatusOK, "user data fetched successfully", map[string]interface{}{
			"user": val,
		})
		return
	}

	ctxx := ctx.Request.Context()
	userDom, statusCode, err := c.usecase.GetByEmail(ctxx, userClaims.Email)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	userResponse := responses.FromV1Domain(userDom)

	go c.ristrettoCache.Set(fmt.Sprintf("user/%s", userClaims.Email), userResponse)

	NewSuccessResponse(ctx, statusCode, "user data fetched successfully", map[string]interface{}{
		"user": userResponse,
	})

}

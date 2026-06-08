package util

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(userID, secret string) (string, error) {
	expirationTime := 24 * time.Hour
	claims := jwt.StandardClaims{
		Subject:   userID,
		ExpiresAt: time.Now().Add(expirationTime).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func GetAuthJWT(ginContext *gin.Context) string {
	userTokenID, exists := ginContext.Get("userID")
	if !exists {
		ginContext.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse{
			Error: "token JWT inexistente",
		})
		return ""
	}

	userID, ok := userTokenID.(string)

	if !ok {
		ginContext.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse{
			Error: "erro no parsing do token",
		})
		return ""
	}

	return userID
}

// func GetAuthJWTWithMember(ginContext *gin.Context) (*database.Member, error) {
// 	memberID := GetAuthJWT(ginContext)
//
// 	ctx := ginContext.Request.Context()
// 	dbQueries := services.GetDatabase()
//
// 	member, err := dbQueries.GetMemberByID(ctx, memberID)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return &member, nil
// }
//
// func IsMemberAdmin(ctx context.Context, memberID string) (bool, error) {
// 	dbQueries := services.GetDatabase()
// 	memberRoles, err := dbQueries.GetMemberRoles(ctx, memberID)
// 	if err != nil {
// 		return false, err
// 	}
//
// 	for _, role := range memberRoles {
// 		if role.Title == "administrator" {
// 			return true, nil
// 		}
// 	}
//
// 	return false, nil
// }
//
// func IsMemberAtLeastManager(ctx context.Context, memberID string) (bool, error) {
// 	dbQueries := services.GetDatabase()
// 	memberRoles, err := dbQueries.GetMemberRoles(ctx, memberID)
// 	if err != nil {
// 		return false, err
// 	}
//
// 	authorizedRoles := []string{
// 		"administrator",
// 		"project_manager",
// 		"coordinator",
// 	}
//
// 	for _, role := range memberRoles {
// 		if slices.Contains(authorizedRoles, role.Title) {
// 			return true, nil
// 		}
// 	}
//
// 	return false, nil
// }

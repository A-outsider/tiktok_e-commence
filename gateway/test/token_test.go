package test

import (
	"github.com/stretchr/testify/assert"
	"gomall/gateway/config"
	"gomall/gateway/utils/token"
	"testing"
)

// 生成 Access Token 的测试
func TestGenerateAccessToken(t *testing.T) {

	config.GetConf().Jwt.AccessSecret = "tiktok_mall_access"
	config.GetConf().Jwt.Issuer = "zty"
	config.GetConf().Jwt.AccessExpireTime = "60m"

	// 生成 token
	userId := "12345"
	tk, err := token.GenerateAccessToken(userId)
	assert.NoError(t, err, "Should generate access token without error")

	// 解析 token
	claims, err := token.ParseToken(tk)
	assert.NoError(t, err, "Should parse token without error")
	assert.Equal(t, userId, claims.UserId, "UserId should match")
	assert.Equal(t, uint(0), claims.TokenType, "TokenType should be access token")
}

// 解析 Refresh Token 的测试
func TestParseToken(t *testing.T) {

	config.GetConf().Jwt.RefreshSecret = "refresh"
	config.GetConf().Jwt.Issuer = "zty"
	config.GetConf().Jwt.RefreshExpireTime = "60m"

	// 生成 token
	userId := "12345"
	tk, err := token.GenerateRefreshToken(userId)
	assert.NoError(t, err, "Should generate access token without error")

	// 解析 token
	claims, err := token.ParseToken(tk)
	assert.NoError(t, err, "Should parse token without error")
	assert.Equal(t, userId, claims.UserId, "UserId should match")
	assert.Equal(t, uint(1), claims.TokenType, "TokenType should be access token")
}

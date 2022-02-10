// Package auth grpc的token校验
package auth

import (
	"context"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/gbrlsnchs/jwt/v3"
	"go.uber.org/dig"
	"google.golang.org/grpc/metadata"

	"github.com/scjtqs2/bot_adapter/config"
)

const (
	// AuthTokenErrorInvalidAppidAppsecret
	AuthTokenErrorInvalidAppidAppsecret = "AUTH_TOKEN_ERROR_INVALID_APPID_APPSECRET"
	// AuthTokenErrorInvalidToken
	AuthTokenErrorInvalidToken = "AUTH_TOKEN_ERROR_INVALID_TOKEN"
	// AuthTokenErrorEmptyAuthorizeMetadata
	AuthTokenErrorEmptyAuthorizeMetadata = "AUTH_TOKEN_ERROR_EMPTY_AUTHORIZE_METADATA"
	TTL                                  = 2 * time.Hour
)

// AuthServer auth的结构体
type AuthServer struct {
	hs   *jwt.HMACSHA
	conf *config.Config
}

// NewAuthServer 初始化 auth
func NewAuthServer(ct *dig.Container) (*AuthServer, error) {
	secret := "secret"
	if os.Getenv("ADAPTER_SECRET") != "" {
		secret = os.Getenv("ADAPTER_SECRET")
	}
	var au AuthServer
	au.hs = jwt.NewHS256([]byte(secret))
	_ = ct.Invoke(func(opt *config.Config) {
		au.conf = opt
	})
	return &au, nil
}

// GetToken 使用appid 和 appsecret 拉取token
func (au *AuthServer) GetToken(appid, appsecret string) (string, error) {
	if f, ok := au.CheckAppidSecret(appid, appsecret); ok {
		now := time.Now()
		pl := jwt.Payload{
			Issuer:         "bot_adapter",
			Subject:        f.PluginName,
			Audience:       jwt.Audience{f.AppSecret},
			ExpirationTime: jwt.NumericDate(now.Add(TTL)),
			IssuedAt:       jwt.NumericDate(now),
			JWTID:          f.AppID,
		}
		token, err := jwt.Sign(pl, au.hs)
		if err != nil {
			return "", err
		}
		return string(token), nil
	}
	return "", errors.New(AuthTokenErrorInvalidAppidAppsecret)
}

// CheckAppidSecret 校验 appid和 appsecret 是否是配置文件中有效部分
func (au *AuthServer) CheckAppidSecret(appid, appsecret string) (*config.PluginConfig, bool) {
	for _, v := range au.conf.Plugins {
		if v.AppID == appid && v.AppSecret == appsecret {
			return v, true
		}
	}
	return nil, false
}

// Verifying 校验token
func (au *AuthServer) Verifying(token string) error {
	var pl jwt.Payload
	_, err := jwt.Verify([]byte(token), au.hs, &pl)
	if err != nil {
		return errors.New(AuthTokenErrorInvalidToken)
	}
	return nil
}

// CheckFormGrpcContext 从grpc的ctx校验token
func (au *AuthServer) CheckFormGrpcContext(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errors.New(AuthTokenErrorEmptyAuthorizeMetadata)
	}
	token, ok := md["authorization"]
	if !ok || len(token) == 0 {
		return errors.New(AuthTokenErrorEmptyAuthorizeMetadata)
	}
	rawToken := strings.Replace(token[0], "Bearer ", "", 1)
	return au.Verifying(rawToken)
}

// GetPayload 解码出payload
func (au *AuthServer) GetPayload(token string) (*jwt.Payload, error) {
	var pl jwt.Payload
	_, err := jwt.Verify([]byte(token), au.hs, &pl)
	return &pl, err
}

// GetPayload 解码出payload
func (au *AuthServer) GetPayloadFromGrpcContex(ctx context.Context) (*jwt.Payload, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New(AuthTokenErrorEmptyAuthorizeMetadata)
	}
	token, ok := md["authorization"]
	if !ok || len(token) == 0 {
		return nil, errors.New(AuthTokenErrorEmptyAuthorizeMetadata)
	}
	rawToken := strings.Replace(token[0], "Bearer ", "", 1)
	var pl jwt.Payload
	_, err := jwt.Verify([]byte(rawToken), au.hs, &pl)
	return &pl, err
}

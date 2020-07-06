package oauth2_logger_proxy

import (
	"context"
	"github.com/go-oauth2/oauth2/v4"
	"log"
)

type proxyManager struct {
	manager oauth2.Manager
}

func ProxyManager(manager oauth2.Manager) oauth2.Manager {
	return &proxyManager{
		manager: manager,
	}
}

func (pm proxyManager) GetClient(ctx context.Context, clientID string) (cli oauth2.ClientInfo, err error) {
	log.Printf("GetClient(%s)", clientID)
	cli, err = pm.manager.GetClient(ctx, clientID)
	log.Printf("GetClient() = %v, %v", cli, err)
	return cli, err
}

func (pm proxyManager) GenerateAuthToken(ctx context.Context, rt oauth2.ResponseType, tgr *oauth2.TokenGenerateRequest) (authToken oauth2.TokenInfo, err error) {
	log.Printf("GenerateAuthToken(%#v, %#v)", rt, tgr)
	authToken, err = pm.manager.GenerateAuthToken(ctx, rt, tgr)
	log.Printf("GenerateAuthToken() = %#v %v", authToken, err)
	return authToken, err
}

func (pm proxyManager) GenerateAccessToken(ctx context.Context, rt oauth2.GrantType, tgr *oauth2.TokenGenerateRequest) (accessToken oauth2.TokenInfo, err error) {
	log.Printf("GenerateAccessToken()")
	accessToken, err = pm.manager.GenerateAccessToken(ctx, rt, tgr)
	log.Printf("GenerateAccessToken() = %v, %v", accessToken, err)
	return accessToken, err
}

func (pm proxyManager) RefreshAccessToken(ctx context.Context, tgr *oauth2.TokenGenerateRequest) (accessToken oauth2.TokenInfo, err error) {
	log.Printf("RefreshAccessToken(%#v)", tgr)
	accessToken, err = pm.manager.RefreshAccessToken(ctx, tgr)
	log.Printf("RefreshAccessToken() = %v, %v", accessToken, err)
	return accessToken, err
}

func (pm proxyManager) RemoveAccessToken(ctx context.Context, access string) (err error) {
	log.Printf("RemoveAccessToken()")
	err = pm.manager.RemoveAccessToken(ctx, access)
	log.Printf("RemoveAccessToken() = %v", err)
	return err
}

func (pm proxyManager) RemoveRefreshToken(ctx context.Context, refresh string) (err error) {
	log.Printf("RemoveRefreshToken(%#v)", refresh)
	err = pm.manager.RemoveRefreshToken(ctx, refresh)
	log.Printf("RemoveRefreshToken() = %v", err)
	return err
}

func (pm proxyManager) LoadAccessToken(ctx context.Context, access string) (ti oauth2.TokenInfo, err error) {
	log.Printf("LoadAccessToken(%#v)", access)
	ti, err = pm.manager.LoadAccessToken(ctx, access)
	log.Printf("LoadAccessToken() = %v, %v", ti, err)
	return ti, err
}

func (pm proxyManager) LoadRefreshToken(ctx context.Context, refresh string) (ti oauth2.TokenInfo, err error) {
	log.Printf("LoadRefreshToken(%#v)", refresh)
	ti, err = pm.manager.LoadRefreshToken(ctx, refresh)
	log.Printf("LoadRefreshToken() = %v, %v", ti, err)
	return ti, err
}

package oauth2

import "time"

// Config authorization configuration parameters
type Config struct {
	// access token expiration time, 0 means it doesn't expire
	AccessTokenExp time.Duration
	// refresh token expiration time, 0 means it doesn't expire
	RefreshTokenExp time.Duration
	// whether to generate the refreshing token
	IsGenerateRefresh bool
}

// RefreshingConfig refreshing token config
type RefreshingConfig struct {
	// access token expiration time, 0 means it doesn't expire
	AccessTokenExp time.Duration
	// refresh token expiration time, 0 means it doesn't expire
	RefreshTokenExp time.Duration
	// whether to generate the refreshing token
	IsGenerateRefresh bool
	// whether to reset the refreshing create time
	IsResetRefreshTime bool
	// whether to remove access token
	IsRemoveAccess bool
	// whether to remove refreshing token
	IsRemoveRefreshing bool
}

// default configs
var (
	DefaultCodeExp               = time.Minute * 10
	DefaultAuthorizeCodeTokenCfg = &Config{AccessTokenExp: time.Hour * 24 * 365, RefreshTokenExp: time.Hour * 24 * 365, IsGenerateRefresh: true}
	DefaultImplicitTokenCfg      = &Config{AccessTokenExp: time.Hour * 24 * 365}
	DefaultPasswordTokenCfg      = &Config{AccessTokenExp: time.Hour * 24 * 365, RefreshTokenExp: time.Hour * 24 * 365, IsGenerateRefresh: true}
	DefaultClientTokenCfg        = &Config{AccessTokenExp: time.Hour * 24 * 365}
	DefaultRefreshTokenCfg       = &RefreshingConfig{IsGenerateRefresh: true, IsRemoveAccess: true, IsRemoveRefreshing: true}
)

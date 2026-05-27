package session

import (
	"net/http"
	"time"

	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/gomodule/redigo/redis"
)

func NewManager(pool *redis.Pool) *scs.SessionManager {
	manager := scs.New()

	manager.Store = redisstore.New(pool)

	manager.Lifetime = 7 * 24 * time.Hour

	manager.Cookie.HttpOnly = true
	manager.Cookie.Secure = true
	manager.Cookie.SameSite = http.SameSiteLaxMode

	return manager
}
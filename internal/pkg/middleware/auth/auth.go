package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"net/http"
)

func AuthCheck(store *sessions.CookieStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := store.Get(c.Request, "J_SESSION")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			return
		}

		if session.IsNew {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if session.Values["userID"] == nil || session.Values["userID"].(int64) < 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if session.Options.MaxAge <= 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

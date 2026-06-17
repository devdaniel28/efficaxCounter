package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Authrequired() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Key")
        tokenValue := os.Getenv("TOKENAPI")

        if token != tokenValue {
            c.AbortWithStatusJSON(403, gin.H{
                "erro":   "acesso negado",
                "codigo": 403,
            })
            return
        }
        c.Next()
    }
}
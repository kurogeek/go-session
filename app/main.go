package main

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name string
	Count int
}

func main() {
	r := gin.Default()
	keys := [][]byte{
		[]byte("secret-authen123"),
		[]byte("secret-encrypt12"),
	}
	store, _ := redis.NewStore(10, "tcp", "redis:6379", "", keys...)
	store.Options(sessions.Options{
		MaxAge: 60 * 60 * 24 * 1,
		// Secure:   true,
		HttpOnly: true,
	})
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/incr", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		fmt.Println("hello world 4")
		c.JSON(200, gin.H{"count": count})
	})
	r.GET("store/struct", func(c *gin.Context) {
		session := sessions.Default(c)
		v := session.Get("user")
		if v == nil {
			user := User{
				Name: "test",
				Count: 0,
			}
		} else {
			user = v.(User{})
			user.Count ++
		}
		session.Set("user", user)
		sess.Save()
		c.JSON(200, gin.H{"user": user})
	})
	r.GET("/decr", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count--
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})
	r.GET("/clear", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Save()
		c.JSON(200, gin.H{"status": "ok"})
	})
	r.Run(":8000")
}

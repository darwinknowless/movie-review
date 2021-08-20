package controller

import (
	"log"
	"movie-review/model"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (db *Gorm) Register(c *gin.Context) {
	var request model.User
	c.Bind(&request)

	// proses encrypt password dari request menggunakan library bcrypt
	hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	// var user berisi data yang sama dengan request hanya saja password yang dimasukkan adalah
	// password yang sudah di encrypt
	user := model.User{
		FullName: request.FullName,
		Role:     request.Role,
		Email:    request.Email,
		Password: string(hash),
	}

	// response juga sama hanya saja kita tidak menyertakan password
	response := model.User{
		FullName: request.FullName,
		Role:     request.Role,
		Email:    request.Email,
	}
	db.DB.Create(&user)
	writeResponse(c, response, nil)
}

func (db *Gorm) Login(c *gin.Context) {
	var (
		userReq model.User
		userDB  model.User
		result  gin.H
	)

	// data dari request body dimasukkan ke var userReq
	if err := c.Bind(&userReq); err != nil {
		log.Println("Format salah", err.Error())
	}

	// data dari DB sesuai dengan email dari request
	// dimasukkan ke var userReq
	db.DB.Where("email=?", userReq.Email).First(&userDB)

	// proses perbandingan password dari DB(hash password) dengan password dari request
	if err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(userReq.Password)); err != nil {
		log.Println("Password salah")
		log.Println(err)
	} else {
		// setelah berhasil login baru kita membuat tokennya

		// struct pembantu untuk membuat payload data
		type authCustomClaims struct {
			Email              string `json:"email"` // optional
			Role               string `json:"role"`  // optional
			jwt.StandardClaims        // data yang harus ada pada payload
		}

		// masukkan data kedalam payload
		claims := &authCustomClaims{
			userDB.Email,
			userDB.Role,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
		}

		// proses pembuatan token menggunakan HS256
		sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
		token, err := sign.SignedString([]byte(os.Getenv("SECRET")))
		if err != nil {
			log.Println("gagal create token")
		} else {
			log.Println("berhasil create token")
			result = gin.H{
				"token": token,
			}
		}

		c.JSON(200, result)
	}
}

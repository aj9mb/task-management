package helpers

import (
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/aj9mb/task-management/repo"
	"github.com/google/uuid"
)

func GenerateUserName(name string) (string, error) {
	possibleUserNames := make([]string, 0)
	strs := strings.Split(name, " ")
	fN := strs[0]
	possibleUserNames = append(possibleUserNames, fN)
	for i := 1; i < len(strs); i++ {
		possibleUserNames = append(possibleUserNames, fN+strs[i])
	}
	possibleUserNames = append(possibleUserNames, fN+"_"+GenerateRandomString(3))
	possibleUserNames = append(possibleUserNames, fN+"_"+GenerateRandomString(4))
	possibleUserNames = append(possibleUserNames, fN+"_"+GenerateRandomString(5))
	existUserNames, err := repo.CheckUserNameExist(possibleUserNames)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	for _, v := range existUserNames {
		for i, vl := range possibleUserNames {
			if v == vl {
				possibleUserNames = append(possibleUserNames[:i], possibleUserNames[i+1:]...)
			}
		}
	}
	return GetRandomVal(possibleUserNames), err
}

func GetRandomVal(arr []string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return arr[r.Intn(len(arr))]
}

func GenerateRandomString(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	charset := "0123456789abcdefghijklmnopqrstuvwxyz"
	var c string
	for i := 0; i < n; i++ {
		c += string(charset[r.Intn(len(charset))])
	}
	return c
}

func GetUuid() string {
	return uuid.NewString()
}

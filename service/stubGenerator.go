package service

import (
	"context"
	"fmt"
	"github.com/dmitrymatviets/myhood/core/contract"
	"github.com/dmitrymatviets/myhood/core/model"
	"github.com/dmitrymatviets/myhood/util"
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"math"
	"math/rand"
	"strings"
	"time"
)

const (
	namesFilePath    = "./mock_data/names.m.txt"
	surnamesFilePath = "./mock_data/surnames.txt"
)

func GetValidSignupDto(pass string) model.SignupDto {
	rand.Seed(time.Now().UnixNano())

	names := getStringsFromFile(namesFilePath)
	surnames := getStringsFromFile(surnamesFilePath)
	interests := []string{"программирование", "теннис", "книги", "кино", "музыка", "охота", "рыбалка", "футбол", "фото"}

	rand.Shuffle(len(interests), func(i, j int) { interests[i], interests[j] = interests[j], interests[i] })
	return model.SignupDto{
		Credentials: model.Credentials{
			Email:    fmt.Sprintf("test%d%d@test%d%d.com", rand.Int(), rand.Int(), rand.Int(), rand.Int()),
			Password: pass,
		},
		Name:        names[rand.Intn(len(names))],
		Surname:     surnames[rand.Intn(len(surnames))],
		DateOfBirth: time.Date(1970+rand.Intn(40), time.Month(rand.Intn(10)+1), rand.Intn(27)+1, 00, 00, 00, 00, time.Local),
		Gender:      "м",
		Interests:   append([]string{}, interests[0:rand.Intn(len(interests)/2)]...),
		CityId:      model.IntId(rand.Intn(4) + 1),
	}
}

func GenerateProfiles(userRepo contract.IUserRepository, pass string, count int, concurrency int) error {
	eg := errgroup.Group{}
	rateLimiter := make(chan struct{}, concurrency)
	for i := 0; i < count; i++ {
		profile := GetValidSignupDto(pass)
		rateLimiter <- struct{}{}
		if i%int(math.Ceil(float64(count)/1000.0)) == 0 {
			fmt.Printf("Generated %d profiles\r\n", i)
		}
		eg.Go(func() error {
			defer func() {
				<-rateLimiter
			}()
			_, _, err := userRepo.SignUp(context.Background(), profile.ToUserWithPassword())
			if err != nil {
				fmt.Println(err)
			}
			return err
		})
	}
	return eg.Wait()
}

func getStringsFromFile(path string) []string {
	path = util.RelPathToAbs(path)
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var result []string

	bufString := string(buf)
	for _, item := range strings.Split(bufString, "\n") {
		item = strings.TrimSpace(item)
		if len(item) > 0 {
			result = append(result, item)
		}
	}
	return result
}

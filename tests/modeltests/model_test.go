package modeltests

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/johnnyaustor/go-crud-jwt/api/controllers"
	"github.com/johnnyaustor/go-crud-jwt/api/models"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

var server = controllers.Server{}
var userInstance = models.User{}
var postInstance = models.Post{}

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	Database()

	os.Exit(m.Run())
}

func Database() {
	var err error

	TestDbDriver := os.Getenv("TestDbDriver")

	if TestDbDriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			os.Getenv("TestDBUser"),
			os.Getenv("TestDbPassword"),
			os.Getenv("TestDbHost"),
			os.Getenv("TestDbPort"),
			os.Getenv("TestDbName"))
		server.DB, err = gorm.Open(TestDbDriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
			log.Fatal("This is the error: ", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", TestDbDriver)
		}
	}

	if TestDbDriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
			os.Getenv("TestDbHost"),
			os.Getenv("TestDbPort"),
			os.Getenv("TestDbUser"),
			os.Getenv("TestDbName"),
			os.Getenv("TestDbPassword"))
		server.DB, err = gorm.Open(TestDbDriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
			log.Fatal("this is the error: ", err)
		} else {
			fmt.Printf("we ar connected to the %s database\n", TestDbDriver)
		}
	}
}

func refreshUserTable() error {
	err := server.DB.DropTableIfExists(&models.User{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.User{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	return nil
}

func seedOneUser() (models.User, error) {

	refreshUserTable()

	user := models.User{
		Nickname: "Pet",
		Email: "pet@jap.com",
		Password: "password",
	}

	err := server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		log.Fatalf("cannot sed users table; %v", err)
	}
	return user, nil
}

func seedUsers() error {
	users := []models.User {
		{
			Nickname: "Ronald Steven",
			Email:    "rsteven@jap.com",
			Password: "password",
		},
		{
			Nickname: "Victor Wotten",
			Email: "victor@jap.com",
			Password: "password",
		},
	}

	for i := range users {
		err := server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func refreshUserAndPostTable() error {
	err := server.DB.DropTableIfExists(&models.User{}, &models.Post{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		return err
	}
	log.Printf("succesfully refreshed tables")
	return nil
}

func seedOneUserAndOnePost() (models.Post, error) {
	err := refreshUserAndPostTable()
	if err != nil {
		return models.Post{}, err
	}

	user := models.User{
		Nickname: "Markus Miller",
		Email: "Mmil@jap.com",
		Password: "password",
	}
	err = server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		return models.Post{}, err
	}

	post := models.Post{
		Title: "judul markus miller",
		Content: "content markus miller",
		AuthorID: user.ID,
	}
	err = server.DB.Model(&models.Post{}).Create(&post).Error
	if err != nil {
		return models.Post{}, err
	}
	return post, nil
}

func seedUsersAndPosts() ([]models.User, []models.Post, error) {
	var err error

	if err != nil {
		return []models.User{}, []models.Post{}, err
	}

	var users = []models.User{
		{
			Nickname: "Pacarong",
			Email: "pacarong@jap.com",
			Password: "password",
		},
		{
			Nickname: "riana angel",
			Email: "rianaange@jap.com",
			Password: "password",
		},
	}
	var posts = []models.Post{
		{
			Title: "title 11",
			Content: "Content 11",
		},
		{
			Title: "title 12",
			Content: "Content 12",
		},
	}

	for i := range users {
		err = server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = server.DB.Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
	return users, posts, nil
}


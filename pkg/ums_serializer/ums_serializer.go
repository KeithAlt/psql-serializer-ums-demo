package ums_serializer

import (
	"bufio"
	"github.com/KeithAlt/psql-serializer-ums-demo/pkg/errorCheck"
	"math/rand"
	"os"
	"strconv"
)

// User profile structure
type profile struct {
	id         string
	name       string
	email      string
	userRole   string
	phone      string
	imageUrl   string
	userSecret string
}

// TODO: go:embed
const (
	fakeEmails = "./data/in/fake_emails"
	fakeJobs   = "./data/in/fake_jobs"
	fakeNames  = "./data/in/fake_names"
	fakePfps   = "./data/in/pfps"
)

// Generate psql insert statement from meta data
func (this *profile) GenPsqlValuesInsert() string {
	var psql string
	psql = `('` + this.id + `', '` + this.name + `', '` + this.email + `', '` + this.userRole + `', '` + this.phone + `', '` + this.imageUrl + `', '` + this.userSecret + `');`
	return psql
}

// Constructor for new profile
func CreateProfile(id, name, email, userRole, phone, imageUrl, userSecret string) profile {
	return profile{
		id:         id,
		name:       name,
		email:      email,
		userRole:   userRole,
		phone:      phone,
		imageUrl:   imageUrl,
		userSecret: userSecret,
	}
}

// Read target file & return contents in string slice
func readFileContents(path string) []string {
	var content []string

	file, err := os.Open(path)
	errorCheck.Fatal(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	errorCheck.Fatal(scanner.Err())
	return content
}

// gen a random phone number
func generateRandomPhone() string {
	phoneString := "("
	for i := 0; i < 3; i++ {
		phoneString += strconv.Itoa(0 + rand.Intn(9-0))
	}
	phoneString += ") "
	for i := 0; i < 3; i++ {
		phoneString += strconv.Itoa(0 + rand.Intn(9-0))
	}
	phoneString += "-"
	for i := 0; i < 4; i++ {
		phoneString += strconv.Itoa(0 + rand.Intn(9-0))
	}
	return phoneString
}

// Generate psql from data input
func Generate() string {
	names := readFileContents(fakeNames)
	emails := readFileContents(fakeEmails)
	jobs := readFileContents(fakeJobs)
	pfps := readFileContents(fakePfps)
	dbTable := "user_member"
	psqlHeader := "INSERT INTO " + dbTable + " (id, name, email, user_role, phone, image_url, user_secret) VALUES\n"
	psqlQuery := psqlHeader

	for i := 0; i < 99; i++ {
		p := CreateProfile(
			strconv.Itoa(i),
			names[i],
			emails[i],
			jobs[i],
			generateRandomPhone(),
			pfps[i],
			"a89vsad0ansaodjsa90vas0dmfnfslalkasmdv",
		)
		psqlQuery = psqlQuery + p.GenPsqlValuesInsert() + "\n"
	}

	return psqlQuery
}

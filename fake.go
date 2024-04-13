package assert

import (
	"fmt"
	"github.com/lucsky/cuid"
	"math/rand"
)

type FakeProfile struct {
	FirstName string
	LastName  string
	FullName  string
	Email     string
}

func FakeFirstName() string {
	randomNames := []string{
		"John", "Jane", "Bob", "Alice", "Charlie", "David", "Eve", "Frank",
		"Grace", "Heidi", "Ivan", "Judy", "Kevin", "Laura", "Michael", "Nancy",
		"Olivia", "Peter", "Quincy", "Rachel", "Steve", "Tina", "Ursula", "Victor",
		"Wendy", "Xander", "Yvonne", "Zach",
	}
	return randomNames[rand.Intn(len(randomNames))]
}

func FakeLastName() string {
	randomLastNames := []string{
		"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller",
		"Davis", "Rodriguez", "Martinez", "Hernandez", "Lopez", "Gonzalez",
		"Wilson", "Anderson", "Thomas", "Taylor", "Moore", "Jackson", "Martin",
		"Lee", "Perez", "Thompson", "White", "Harris", "Sanchez", "Clark",
		"Ramirez", "Lewis", "Robinson", "Walker", "Young", "Allen", "King",
		"Wright", "Scott", "Torres", "Nguyen", "Hill", "Flores", "Green",
		"Adams", "Nelson", "Baker", "Hall", "Rivera", "Campbell", "Mitchell",
		"Carter", "Roberts",
	}
	return randomLastNames[rand.Intn(len(randomLastNames))]
}

func GetFakeProfile() *FakeProfile {
	firstName := FakeFirstName()
	lastName := FakeLastName()
	domains := []string{"gmail.com", "yahoo.com", "outlook.com", "aol.com", "protonmail.com"}

	return &FakeProfile{
		FirstName: firstName,
		LastName:  lastName,
		FullName:  fmt.Sprintf("%s %s", firstName, lastName),
		Email: fmt.Sprintf(
			"%s-%s-%s@%s",
			firstName,
			lastName,
			cuid.New(),
			domains[rand.Intn(len(domains))],
		),
	}
}

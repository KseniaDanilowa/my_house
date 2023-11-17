package family

type Family struct {
	Name   string
	Age    int
	Status string
}

func CreateFamily() []Family {
	familyMembers := []Family{
		{Name: "Elena", Age: 45, Status: "Mother"},
		{Name: "Ksenia", Age: 20, Status: "Daughter"},
	}
	return familyMembers
}

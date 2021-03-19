package address

import "testing"

type AddressTest struct {
	Address        string
	ExpectedReturn string
}

// go test ./...  = run all tests
// go test --cover = check how many percent your test is covered
// go test --coverprofile result.txt = generate text file your test is covered
// go tool cover --func=result.txt = read file your test is covered
// go tool cover --html=result.txt = read file your test is covered in browser

func TestTypeAddress(t *testing.T) {
	t.Parallel() //RUN TEST PARALLEL

	adressesTest := []AddressTest{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça das Rosas", "Invalid type"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		{"", "Invalid type"},
	}

	for _, address := range adressesTest {
		returnReceived := TypeAddress(address.Address)
		if returnReceived != address.ExpectedReturn {
			t.Errorf("The type received %s is different from the expected %s",
				returnReceived,
				address.ExpectedReturn,
			)
		}
	}
}

func TestAny(t *testing.T) {
	t.Parallel()

	numberOne := 1
	numberTwo := 2

	if numberOne > numberTwo {
		t.Errorf("No!")
	}
}

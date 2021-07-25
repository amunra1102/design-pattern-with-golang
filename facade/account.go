package facade

import "fmt"

type account struct {
	name string
}

func newAccount(name string) *account {
	return &account{name: name}
}

func (a *account) checkAccount(name string) error {
	if a.name != name {
		return fmt.Errorf("account %s is incorrect", name)
	}

	fmt.Println("Account verified")
	return nil
}

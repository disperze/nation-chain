package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Person is a struct that contains all dni info
type Person struct {
	Name       string         `json:"name"`
	MiddleName string         `json:"middle_name"`
	Surname1   string         `json:"surname1"`
	Surname2   string         `json:"surname2"`
	Creator    sdk.AccAddress `json:"creator"`
}

// NewPerson returns a new Person
func NewPerson() Person {
	return Person{}
}

// implement fmt.Stringer
func (w Person) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Creator: %s
FirstName: %s
MiddleName: %s
First Surname: %s
Second Surname: %s`, w.Creator, w.Name, w.MiddleName, w.Surname1, w.Surname2))
}

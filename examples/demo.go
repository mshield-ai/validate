package main

import (
	"fmt"
	"github.com/gookit/validate"
)

type UserForm struct {
	Name string `json:"name" validate:"required,minLen(2),customValidator"`
}

func (f UserForm) CustomValidator(val string) {

}

func main() {
	myV := &validate.Validator{
		Name: "test",
		Func: func() error {
			return nil
		},
	}

	fmt.Printf("%v\n", myV)

	mData := map[string]interface{}{
		"name": "inhere",
		"age":  28,
	}

	v := validate.New()
	v.SetRules(validate.Rules{
		"name": "required,minSize(2),customValidator",
	})

	validate.Map(mData)
}

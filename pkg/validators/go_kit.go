package validators

import (
	"regexp"

	"github.com/gookit/validate"
)

func NewCheckerGoKit(vPtr any) error {
	v := validate.New(vPtr)

	//
	v.SkipOnEmpty = false // should be false
	v.StopOnError = true
	initGlobalValidator()
	initGlobalMsg(v)
	//

	v.Validate()

	if v.Errors.Empty() {
		return nil
	}

	return v.Errors
}

func TranslatorGoKit(err error) string {
	if err != nil {
		val, ok := err.(validate.Errors)
		if ok {
			return string(val.One())
		}
		return err.Error()
	}
	return ""
}

func initGlobalMsg(v *validate.Validation) {
	v.AddMessages(map[string]string{
		"minLength":         "{field} min length is %d",
		"isE164PhoneNumber": "Invalid Phone Number Format.",
	})
}

func initGlobalValidator() {
	validate.AddValidator(
		"isE164PhoneNumber",
		func(val string) bool {
			return regexp.MustCompile(`^\+?[1-9]\d{1,14}$`).Match([]byte(val))
		},
	)
}

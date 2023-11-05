package validator

import "errors"

type ProfileValidator struct {
}

func NewProfileValidator() *ProfileValidator {
	return &ProfileValidator{}
}

func (e ProfileValidator) ValidEmail(email string) (err error) {
	if len(email) < 5 {
		err = errors.New("email is too short")
	}

	isDog := false
	isDot := false
	for _, character := range email {
		if 64 == character {
			isDog = true
		} else if 46 == character {
			isDot = true
		}
		if isDog && isDot {
			return
		}
	}
	err = errors.New("the email is not correct")
	return
}

func (p ProfileValidator) ValidPassword(passwordHash string) (err error) {
	if len(passwordHash) < 8 {
		err = errors.New("password is too short")
	}
	isNumbers := false
	isCharacter := false
	for _, character := range passwordHash {
		if 48 <= character && character <= 57 {
			isNumbers = true
		} else {
			isCharacter = true
		}
		if isNumbers && isCharacter {
			return
		}
	}
	err = errors.New("the password is too simple: you can add a character and number")
	return
}

func (p ProfileValidator) ValidPhone(phone uint64) (err error) {
	if phone/10_000_000_000 == 0 {
		err = errors.New("phone is too short")
	}
	return
}

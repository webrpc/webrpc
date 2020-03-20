package newerr

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"testing"
)

func TestErr(t *testing.T) {

	// handler responds with this..
	appErr := ErrNotFound
	// customErr := errors.New("some custom error..")

	if errors.Is(appErr, ErrNotFound) {
		fmt.Println("yes!", appErr)
	}

	appErr2 := fmt.Errorf("%w: yaww who knows..", ErrNotFound)
	if errors.Is(appErr2, ErrNotFound) {
		fmt.Println("still yes!", appErr2)
	}

	notfoundErr := ErrorNotFound("can't find you")
	if errors.Is(notfoundErr, ErrNotFound) {
		fmt.Println("who you..?", notfoundErr)
	}

	fmt.Println("")

	cause := errors.New("because!")
	notFoundErrWithCause := Errorf(ErrNotFound, cause, "can't find you")
	if errors.Is(notFoundErrWithCause, cause) {
		fmt.Println("yes cause")
	}
	if errors.Is(notfoundErr, notFoundErrWithCause) {
		fmt.Println("yes notfoundeErr")
	}

	// nested..
	fmt.Println("")
	fmt.Println("--nested--")
	fmt.Println("")

	wrappedErr := Errorf(ErrCanceled, notFoundErrWithCause, "ooooooooooo")
	if errors.Is(wrappedErr, ErrCanceled) {
		fmt.Println("cancelled!")
	}
	if errors.Is(wrappedErr, notFoundErrWithCause) {
		fmt.Println("roger that, good.")
	}
	if errors.Is(wrappedErr, ErrNotFound) {
		fmt.Println("oooooooo")
	}
	fmt.Println("ha?", wrappedErr)

	data, err := json.Marshal(wrappedErr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("====>", string(data))

	var e Error
	err = json.Unmarshal(data, &e)
	if err != nil {
		panic(err)
	}
	fmt.Println("??????????", e)

	xx := GetErrorStack(wrappedErr)
	for _, e := range xx {
		fmt.Println("-->", e)
	}

}

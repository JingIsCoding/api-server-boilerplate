package exceptions

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ErrorTestSuite struct {
	suite.Suite
}

type unitTest struct {
	name string
	run  func()
}

func (suite *ErrorTestSuite) TestIsError() {
	tests := []unitTest{
		unitTest{
			name: "should be the same error if it is same instance",
			run: func() {
				suite.True(errors.Is(DatbaseConnectionFailed, DatbaseConnectionFailed), "not the same")
			},
		},
		unitTest{
			name: "should be the same error even if it wrap differnt error",
			run: func() {
				err1 := DatbaseConnectionFailed.Wrap(errors.New("some error"))
				err2 := DatbaseConnectionFailed.Wrap(errors.New("some other error"))
				suite.True(errors.Is(err1, err2), "not the same")
			},
		},
		unitTest{
			name: "should be the same error even if it has differnt message",
			run: func() {
				err1 := DatbaseConnectionFailed.SetMessage("some message")
				err2 := DatbaseConnectionFailed.SetMessage("some other message")
				suite.True(errors.Is(err1, err2), "not the same")
			},
		},
		unitTest{
			name: "should not be the same error if is differnt error",
			run: func() {
				err1 := UserCreateFaled.SetMessage("some message")
				err2 := UserAlreadyExist.SetMessage("some other message")
				suite.False(errors.Is(err1, err2), "the same")
			},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, test.run)
	}
}

func TestErrorTestSuite(t *testing.T) {
	suite.Run(t, new(ErrorTestSuite))
}

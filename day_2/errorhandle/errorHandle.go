package errorHandle

import "errors"

func ErrorHandle(a, b int) (int, error) {
	if a == b {
		return 0, errors.New("this are equal")
	}
	 if  b==0 {
		return  0, errors.New("cannot divide by zero")
	}
	 if a/b < 0 {
		return 0, errors.New("resuslt is negative")
	}
	return a / b, nil
}

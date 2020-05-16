package service

import "strconv"

type status int

func (i status) Error() string {
	return strconv.Itoa(int(i))
}

package util

import "regexp"

const (
	emailRegex = "[a-zA-Z0-9]{1,}@[a-zA-Z0-9]{1,}\\.[a-z]{1,}"
	iinRegex   = `(\d{2}(0[1-9]|1[0-2])(0[1-9]|[12][0-9]|3[01]))[0-9]{6}`
)

//Increment ...
func Increment(num uint64) uint64 {
	return num + 1
}

//Decrement ...
func Decrement(num uint64) uint64 {
	if num == 0 {
		return 0
	}
	return num - 1
}

//RandomStringList ...
func RandomStringList() []string{
	return []string{}
}

//FindEmailFromText ...
func FindEmailFromText(text string) string {
	emailRegexp:=regexp.MustCompile(emailRegex)
	first:=emailRegexp.FindString(text)
	return first
}

//FindAllEmailFromText ...
func FindAllEmailFromText(text string) []string {
	emailRegexp := regexp.MustCompile(emailRegex)
	all := emailRegexp.FindAllString(text, -1)
	return all
}

//FindIinFromText ...
func FindIinFromText(text string) string{
	iinRegex:= regexp.MustCompile(iinRegex)
	first:=iinRegex.FindString(text)
	return first
}
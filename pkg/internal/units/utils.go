package units

import (
	"regexp"
	"strconv"
	"strings"
)

func ExtractInt(arg string) (num int, str string) {
	re := regexp.MustCompile("[0-9]+")
	temp := re.FindAllString(arg, -1)
	var nums []int
	for _, v := range temp {
		i, parseErr := strconv.Atoi(v)
		if parseErr != nil {
			continue
		}
		nums = append(nums, i)
	}
	str = arg
	for _, v := range temp {
		str = strings.ReplaceAll(str, v, "")
	}

	var nStr string
	for i := range nums {
		nStr += strconv.Itoa(nums[i])
	}
	num, _ = strconv.Atoi(nStr)
	return
}

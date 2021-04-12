package stringutils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*
	公民身份证号
	xxxxxx yyyy MM dd 375 0     十八位
	xxxxxx   yy MM dd  75 0     十五位

	地区：[1-9]\d{5}
	年的前两位：(18|19|([23]\d))      1800-2399
	年的后两位：\d{2}
	月份：((0[1-9])|(10|11|12))
	天数：(([0-2][1-9])|10|20|30|31) 闰年不能禁止29+

	三位顺序码：\d{3}
	两位顺序码：\d{2}
	校验码：   [0-9Xx]

	十八位：^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$
	十五位：^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}$

	总：
	(^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$)|(^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}$)
*/
func IsIDCard(data string) bool {
	if len(data) == 18 {
		return checkIDCardLast(data) && isIDCard(data)
	} else {
		return isIDCard(data)
	}
}

func isIDCard(data string) bool {
	return MatchString(`(^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$)|(^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{2}$)`, data)
}

//第二代身份证校验码（GB 11643-1999）
func checkIDCardLast(data string) bool {
	cardNo := strings.ToUpper(data)
	checks := []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}
	var sum int
	for i := 17; i > 0; i-- {
		n, _ := strconv.Atoi(cardNo[17-i : 18-i])
		p := int(math.Pow(2, float64(i))) % 11
		sum += n * p
	}
	return cardNo[17:] == checks[sum%11]
}

//银行卡号
func IsBankCard(data string) bool {
	return MatchString(`^[0-9]{13,19}$`, data) && checkLuhn(data)
}

//银行卡号校验：Luhn算法（模10算法）
//Luhn校验规则：16位银行卡号（19位通用）:
//1.将未带校验位的 15（或18）位卡号从右依次编号 1 到 15（18），位于奇数位号上的数字乘以 2。
//2.将奇位乘积的个十位全部相加，再加上所有偶数位上的数字。
//3.将加法和加上校验位能被 10 整除。
func checkLuhn(bankNo string) bool {
	var sum int
	oddSum := []int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}
	odd := len(bankNo) & 1
	for i, c := range bankNo {
		if c < '0' || c > '9' {
			return false
		}
		if i&1 == odd {
			sum += oddSum[c-'0']
		} else {
			sum += int(c - '0')
		}
	}
	return sum%10 == 0
}

/*
	验证所给手机号码是否符合手机号的格式.
	移动: 134、135、136、137、138、139、150、151、152、157、158、159、182、183、184、187、188、178(4G)、147(上网卡)；
	联通: 130、131、132、155、156、185、186、176(4G)、145(上网卡)、175；
	电信: 133、153、180、181、189 、177(4G)；
	卫星通信:  1349
	虚拟运营商: 170、173
	2018新增: 16x, 19x
*/
func IsMobile(data string) bool {
	return MatchString(`^13[\d]{9}$|^14[5,7]{1}\d{8}$|^15[^4]{1}\d{8}$|^16[\d]{9}$|^17[0,3,5,6,7,8]{1}\d{8}$|^18[\d]{9}$|^19[\d]{9}$`, data)
}

//国内座机电话号码："XXXX-XXXXXXX"、"XXXX-XXXXXXXX"、"XXX-XXXXXXX"、"XXX-XXXXXXXX"、"XXXXXXX"、"XXXXXXXX"
func IsTel(data string) bool {
	return MatchString(`^((\d{3,4})|\d{3,4}-)?\d{7,8}$`, data)
}

//Email地址
func IsEmail(data string) bool {
	//return MatchString(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`, data)
	return MatchString(`^[a-zA-Z0-9_\-\.]+@[a-zA-Z0-9_\-]+(\.[a-zA-Z0-9_\-]+)+$`, data)
}

//URL地址
func IsURL(data string) bool {
	return MatchString(`(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`, data)
}

//Mac地址
func IsMac(data string) bool {
	return MatchString(`^([0-9A-Fa-f]{2}[\-:]){5}[0-9A-Fa-f]{2}$`, data)
}

//腾讯QQ号，从10000开始
func IsQQ(data string) bool {
	return MatchString(`^[1-9][0-9]{4,}$`, data)
}

//邮政编码
func IsPostCode(data string) bool {
	return MatchString(`^\d{6}$`, data)
}

//是日期，日月格式必须为01
func IsDateFormat(data string, sep string) bool {
	var pattern string
	if sep == "" {
		pattern = `^(18|19|2\d|3\d)\d{2}((0[1-9])|10|11|12)((0[1-9])|([12][0-9])|30|31)$`
	} else if sep == "年" {
		pattern = `^(18|19|2\d|3\d)\d{2}[年]((0[1-9])|10|11|12)[月]((0[1-9])|([12][0-9])|30|31)[日]$`
	} else {
		pattern = `^(18|19|2\d|3\d)\d{2}[` + sep + `]((0[1-9])|10|11|12)[` + sep + `]((0[1-9])|([12][0-9])|30|31)$`
	}
	return MatchString(pattern, data)
}

//检查账号（字母开头，数字字母下划线）
// @param length 长度验证： 1个值时为指定长度；2个值时分别为 min 和 max
func IsAccount(data string, length ...uint) bool {
	var lengthStr string
	if len(length) == 1 && length[0] > 0 {
		lengthStr = fmt.Sprintf("{%d}", length[0]-1)
	} else if len(length) == 2 && length[0] <= length[1] && length[0] > 0 {
		lengthStr = fmt.Sprintf("{%d,%d}", length[0]-1, length[1]-1)
	} else {
		lengthStr = "{5,19}"
	}
	return MatchString(fmt.Sprintf(`^[A-Za-z]{1}[0-9A-Za-z_]%s$`, lengthStr), data)
}

//检查密码
// @param length 长度验证： 1个值时为指定长度；2个值时分别为 min 和 max
func IsPwd(data string, level uint, length ...uint) bool {
	var lengthStr string
	if len(length) == 1 && length[0] > 0 {
		lengthStr = fmt.Sprintf("{%d}", length[0])
	} else if len(length) == 2 && length[0] <= length[1] && length[0] > 0 {
		lengthStr = fmt.Sprintf("{%d,%d}", length[0], length[1])
	} else {
		lengthStr = "{6,20}"
	}
	switch level {
	case 1: //包含数字、字母
		return MatchString(fmt.Sprintf(`^[\w\S]%s$`, lengthStr), data) && HasNumber(data) && HasAlpha(data)
	case 2: //包含数字、字母、下划线
		return MatchString(fmt.Sprintf(`^[\w\S]%s$`, lengthStr), data) && HasNum_Alpha(data)
	case 3: //包含数字、字母、特殊字符
		return MatchString(fmt.Sprintf(`^[\w\S]%s$`, lengthStr), data) && HasNumber(data) && HasAlpha(data) && HasChar(data)
	case 4: //包含数字、大小写字母
		return MatchString(fmt.Sprintf(`^[\w\S]%s$`, lengthStr), data) && HasNumber(data) && HasUpper(data) && HasLower(data)
	case 5: //包含数字、大小写字母、下划线
		return MatchString(fmt.Sprintf(`^[\w\S]%s$`, lengthStr), data) && HasNumber(data) && HasUpper(data) && HasLower(data) && MatchString("[_]", data)
	case 6: //包含数字、大小写字母、特殊字符
		return MatchString(fmt.Sprintf(`^[\w\S]%s$`, lengthStr), data) && HasNumber(data) && HasUpper(data) && HasLower(data) && HasChar(data)
	default:
		return MatchString(fmt.Sprintf(`^[\w\S]%s$`, lengthStr), data)
	}
}

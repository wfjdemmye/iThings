package utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/carlmjohnson/versioninfo"
	"math/big"
	"net"
	"net/http"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

func init() {
	PrintVersion()
}

func MD5V(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}

/*
检测用户名是否符合规范 只可以使用字母数字及下划线 最多30个字符
*/
func CheckUserName(name string) error {
	if len(name) > 30 {
		return errors.New("pwd len more than 30")
	}
	if IsMobile(name) {
		return errors.New("pwd can't be phone number")
	}
	if IsEmail(name) {
		return errors.New("pwd can't be email")
	}
	return nil
}

/*
检测密码是否符合规范 需要至少8位 并且需要包含数字和字母
*/
//密码强度必须为字⺟⼤⼩写+数字+符号，9位以上
func CheckPasswordLever(ps string) int32 {
	level := int32(0)
	if len(ps) < 8 {
		return 0
	}
	num := `[0-9]{1}`
	a_z := `[a-z]{1}`
	A_Z := `[A-Z]{1}`
	symbol := `[!@#~$%^&*()+|_]{1}`

	if b, err := regexp.MatchString(num, ps); b && err == nil {
		level++
	}
	if b, err := regexp.MatchString(a_z, ps); b && err == nil {
		level++
	}
	if b, err := regexp.MatchString(A_Z, ps); b && err == nil {
		level++
	}
	if b, err := regexp.MatchString(symbol, ps); b && err == nil {
		level++
	}
	return level
}

// 识别手机号码
func IsMobile(mobile string) bool {
	result, _ := regexp.MatchString(`^(1[0-9][0-9]\d{4,8})$`, mobile)
	if result {
		return true
	} else {
		return false
	}
}

func IsEmail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

/*
将密码的md5和uid进行md5
*/
func MakePwd(pwd string, uid int64, isMd5 bool) string {
	if isMd5 == false {
		pwd = MD5V([]byte(pwd))
	}
	strUid := strconv.FormatInt(uid, 8)
	return MD5V([]byte(pwd + strUid + "god17052709767"))
}

// 获取正在运行的函数名
func FuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	funcs := strings.Split(f.Name(), "/")
	if len(funcs) > 0 {
		return funcs[len(funcs)-1]
	}
	return f.Name()
}

//func GetPos()string{
//	pc, file, line, _ := runtime.Caller(2)
//	f := runtime.FuncForPC(pc)
//
//	fmt.Sprintf("%s:%d:%s\n\n\n",file,line,f.Name())
//	return fmt.Sprintf("%s:%d:%s",file,line,f.Name())
//}

func Ip2binary(ip string) string {
	str := strings.Split(ip, ".")
	var ipstr string
	for _, s := range str {
		i, err := strconv.ParseUint(s, 10, 8)
		if err != nil {
			fmt.Println(err)
		}
		ipstr = ipstr + fmt.Sprintf("%08b", i)
	}
	return ipstr
}

// MatchIP 测试IP地址和地址端是否匹配 变量ip为字符串，例子"192.168.56.4" iprange为地址端"192.168.56.64/26"
func MatchIP(ip, iprange string) bool {
	ipb := Ip2binary(ip)
	if strings.Contains(iprange, "/") { //如果是ip段
		ipr := strings.Split(iprange, "/")
		masklen, err := strconv.ParseUint(ipr[1], 10, 32)
		if err != nil {
			return false
		}
		iprb := Ip2binary(ipr[0])
		return strings.EqualFold(ipb[0:masklen], iprb[0:masklen])
	} else {
		return ip == iprange
	}

}

// @Summary 获取真实的源ip
func GetIP(r *http.Request) (string, error) {
	ip := r.Header.Get("X-Real-IP")
	if net.ParseIP(ip) != nil {
		return ip, nil
	}

	ip = r.Header.Get("X-Forward-For")
	for _, i := range strings.Split(ip, ",") {
		if net.ParseIP(i) != nil {
			return i, nil
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	if net.ParseIP(ip) != nil {
		return ip, nil
	}

	return "", errors.New("no valid ip found")
}

func MethodToNum(methond string) string {
	switch methond {
	case "GET":
		return "1"
	case "POST":
		return "2"
	case "HEAD":
		return "3"
	case "OPTIONS":
		return "4"
	case "PUT":
		return "5"
	case "DELETE":
		return "6"
	case "TRACE":
		return "7"
	case "CONNECT":
		return "8"
	default:
		return "-1"
	}
}

func PrintVersion() {
	fmt.Printf("gitInfo: lastCommitTime:%v,lastCommitHash:%v\n",
		versioninfo.LastCommit, versioninfo.Revision)
}

// IP_int64转stringx.x.x.x
func InetNtoA(ip int64) string {
	return fmt.Sprintf("%d.%d.%d.%d", byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}

// string转int64   x.x.x.x->int64
func InetAtoN(ip string) int64 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return ret.Int64()
}

import "strings"
import "strconv"

const (
    IPv4 = "IPv4"
    IPv6 = "IPv6"
    Neither = "Neither"
)

func validIPAddress(IP string) string {
    if len(IP) < 7 || len(IP) > 4*8+7 { return Neither }
    for i := 0; i < 5; i++ {
        if IP[i] == '.' { if validIPv4(IP) {return IPv4} else { return Neither} }
        if IP[i] == ':' { if validIPv6(IP) {return IPv6} else { return Neither} }
    }
    return Neither
}

func validIPv4(IP string) bool {
    groups := strings.Split(IP, ".")
    if len(groups) != 4 { return false }
    for i := 0; i < 4; i++ {
        numStr := groups[i]
        if len(numStr) > 3 || len(numStr) == 0 || numStr[0] == '-' { return false }
        num, err := strconv.Atoi(numStr)
        if err != nil || len(numStr) > 1 && numStr[0] == '0' || num > 255 { return false }
    }
    return true
}

func validIPv6(IP string) bool {
    groups := strings.Split(IP, ":")
    if len(groups) != 8 { return false }
    for i := 0; i < 8; i++ {
        numStr := groups[i]
        if len(numStr) > 4 || len(numStr) == 0 || numStr[0] == '-' { return false }
        num32, err := strconv.ParseInt(numStr, 16, 32)
        if err != nil || num32 > 0xFFFF { return false }
        if len (numStr) >= 3 && numStr[0] == 0 && numStr[1] == 0 { return false }
    }
    return true
}

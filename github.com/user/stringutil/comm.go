package stringutil

import (
    "strings"
    "bytes"
)

func Basename(s string) string {
    //Discard last '/' and everyting before.
    for i := len(s)-1; i >= 0; i--{
        if s[i] == '/'{
            s = s[i+1:]
            break
        }
    }

    //Preserve everything before last '.'
    for i := len(s)-1; i >=0; i--{
        if s[i] == '.'{
            s = s[:i]
            break
        }
    }

    return s
}

func Basename2(s string) string {
    slash := strings.LastIndex(s,"/")
    s = s[slash+1:]
    if dot := strings.LastIndex(s,"."); dot>=0{
        s = s[:dot]
    }

    return s
}


func Reverse(s string) string{
    r := []rune(s)
    for i,j:=0, len(r)-1; i<len(r)/2; i,j=i+1,j-1{
        r[i],r[j] = r[j],r[i]
    }
    
    return string(r)
}

func ReverseInt(s []int){
    for i,j := 0,len(s)-1; i < j; i,j=i+1,j-1{
        s[i],s[j] = s[j],s[i]
    }
}


func HasPrefix(s, prefix string) bool {
    return len(s)>=len(prefix) && s[:len(prefix)]==prefix
}

func HasSuffix(s, suffix string) bool{
    return len(s)>=len(suffix)&& s[len(s)-len(suffix):]==suffix
}

func Contains(s, substr string) bool{
    for i:=0; i<len(s); i++{
        if HasPrefix(s[i:],substr){
            return true
        }
    }
    return false
}


func Comma(s string) string{
    i := 1
    dot := len(s)-i*3-(i-1)
    for ; dot>0 ;{
        s = s[:dot] + "," + s[dot:]
        i++
        dot = len(s)-i*3-(i-1)
    }

    return s
}

func Comma2(s string) string{
    if len(s)<=3{
        return s
    }

    var buf bytes.Buffer
    
    i := len(s)%3
    if i>0{
        buf.WriteString(s[:i])
    }

    for k:=i; i<len(s); i++{
        if (i-k)%3==0 {
            buf.WriteString(",")
        }
        buf.WriteByte(s[i])
    }

    return buf.String()
}

func Comma3Helper(s string) string{
    if len(s)<=3{
        return s
    }

    var buf bytes.Buffer
    for i:=0;i<len(s);i++{
        if i!=0 && i%3==0{
            buf.WriteString(",")
        }
        buf.WriteByte(s[i])
    }

    return buf.String()
}

func Comma3(s string) string{
    
    if len(s)<=3{
        return s
    }

    var buf bytes.Buffer
    flag := 0
    if s[0]=='+' || s[0]=='-'{
        buf.WriteByte(s[0])
        flag = 1
    }

    dot := strings.LastIndex(s,".")
    if dot==-1{
        buf.WriteString(Comma2(s[flag:]))
    }else{
        buf.WriteString(Comma2(s[flag:dot]))
        buf.WriteByte('.')
        buf.WriteString(Comma3Helper(s[dot+1:]))
    }

    return buf.String()
}
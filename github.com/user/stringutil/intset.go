package stringutil

import (
    "bytes"
    "fmt"
)

type IntSet struct{
    words []uint64
}

func (s *IntSet) Has(x int) bool{
    word,bit := x/64,uint(x%64)
    return word<len(s.words) && s.words[word]&(1<<bit)!=0
}

func (s *IntSet) Add(x int){
    word,bit := x/64,uint(x%64)
    for word >= len(s.words){
        s.words = append(s.words,0)
    }
    s.words[word] |= (1<<bit)
}

func (s *IntSet) String() string {
    var buf bytes.Buffer
    buf.WriteByte('{')
    for i,word := range s.words{
        if word==0{
            continue
        }

        for j:=0;j<64;j++{
            if word&(1<<uint(j)) != 0{
                if buf.Len()> len("{"){
                    buf.WriteByte(' ')
                }
                fmt.Fprintf(&buf, "%d", 64*i+j)
            }
        }
    } 
    buf.WriteByte('}')
    return buf.String()
}

func (s* IntSet) Len() int{
    cnt:=0

    for _,word := range s.words{
        if word==0{
            continue
        }

        for j:=0;j<64;j++{
            if word&(1<<uint(j)) != 0{
                cnt++
            }
        }
    } 

    return cnt
} 

func (s *IntSet) Remove(x int){
    word,bit := x/64,uint(x%64)
    if word<len(s.words) && s.words[word]&(1<<bit)!=0{
        s.words[word] = s.words[word] &^ (1<<bit)
    }
}

func (s *IntSet) Clear(){
    s.words = s.words[:0]
}

func (s *IntSet) Copy() *IntSet{
    // words := make([]uint64,len(s.words),cap(s.words))
    // copy(words, s.words)

    words := s.words
    var ret IntSet
    ret.words = words
    return &ret
}

func (s *IntSet) AddAll(vals...int){
    for _,val := range vals{
        s.Add(val)
    }
}

func (s *IntSet) UnionWith(t *IntSet) {
    for i,tword := range t.words{
        if i<len(s.words){
            s.words[i] |= tword
        }else{
            s.words = append(s.words,tword)
        }
    }
}

func (s *IntSet) IntersectWith(t *IntSet){
    for i,word := range s.words{
        if word==0{
            continue
        }

        for j:=0;j<64;j++{
            if (word&(1<<uint(j)) != 0) && !(t.Has(64*i+j)) {
                s.Remove(64*i+j)
            }
        }
    } 
}

func (s *IntSet) DifferenceWith(t *IntSet) {
    for i,word := range s.words{
        if word==0{
            continue
        }

        for j:=0;j<64;j++{
            if (word&(1<<uint(j)) != 0) && (t.Has(64*i+j)) {
                s.Remove(64*i+j)
            }
        }
    }
}
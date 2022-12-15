package slice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sync"
)

type Slice struct {
	sl    []any
	mutex sync.Mutex
}

func (s *Slice) Add(data any) (index int64) {
	v := append(s.sl, data)
	s.sl = v
	return int64(len(s.sl)) - 1
}

func (s *Slice) Delete(index int64) (ok bool) {
	if index >= int64(len(s.sl)) {
		return false
	}
	s.sl = append(s.sl[:index], s.sl[index+1:]...)
	return true
}

func (s *Slice) Get(index int64) (val any) {
	if index > int64(len(s.sl)) {
		return nil
	} else {
		return s.sl[index]
	}
}

func (s *Slice) Print() {
	fmt.Println("Length", len(s.sl))
	//fmt.Println(s.sl)
	for i := 0; i < len(s.sl); i++ {
		fmt.Println(s.sl[i])
	}

}

func aGrB(a any, b any) bool {
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(a)
	aBytes := reqBodyBytes.Bytes()

	reqBodyBytes = new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(b)
	bBytes := reqBodyBytes.Bytes()
	if len(aBytes) > len(bBytes) {
		return true
	}
	if len(aBytes) < len(bBytes) {
		return false
	}
	for i := 0; i < len(bBytes); i++ {
		if aBytes[i] > bBytes[i] {
			return true
		}
	}
	return false
}
func NewSlice() *Slice {
	l := &Slice{}
	return l
}
func (s *Slice) SortIncrease() {

	s.Sort(true)
}

func (s *Slice) SortDecrease() {
	s.Sort(false)
}

func (s *Slice) Len() int64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return int64(len(s.sl))
}

func (s *Slice) String() string {
	return fmt.Sprintf("%+v", s.sl)
}

func (s *Slice) Sort(increase bool) {

	switch increase {
	case true:
		for j := 1; j < len(s.sl); j++ {
			for i := 1; i < len(s.sl); i++ {
				if aGrB(s.sl[i-1], s.sl[i]) {
					a := s.sl[i-1]
					s.sl[i-1] = s.sl[i]
					s.sl[i] = a
				}
			}
		}
	case false:
		for j := 1; j < len(s.sl); j++ {
			for i := 1; i < len(s.sl); i++ {
				if !aGrB(s.sl[i-1], s.sl[i]) {
					a := s.sl[i-1]
					s.sl[i-1] = s.sl[i]
					s.sl[i] = a
				}
			}
		}
	}

}

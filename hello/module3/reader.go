package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte)(n int , e error){
	n,e = rot.r.Read(b)	
	for i := 0 ; i<len(b) ; i++ {
		if ( b[i] >= 'a' && b[i] < 'n' ) || (b[i] >= 'A' && b[i] < 'N'){
			b[i] += 13
		}else if ( b[i] > 'm' && b[i] <= 'z' ) || (b[i] > 'M' && b[i] <= 'Z'){
			b[i] -= 13
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
//	fmt.Println(s)
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}


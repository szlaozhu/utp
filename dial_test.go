package utp

import "testing"

func TestDial(t *testing.T) {
	addr, err := ResolveAddr("utp", ":0")
	if err != nil {
		t.Fatal(err)
	}

	l, err := Listen("utp", addr)
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()

	ch := make(chan struct{})
	go func() {
		l.Accept()
		close(ch)
	}()

	c, err := DialUTP("utp", nil, l.Addr().(*Addr))
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	<-ch
}

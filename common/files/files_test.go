package files

import "testing"

import (
	"io/ioutil"
	"os"
	"os/exec"
	"fmt"
	//"io"
)

func TestXopenFiles(t *testing.T) {
	f, _ := ioutil.TempFile(".", "__")
	s    := "This is great stuff"
	f.Write([]byte(s))

	// Testing reading a regular file
	data   := make([]byte, len(s))
	xf, xr := Xopen(f.Name())
	if i, err := xr.Read(data); err != nil || i < len(s) || string(data) != s {
		t.Errorf("Problems reading a regular file with Xopen()")
	}
	xf.Close()

	// Testing bzip gzip2
	bzGzipTesting := func(gzipOrBzip, ext string) {
		for i, _ := range data { data[i] = 0 } // let's clean up data
		cmd := exec.Command(gzipOrBzip, f.Name())
		if err := cmd.Run(); err != nil {
			t.Errorf("Problems compressing with gzip")
		}
		xf, xr = Xopen(fmt.Sprintf("%s.%s", f.Name(), ext) )
		if i, err := xr.Read(data); err != nil || i < len(s) || string(data) != s {
			t.Errorf("Problems reading a regular file with Xopen()")
		}
		xf.Close()
		cmd = exec.Command(gzipOrBzip, "-d", fmt.Sprintf("%s.%s", f.Name(), ext) )
		if err := cmd.Run(); err != nil {
			t.Errorf("Problems de-compressing with gzip", err)
		}
	}

	bzGzipTesting("gzip", "gz")
	bzGzipTesting("bzip2", "bz2")

	os.Remove(f.Name())
	os.Remove(fmt.Sprintf("%s.gz" ,f.Name()))
	os.Remove(fmt.Sprintf("%s.bz2" ,f.Name()))
	f.Close()
}

// TODO: Test stdin case
/*
func TestXopenStdin(t *testing.T) {
	s := "This is great stuff"
	os.Stdin.Write([]byte(s));
	xf, xr := Xopen("-")
	data := make([]byte, len(s))
	if i, err := xr.Read(data); err != io.EOF || string(data) != s {
		t.Errorf("Problems reading a regular file with Xopen('-')", string(data) != s, i, len(s), err)
	}
	xf.Close()
}
*/

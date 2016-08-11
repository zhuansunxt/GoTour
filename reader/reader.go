// Implement a Reader type that emits an 
// infinite stream of the ASCII character 'A'.

// The purpose this practice is to get familiar with the io.Reader interface
// and to know the Read() method defined by that interface.

package reader

//import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (MyReader) Read (b []byte) (n int, e error) {
	b[0] = 'A'
	return 1, nil;
}
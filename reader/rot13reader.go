// A common pattern is an io.Reader that wraps another io.Reader, 
// modifying the stream in some way.

// Implement a rot13Reader that implements io.Reader and reads from an io.Reader, 
// modifying the stream by applying the rot13 substitution cipher to all 
// alphabetical characters.

package reader 

import "io"

// customized type wrap a io.Reader type 
type rot13reader struct {
	r io.Reader
}

func (rot rot13reader) Read (b []byte) (n int, e error) {
	n, e = rot.r.Read(b)

	// Apply rot13 substitution.
	for i:= 0; i < len(b); i++ {
		if (b[i] >= 'A' && b[i] < 'N') || (b[i] >= 'a' && b[i] < 'n') {
			b[i] += 13
		} else if (b[i] > 'M' && b[i] <= 'Z') || (b[i] > 'm' && b[i] <= 'z') {
			b[i] -= 13
		}
	}

	return ;
}



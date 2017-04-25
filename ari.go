package ReadablityTools

import (
	"bytes"
	"fmt"
)

// @param text: what you want to caculate the readablity souce
// @return source
func ARI(text []byte) (source float32) {
	words := bytes.Split(text,[]byte{' ','\t','.','!','?',','})
	if len(words) < 100 {
		fmt.Errorf("Too few words\n")
		return
	}
	sentence := bytes.Split(text,[]byte{'.','!','?'})
	var charCount = 0
	for _, word := range words {
		charCount += len(word)
	}
	source = 4.71 * float32(charCount)/float32(len(words))
	source += 0.5 * float32(len(words))/float32(len(sentence)) - 21.34
	return
}

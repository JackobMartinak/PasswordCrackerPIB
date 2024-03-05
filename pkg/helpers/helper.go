package helpers

import (
	"bytes"
	"crypto/sha256"
	"log"
	"os"
)

type PasswordAndHash struct {
  Hash [32]byte
  Pass string
  Found bool
} 

// Konverzia hesla do 256Sha
func ConvertToHash(pass string) <-chan PasswordAndHash {
  out := make(chan PasswordAndHash)
  go func() {
    n := sha256.Sum256([]byte(pass))
    st := PasswordAndHash{Hash: n, Pass: pass, Found: false}
    out <- st

    close(out)
  }()
  return out
}

// Porovnavanie Hashov
func CompareHash(ph <-chan PasswordAndHash, passhash []byte) {
  
  for n := range ph {
    if bytes.Equal(n.Hash[:], passhash) {
      log.Printf("Found Password: %s", n.Pass)
      os.Exit(1)
    }
  }
}

//chocolate TODO: Spravit funkciu na rozoznavanie zakladnych hashovacich algoritmov a potom tomu prisposobit crackovanie 

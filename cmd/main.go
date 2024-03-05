package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
	"github.com/JackobMartinak/PasswordCrackerPIB/pkg/helpers"
)

func main()  {
  var wordlist string
  var hashInput string
  var passwordHash []byte
  
  // Nacitanie wordlistu a hashu na cracknutie

  flag.StringVar(&wordlist, "wordlist", "", "list of words")
  flag.StringVar(&hashInput, "hash", "", "password hash for cracking")
  flag.Parse()
  
  // Decodovanie na mozne pouzitie pri porovnavani
  passwordHash, err := hex.DecodeString(hashInput)
  if err != nil {
    fmt.Println("Error decoding hash:", err)
    os.Exit(1)
  }

  if len(wordlist) == 0 || len(hashInput) == 0 {
    fmt.Println("Usage: main.go -wordlist [path] -hash [hash]")
    os.Exit(1)
  }

  file, err := os.Open(wordlist)
  if err != nil {
    log.Fatalln("Error opening the wordlist, Exiting", err)
    os.Exit(1)

  }

  defer file.Close()
  
  // Scanner daneho wordlistu po riadkoch
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    if len(scanner.Text()) == 0 {
      continue
    }
    // Concurency - Go routines  
    out := helpers.ConvertToHash(scanner.Text())
    helpers.CompareHash(out, passwordHash)
  }

  if err := scanner.Err(); err != nil {
    log.Fatalln(err)
  }

}

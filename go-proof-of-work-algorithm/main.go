package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"time"
)

func main() {
    
	hash_result := ""

	for difficulty_bits := 0; difficulty_bits < 32; difficulty_bits++ {
		difficulty := math.Pow(2, float64(difficulty_bits))
		fmt.Printf("Difficulty: %v (%v)\n", int(difficulty), difficulty_bits)
		fmt.Printf("Searching...\n")

		new_block := "This is a test block" + hash_result

		start_time := time.Now()

	    proof_of_work(new_block, difficulty_bits)

		
		elapsed_time := time.Since(start_time) * 1000
		fmt.Printf("Elaspsed time: %v\n", elapsed_time)

	}
}

func proof_of_work(header string, difficulty_bits int) (string, int) {
	max_nonce := math.Pow(2, 32)
	difficulty_target := math.Pow(2, float64(256-difficulty_bits))
	fmt.Printf("Difficulty target is %v\n", difficulty_target)

	for nonce := 0; nonce <= int(max_nonce); nonce++ {
		text := header + strconv.Itoa(nonce)
		hash := sha256.New()
		hash.Write([]byte(text))
		hash_result := hex.EncodeToString((hash.Sum(nil)))		

        // Convert hash_result to bigint
		ch := big.NewInt(0)
        ch.SetString(hash_result, 16)
        converted_hash := ch.String()
        
		// Convert difficulty target to a string, then to bigint to aid comparison with conerted hash_result
		ct := big.NewInt(0)
        ct.SetString(strconv.Itoa(int(difficulty_target)), 16)
        converted_target := ct.String()
	
		if converted_hash <= (converted_target) {
			fmt.Printf("Success!\n")
			fmt.Printf("Hash target is %v\n", hash_result)
			fmt.Printf("Nonce is %v\n", nonce)
			return hash_result, nonce
		}
	}
	return "", 0
}

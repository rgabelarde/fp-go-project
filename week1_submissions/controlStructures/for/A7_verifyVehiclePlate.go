package main

import (
	"fmt"
	"strings"
	"unicode"
)

func letterToNumber(letter rune) int {
    return int(unicode.ToUpper(letter)) - int('A') + 1
}

func processPrefix(prefix string) (int, int) {
    switch len(prefix) {
    case 3:
        // Use last 2 letters for 3-letter prefix
        return letterToNumber(rune(prefix[1])), letterToNumber(rune(prefix[2]))
    case 2:
        // Use both letters for 2-letter prefix
        return letterToNumber(rune(prefix[0])), letterToNumber(rune(prefix[1]))
    case 1:
        // Use 0 for first position and the letter for second position
        return 0, letterToNumber(rune(prefix[0]))
    default:
        return 0, 0
    }
}

func processNumbers(numbers string) []int {
    // Pad with leading zeros if less than 4 digits
    for len(numbers) < 4 {
        numbers = "0" + numbers
    }
    
    result := make([]int, 4)
    for i := 0; i < 4; i++ {
        result[i] = int(numbers[i] - '0')
    }
    return result
}

func calculateChecksum(numbers []int) string {
    // Fixed multipliers
    multipliers := []int{9, 4, 5, 4, 3, 2}
    
    // Checksum letters
    checksumLetters := []string{"A", "Z", "Y", "X", "U", "T", "S", "R", "P", "M", "L", "K", "J", "H", "G", "E", "D", "C", "B"}
    
    sum := 0
    for i := 0; i < len(numbers); i++ {
        sum += numbers[i] * multipliers[i]
    }
    
    remainder := sum % 19
    return checksumLetters[remainder]
}

func main() {
    var plateNumber string
    fmt.Print("Enter vehicle plate number: ")
    fmt.Scanln(&plateNumber)
    
    // Split the plate number into prefix and numbers
    var prefix, numbers, checkLetter string
    
    // Find the position where numbers start
    numberStart := -1
    for i, char := range plateNumber {
        if unicode.IsDigit(char) {
            numberStart = i
            break
        }
    }
    
    if numberStart == -1 {
        fmt.Println("Invalid plate number format")
        return
    }
    
    prefix = plateNumber[:numberStart]
    numbers = ""
    checkLetter = string(plateNumber[len(plateNumber)-1])
    
    // Extract numbers (excluding the last character which is the check letter)
    for i := numberStart; i < len(plateNumber)-1; i++ {
        if unicode.IsDigit(rune(plateNumber[i])) {
            numbers += string(plateNumber[i])
        }
    }
    
    // Process prefix
    num1, num2 := processPrefix(prefix)
    
    // Process numbers
    numArray := processNumbers(numbers)
    
    // Combine all numbers
    allNumbers := []int{num1, num2}
    allNumbers = append(allNumbers, numArray...)
    
    // Calculate checksum
    calculatedChecksum := calculateChecksum(allNumbers)
    
    // Compare with given check letter
    if strings.ToUpper(checkLetter) == calculatedChecksum {
        fmt.Println("The vehicle plate given is correct!")
    } else {
        fmt.Println("The vehicle plate given is not correct!")
    }
}

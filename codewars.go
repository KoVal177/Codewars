package main

import "fmt"
import "strconv"

import "strings"

//import "bufio"
//import "os"
import "math"

//import "reflect"

func countSheep(num int) string {
  var output string
  for i := 1; i <= num; i++ {
    output += strconv.Itoa(i) + " sheep..."
  }
  return output
}

func Multiply(a int, b int) int {
  result := a * b
  return result
}

func Summation(n int) int {
  var output int
  for i := 1; i <= n; i++ {
    output += i
  }
  return output
}

func ToAlternatingCase(str string) (output string) {
  tmp := []rune(str)
  for i, letter := range str {
    if letter >= 97 && letter <= 122 {
      tmp[i] = letter - 32
    } else if letter >= 65 && letter <= 90 {
      tmp[i] = letter + 32
    } else {
      tmp[i] = letter
    }
  }
  output = string(tmp)
  return
}

func Maps(x []int) []int {
  for i := 0; i < len(x); i++ {
    x[i] += x[i]
  }
  return x
}

func GetSize(w, h, d int) [2]int {
  var output [2]int
  output[0] = (w*h + w*d + h*d) * 2
  output[1] = w * h * d
  return output
}

func Points(games []string) int {
  var score int
  for _, j := range games {
    if j[0] > j[2] {
      score += 3
    } else if j[0] == j[2] {
      score += 1
    }
    //score = (j[0] - 48) +
    //score, _ = strconv.Atoi(string(j[0]))
  }
  return score
}

func century(year int) int {
  cent := year / 100
  if cent*100 < year {
    cent++
  }
  return cent
}

func DNAStrand(dna string) string {
  dna_compl := []rune(dna)
  for i, j := range dna {
    switch j {
    case 'A':
      dna_compl[i] = 'T'
    case 'T':
      dna_compl[i] = 'A'
    case 'C':
      dna_compl[i] = 'G'
    case 'G':
      dna_compl[i] = 'C'
    }
  }
  return string(dna_compl)
}

func SumOfIntegersInString(strng string) (output int) {
  var tmp int
  for _, j := range strng {
    if j >= '0' && j <= '9' {
      tmp = tmp*10 + int(j) - '0'
    } else if tmp > 0 {
      output += tmp
      tmp = 0
    }
  }
  output += tmp
  return output
}

func Evaporator(content float64, evapPerDay int, threshold int) int {
  var day int
  evap_coeff := (100 - float64(evapPerDay)) / 100
  thresh_ml := (content * float64(threshold)) / 100
  for content > thresh_ml {
    content *= evap_coeff
    day++
    fmt.Println(content)
  }
  return day
}

func ContainAllRots(strng string, arr []string) bool {
  var res bool
  for _, letter := range strng {
    strng = strng[1:] + string(letter)
    res = false
    for _, element := range arr {
      if strng == element {
        res = true
      }
    }
    if !res {
      return res
    }

  }
  return true
}

type Tuple struct {
  Char  rune
  Count int
}

func OrderedCount(text string) []Tuple { //(output []Tuple) {
  var output = []Tuple{}
  var tmp = map[rune]int{}
  var queue string
  for _, letter := range text {
    tmp[letter] += 1
    if tmp[letter] == 1 {
      queue += string(letter)
    }
  }
  for _, letter := range queue {
    output = append(output, Tuple{Char: letter, Count: tmp[letter]})
  }
  return output
}

func MxDifLg(a1 []string, a2 []string) int {
  var max_a, min_a, max_b, min_b int
  if (len(a1) == 0) || (len(a2) == 0) {
    return -1
  }
  for _, string := range a1 {
    switch {
    case min_a == 0:
      min_a = len(string)
    case len(string) > max_a:
      max_a = len(string)
    case len(string) < min_a:
      min_a = len(string)
    }
    for _, string := range a2 {
      switch {
      case min_b == 0:
        min_b = len(string)
      case len(string) > max_b:
        max_b = len(string)
      case len(string) < min_b:
        min_b = len(string)
      }
    }
  }
  max_a = max_a - min_b
  if max_a < 0 {
    max_a = -max_a
  }
  max_b = max_b - min_a
  if max_b < 0 {
    max_b = -max_b
  }
  if max_b > max_a {
    max_a = max_b
  }
  return max_a
}

func WallPaper(l, w, h float64) string {
  var numbers = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen", "twenty"}
  var resI int
  if l*w*h == 0 {
    return "zero"
  }
  resF := (l + w) * 2 * h * 1.15 / 5.2
  resI = int(resF)
  if resF > float64(resI) {
    resI++
  }
  return numbers[resI]
}

func NewAvg(arr []float64, navg float64) int64 {
  var sum float64
  for _, i := range arr {
    sum += i
  }
  sum = float64(len(arr)+1)*navg - sum
  switch {
  case sum < 0:
    return -1
  case float64(int64(sum)) < sum:
    sum++
  }
  return int64(sum)
}

func Add(x int) func(int) int {
  return func(y int) int {
    return x + y
  }
}

func HighAndLow(in string) string {
  var min, max, num int
  tmp := strings.Fields(in)
  min, _ = strconv.Atoi(tmp[0])
  max = min
  for _, st := range tmp {
    num, _ = strconv.Atoi(st)
    switch {
    case num > max:
      max = num
    case num < min:
      min = num
    }

  }
  in = fmt.Sprintf("%d %d", max, min)
  return in
}

func Arithmetic(a int, b int, operator string) int {
  switch operator {
  case "add":
    return a + b
  case "subtract":
    return a - b
  case "multiply":
    return a * b
  case "divide":
    return a / b
  }
  return 0
}

func CartesianNeighbor(x, y int) [][]int {
  var res [][]int
  for i := x - 1; i <= x+1; i++ {
    for j := y - 1; j <= y+1; j++ {
      if (i != x) || (j != y) {
        res = append(res, []int{i, j})
      }
    }
  }
  return res
}

//func Solvegcd(s int, g int) []int {
func Solve(s int, g int) []int {
  if s%g != 0 {
    return []int{-1, -1}
  } else {
    return []int{g, s - g}
  }
}

func Divisors(n int) int {
  var sum int
  limit := int(math.Sqrt(float64(n)))
  for i := 1; i <= limit; i++ {
    if n%i == 0 {
      sum++
    }
  }
  sum *= 2
  if n == limit*limit {
    sum--
  }
  return sum
}

func Dative(word string) string {
  wordSl := []rune(word)
  s1 := "eéiíöőüű"
  s2 := "aáoóuú"
  for i := len(wordSl) - 1; i >= 0; i-- {
    for _, j := range s1 {
      if wordSl[i] == j {
        return word + "nek"
      }
    }
    for _, j := range s2 {
      if wordSl[i] == j {
        return word + "nak"
      }
    }
  }
  return "false"
}

func Angle(n int) int {
  return (n*180 - 360)
}

func Factorial(n int) int {
  var factorial = 1
  for i := 1; i <= n; i++ {
    factorial *= i
  }
  return factorial
}

func FindShort(s string) int {
  //fmt.Println(s)
  var start, res int
  res = len(s)
  for i, j := range s {
    switch j {
    case ' ':
      if i > start {
        if res > i-start {
          res = i - start
        }
        start = i + 1
      }
    // case of "'s"
    case '\'':
      if res > i-start {
        res = i - start
      }
      start = i + 3
    }
  }
  // check length of last word
  if res > (len(s) - start) {
    res = len(s) - start
  }
  return res
}

type NBAPlayer struct {
  Team string
  Ppg  float64
}

func SumPpg(playerOne, playerTwo NBAPlayer) float64 {
  return playerOne.Ppg + playerTwo.Ppg
}

type PosPeaks struct {
  Pos   []int
  Peaks []int
}

/*
func PickPeaks(array []int) PosPeaks {
  var output = PosPeaks{[]int{}, []int{}}
  var diffs = []int{}
  var peaks = []int{}
  for i, elem := range array {
    switch {
    case i == 0:
      diffs = append(diffs, 0)
      continue
    case array[i-1] > elem:
      diffs = append(diffs, -1)
    case array[i-1] == elem:
      diffs = append(diffs, 0)
    case array[i-1] < elem:
      diffs = append(diffs, 1)
    }
  }
  flag_peak := false
  start := 0
  for i, elem := range diffs {
    switch {
    case elem == -1 && flag_peak:
      peaks = append(peaks, start)
      flag_peak = false
    case elem == 1:
      start = i
      flag_peak = true
    }
  }
  for _, i := range peaks {
    output.Pos = append(output.Pos, i)
    output.Peaks = append(output.Peaks, array[i])
  }
  return output
}
*/

func PickPeaks(array []int) PosPeaks {
  // algorithm consider every "up" movement as possible peak starting point and waits for "down" movement for confirmation. then repeats
  var output = PosPeaks{[]int{}, []int{}}
  flag_peak := false // flag signals we are in possible peak phase. can be removed and instead varible "start" used as condition (0 or not)
  start := 0         // starting point of possible peak
  for i := 1; i < len(array); i++ {
    switch {
    case array[i-1] < array[i]: // movement up - possible peak
      start = i
      flag_peak = true
    case array[i-1] > array[i] && flag_peak: //movement down - peak confirmed
      output.Pos = append(output.Pos, start)
      output.Peaks = append(output.Peaks, array[start])
      flag_peak = false
    }
  }
  return output
}

func BouncingBall(h, bounce, window float64) int {
  switch {
  case !(bounce > 0 && bounce < 1):
    return -1
  case h <= window:
    return -1
  }
  count := 1
  for i := h * bounce; i > window; i *= bounce {
    count += 2
  }
  return count
}

func NbMonths(startPriceOld, startPriceNew, savingperMonth int, percentLossByMonth float64) [2]int {
  var sum float64
  pO := float64(startPriceOld)
  pN := float64(startPriceNew)
  sM := float64(savingperMonth)
  pLoss := percentLossByMonth / 100
  j := 0
  sum = pO - pN + sM*float64(j)

  for sum < 0 {
    pLoss += 0.005 * float64(j%2)
    j++
    pO *= (1 - pLoss)
    pN *= (1 - pLoss)
    sum = pO - pN + sM*float64(j)
    fmt.Println(sum, pLoss)
  }
  return ([2]int{int(j), int(sum)})
}

func Gcdi(x, y int) int {
  switch {
  case x < 0:
    x = -x
  case y < 0:
    y = -y
  }
  for i := Mini(x, y); i >= 1; i-- {
    if x%i == 0 && y%i == 0 {
      return i
    }
  }
  return x
}

func Som(x, y int) int {
  return x + y
}
func Maxi(x, y int) int {
  if x >= y {
    return x
  } else {
    return y
  }
}
func Mini(x, y int) int {
  if x <= y {
    return x
  } else {
    return y
  }
}
func Lcmu(x, y int) int {
  switch {
  case x < 0:
    x = -x
  case y < 0:
    y = -y
  }
  for i := Maxi(x, y); true; i++ {
    if i%x == 0 && i%y == 0 {
      return i
    }
  }
  return x
}

type FParam func(int, int) int

func OperArray(f FParam, arr []int, init int) []int {
  output := []int{}
  output = append(output, f(init, arr[0]))
  for i, j := range arr[1:] {
    output = append(output, f(output[i], j))
  }
  return output
}

func main() {
  var a = map[string]string{}
  a["a"] = "df"
  fmt.Println(a)

}

/*func main() {
  var a = map[string]int{"a": 5, "df": 6}
  //a := make(map[string]int)
  a["a"] = 5
  a["b"] = 1231
  fmt.Println(a)

  //  a := Evaporator(10, 10, 5)
  //  fmt.Println(a)
}
*/

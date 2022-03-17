package main
import "fmt"

func convert(num int){
   var biner []int

   for num !=0 {
      biner = append(biner, num%2)
      num = num / 2
   }
   if len(biner)==0{
      fmt.Printf("%d\n", 0)
   } else {
      for i:=len(biner)-1; i>=0; i--{
         fmt.Printf("%d", biner[i])
      }
      fmt.Println()
   }
}

func main() {
	var n int
	fmt.Scanln(&n)
	convert(n)
}
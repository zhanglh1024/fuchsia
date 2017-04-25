package main
import(
	"fmt"
)
func main(){
	tim:=[]int{2,3,5,33,35,22,34,21}
	for i,j:=range tim{
		fmt.Println(i,"-----",j)
	}
}

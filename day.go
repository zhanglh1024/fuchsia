package main
import (
	"fmt"
	"github.com/jinzhu/now"
	"time"
)
func main(){
	day:=time.Now().Format("2006-01-02")+" 23:59:59"
	dayTail,_:=time.Parse("2006-01-02 15:04:05",day)
	datData:=dayTail.AddDate(0,0,5).Unix()
	endDay:=now.EndOfMonth().Unix()
	tm:=time.Unix(endDay,0)
	fmt.Println(day,"----",tm.Format("2006-01-02 15:04:05"),"------",dayTail,endDay,datData)
}

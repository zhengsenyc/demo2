package main
import(
	"fmt"
	"os"
	"strings"
)

func main(){
	fmt.Println(strings.Join(os.Args[:],"\n"))
	fmt.Println(os.Args)
}

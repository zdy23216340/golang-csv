## CSV handler For golang.

## Usage
### Is very simple.

``` 
package main

import "github.com/zdy23216340/golang-csv"
import "fmt"

func main(){

	myCsv := csv.New()
	content := myCsv.WriteTitle([]string{"id","name","age"}).WriteRow([]string{"1","John","35"}).String()//.ToFile("./a.csv") 
	fmt.Println(content)
	// Then you can write to file.

	myLocalCsv := csv.Read("./a.csv")
}
```

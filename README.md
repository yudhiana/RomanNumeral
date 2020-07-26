# Roman Numeral Converter

## architecture code:
- uncle bob
    - model
    - repository
    - usecase

## Testing
- Golang package testing

---
Example: 
``` 
import (
	"fmt"
	"romans_numeric/usecase"
)

func main() {
	cases := usecase.NewRoman()
	fmt.Println(cases.RomanToNumber("MCMXLV"))
	fmt.Println(cases.NumberToRoman(1945))
}
```

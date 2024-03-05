package webserver

import (
	"fmt"
	"math"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	jariJari := 7
	tinggi := 10
	strOutput := fmt.Sprintf(`jariJari : %d, tinggi: %d, volume : %f, luas alas: %f, keliling alas: %f`, jariJari, tinggi, math.Pi*math.Pow(float64(jariJari), 2)*float64(tinggi), math.Pi*math.Pow(float64(jariJari), 2), 2*math.Pi*float64(jariJari))
	fmt.Fprintln(w, strOutput)
}

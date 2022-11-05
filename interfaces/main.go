package main

import "fmt"

// interface contract data for struct (interface keyword)
type KalkulasiGaji interface {
	gaji() int
}

type pegawaiTetap struct {
	gajipokok    int
	gajitambahan int
}

type pegawaiMagang struct {
	gajipokok int
}

type freelacer struct {
	gajipokok int
	jamker    int
}

// init an interface in func of struct
func (pt pegawaiTetap) gaji() int {
	return pt.gajipokok + pt.gajitambahan
}

func (pm pegawaiMagang) gaji() int {
	return pm.gajipokok
}

func (f freelacer) gaji() int {
	return f.gajipokok * f.jamker
}

// calculate total with interface
func totalGaji(kalk []KalkulasiGaji) int {
	total := 0
	for _, t := range kalk {
		total += t.gaji() // method of interface
	}
	return total
}

func main() {
	pt1 := pegawaiTetap{4000000, 2000000}
	pt2 := pegawaiTetap{4000000, 1000000}

	pm := pegawaiMagang{4000000}

	f1 := freelacer{100000, 5}
	f2 := freelacer{150000, 5}

	// calc all
	total := totalGaji([]KalkulasiGaji{pt1, pt2, pm, f1, f2})

	fmt.Println(total)
}

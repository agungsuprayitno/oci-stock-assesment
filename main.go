package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

type Cicilan struct {
	AngsuranKe        int    `json:"angsuranKe"`
	Tanggal           string `json:"tanggal"`
	TotalAngsuran     string `json:"totalAngsuran"`
	AngsuranPokok     string `json:"angsuranPokok"`
	AngsuranBunga     string `json:"angsuranBunga"`
	SisaAngsuranPokok string `json:"sisaAngsuranPokok"`
}

func main() {
	http.HandleFunc("/hitung-cicilan", func(w http.ResponseWriter, r *http.Request) {
		plafon := 5000000
		sukuBunga := 8
		lamaPinjaman := 6
		sisaPlafon := float64(plafon)

		var i = 1
		var semuaCicilan []Cicilan
		for i <= lamaPinjaman {
			angsuran := hitungAngsuranTetap(plafon, sukuBunga, lamaPinjaman)
			angsuranBunga := hitungBungaPerBulan(sisaPlafon, sukuBunga)
			angsuranPokok := hitungAngsuranPokokPerBulan(angsuran, angsuranBunga)
			sisaPlafon = hitungSisaPlafon(sisaPlafon, angsuran)

			cicilan := Cicilan{
				AngsuranKe:        i,
				Tanggal:           "tanggal",
				TotalAngsuran:     fmt.Sprintf("%.2f", angsuran),
				AngsuranPokok:     fmt.Sprintf("%.2f", angsuranPokok),
				AngsuranBunga:     fmt.Sprintf("%.2f", angsuranBunga),
				SisaAngsuranPokok: fmt.Sprintf("%.2f", sisaPlafon),
			}
			semuaCicilan = append(semuaCicilan, cicilan)

			i++
		}

		cicilanResponse, err := json.Marshal(semuaCicilan)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(cicilanResponse)
	})

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func hitungBungaPerBulan(plafon float64, sukuBunga int) float64 {
	var sukuBungaPerbulan float64 = (float64(sukuBunga) / 100) / 12
	bungaPerBulan := float64(plafon) * sukuBungaPerbulan
	return bungaPerBulan
}

func hitungAngsuranTetap(plafon int, sukuBunga int, lamaPinjaman int) float64 {
	var sukuBungaPerbulan float64 = (float64(sukuBunga) / 100) / 12
	angsuranTetap := float64(plafon) * (sukuBungaPerbulan / (1 - math.Pow((1+sukuBungaPerbulan), float64(-lamaPinjaman))))
	return angsuranTetap
}

func hitungAngsuranPokokPerBulan(angsuranTetap float64, bungaPerBulan float64) float64 {
	angsuranPokok := angsuranTetap - bungaPerBulan
	return angsuranPokok
}

func hitungSisaPlafon(plafon float64, angsuranTetap float64) float64 {
	sisaPlafon := fmt.Sprintf("%.2f", float64(plafon)-angsuranTetap)
	floatsisaPlafon, err := strconv.ParseFloat(sisaPlafon, 64)

	if err != nil {
		fmt.Println("Error:", err)
	}

	if floatsisaPlafon < 0 {
		return 0
	} else {
		return floatsisaPlafon
	}

}

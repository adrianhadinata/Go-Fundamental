package main
 
import (
    "fmt"
)
 
func main() {
	//Tulis kode disini

	var jumlahBambu int
	var jumlahPotong int

	fmt.Scan(&jumlahBambu)

	sekatInput := 0
	var sekat []int 

	for i := 1; i <= jumlahBambu; i++ {
		fmt.Scan(&sekatInput)
		sekat = append(sekat, sekatInput)
	}
	
	fmt.Scan(&jumlahPotong)

	for i := 1; i <= jumlahPotong; i++ {
		fmt.Printf("Potongan ke- %d\n", i)
		
		for j := 0; j < jumlahBambu; j++ {
			sekat[j] = sekat[j] - 1
			
			hasil := sekat[j]

			if hasil < 0 {
				hasil = 0
			}
			
			fmt.Println(hasil)
		}
	}
}
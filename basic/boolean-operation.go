package basic

import "fmt"

func MainBooleanOperation() {
	var (
		ujian        = 80
		absensi      = 80
		lulusUjian   = ujian >= 80
		lulusAbsensi = absensi >= 80
	)

	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	lulus := lulusAbsensi && lulusUjian
	fmt.Println(lulus)
}

package money_text

var units = []string{"", "satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan"}
var teens = []string{"sepuluh", "sebelas", "dua belas", "tiga belas", "empat belas", "lima belas", "enam belas", "tujuh belas", "delapan belas", "sembilan belas"}
var tens = []string{"", "sepuluh", "dua puluh", "tiga puluh", "empat puluh", "lima puluh", "enam puluh", "tujuh puluh", "delapan puluh", "sembilan puluh"}

func konversiRatusan(n int) string {
	str := ""
	if n >= 100 {
		str += units[n/100] + " ratus"
		n %= 100
		if n > 0 {
			str += " "
		}
	}
	if n >= 20 {
		str += tens[n/10]
		n %= 10
		if n > 0 {
			str += " "
		}
	}
	if n >= 10 {
		str += teens[n-10]
		n = 0
	}
	if n > 0 {
		str += units[n]
	}
	return str
}

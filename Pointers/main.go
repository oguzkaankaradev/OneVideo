package main

import "fmt"

//pointerlar birçok dilde vardır, java ve c# ta referans tip olarak karşımıza çıkmaktadır.
//c# ve java daha üst seviye diller olduğu için direkt olarak pointer diye kullanamıyoruz.
//c ve c++ ta pointerlar çok kez kullanıyoruz, yeri geldikçe godada kullanacağız.
//pontersa veritiplerinin başına yıldız(*) değeri koyacağız.
// Pointerlerde bir değişkendir, ancak görevleri adres tutmaktır.
//adres değerindeki değeri almak istiyorsak pointer değişkenimizin başıona * koymamız yeterli.

//int bool string gibi veritipleri golangda ilkel veri tipleri olarak geçmektedir. yani değer tiplidir. Fonksiyonla değiştirmek istiyorsak adres değeri üzerinden müdahalede bulunlıyız.
//array ve sliceler referans tiplidir, zaten kendileri pointerdir. fonksiyonlar ile direkt değerini değiştirebiliriz.

//biz bir değişkeni pointer ile kullanıyorsak onun referansını kullanıyoruz.
// var a int = 10, 10 bizim referansımız. a da bizim değerimiz.

// referans == adres.
func main() {
	// var a int

	// var p *int

	// a = 10

	// a = 200000000000000000

	// p = &a //ampersant & ramdeki adres değerini almanızı sağlar.

	// *p = 20 //bu aslında a=20
	// //değişken değerini yazdırdık.
	// fmt.Println(a)

	// //ADRES DEĞERİNİ YAZdırdık
	// fmt.Println(p)

	// //adres değerinin içindeki değeri yazdırdık.
	// fmt.Println(*p)

	// //yazılımı sistem üstten alta okur.

	// var a = 10 //20
	// var b int
	// var p *int

	// b = a //b = 10,
	// p = &a //gvfblmkdfomdgomkf
	// *p = 20

	// fmt.Println(a, b)
	// var a int = 10
	// fmt.Println(a) //10

	// add12(a)
	// fmt.Println(a)

	// add12pointer(&a) //adres bilgisini
	// //fonksiyonlarımızın kendi değişkenlerimiz üzerinde modifiye yapmasını istiyorsak pointerları kullanacağız.
	// fmt.Println(a)

	//aray ve slicelerde örnek.
	// var numbers = []int{1, 2, 3, 4, 5, 6}
	// fmt.Println(numbers)

	// changeIndex(numbers)
	// fmt.Println(numbers)

	var stringExample string
	stringExample = "Ben bir denemyim amacım : "
	fmt.Println(stringExample)

	stringChange(&stringExample)
	fmt.Println(stringExample)
}

//biz add12 fonksiyonuna a yı gönderirken anın sadece değerini gönderiyoruz.
// func add12(x int) {
// 	x = x + 12
// 	//yeni x imizin değeri 22
// }

// func add12pointer(x *int) {
// 	*x = *x + 12
// }

// func changeIndex(tasiyiciSlice []int) {
// 	tasiyiciSlice[0] = 1000
// }

//Ödev,
func stringChange(tasiyiciString *string) {
	*tasiyiciString = *tasiyiciString + "Değer tiplimi referans tiplimi öğrenmek."
}

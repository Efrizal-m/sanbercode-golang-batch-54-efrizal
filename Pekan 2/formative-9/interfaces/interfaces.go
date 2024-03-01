package interfaces

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

//---------------------------------------------------//

type PhoneInfo interface {
	Show() string
}

//---------------------------------------------------//
//---------------------------------------------------//

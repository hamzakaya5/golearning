package intry

import "fmt"

type spaceInfos struct {
	constant    float64
	forcePlanet planet
}

type planet struct {
	mass        float64
	distToEarth float64
}

type kutlecekimi interface {
	kutleCekimKuvveti() float64
}

func (s spaceInfos) kutleCekimKuvveti() float64 {
	return s.constant * s.forcePlanet.mass * 150 / s.forcePlanet.distToEarth
}

func cekimhesabi(a kutlecekimi) {
	fmt.Println(a.kutleCekimKuvveti())
}

func KCH() {
	kcForce := spaceInfos{constant: 15, forcePlanet: planet{distToEarth: 1800, mass: 100}}
	cekimhesabi(kcForce)
}

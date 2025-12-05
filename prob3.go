package main

import (
	"fmt"
	"log"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func prob3() {
	N := 10  //num de divisoes
	L := 2.0 //tamanho da barra
	numIteracoes := 10000
	dx := L / float64(N)
	dt := 0.2 * dx * dx
	numGraficos := 30

	temps := make([]float64, N)
	for i := 1; i < N-1; i++ {
		temps[i] = math.Sin(math.Pi / 2 * dx * float64(i))
	}
	temps[0] = 0
	temps[N-1] = 0

	tempsNovas := make([]float64, N)
	copy(tempsNovas, temps)

	fmt.Println(temps)

	for k := 1; k < numIteracoes; k++ {
		for i := 1; i < N-1; i++ {
			tempsNovas[i] = temps[i] + (dt/(dx*dx))*(temps[i+1]-2*temps[i]+temps[i-1])
		}

		copy(temps, tempsNovas)

		pts := make(plotter.XYs, len(temps))
		for i, t := range temps {
			pts[i].X = float64(i) * dx
			pts[i].Y = t
		}

		p := plot.New()

		p.Title.Text = fmt.Sprintf("t = %v", float64(k)*dt)
		p.X.Label.Text = "X (m)"
		p.Y.Label.Text = "T (°C)"

		xPad := 0.03 * L
		p.X.Max = L + xPad
		p.Y.Max = 1.0

		line, err := plotter.NewLine(pts)
		if err != nil {
			log.Fatalf("NewLine: %v", err)
		}

		line.Color = plotutil.Color(0)
		p.Add(line)

		if k%(numIteracoes/numGraficos) == 0 || k < 1000 && k%5 == 0 {
			if err := p.Save(6*vg.Inch, 4*vg.Inch, fmt.Sprintf("prob3_%v.png", k)); err != nil {
				log.Fatalf("Save: %v", err)
			}
		}
	}

}

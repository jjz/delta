package main

/*
 Implement black scholes so we can get the greeks.
*/

import (
	"fmt"
	"github.com/chobie/go-gaussian"
	"math"
	"time"
)

type Option struct {
	K         float64 // strike price
	S0        float64 // strike at time 0
	r         float64 // risk free rate
	sigma     float64 // volatility
	eval_date string  // current time
	exp_date  string  // expiration date
	T         float64 // distance between
	right     string  // 'C' = call, 'P' = put

	price float64
	delta float64
	theta float64
	gamma float64
}

// if you don't know the price, pass -1.0, but you have to know volatility.
// if you don't know volatility,  pass -1.0, but you have to know price.
func NewOption(right string, S0 float64, K float64, eval_date string, exp_date string, r float64, sigma float64, price float64) *Option {
	o := &Option{
		K:         K,
		S0:        S0,
		r:         r,
		eval_date: eval_date,
		exp_date:  exp_date,
		T:         1,
		right:     right,
		sigma:     sigma,
		price:     price,
	}
	o.Initialize()
	return o
}

func calculateT(eval_date string, exp_date string) float64 {
	dtfmt := "20060102"
	evalDt, _ := time.Parse(dtfmt, eval_date)
	expDt, _ := time.Parse(dtfmt, exp_date)
	return (expDt.Sub(evalDt).Hours() / 24) / 365.0
}

func d1(S0 float64, K float64, T float64, r float64, sigma float64) float64 {
	return (math.Log(S0/K) + (r+math.Pow(sigma, 2)/2)*T) / (sigma * math.Sqrt(T))
}

func d2(S0 float64, K float64, T float64, r float64, sigma float64) float64 {
	return (math.Log(S0/K) + (r-math.Pow(sigma, 2)/2)*T) / (sigma * math.Sqrt(T))
}

const PI float64 = 3.14159265359

// calculate Black Scholes price and greeks
func (self *Option) Initialize() {
	norm := gaussian.NewGaussian(0.0, 1.0)

	if self.sigma < 0 {
		self.sigma = self.impliedVol()
	}

	td1 := d1(self.S0, self.K, self.T, self.r, self.sigma)
	td2 := d2(self.S0, self.K, self.T, self.r, self.sigma)

	// we know volatility and want a price, or we're guessing at volatility and we want a price.
	if self.price < 0 {
		if self.right == "C" {
			fmt.Println(td1)
			fmt.Println(td2)
			fmt.Println((self.S0 * math.Exp(-0 * 1)) *norm.Cdf(td1))
			fmt.Println(self.K*math.Exp(-self.r*self.T)*norm.Cdf(td2))
			self.price =(self.S0 * math.Exp(-0 * 1)) *norm.Cdf(td1) - self.K*math.Exp(-self.r*self.T)*norm.Cdf(td2)
		} else if self.right == "P" {
			self.price = self.K*math.Exp(-self.r*self.T)*norm.Cdf(-td2) - self.S0*norm.Cdf(-td1)
		}
	}

	nPrime := math.Pow((2*PI), -(1/2)) * math.Exp(-0.5*math.Pow(td1, 2))

	// handle the rest of the greeks now that we know everything else.
	if self.right == "C" {
		self.delta = norm.Cdf(td1)
		self.gamma = (nPrime / (self.S0 * self.sigma * math.Pow(self.T, (1/2))))
		self.theta = (nPrime)*(-self.S0*self.sigma*0.5/math.Sqrt(self.T)) - self.r*self.K*math.Exp(-self.r*self.T)*norm.Cdf(td2)
	} else if self.right == "P" {
		self.delta = norm.Cdf(td1) - 1
		self.gamma = (nPrime / (self.S0 * self.sigma * math.Pow(self.T, (1/2))))
		self.theta = (nPrime)*(-self.S0*self.sigma*0.5/math.Sqrt(self.T)) + self.r*self.K*math.Exp(-self.r*self.T)*norm.Cdf(-td2)
	}
}

// use newton raphson method to find volatility
func (self *Option) impliedVol() float64 {
	norm := gaussian.NewGaussian(0, 1)
	//v := math.Sqrt(2*PI/self.T) * self.price / self.S0
	v := 0.5

	for i := 0; i < 100; i++ {
		d1 := (math.Log(self.S0/self.K) + (self.r+0.5*math.Pow(v, 2))*self.T) / (v * math.Sqrt(self.T))
		d2 := d1 - v*math.Sqrt(self.T)
		vega := self.S0 * norm.Pdf(d1) * math.Sqrt(self.T)
		cp := 1.0
		if self.right == "P" {
			cp = -1.0
		}
		price0 := cp*self.S0*norm.Cdf(cp*d1) - cp*self.K*math.Exp(-self.r*self.T)*norm.Cdf(cp*d2)
		v = v - (price0-self.price)/vega
		//fmt.Printf("-- next vol %v : %v  / %v \n", i, v, math.Pow(10, -25))
		if math.Abs(price0-self.price) < math.Pow(10, -25) {
			break
		}
	}
	return v
}

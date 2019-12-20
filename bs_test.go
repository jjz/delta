package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCallOption(t *testing.T) {
	S0 := 50.0
	K := 100.0
	right := "C"
	price := -1.0
	vol := 0.25
	r := 0.05 // risk free rate
	eval_date := "20150115"
	exp_date := "20160115"

	opt := NewOption(right, S0, K, eval_date, exp_date, r, vol, price)

	assert.Equal(t, opt.right, "C","option rigth")
	assert.Equal(t, opt.T, 1.0,"option time")
	assert.Equal(t, opt.price, 0.027352493394418753,"option price")
	assert.Equal(t, opt.delta, 0.007190785164003985,"delta")
	assert.Equal(t, opt.theta, -0.32923029818694033,"theta")
	assert.Equal(t, opt.gamma, 0.004001548287317137,"gamma")
	assert.Equal(t, opt.sigma, 0.25,"sigma")

}

func TestPutOption(t *testing.T){
	S0 := 50.0
	K := 60.0
	right := "P"
	price := -1.0
	vol := 0.25
	r := 0.05 // risk free rate
	eval_date := "20150115"
	exp_date := "20160115"

	opt := NewOption(right, S0, K, eval_date, exp_date, r, vol, price)

	assert.Equal(t, opt.right, "P","option rigth")
	assert.Equal(t, opt.T, 1.0,"option time")
	assert.Equal(t, opt.price, 9.586472902449302,"option price")
	assert.Equal(t, opt.delta, -0.656998882557075,"delta")
	assert.Equal(t, opt.theta, -3.637720165611989,"theta")
	assert.Equal(t, opt.gamma, 0.07372212501922741,"gamma")
	assert.Equal(t, opt.sigma, 0.25,"sigma")
}
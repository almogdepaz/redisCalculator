package main

import (
	"strconv"

	"redisCalculator/commons"
)



type Handler interface {
	apply(data commons.Input) string
	applysTo() string
}


type Add struct {

}

func (r Add) apply(data commons.Input) string {
	x, xErr := strconv.ParseInt(data.X, 10, 64)
	y, yErr := strconv.ParseInt(data.Y, 10, 64)

	if(xErr  != nil ||yErr  != nil){
		return "error parsing data"}

	return strconv.FormatInt(x+y, 10)
}

func (r Add) applysTo() string {
	return "+"
}


type subtract struct {

}

func (r subtract) apply(data commons.Input) string {
	x, xErr := strconv.ParseInt(data.X, 10, 64)
	y, yErr := strconv.ParseInt(data.Y, 10, 64)

	if(xErr  != nil ||yErr  != nil){
		return "error parsing data"}

	return strconv.FormatInt(x-y, 10)
}

func (r subtract) applysTo() string {
	return "-"
}


type multiply struct {

}

func (r multiply) apply(data commons.Input) string {
	x, xErr := strconv.ParseInt(data.X, 10, 64)
	y, yErr := strconv.ParseInt(data.Y, 10, 64)

	if(xErr  != nil ||yErr  != nil){
		return "error parsing data"}

	return strconv.FormatInt(x*y, 10)
}

func (r multiply) applysTo() string {
	return "*"
}


type devide struct {

}

func (r devide) apply(data commons.Input) string {

	x, xErr := strconv.ParseFloat(data.X,64)
	y, yErr := strconv.ParseFloat(data.Y,  64)

	if(xErr  != nil ||yErr  != nil){
		return "error parsing data"}

	return strconv.FormatFloat(x/y, 'f', -1, 64)


}

func (r devide) applysTo() string {
	return "/"
}

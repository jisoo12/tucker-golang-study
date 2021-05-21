package main

import (
	"fmt"
	"time"
)

// 택배회사 구조체
type Courier struct {
	Name string
}

// 물품 구조체
type Product struct {
	Name	string
	Price	int
	ID		int
}

// 소포 구조체
type Parcel struct {
	Pdt						*Product
	ShippedTime		time.Time
	DeliveredTime	time.Time
}

func (c *Courier) SendProduct(pdt *Product) *Parcel {
	p := &Parcel{}
	p.ShippedTime = time.Now()
	p.Pdt = pdt
	return p
}

func (p *Parcel) Delivered() *Product {
	p.DeliveredTime = time.Now()
	return p.Pdt
}
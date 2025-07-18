package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second * 3)
	defer cancel()

	bookHotel(ctx)
}

func bookHotel(ctx context.Context){
	select {
		case <- ctx.Done():
			fmt.Println("Hotel booked canceled. Timeout reached.")
		case <- time.After(time.Second * 1	):
		// case <- time.After(time.Second * 5	):
			fmt.Println("Hotel booked")
	}
}
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type abc struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

//var wg = sync.WaitGroup{}

func SimulateContext(ctx context.Context) error {
	// defer wg.Done()
	ticker := time.Tick(1 * time.Second)
	for {

		select {
		case <-ctx.Done():
			fmt.Println("Context Cancelled")
			return ctx.Err()

		case <-ticker:
			val := ctx.Value("Uname")

			if val != nil {
				fmt.Println("Value is :", val)

			} else {
				fmt.Println("No value recived")
			}
		default:
			fmt.Println("Working Fine")
			time.Sleep(1 * time.Second)
		}

	}
}

func Withval(xc context.Context) {
	// time.Sleep(5 * time.Second)
	// fmt.Println("We are here")
	// _, ok := xc.Deadline()
	// fmt.Println("DEADLIN() => ", ok)
	// // defer wg.Done()
	val := xc.Value("Uname")

	// if val != nil {
	// 	fmt.Println("Value is :", val)

	// } else {
	// 	fmt.Println("No value recived")
	// }

	mar, err := json.Marshal(val) // marshal convert the data into streams of bytes
	fmt.Println("Marshales value", mar)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		var name abc
		err := json.Unmarshal(mar, &name) // unmarshal converts the bytes into the desired data used usually when the go application intracts the api's
		fmt.Println(name.Name)
		if err != nil {
			fmt.Println("Error during unmarshal")
		}

	}
}

func conxt() {
	// var t abc
	t := abc{"raju", 1}
	ctx := context.WithValue(context.Background(), "Uname", t)
	//ctx, cnlFn := context.WithDeadline(ctx, time.Now().Add(20*time.Second))
	Withval(ctx)
	// defer cancel()
	// a := time.Now()
	// wg.Add(1)
	//err := SimulateContext(ctx)
	// if err != nil {
	// 	fmt.Printf("ERRR %v", err)
	// 	defer cnlFn()
	// }
	// time.Sleep(5 * time.Second)
	// Withval(ctx)
	// wg.Wait()
	fmt.Println("Cancelling context")
	// b := time.Now()

	// fmt.Println(a)
	// fmt.Println(b)
	// l := b.Sub(a)
	// fmt.Println(l)
	//c := b.Second() - a.Second()

	// time.Sleep(5 * time.Second)

}

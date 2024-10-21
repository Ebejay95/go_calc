/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jonathaneberle <jonathaneberle@student.    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/10/21 21:58:18 by jonathanebe       #+#    #+#             */
/*   Updated: 2024/10/21 22:07:52 by jonathanebe      ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

const (
	D = "\033[0m"
	R = "\033[31m"
	G = "\033[32m"
	Y = "\033[33m"
	B = "\033[34m"
	P = "\033[35m"
	C = "\033[36m"
	W = "\033[37m"
)

func main() {

	var currentVal float64 = 0
	fmt.Println(C + "\n████████████████████████████████████████" + D)
	fmt.Println(C + "██                                    ██" + D)
	fmt.Println(C + "██    Welcome to the Go Calculator    ██" + D)
	fmt.Println(C + "██                                    ██" + D)
	fmt.Println(C + "████████████████████████████████████████" + D)
	
	for {
		var price string
		fmt.Print(currentVal, B + ": " + D)
		fmt.Scan(&price)
		if price == "EXIT" {
			break
		}
	}
}
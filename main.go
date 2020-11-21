package main

import (
	"bufio"
	"fmt"
	"os"
	"stografo/bankAccounts-GoProyect/accounts"
	"strconv"
)

func main() {
	//importar scanner
	scanner := bufio.NewScanner(os.Stdin)

	account := accounts.NewAccount("Alberto")

	//depositar
	println("Cuanto dinero quieres poner?")
	scanner.Scan()
	deposito, _ := strconv.Atoi(scanner.Text())
	account.Deposit(deposito)
	fmt.Println("Tu sldo actual es de ", account.Balance())

	//retirar dinero
	println("Cuanto dinero quieres retirar")
	scanner.Scan()
	retiro, _ := strconv.Atoi(scanner.Text())
	err := account.WithDraw(retiro)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Tu sldo actual es de ", account.Balance())

}

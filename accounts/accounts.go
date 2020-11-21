package accounts

import "errors"

//Account es la estructura que tiene las cuentas
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Tu saldo actual no te permite retirar el monto espesificado")

//NewAccount es el constructor, va a crear las cuentas
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//Deposit es un metodo
func (a *Account) Deposit(amount int) {
	a.balance += amount

}

//Balance nos muestra el balance de la cuenta
func (a Account) Balance() int {
	return a.balance
}

//WithDraw retira fondos de tu cuenta
func (a *Account) WithDraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

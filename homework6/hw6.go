package main

import (
	"fmt"
)

func main() {

	//1
	balance := 100000
	add_sum(&balance, 10000)
	add_sum(&balance, 10000)
	add_sum(&balance, 10000)
	add_sum(&balance, 10000)
	add_sum(&balance, 10000)
	add_sum(&balance, 300000)

	//2
	loan_amount := 3000000
	loan_payment(&loan_amount, 200000)
	loan_payment(&loan_amount, 200000)
	loan_payment(&loan_amount, 200000)
	loan_payment(&loan_amount, 200000)

	//3
	account_balance := 500000
	transaction(&account_balance, 100000)
	transaction(&account_balance, 120000)
	transaction(&account_balance, 20000)

	//4
	investment_amount := 100000 
	accrual(&investment_amount)
	accrual(&investment_amount, 10000)
	accrual(&investment_amount)
	accrual(&investment_amount, 50000)
	accrual(&investment_amount, 50000)

	//6
	investment_amount = 200000
	compound_interest(&investment_amount)
	compound_interest(&investment_amount)
	compound_interest(&investment_amount)
	compound_interest(&investment_amount)
	compound_interest(&investment_amount)
	compound_interest(&investment_amount)

	//5
	course := 74.5
	fmt.Println(convert(&course, 1000))
	update_course(&course, 71)
	fmt.Println(convert(&course, 2000))

	//7
	daily_balance := 5000 
	expenditure := 0
	spend_daily(&daily_balance, &expenditure, 1000)
	spend_daily(&daily_balance, &expenditure, 200)
	spend_daily(&daily_balance, &expenditure, 4000)
	spend_daily(&daily_balance, &expenditure, 300)

	//8
	deposit := 500000 
	annual_accrual(&deposit)
	annual_accrual(&deposit)
	annual_accrual(&deposit)
	annual_accrual(&deposit)
	annual_accrual(&deposit)
	annual_accrual(&deposit)
	annual_accrual(&deposit)
	annual_accrual(&deposit)
	annual_accrual(&deposit)
	annual_accrual(&deposit)
	annual_accrual(&deposit)

	//9
	var transaction_balance float64 = 100000 
	commisioned_trans(&transaction_balance, 10000)
	commisioned_trans(&transaction_balance, 5000)
	commisioned_trans(&transaction_balance, 2000)
	commisioned_trans(&transaction_balance, 11000)

	//10
	start_investment := 300000 
	investment_income(&start_investment)
	investment_income(&start_investment)
	investment_income(&start_investment)
	investment_income(&start_investment)
	investment_income(&start_investment)
	investment_income(&start_investment)
	investment_income(&start_investment)

	//11
	bank_balance := 10000
	add(&bank_balance, 1000)
	add(&bank_balance, 10000)
	add(&bank_balance, 1000)
	add(&bank_balance, 25000)
	add(&bank_balance, 10040)

	//12
	bank_balance = 20000
	spend(&bank_balance, 10000)
	spend(&bank_balance, 5000)
	spend(&bank_balance, 6000)

	//13
	bank_balance = 15000 
	check_balance(&bank_balance)

	//14
	bank_balance = 50000 
	interest2(&bank_balance)
	interest2(&bank_balance)
	interest2(&bank_balance)
	interest2(&bank_balance)
	interest2(&bank_balance)
	interest2(&bank_balance)

	//15
	account1, account2 := 30000, 15000
	transaction_between(&account1, &account2, 5000)
	transaction_between(&account1, &account2, 7000)
	transaction_between(&account1, &account2, 3000)
	transaction_between(&account1, &account2, 8000)
	transaction_between(&account1, &account2, 12000)

	//16
	bank_balance = 40000
	max_withdrawal := 10000
	withdraw_with_limit(&bank_balance, &max_withdrawal, 5000)
	withdraw_with_limit(&bank_balance, &max_withdrawal, 11000)


}

//1
func add_sum(balancePtr *int, added_sum int) {
	if *balancePtr > 500000 {
		fmt.Println("The upper maximum is reached")
	} else {
		*balancePtr += added_sum
		fmt.Println("Current balance is ", *balancePtr)
	}
}

//2
func loan_payment(loanPtr *int, payment int) {
	*loanPtr -= payment
	*loanPtr = int(float64(*loanPtr) * 1.10)
	fmt.Println("Current loan balance is ", *loanPtr)
	if *loanPtr < 500000 {
		fmt.Println("The loan is almost payed out")
	}
}

//3
func transaction(account *int, trans_sum int) {
	if trans_sum > 100000 {
		fmt.Println("Сумма перевода превышает лимит")
	} else {
		*account -= trans_sum
		fmt.Println("Остаток на счете ", *account)
	}
}

//4
func accrual(bank_account *int, add_invest ...int) {
	if *bank_account > 150000 {
		fmt.Println("Достигнут лимит вклада")
	} else {
		*bank_account = int(float64(*bank_account) * 1.05)
		if len(add_invest) > 0 {
			*bank_account += add_invest[0]
		}
		fmt.Println("Вклад равен ", *bank_account)
	}
}

//5
func update_course(old_course *float64, new_course float64) {
	if new_course < 70 {
		fmt.Println("Too low")
	} else {
		*old_course = new_course
	}
}

func convert(course *float64, usd float64) float64{
	return usd * (*course)
}

//6
func compound_interest(bank_accountP *int) {
	if *bank_accountP > 300000 {
		fmt.Println("Достигнут лимит начислений")
	} else {
		*bank_accountP = int(float64(*bank_accountP)*1.05)
		fmt.Println("баланс: ", *bank_accountP)
	}
}

func spend_daily(limitPtr *int, expenditurePtr *int, expense int) {
	if *limitPtr - expense >= 0 {
		*limitPtr -= expense
		*expenditurePtr += expense
		fmt.Println("Сумма расходов: ", *expenditurePtr, "Оставшаяся сумма: ", *limitPtr)
	} else {
		fmt.Println("Превышен дневной лимит")
	}
}

//8
func annual_accrual(depositPtr *int) {
	if *depositPtr > 700000 {
		fmt.Println("Достигнут лимит начислений")
	} else {
		*depositPtr = int(float64(*depositPtr) * 1.06)
		fmt.Println("Баланс по депозиту ", *depositPtr)
	}

}

//9
func commisioned_trans(balancePtr *float64, trans_sum float64) {
	if *balancePtr < 50000  {
		fmt.Println("Баланс ниже допустимого уровня")
	} else {
		*balancePtr = *balancePtr - (trans_sum * 1.01)
		fmt.Println("Current balance is ", *balancePtr)
	}
}

//10
func investment_income(investmentPtr *int) {
	if *investmentPtr > 400000 {
		fmt.Println("Достигнут лимит дохода")
	} else {
		*investmentPtr = int(float64(*investmentPtr) * 1.07)
		fmt.Println("Current balance is ", *investmentPtr)
	}
}

//11
func add(balancePtr *int, added_amount int) {
	*balancePtr += added_amount	
	fmt.Println("Current balance is ", *balancePtr)
}

//12
func spend(bank_accountPtr *int, to_spend int) {
	if *bank_accountPtr - to_spend < 0 {
		fmt.Println("Недостаточно средств")
	} else {
		*bank_accountPtr -= to_spend
	}
}

//13
func check_balance(bank_accountPtr *int) {
	if *bank_accountPtr < 5000 {
		fmt.Println("Баланс низкий")
	} else {
		fmt.Println("Баланс составляет ", *bank_accountPtr)
	}
}

//14
func interest2(bank_accountPtr *int) {
	*bank_accountPtr = int(float64(*bank_accountPtr)*1.02)
	fmt.Println("Current balance is ", *bank_accountPtr)
}

func transaction_between(acc1, acc2 *int, trans_amount int) {
	if *acc1 - trans_amount >= 0 {
		*acc1 -= trans_amount
		*acc2 += trans_amount
		fmt.Println("1 account balance: ", *acc1)
		fmt.Println("2 account balance: ", *acc2)
	} else {
		fmt.Println("The balance on account 1 is insufficient")
	}
}

//16
func withdraw_with_limit(bank_accountPtr, max_withdrawal *int,  withdrawal_amount int) {
	if withdrawal_amount > *max_withdrawal {
		fmt.Println("Превышен лимит снятия")
	} else {
		*bank_accountPtr -= withdrawal_amount
	}
}
package main

import (
	"errors"
	"fmt"
)

const (
	goodCreditScore = 450
	lowScoreRatio   = 10
	goodScoreRatio  = 20
)

var (
	ErrCreditScore = errors.New("invalid credit score")
	ErrIncome      = errors.New("income invalid")
	ErrLoanAmount  = errors.New("loan amount invalid")
	ErrLoanTerm    = errors.New("loan term not multiple of 12")
)

func checkLoan(creditScore int, income float64, loanAmount float64, loanTerm float64) error {
	// Good Credit Score
	interest := 20.0
	if creditScore >= goodCreditScore {
		interest = 15.0
	}

	// Validate score
	if creditScore < 1 {
		return ErrCreditScore
	}

	// Validate Income
	if income < 1 {
		return ErrIncome
	}

	// Validate loanAmount
	if loanAmount < 1 {
		return ErrLoanAmount
	}

	// Validate Term
	if loanTerm < 1 || int(loanTerm)%12 != 0 {
		return ErrLoanTerm
	}

	rate := interest / 100
	payment := ((loanAmount * rate) / loanTerm) + (loanAmount / loanTerm)

	// Total cost of the loan
	totalInterest := (payment * loanTerm) - loanAmount

	// Can they afford the according to the rules?
	approved := false
	if income > payment {
		ratio := (payment / income) * 100
		if creditScore >= goodCreditScore && ratio < goodScoreRatio {
			approved = true
		} else if ratio < lowScoreRatio {
			approved = false
		}
	}

	fmt.Println("Credit Score    :", creditScore)
	fmt.Println("Income          :", income)
	fmt.Println("Loan Amount     :", loanAmount)
	fmt.Println("Loan Term       :", loanTerm)
	fmt.Println("Monthly Payment :", payment)
	fmt.Println("Rate            :", interest)
	fmt.Println("Total Cost      :", totalInterest)
	fmt.Println("Approved        :", approved)
	fmt.Println("")

	return nil
}

func main() {
	// Approved
	fmt.Println("Applicant 1\n-----------")
	err := checkLoan(500, 1000, 1000, 24)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Denied
	fmt.Println("Applicant 2\n-----------")
	err = checkLoan(350, 1000, 10000, 12)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

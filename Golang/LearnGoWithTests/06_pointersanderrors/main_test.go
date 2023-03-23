package main

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	//assertError := func(t testing.TB, err error) {
	//	t.Helper()
	//	if err == nil {
	//		t.Error("wanted an error but didn't get one")
	//	}
	//}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(5))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(15))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(40))
		assertError(t, err, errInsufficientFunds)
		if err == nil {
			t.Errorf("Did not get an error trying to withdraw insufficient funds :o")
		}
		assertBalance(t, wallet, Bitcoin(20))
	})
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

// I would've never thought of this, this is so cool.
func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatalf("didn't get an error but wanted one. Error: %s", got)
	}
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

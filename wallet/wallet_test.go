package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{
			balance: Bitcoin(20),
		}
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw with insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{
			balance: startingBalance,
		}
		// To indicate an error we return an error
		err := wallet.Withdraw(Bitcoin(200))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})

}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Errorf("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()

	if got == nil {
		// Stops the test
		t.Fatal("didn't get an error but wanted one")
	}

	// %q will double quote strings
	if got.Error() != want.Error() {
		t.Errorf("got %q, want %q", got, want)
	}
}
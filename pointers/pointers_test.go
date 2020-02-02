package pointers

import "testing"

func TestWallet(t *testing.T) {
	walletAssertion := func(t *testing.T, got Wallet, want Bitcoin) {
		t.Helper()
		if got.Balance() != want {
			t.Errorf("Wrong balance. Want: %s got %s", want, got.Balance())
		}
	}
	errAssertion := func(t *testing.T, err error, want error) {
		t.Helper()
		if err == nil {
			t.Fatalf("Wanted error, didn't get one")
		}
		if err != want {
			t.Errorf("got: %s want: %s", err.Error(), want.Error())
		}
	}
	noErrAssert := func(t *testing.T, err error) {
		t.Helper()
		if err != nil {
			t.Errorf("Wanted no err, but got %s", err.Error())
		}
	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		walletAssertion(t, wallet, Bitcoin(10))
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 30}
		err := wallet.Withdraw(10)
		noErrAssert(t, err)
		walletAssertion(t, wallet, Bitcoin(20))
	})
	t.Run("Withdraw with insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(100)
		walletAssertion(t, wallet, Bitcoin(20))
		errAssertion(t, err, InsufficentError)
	})
}

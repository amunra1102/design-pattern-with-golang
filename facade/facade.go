package facade

import "fmt"

type WalletFacade struct {
	account      *account
	wallet       *wallet
	securityCode *securityCode
	notification *notification
}

func NewWalletFacade(accountID string, code int) *WalletFacade {
	fmt.Println("Starting create account")
	return &WalletFacade{
		account:      newAccount(accountID),
		wallet:       newWallet(),
		securityCode: newSecurityCode(code),
		notification: &notification{},
	}
}

func (w *WalletFacade) AddMoneyToWallet(accountID string, code int, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(code)
	if err != nil {
		return err
	}

	w.wallet.creditBalance(amount)
	w.notification.sendWalletCreditNotification()
	return nil
}

func (w *WalletFacade) DeductMoneyFromWallet(accountID string, code int, amount int) error {
	fmt.Println("Starting debit money to wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(code)
	if err != nil {
		return err
	}

	err = w.wallet.debitBalance(amount)
	if err != nil {
		return err
	}

	w.notification.sendWalletDebitNotification()
	return nil
}
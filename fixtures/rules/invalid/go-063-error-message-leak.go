package invalid

import "log"

// Bypass attempt: sensitive data leaked via error message wrapping
// ruleid: go.security.no-sensitive-data-in-logs
func processPayment(token string) error {
	err := chargeCard(token)
	if err != nil {
		log.Printf("payment failed for token %s: %v", token, err)
	}
	return err
}

func chargeCard(token string) error { return nil }

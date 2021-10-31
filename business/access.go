package business

func HasAccess(i, p int) bool {
	return i&p == p
}

func GrantAccess(i, p int) int {
	return i | p
}

func RevokeAccess(i, p int) int {
	return i &^ p
}

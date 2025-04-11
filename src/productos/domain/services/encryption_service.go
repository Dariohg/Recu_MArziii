package services

type EncryptionService interface {
	Encrypt(plainText string) (string, error)
	Verify(hashedText, plainText string) bool
}

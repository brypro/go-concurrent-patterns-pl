package main

import "fmt"

/*
strategy pattern
is a behavioral design pattern that lets
you define a family of algorithms,
put each of them into a separate class,
and make their objects interchangeable.

*/

func main() {
	sha := &SHA{}
	md5 := &MD5{}
	passwordProtector := NewPasswordProtector("user", "password", sha)
	passwordProtector.Hash()                // Hashing password for user with SHA algorithm
	passwordProtector.SetHashAlgorithm(md5) // change algorithm
	passwordProtector.Hash()                // Hashing password for user with MD5 algorithm
}

type PasswordProtector struct {
	user          string
	passwordName  string
	hashAlgorithm HashAlgorithm
}
type HashAlgorithm interface {
	Hash(p *PasswordProtector)
}

func NewPasswordProtector(user, passwordName string, hashAlgorithm HashAlgorithm) *PasswordProtector {
	return &PasswordProtector{
		user:          user,
		passwordName:  passwordName,
		hashAlgorithm: hashAlgorithm,
	}
}

func (p *PasswordProtector) SetHashAlgorithm(hashAlgorithm HashAlgorithm) {
	p.hashAlgorithm = hashAlgorithm
}

func (p *PasswordProtector) Hash() {
	p.hashAlgorithm.Hash(p)
}

type SHA struct{}

func (s *SHA) Hash(p *PasswordProtector) {
	// hash
	fmt.Printf("Hashing password %s for user %s with SHA algorithm \n", p.passwordName, p.user)
}

type MD5 struct{}

func (m *MD5) Hash(p *PasswordProtector) {
	// hash
	fmt.Printf("Hashing password %s for user %s with MD5 algorithm \n", p.passwordName, p.user)
}

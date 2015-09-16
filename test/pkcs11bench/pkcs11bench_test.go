package pkcs11bench

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"flag"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/letsencrypt/boulder/Godeps/_workspace/src/github.com/cloudflare/cfssl/crypto/pkcs11key"
	"github.com/letsencrypt/boulder/Godeps/_workspace/src/github.com/miekg/pkcs11"
)

var module = flag.String("module", "", "Path to PKCS11 module")
var tokenLabel = flag.String("tokenLabel", "", "Token label")
var pin = flag.String("pin", "", "PIN")
var privateKeyLabel = flag.String("privateKeyLabel", "", "Private key label")
var slotID = flag.Int("slotID", -1, "Slot")

func initPKCS11Context(modulePath string) (*pkcs11.Ctx, error) {
	context := pkcs11.New(modulePath)

	// Set up a new pkcs11 object and initialize it
	if context == nil {
		return nil, fmt.Errorf("unable to load PKCS#11 module")
	}

	err := context.Initialize()
	return context, err
}

var context *pkcs11.Ctx
func init() {
	flag.Parse()
	var err error
	context, err = initPKCS11Context(*module)
	if err != nil {
		panic("Failed to init PKCS11 module: " + err.Error())
	}
}

// BenchmarkPKCS11 signs a certificate repeatedly using a PKCS11 token and
// measures speed. To run:
// go test -bench=. -benchtime 1m ./test/pkcs11bench/ \
//   -module /usr/lib/softhsm/libsofthsm.so -token-label "softhsm token" \
//   -pin 1234 -private-key-label "my key" -slot-id 7 -v
// You can adjust benchtime if you want to run for longer or shorter.
// TODO: Parallel benchmarking. Currently if you try this with a Yubikey Neo,
// you will get a bunch of CKR_USER_ALREADY_LOGGED_IN errors. This is because
// pkcs11key logs into the token before each signing operation (which is probably a
// performance bug). Also note that some PKCS11 modules (opensc) are not
// threadsafe.
func BenchmarkPKCS11(b *testing.B) {
	if *module == "" || *tokenLabel == "" || *pin == "" || *privateKeyLabel == "" || *slotID == -1 {
		b.Fatal("Must pass all flags: module, tokenLabel, pin, privateKeyLabel, and slotID")
		return
	}

	N := big.NewInt(1)
	N.Lsh(N, 6000)
	// A minimal, bogus certificate to be signed.
	template := x509.Certificate{
		SerialNumber:       big.NewInt(1),
		PublicKeyAlgorithm: x509.RSA,
		NotBefore:          time.Now(),
		NotAfter:           time.Now(),

		PublicKey: &rsa.PublicKey{
			N: N,
			E: 1 << 17,
		},
	}

	// Reset the benchmarking timer so we don't include setup time.
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		p, err := pkcs11key.New(context, *tokenLabel, *pin, *privateKeyLabel, *slotID)
		if err != nil {
			b.Fatal(err)
			return
		}
		defer p.Destroy()

		for pb.Next() {
			_, err = x509.CreateCertificate(rand.Reader, &template, &template, template.PublicKey, p)
			if err != nil {
				b.Fatal(err)
				return
			}
		}
	})
}

// Dummy test to avoid getting "warning: no tests found"
func TestNothing(t *testing.T) {
}

//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE_test --destination=mock_$GOFILE // want "incorrect mockgen options, expected: //go:generate mockgen -source=\\$GOFILE -package=\\$GOPACKAGE --destination=mock_\\$GOFILE"
package c

type c interface {
	C()
}

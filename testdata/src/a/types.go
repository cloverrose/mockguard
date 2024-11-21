//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE --destination=mock_$GOFILE // OK
package a

type a interface {
	A()
}

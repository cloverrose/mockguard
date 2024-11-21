//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE --destination=mock_$GOFILE // want "incorrect mockgen file, expected: types.go"
package b

type b interface {
	B()
}

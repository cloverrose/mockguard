package mockguard

// FileName is configuration what file can contain mockgen comment.
var FileName = "types.go"

// Content is configuration what comment content is valid as mockgen comment.
var Content = "//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE --destination=mock_$GOFILE"

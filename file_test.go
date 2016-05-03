package keyring

import (
	"os"
	"testing"
)

func TestFileKeyringSetWhenEmpty(t *testing.T) {
	k := &fileKeyring{
		Dir: os.TempDir(),
		PasswordFunc: passwordFunc(func(string) (string, error) {
			return "no more secrets", nil
		}),
	}
	item := Item{Key: "llamas", Data: []byte("llamas are great")}

	if err := k.Set(item); err != nil {
		t.Fatal(err)
	}

	foundItem, err := k.Get("llamas")
	if err != nil {
		t.Fatal(err)
	}

	if string(foundItem.Data) != "llamas are great" {
		t.Fatalf("Value stored was not the value retrieved: %q", foundItem.Data)
	}

	if foundItem.Key != "llamas" {
		t.Fatalf("Key wasn't persisted: %q", foundItem.Key)
	}
}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip19"
)

func main() {
	for _, arg := range os.Args[1:] {
		var sk string
		if _, s, err := nip19.Decode(arg); err != nil {
			log.Fatal(err)
		} else {
			sk = s.(string)
		}
		if pub, err := nostr.GetPublicKey(sk); err == nil {
			if pp, err := nip19.EncodePublicKey(pub); err != nil {
				log.Fatal(err)
			} else {
				fmt.Println(pp)
			}
		} else {
			log.Fatal(err)
		}
	}
}

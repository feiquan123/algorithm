package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/feiquan123/algorithm/rsa"
)

var (
	e   = flag.Bool("e", true, "is encode")
	pub = flag.String("pub", "1710503,653797", "public key")
	pri = flag.String("pri", "1710503,1091533", "private key")
	gen = flag.Bool("gen", false, "generate public & private keys")
	f   = flag.String("f", "", "source or encode file")
	s   = flag.String("s", "", "encode or decode string")
)

func parse(s string, pub bool) (n uint64, m uint64, err error) {
	st := "pivate"
	if pub {
		st = "public"
	}
	arr := strings.Split(s, ",")
	if len(arr) != 2 {
		return 0, 0, fmt.Errorf("parse %s key error", st)
	}
	n, err = strconv.ParseUint(arr[0], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	m, err = strconv.ParseUint(arr[1], 10, 64)
	if err != nil {
		return 0, 0, err
	}

	return n, m, nil
}

var r = rsa.New(3000)

func init() {
	flag.Parse()
}

func main() {
	if *gen {
		fmt.Println("public  key:", r.PublicKey())
		fmt.Println("private key:", r.PrivateKey())
		return
	}

	var file = os.Stdin
	var err error

	if len(*f) != 0 {
		file, err = os.Open(*f)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	if *e {
		// encode
		if len(*pub) == 0 {
			fmt.Println("no public key")
			return
		}

		pubN, pubE, err := parse(*pub, *e)
		if err != nil {
			fmt.Println(err)
			return
		}
		pubKey := rsa.NewPublicKey(pubN, pubE)
		var b []byte
		b = []byte(*s)

		if len(b) == 0 {
			b, err = ioutil.ReadAll(file)
			if err != nil {
				fmt.Println("read source file failed", err)
				return
			}
		}

		fmt.Printf("%s", r.Encode(b, pubKey))
	} else {
		// decode
		if len(*pri) == 0 {
			fmt.Println("no private key")
			return
		}

		priN, priB, err := parse(*pri, !*e)
		if err != nil {
			fmt.Println(err)
			return
		}
		priKey := rsa.NewPrivateKey(priN, priB)
		var b []byte
		b = []byte(*s)

		if len(b) == 0 {
			b, err = ioutil.ReadAll(file)
			if err != nil {
				fmt.Println("read encode file failed", err)
				return
			}
		}

		sb, err := r.Decode(b, priKey)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%s", sb)
	}
}

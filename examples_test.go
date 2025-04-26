package sbargp

import (
	"flag"
	"fmt"
	"reflect"
)

type exampleConf struct {
	StoreName    string
	NumItems     int
	NumLocations int
	HasFreeFood  bool
	Department   struct {
		Name      string
		Specialty string
	}
	Vegies []string
	Fruits map[string]struct {
		Yummy bool
		Color string
	}
}

func Example_simple() {
	var c exampleConf
	err := Parse(
		&c,
		[]string{
			// Normally these would be os.Args[1:] but for the example we
			// provide a list of strings.
			"-StoreName", "kings",
			"-NumLocations", "3",
			"-conf", "./bs/example.toml",
		},
		ParserOpts[exampleConf]{
			ProgName:     "example",
			RequiredArgs: []string{"StoreName"},
			ArgDefsSetter: func(conf *exampleConf, fs *flag.FlagSet) error {
				// It is generally recommented (though not enforced) to name
				// CMD line arguments the same as the keys in the TOML file.
				fs.StringVar(&c.StoreName, "StoreName", "", "a string")
				fs.IntVar(&c.NumItems, "NumItems", 1, "a int")
				fs.IntVar(&c.NumItems, "NumLocations", 0, "a int")

				// Defaults for non-FlagSet arguments can be set here as well
				c.HasFreeFood = true // :)
				return nil
			},
		},
	)

	fmt.Println("Parsing error: ", err)

	fmt.Println("Parsed conf:")
	typ, val := reflect.TypeOf(c), reflect.ValueOf(c)
	for i := 0; i < typ.NumField(); i++ {
		fmt.Printf("  %-20s → %v\n", typ.Field(i).Name, val.Field(i).Interface())
	}

	//Output:
	// Parsing error:  <nil>
	// Parsed conf:
	//   StoreName            → kings
	//   NumItems             → 3
	//   NumLocations         → 0
	//   HasFreeFood          → true
	//   Department           → {produce fruit and vegetables}
	//   Vegies               → [potato carrot chicken]
	//   Fruits               → map[apple:{true red} banana:{true yellow} orange:{true orange}]
}

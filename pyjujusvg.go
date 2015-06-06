package main

import (
    "io/ioutil"
    "log"
    "os"
    "strings"

    "gopkg.in/errgo.v1"
    "gopkg.in/juju/charm.v5"

    // Import the jujusvg library and the juju charm library
    "github.com/juju/jujusvg"
)

// iconURL takes a reference to a charm and returns the URL for that charm's icon.
// In this case, we're using the api.jujucharms.com API to provide the icon's URL.
func iconURL(ref *charm.Reference) string {
    if ref.Schema == "local" {
        return "http://marcanonical.com/icon.svg"
    }

    return "https://api.jujucharms.com/v4/" + ref.Path() + "/icon.svg"
}

func generate(bundle_file string) (*jujusvg.Canvas, error) {
    // First, we need to read our bundle data into a []byte
    bundle_data, err := ioutil.ReadFile(bundle_file)
    if err != nil {
        return nil, errgo.Newf("Error reading bundle: %s", err)
    }

    // Next, generate a charm.Bundle from the bytearray by passing it to ReadNewBundleData.
    // This gives us an in-memory object representation of the bundle that we can pass to jujusvg
    bundle, err := charm.ReadBundleData(strings.NewReader(string(bundle_data)))
    if err != nil {
        return nil, errgo.Newf("Error parsing bundle: %s\n", err)
    }


    fetcher := &jujusvg.HTTPFetcher{
        IconURL: iconURL,
    }

    // Next, build a canvas of the bundle.  This is a simplified version of a charm.Bundle
    // that contains just the position information and charm icon URLs necessary to build
    // the SVG representation of the bundle
    canvas, err := jujusvg.NewFromBundle(bundle, iconURL, fetcher)
    if err != nil {
        return nil, errgo.Newf("Error generating canvas: %s\n", err)
    }

    return canvas, nil
}

func main() {
    if len(os.Args) != 2 {
        log.Fatalf("Please provide the path of a bundle as the first argument")
    }
    canvas, err := generate(os.Args[1])

    if err != nil {
        log.Fatalf("FAIL: %s\n", err)
    }

    // Finally, marshal that canvas as SVG to os.Stdout; this will print the
    // SVG data required to generate an image of the bundle.
    canvas.Marshal(os.Stdout)
}

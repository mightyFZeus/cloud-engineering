package main

import (
	"fmt"
	"net/netip"
	"os"
)

type subnet struct {
	name string
	az   string
	cidr string
}

func main() {
	vpc := netip.MustParsePrefix("10.20.0.0/16").Masked()

	subnets := []subnet{
		{name: "public-a", az: "eu-west-1a", cidr: "10.20.0.0/24"},
		{name: "private-a", az: "eu-west-1a", cidr: "10.20.1.0/24"},
		{name: "public-b", az: "eu-west-1b", cidr: "10.20.2.0/24"},
		{name: "private-b", az: "eu-west-1b", cidr: "10.20.3.0/24"},
	}

	var parsed []netip.Prefix
	failed := false

	for _, s := range subnets {
		prefix, err := netip.ParsePrefix(s.cidr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "FAIL: %s has invalid CIDR %q\n", s.name, s.cidr)
			failed = true
			continue
		}

		prefix = prefix.Masked()

		if prefix.Bits() != 24 {
			fmt.Fprintf(os.Stderr, "FAIL: %s must use a /24 CIDR\n", s.name)
			failed = true
		}

		if prefix.Bits() < vpc.Bits() || !vpc.Contains(prefix.Addr()) {
			fmt.Fprintf(
				os.Stderr,
				"FAIL: %s (%s) is outside VPC %s\n",
				s.name,
				prefix,
				vpc,
			)
			failed = true
		}

		for index, previous := range parsed {
			if prefix.Contains(previous.Addr()) || previous.Contains(prefix.Addr()) {
				fmt.Fprintf(
					os.Stderr,
					"FAIL: %s (%s) overlaps %s (%s)\n",
					s.name,
					prefix,
					subnets[index].name,
					previous,
				)
				failed = true
			}
		}

		parsed = append(parsed, prefix)
		fmt.Printf("CHECKED: %s %s %s\n", s.name, s.az, prefix)
	}

	if failed {
		os.Exit(1)
	}

	fmt.Printf("PASS: %d non-overlapping subnets are inside %s\n", len(parsed), vpc)
}

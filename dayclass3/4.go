package main

import (
	"fmt"
)

func main() {
	 x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"]) // 10
	fmt.Println(x) // map[key:10]
	delete(x, "key")
	fmt.Println(x) // map[]	
    elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	fmt.Println(elements["Li"]) // Lithium
	fmt.Println(elements["Un"])
	name,ok := elements["Un"]
	fmt.Println(name,ok)
	z := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B": "Boron",
		"C": "Carbon",
		"N": "Nitrogen",
		"O": "Oxygen",
		"F": "Fluorine",
		"Ne":"Neon",
	}
	fmt.Println(z) // map[B:Boron Be:Beryllium C:Carbon F:Fluorine H:Hydrogen He:Helium Li:Lithium N:Nitrogen Ne:Neon O:Oxygen]
	family := map[string]map[string]string{
		"me": map[string]string{
			"frist_name":"Jhon",
			"last_name": "Deo",
		},
		"father": map[string]string{
			"frist_name":"Jony",
			"last_name": "Saha",
		},
		"mother": map[string]string{
			"frist_name":"Mim",
			"last_name": "Saha",
		},
	}
	if info,ok := family["me"];ok{
		fmt.Println(info["frist_name"],info["last_name"]) // Jhon Deo
	}
}

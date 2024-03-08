package controller

import (
	"encoding/json"
	"example/gymshark/models"
	"fmt"
	"log"
	"net/http"
	"slices"
	"sort"
)

func ComputePackages(w http.ResponseWriter, r *http.Request) {
	var sizes models.Sizes

	// Decodes request body
	decoder := json.NewDecoder(r.Body)
	// Tries to parse the request body into a Sizes object
	err := decoder.Decode(&sizes)

	// If there is an error, log it out
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println("JSON Sizes are: ", sizes.Sizes)
	fmt.Println("JSON Capacity is:", sizes.Capacity)

	// Prepare the ascending and descending slices for package sizes
	var ascendingSizes []int = sizes.Sizes[:]
	sort.Ints(ascendingSizes)

	fmt.Println("Ascending sizes are", ascendingSizes)

	// Prepare answer map. Initialise all to 0
	answer := make(map[int]int)
	for i := 0; i < len(ascendingSizes); i++ {
		answer[ascendingSizes[i]] = 0
	}

	// Computes packages

	var capacityCounter int = sizes.Capacity
	var numPackages int
	var finalIndex int = len(ascendingSizes) - 1

	// Capacity < Smallest Package Size Scenario
	if capacityCounter < ascendingSizes[0] {
		answer[ascendingSizes[0]] = 1
	} else {

		// No Remainder Scenario (e.g. perfect distribution)
		for i := 0; i <= finalIndex; i++ {
			numPackages = capacityCounter / ascendingSizes[finalIndex-i]
			if numPackages > 0 {
				capacityCounter = capacityCounter - ascendingSizes[finalIndex-i]*numPackages
				answer[ascendingSizes[finalIndex-i]] = numPackages
			}
		}

		fmt.Println("Pre-Remainder Package Sizes", answer)
		fmt.Println("Remaining Capacity:", capacityCounter)

		// Deal with remainder
		if capacityCounter > 0 {

			// Calculate waste from sending an extra smallest package
			// The extra will never be more than the size of the smallest package
			var excessCapacity int = ascendingSizes[0] - capacityCounter
			fmt.Println("Wastage from adding an additional smallest size package", excessCapacity)

			// Calculate waste from making an existing package bigger
			var computedPackageSizes []int
			for key, value := range answer {
				if value > 0 {
					computedPackageSizes = append(computedPackageSizes, key)
				}
			}
			var wastageFromEnlargement []int
			sort.Ints(computedPackageSizes)
			fmt.Println("Computed package sizes detected", computedPackageSizes)

			wasteValueToPackSizeIndex := make(map[int]int)

			for _, packSize := range computedPackageSizes {
				if packSize != ascendingSizes[finalIndex] {
					var packSizeIndex int = sort.SearchInts(ascendingSizes, packSize)
					fmt.Println("index ", packSizeIndex)
					var waste int = ascendingSizes[packSizeIndex+1] - ascendingSizes[packSizeIndex] - capacityCounter
					fmt.Println("Respective waste", waste)
					wastageFromEnlargement = append(wastageFromEnlargement, waste)
					wasteValueToPackSizeIndex[waste] = packSizeIndex
				}
			}

			fmt.Println("Waste Map Waste: ", wasteValueToPackSizeIndex)

			var excessCapacityTwo int = slices.Min(wastageFromEnlargement)

			fmt.Println("Least Waste: ", excessCapacityTwo)

			if excessCapacityTwo > excessCapacity {
				answer[ascendingSizes[0]] = answer[ascendingSizes[0]] + 1
			} else {
				answer[ascendingSizes[wasteValueToPackSizeIndex[excessCapacityTwo]]] = answer[ascendingSizes[wasteValueToPackSizeIndex[excessCapacityTwo]]] - 1
				answer[ascendingSizes[wasteValueToPackSizeIndex[excessCapacityTwo]+1]] = answer[ascendingSizes[wasteValueToPackSizeIndex[excessCapacityTwo]+1]] + 1
			}
		}
	}

	fmt.Println("Final Answer:", answer)

	// HTTP stuff
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	jsonString, _ := json.Marshal(answer)
	w.Write([]byte(jsonString))

}

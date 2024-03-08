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

	var orderSpecification models.OrderSpecification

	// Decodes request body
	decoder := json.NewDecoder(r.Body)
	// Attempt to parse the request body into a Sizes object
	err := decoder.Decode(&orderSpecification)
	// If there is an error, log it out
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println("Order specification package sizes:", orderSpecification.Sizes)
	fmt.Println("Order specification desired capacity:", orderSpecification.Capacity)

	// Prepares the sorted slice for package sizes
	var sortedSizes []int = orderSpecification.Sizes[:]
	sort.Ints(sortedSizes)

	fmt.Println("Sorted package sizes:", sortedSizes)

	// Prepares the answer map with the package size as keys. Initialise values to 0
	answer := make(map[int]int)
	for i := 0; i < len(sortedSizes); i++ {
		answer[sortedSizes[i]] = 0
	}

	// Computes number of packages for each package size

	var capacityCounter int = orderSpecification.Capacity
	var finalIndex int = len(sortedSizes) - 1

	// Base case where the desired capacity is smaller than the smallest package size
	if capacityCounter < sortedSizes[0] {
		answer[sortedSizes[0]] = 1
	} else {

		// No remainder scenario (e.g. perfect distribution)
		// Allocates from the biggest package size to the smallest
		for i := 0; i <= finalIndex; i++ {
			numPackages := capacityCounter / sortedSizes[finalIndex-i]
			if numPackages > 0 {
				capacityCounter = capacityCounter - sortedSizes[finalIndex-i]*numPackages
				answer[sortedSizes[finalIndex-i]] = numPackages
			}
		}

		fmt.Println("Interim answer before dealing with the remaining capacity:", answer)
		fmt.Println("Remaining capacity:", capacityCounter)

		// Deal with the remaining capacity
		if capacityCounter > 0 {

			// Calculate waste from sending an extra smallest package
			// The extra will never be more than the size of the smallest package
			excessCapacityFromAddingAPackage := sortedSizes[0] - capacityCounter
			fmt.Println("Excess items from adding an additional smallest size package:", excessCapacityFromAddingAPackage)

			// Calculate waste from making an existing package larger
			var computedPackageSizes []int
			// Identifies package sizes in interim solution
			for key, value := range answer {
				if value > 0 {
					computedPackageSizes = append(computedPackageSizes, key)
				}
			}

			fmt.Println("Interim solution package sizes:", computedPackageSizes)

			// Prepares map to identify which package size to make bigger later on
			wasteValueToPackSizeIndex := make(map[int]int)
			var wastageFromEnlargement []int

			for _, packSize := range computedPackageSizes {
				// Can't make the largest package size any bigger
				if packSize != sortedSizes[finalIndex] {
					var packSizeIndex int = sort.SearchInts(sortedSizes, packSize)
					var waste int = sortedSizes[packSizeIndex+1] - sortedSizes[packSizeIndex] - capacityCounter
					fmt.Println("Excess items from making the", sortedSizes[packSizeIndex], "item package size bigger:", waste)
					wastageFromEnlargement = append(wastageFromEnlargement, waste)
					wasteValueToPackSizeIndex[waste] = packSizeIndex
				}
			}

			excessCapacityFromBiggerPackage := slices.Min(wastageFromEnlargement)
			fmt.Println("Least amount of waste from making a package size bigger: ", excessCapacityFromBiggerPackage)

			if excessCapacityFromBiggerPackage > excessCapacityFromAddingAPackage {
				fmt.Println("More efficient to add an extra package of the smallest size")
				answer[sortedSizes[0]] = answer[sortedSizes[0]] + 1
			} else {
				fmt.Println("More efficient to make an existing package bigger")
				answer[sortedSizes[wasteValueToPackSizeIndex[excessCapacityFromBiggerPackage]]] = answer[sortedSizes[wasteValueToPackSizeIndex[excessCapacityFromBiggerPackage]]] - 1
				answer[sortedSizes[wasteValueToPackSizeIndex[excessCapacityFromBiggerPackage]+1]] = answer[sortedSizes[wasteValueToPackSizeIndex[excessCapacityFromBiggerPackage]+1]] + 1
			}
		}
	}

	fmt.Println("Final Answer:", answer)

	// Sends HTTP response back
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	jsonString, _ := json.Marshal(answer)
	w.Write([]byte(jsonString))

}

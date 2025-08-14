package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPrepTime int) int {
    if avgPrepTime < 1{
        return 2 * len(layers)
    }
    return avgPrepTime * len(layers)
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    amtNoodles, amtSauce := 0, 0.0
    for i := 0;i < len(layers); i++{
        if layers[i] == "sauce"{
            amtSauce += 0.2
        } else if layers[i] == "noodles"{
            amtNoodles += 50
        }
    }
    return amtNoodles, amtSauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
    myList[len(myList) - 1] = friendsList[len(friendsList) - 1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, numPortions int) []float64 {
    var scaledQuantities []float64
    mul := float64(numPortions) / 2
    for i := 0; i < len(quantities); i++ {
        scaledQuantities = append(scaledQuantities, (float64(mul) * quantities[i]))
    }
    return scaledQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.

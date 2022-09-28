package main

// type aliases
type (
	Name          string
	Breed         int
	Gender        int
	NameToDogFunc func(Name) Dog
)

// define possible breeds
const (
	Bulldog Breed = iota
	Havanese
	Cavalier
	Poodle
)

// define possible genders
const (
	Male Gender = iota
	Female
)

var (
	maleHavaneseSpawner = DogSpawner(Havanese, Male)
	femalePoodleSpawner = DogSpawner(Poodle, Female)
)

type Dog struct {
	Name   Name
	Breed  Breed
	Gender Gender
}

/*
DogSpawnerCurry is a function which takes a Breed as input,
and returns a function which takes a gender as input
which in turn returns a function which takes a name as input.

the final output is a Dog
*/
func DogSpawnerCurry(breed Breed) func(Gender) NameToDogFunc {
	return func(gender Gender) NameToDogFunc {
		return func(name Name) Dog {
			return Dog{
				Breed:  breed,
				Gender: gender,
				Name:   name,
			}
		}
	}
}

func DogSpawner(breed Breed, gender Gender) NameToDogFunc {
	return func(n Name) Dog {
		return Dog{
			Breed:  breed,
			Gender: gender,
			Name:   n,
		}
	}
}

func main() {
	bucky := DogSpawnerCurry(Havanese)(Male)("Bucky")

	havaneseSpawner := DogSpawnerCurry(Havanese)
	rocky := havaneseSpawner(Male)("Rocky")

	femaleHavanese := havaneseSpawner(Female)
	lola := femaleHavanese("Lola")
	dotty := femaleHavanese("Dotty")

	rocky := maleHavaneseSpawner("rocky")
	tipsy := femalePoodleSpawner("tipsy")

}

func createDogsWithoutPartialApplication() {
	bucky := Dog{
		Name:   "Bucky",
		Breed:  Havanese,
		Gender: Male,
	}

	rocky := Dog{
		Name:   "Rocky",
		Breed:  Havanese,
		Gender: Male,
	}

	tipsy := Dog{
		Name:   "Tipsy",
		Breed:  Poodle,
		Gender: Female,
	}
	_ = bucky
	_ = rocky
	_ = tipsy
}

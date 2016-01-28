package main

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
    "os"
    "log"
    "net/http"
    "io"
)

type Config struct {
	Code    string `json:"code"`
	Product struct {
		_id                                 string        `json:"_id"`
		_keywords                           []string      `json:"_keywords"`
		AdditivesDebugTags                  []interface{} `json:"additives_debug_tags"`
		AdditivesOldTags                    []interface{} `json:"additives_old_tags"`
		AdditivesPrevTags                   []interface{} `json:"additives_prev_tags"`
		AdditivesTags                       []interface{} `json:"additives_tags"`
		Allergens                           string        `json:"allergens"`
		AllergensHierarchy                  []interface{} `json:"allergens_hierarchy"`
		AllergensTags                       []interface{} `json:"allergens_tags"`
		Brands                              string        `json:"brands"`
		BrandsTags                          []interface{} `json:"brands_tags"`
		Categories                          string        `json:"categories"`
		CategoriesHierarchy                 []string      `json:"categories_hierarchy"`
		CategoriesTags                      []string      `json:"categories_tags"`
		CheckersTags                        []interface{} `json:"checkers_tags"`
		CitiesTags                          []interface{} `json:"cities_tags"`
		Code                                string        `json:"code"`
		CodesTags                           []string      `json:"codes_tags"`
		Complete                            int           `json:"complete"`
		CorrectorsTags                      []string      `json:"correctors_tags"`
		Countries                           string        `json:"countries"`
		CountriesHierarchy                  []string      `json:"countries_hierarchy"`
		CountriesTags                       []string      `json:"countries_tags"`
		CreatedT                            int           `json:"created_t"`
		Creator                             string        `json:"creator"`
		EditorsTags                         []string      `json:"editors_tags"`
		EmbCodes                            string        `json:"emb_codes"`
		EmbCodesOrig                        string        `json:"emb_codes_orig"`
		EmbCodesTags                        []interface{} `json:"emb_codes_tags"`
		EntryDatesTags                      []string      `json:"entry_dates_tags"`
		ExpirationDate                      string        `json:"expiration_date"`
		GenericName                         string        `json:"generic_name"`
		ID                                  string        `json:"id"`
		InformersTags                       []string      `json:"informers_tags"`
		Ingredients                         []interface{} `json:"ingredients"`
		IngredientsDebug                    []interface{} `json:"ingredients_debug"`
		IngredientsFromPalmOilTags          []interface{} `json:"ingredients_from_palm_oil_tags"`
		IngredientsIdsDebug                 []interface{} `json:"ingredients_ids_debug"`
		IngredientsTags                     []interface{} `json:"ingredients_tags"`
		IngredientsText                     string        `json:"ingredients_text"`
		IngredientsTextDebug                string        `json:"ingredients_text_debug"`
		IngredientsTextWithAllergens        string        `json:"ingredients_text_with_allergens"`
		IngredientsThatMayBeFromPalmOilTags []interface{} `json:"ingredients_that_may_be_from_palm_oil_tags"`
		InterfaceVersionCreated             string        `json:"interface_version_created"`
		InterfaceVersionModified            string        `json:"interface_version_modified"`
		Labels                              string        `json:"labels"`
		LabelsHierarchy                     []interface{} `json:"labels_hierarchy"`
		LabelsTags                          []interface{} `json:"labels_tags"`
		Lang                                string        `json:"lang"`
		LastEditDatesTags                   []string      `json:"last_edit_dates_tags"`
		LastModifiedBy                      string        `json:"last_modified_by"`
		LastModifiedT                       int           `json:"last_modified_t"`
		Lc                                  string        `json:"lc"`
		Link                                string        `json:"link"`
		ManufacturingPlaces                 string        `json:"manufacturing_places"`
		ManufacturingPlacesTags             []interface{} `json:"manufacturing_places_tags"`
		NoNutritionData                     interface{}   `json:"no_nutrition_data"`
		NutrientLevels                      struct {
			Fat string `json:"fat"`
		} `json:"nutrient_levels"`
		NutrientLevelsTags []string `json:"nutrient_levels_tags"`
		Nutriments         struct {
			Carbohydrates        string `json:"carbohydrates"`
			Carbohydrates100g    string `json:"carbohydrates_100g"`
			CarbohydratesServing string `json:"carbohydrates_serving"`
			CarbohydratesUnit    string `json:"carbohydrates_unit"`
			CarbohydratesValue   string `json:"carbohydrates_value"`
			Fat                  string `json:"fat"`
			Fat100g              string `json:"fat_100g"`
			FatServing           string `json:"fat_serving"`
			FatUnit              string `json:"fat_unit"`
			FatValue             string `json:"fat_value"`
			Proteins             string `json:"proteins"`
			Proteins100g         string `json:"proteins_100g"`
			ProteinsServing      string `json:"proteins_serving"`
			ProteinsUnit         string `json:"proteins_unit"`
			ProteinsValue        string `json:"proteins_value"`
		} `json:"nutriments"`
		NutritionDataPer     string        `json:"nutrition_data_per"`
		NutritionGradesTags  []string      `json:"nutrition_grades_tags"`
		NutritionScoreDebug  string        `json:"nutrition_score_debug"`
		Origins              string        `json:"origins"`
		OriginsTags          []interface{} `json:"origins_tags"`
		Packaging            string        `json:"packaging"`
		PackagingTags        []interface{} `json:"packaging_tags"`
		PhotographersTags    []interface{} `json:"photographers_tags"`
		PnnsGroups1          string        `json:"pnns_groups_1"`
		PnnsGroups1Tags      []string      `json:"pnns_groups_1_tags"`
		PnnsGroups2          string        `json:"pnns_groups_2"`
		PnnsGroups2Tags      []string      `json:"pnns_groups_2_tags"`
		ProductName          string        `json:"product_name"`
		PurchasePlaces       string        `json:"purchase_places"`
		PurchasePlacesTags   []interface{} `json:"purchase_places_tags"`
		Quantity             string        `json:"quantity"`
		Rev                  int           `json:"rev"`
		ServingQuantity      float64       `json:"serving_quantity"`
		ServingSize          string        `json:"serving_size"`
		Sortkey              int           `json:"sortkey"`
		States               string        `json:"states"`
		StatesHierarchy      []string      `json:"states_hierarchy"`
		StatesTags           []string      `json:"states_tags"`
		Stores               string        `json:"stores"`
		StoresTags           []interface{} `json:"stores_tags"`
		Traces               string        `json:"traces"`
		TracesHierarchy      []interface{} `json:"traces_hierarchy"`
		TracesTags           []interface{} `json:"traces_tags"`
		UnknownNutrientsTags []interface{} `json:"unknown_nutrients_tags"`
	} `json:"product"`
	Status        int    `json:"status"`
	StatusVerbose string `json:"status_verbose"`
}

//http://world.openfoodfacts.org/api/v0/product/748927028669.json

func main(){

	 barcode := os.Args[1]
	 fmt.Println(barcode)

	 url := "http://world.openfoodfacts.org/api/v0/product/"+barcode+".json"
    // don't worry about errors
    response, e := http.Get(url)
    if e != nil {
        log.Fatal(e)
    }

    defer response.Body.Close()

    //open a file for writing
    file, err := os.Create("/Users/MattWilfert/Desktop/"+barcode+".json")
    if err != nil {
        log.Fatal(err)
    }
    // Use io.Copy to just dump the response body to the file. This supports huge files
    _, err = io.Copy(file, response.Body)
    if err != nil {
        log.Fatal(err)
    }
    file.Close()
    fmt.Println("Success!")

    content, err := ioutil.ReadFile("/Users/MattWilfert/Desktop/"+barcode+".json")
    if err!=nil{
        fmt.Print("Error:",err)
    }
    var conf Config
    err=json.Unmarshal(content, &conf)
    if err!=nil{
        fmt.Print("Error:",err)
    }

    fmt.Println(conf.Product.ProductName)
    fmt.Println("Protein: "+conf.Product.Nutriments.Proteins)
    fmt.Println("Carbohydrates: "+conf.Product.Nutriments.Carbohydrates100g)
    fmt.Println("Fats: "+conf.Product.Nutriments.Fat100g)
}


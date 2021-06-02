package main

import (
	"bufio"
	"fmt"
	"math/rand"

	"os"
	"sort"
	"strings"
	"time"
)

type Product struct {
	Name     string
	Category string
	Price    int
}

func randomInt(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

func findProductByName(name string, products []Product) []Product {
	result := []Product{}
	for _, product := range products {
		if product.Name == name {
			result = append(result, product)
		}
	}
	return result
}

func findProductByCategory(category string, products []Product) []Product {
	result := []Product{}
	for _, product := range products {
		if product.Category == category {
			result = append(result, product)
		}
	}
	return result
}

func findProductByPrice(min int, max int, products []Product) []Product {
	result := []Product{}
	for _, product := range products {
		if product.Price >= min && product.Price <= max {
			result = append(result, product)
		}
	}
	return result
}

func main() {
	categories := [4]string{"fashion", "electronics", "sport", "food"}
	products := [20]Product{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(products); i++ {
		products[i] = Product{
			fmt.Sprintf("%s %d", "Product", i),
			categories[rand.Intn(len(categories))],
			randomInt(100, 200),
		}
	}
	fmt.Println("All products")
	for _, product := range products {
		fmt.Println(product)
	}
	productsSlice := products[:]
	sort.Slice(productsSlice, func(i, j int) bool {
		return products[i].Price < products[j].Price
	})
	fmt.Println(productsSlice)
	fmt.Println("Top 5 price")
	for i := 0; i < 5; i++ {
		fmt.Println(productsSlice[i])
	}
	myMap := make(map[string]int)
	for _, product := range products {
		myMap[product.Category]++
	}
	reader := bufio.NewReader(os.Stdin)
	// Input name of Product
	fmt.Println("Enter name of product: ")
	name, _ := reader.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)
	fmt.Println("Searching result: ")
	resultName := findProductByName(name, productsSlice)
	for _, result := range resultName {
		fmt.Printf("%s %s %d", result.Name, result.Category, result.Price)
		fmt.Println()
	}
	// Input name of Category
	fmt.Println("Enter name of category: ")
	c, _ := reader.ReadString('\n')
	c = strings.Replace(c, "\n", "", -1)
	resultCategory := findProductByCategory(c, productsSlice)
	fmt.Printf("Product of %s", c)
	fmt.Println()
	for _, result := range resultCategory {
		fmt.Printf("Name: %s, Price: %d", result.Name, result.Price)
		fmt.Println()
	}
	// Input min and max price values
	fmt.Println("Input min max value: ")
	var min, max int
	_, err := fmt.Scan(&min, &max)
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Println(min, max)
	if min > max {
		tmp := min
		min = max
		max = tmp
	}
	resultPrice := findProductByPrice(min, max, productsSlice)
	for _, result := range resultPrice {
		fmt.Println(result)
	}
}

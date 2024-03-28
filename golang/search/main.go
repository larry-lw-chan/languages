package main

import (
	"fmt"

	"github.com/blevesearch/bleve/v2"
)

func main() {
	// createIndex()
	// addAnotherIndex()
	searchQuery()
	// addDocumentMappingAndCreateIndex()
}

func createIndex() {
	// define some data
	data := struct {
		Id      string
		Content string
		Url     string
	}{
		"example",
		"text is good",
		"http://example.com",
	}

	// open a new index
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("example/bleve", mapping)
	if err != nil {
		panic(err)
	}
	index.Index(data.Id, data)
}

func searchQuery() {
	index, _ := bleve.Open("example/bleve")
	query := bleve.NewQueryStringQuery("text is good")
	searchRequest := bleve.NewSearchRequest(query)
	searchRequest.Fields = []string{"Id", "Content", "Url"} // Add this line
	searchResult, _ := index.Search(searchRequest)

	fmt.Println(searchResult.Hits)
	fmt.Println(searchResult.Hits[0].Fields["Url"])
}

func addAnotherIndex() {
	// define some data
	data := struct {
		Id      string
		Content string
		Url     string
	}{
		"example2",
		"text is good",
		"http://example2.com",
	}

	// open a new index
	// mapping := bleve.NewIndexMapping()
	index, err := bleve.Open("example/bleve")
	if err != nil {
		panic(err)
	}
	index.Index(data.Id, data)
}

func addDocumentMappingAndCreateIndex() {
	// Create a new index mapping
	indexMapping := bleve.NewIndexMapping()

	// Create a new document mapping
	docMapping := bleve.NewDocumentMapping()

	// Add the document mapping to the index mapping
	indexMapping.AddDocumentMapping("doc", docMapping)

	// Create a new index with the mapping
	index, err := bleve.New("example/bleve", indexMapping)
	if err != nil {
		panic(err)
	}

	// Now you can add documents to the index
	data := struct {
		Id      string
		Content string
		Url     string
	}{
		"example",
		"text is good",
		"http://example.com",
	}
	index.Index(data.Id, data)
}

// func addDocument() {
// 	indexMapping := bleve.NewIndexMapping()
// 	blogMapping := bleve.NewDocumentMapping()
// 	indexMapping.AddDocumentMapping("blog", blogMapping)

// 	nameFieldMapping := bleve.NewTextFieldMapping()
// 	nameFieldMapping.Analyzer = "en"
// 	blogMapping.AddFieldMappingsAt("name", nameFieldMapping)
// }

// // search for some text
// query := bleve.NewMatchQuery("text")
// search := bleve.NewSearchRequest(query)
// searchResults, err := index.Search(search)
// if err != nil {
// 	fmt.Println(err)
// 	return
// }

// // searchResults is an iterator
// fmt.Println(searchResults)

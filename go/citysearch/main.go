package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/RediSearch/redisearch-go/redisearch"
)

type CityDetail struct {
	Id          int64
	Name        string
	Country     string
	Latitude    float64
	Longitude   float64
	Population  int64
	CountryCode string
}

/*
FT.CREATE idx:cities ON HASH PREFIX 1 cities:
SCHEMA id NUMERIC SORTABLE name TEXT SORTABLE country TEXT SORTABLE latitude NUMERIC longitude NUMERIC population NUMERIC SORTABLE country_code TEXT
*/

type Service struct {
	redisClient *redisearch.Client
}

func NewService(rc *redisearch.Client) Service {
	return Service{
		redisClient: rc,
	}
}

func main() {
	const fileName string = "worldcities_large.csv"
	hasBeenExported := true

	//alternative --> NewClientFromPool if you already have a connection pool
	client := redisearch.NewClient("127.0.0.1:6379", "idx:cities") // lenovo : 192.168.100.9:6379
	s := NewService(client)

	if !hasBeenExported {
		csvFile, err := os.Open(fileName)
		if err != nil {
			log.Fatalf("ERROR Opening dataset file %s. Detail :%s\n", fileName, err.Error())
		}
		defer csvFile.Close()

		csvLines, err := csv.NewReader(csvFile).ReadAll()
		if err != nil {
			log.Fatalf("ERROR Reading dataset file. Detail :%s\n", err.Error())
		}

		cityDetailList := parseCsv(csvLines)

		for _, cityData := range cityDetailList {
			err = s.IndexNewCity(cityData)
			if err != nil {
				break
			}
		}
	}
	/*
		err := s.IndexNewCity(CityDetail{Name: "testaja", Population: 666, Id: 666, Country: "Retroville", CountryCode: "RTV", Latitude: -4.666, Longitude: 4.666})
		if err != nil {
			log.Println("error waktu indexing ", err.Error())
		}
	*/

	//cities := s.searchByName("test")
	//printCities(cities)

	/*
		cityDetails := s.SearchTropicalCitySortedByPopulation()
		printCities(cityDetails)
	*/

	rawResult := s.GetNumberOfCitiesInEquatorByCountry()
	printAggregateResult(rawResult)
}

func (s *Service) GetNumberOfCitiesInEquatorByCountry() [][]string {
	filterStatement := "@latitude:[-23.5 23.5]"

	q := redisearch.
		NewAggregateQuery().
		SetQuery(redisearch.NewQuery(filterStatement)).
		GroupBy(
			*redisearch.
				NewGroupBy().
				AddFields("@country").
				Reduce(*redisearch.NewReducerAlias(redisearch.GroupByReducerCount, []string{}, "city_count")),
		)

	rawResult, _, err := s.redisClient.Aggregate(q)
	if err != nil {
		log.Println("ada error waktu aggregate cuy : ", err.Error())
	}

	return rawResult
}

func (s *Service) IndexNewCity(city CityDetail) error {
	doc := buildDocument(city)
	err := s.redisClient.Index(doc)
	if err != nil {
		log.Printf("ada error waktu indexing cuy : %s", err.Error())
	}
	return err
}

func (s *Service) SearchTropicalCitySortedByPopulation() []CityDetail {
	q := redisearch.
		NewQuery("@latitude:[-23.5 23.5]").Limit(0, 100).
		SetReturnFields("id", "name", "country", "country_code", "population", "latitude", "longitude")
	docs, _, err := s.redisClient.Search(q)
	if err != nil {
		log.Println("error search cuy ", err.Error())
		return []CityDetail{}
	}
	return buildCityDetails(docs)
}

func (s *Service) searchByName(str string) []CityDetail {
	q := redisearch.
		NewQuery(fmt.Sprintf("@name:*%s*", str)).Limit(0, 100).
		SetReturnFields("id", "name", "country", "country_code", "population", "latitude", "longitude")
	docs, _, err := s.redisClient.Search(q)
	if err != nil {
		log.Println("error search cuy ", err.Error())
		return []CityDetail{}
	}
	return buildCityDetails(docs)
}

func buildKey(cityData CityDetail) string {
	return fmt.Sprintf("cities:%d", cityData.Id)
}

func printCities(cities []CityDetail) {
	for _, city := range cities {
		fmt.Println(city)
	}
}

func printAggregateResult(rawResult [][]string) {
	s := ""
	for _, line := range rawResult {
		s = strings.Join(line, "\t\t")
		fmt.Println(s)
	}
}

func buildDocument(cityData CityDetail) redisearch.Document {
	doc := redisearch.NewDocument(buildKey(cityData), 1.0)
	doc.
		Set("id", cityData.Id).
		Set("name", cityData.Name).
		Set("latitude", cityData.Latitude).
		Set("longitude", cityData.Longitude).
		Set("population", cityData.Population).
		Set("country", cityData.Country).
		Set("country_code", cityData.CountryCode)

	return doc
}

func buildCityDetail(doc redisearch.Document) CityDetail {
	var res CityDetail
	props := doc.Properties

	if val, ok := props["id"].(string); ok {
		valueInt, _ := strconv.ParseInt(val, 10, 64)
		res.Id = valueInt
	}

	if val, ok := props["name"].(string); ok {
		res.Name = val
	}

	if val, ok := props["country"].(string); ok {
		res.Country = val
	}

	if val, ok := props["country_code"].(string); ok {
		res.Country = val
	}

	if val, ok := props["population"].(string); ok {
		valueInt, _ := strconv.ParseInt(val, 10, 64)
		res.Population = valueInt
	}

	if val, ok := props["latitude"].(string); ok {
		valueFloat, _ := strconv.ParseFloat(val, 10)
		res.Latitude = valueFloat
	}

	if val, ok := props["longitude"].(string); ok {
		valueFloat, _ := strconv.ParseFloat(val, 10)
		res.Longitude = valueFloat
	}

	return res
}

func buildCityDetails(docs []redisearch.Document) []CityDetail {
	res := make([]CityDetail, len(docs))
	for idx, doc := range docs {
		res[idx] = buildCityDetail(doc)
	}
	return res
}

func parseCsv(csvLines [][]string) []CityDetail {
	resultList := make([]CityDetail, len(csvLines))
	var latitude, longitude float64
	var population, id int64

	for i, line := range csvLines {
		latitude, _ = strconv.ParseFloat(line[2], 64)
		longitude, _ = strconv.ParseFloat(line[3], 64)
		population, _ = strconv.ParseInt(line[9], 10, 64)
		id, _ = strconv.ParseInt(line[10], 10, 64)
		resultList[i] = CityDetail{
			Id:          id,
			Name:        line[1],
			Country:     line[4],
			Latitude:    latitude,
			Longitude:   longitude,
			Population:  population,
			CountryCode: line[6],
		}
	}

	return resultList
}

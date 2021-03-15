package hw10

import (
	"bufio"
	"io"
	"log"
	"strings"
)

type User struct {
	Email string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	u := getUsers(r, domain)
	return countDomains(u)
}

func getUsers(r io.Reader, domain string) (result []string) {
	scanner := bufio.NewScanner(r)
	result = make([]string, 0)
	var user User
	recvuser := &user
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "."+domain) {
			err := recvuser.UnmarshalJSON(scanner.Bytes())
			if err != nil {
				log.Println(err)
			}
			result = append(result, recvuser.Email)
		}
	}
	return result
}

func countDomains(u []string) (DomainStat, error) {
	result := make(DomainStat)
	for _, ui := range u {
		x := strings.ToLower(strings.Split(ui, "@")[1])
		result[x]++
	}
	return result, nil
}

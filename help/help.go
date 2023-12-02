package help

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const sessionId = ""

func GetInput(day int) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/2023/day/%d/input", day), nil)
	req.Header.Set("Cookie", fmt.Sprintf("session=%s", sessionId))
	res, _ := client.Do(req)

	body, _ := io.ReadAll(res.Body)
	res.Body.Close()
	return strings.TrimRight(string(body), "\n")
}

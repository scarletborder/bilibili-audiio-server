package media_handler

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func getBiliM4s(url string) ([]byte, error) {
	// Create a new HTTP client and set the necessary headers
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Adding the headers
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
	req.Header.Add("Referer", "https://www.bilibili.com")

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Write the response body to the file
	return io.ReadAll(resp.Body)
}

func ProxyHandler(c *fiber.Ctx) error {
	// 先找查询参数
	var url string
	var ok bool
	url = c.Query("url")
	if url == "" {
		// 没有,body里找
		// json 解析
		var body map[string]interface{}
		err := c.BodyParser(&body)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "invalid json",
			})
		}
		url, ok = body["url"].(string)
		if !ok {
			return c.Status(400).JSON(fiber.Map{
				"error": "invalid json",
			})
		}
	} else {
		// 有， base64解码
		tmp, err := base64.StdEncoding.DecodeString(url)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "invalid base64 request",
			})
		}
		url = string(tmp)
	}

	bytes, err := getBiliM4s(url)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	setResponseHeader(c, "")
	return c.Send(bytes)
}

func setResponseHeader(c *fiber.Ctx, name string) {
	if name == "" {
		name = uuid.NewString()
	}

	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Max-Age", "86400")
	c.Set("Content-Type", "audio/mp4")
	c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fmt.Sprintf("%s.m4s", name)))
}

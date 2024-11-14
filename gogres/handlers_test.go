package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TomiwaAribisala-git/gogres/models"
)

func TestGetAllStock(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/api/stock", nil)
	if err != nil {
		t.Fatal(err)
	}

	GetAllStock(rr, r)

	rs := rr.Result()

	if rs.StatusCode != 200 {
		t.Errorf("got %q; want %q", rs.StatusCode, 200)
	}

	defer rs.Body.Close()

	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	bytes.TrimSpace(body)

	var allstock []models.Stock

	err = json.Unmarshal([]byte(body), &allstock)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(allstock)
}

func TestGetStock(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/api/stock/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	GetStock(rr, r)

	rs := rr.Result()

	if rs.StatusCode != 200 {
		t.Errorf("got %q; want %q", rs.StatusCode, 200)
	}

	defer rs.Body.Close()

	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	bytes.TrimSpace(body)

	var stock models.Stock

	err = json.Unmarshal([]byte(body), &stock)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(stock)
}

func TestUpdateStock(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodPut, "/api/stock/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	UpdateStock(rr, r)

	rs := rr.Result()

	if rs.StatusCode != 200 {
		t.Errorf("got %q; want %q", rs.StatusCode, 200)
	}

	defer rs.Body.Close()

	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	bytes.TrimSpace(body)

	var res Response

	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}

func TestDeleteStock(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodDelete, "/api/stock/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	DeleteStock(rr, r)

	rs := rr.Result()

	if rs.StatusCode != 200 {
		t.Errorf("got %q; want %q", rs.StatusCode, 200)
	}

	defer rs.Body.Close()

	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	bytes.TrimSpace(body)

	var res Response

	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}

func TestCreateStock(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodPost, "/api/newstock", nil)
	if err != nil {
		t.Fatal(err)
	}

	CreateStock(rr, r)

	rs := rr.Result()

	if rs.StatusCode != 200 {
		t.Errorf("got %q; want %q", rs.StatusCode, 200)
	}

	defer rs.Body.Close()

	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	bytes.TrimSpace(body)

	var res Response

	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}

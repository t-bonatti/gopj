package parser

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetRfbFiles(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.URL.String(), "/CNPJ")
		data, _ := os.ReadFile("../../../test/data/pagerfb.html")
		rw.Write(data)
	}))
	defer server.Close()

	files := []string{server.URL + "/CNPJ/F.K03200$W.SIMPLES.CSV.D20312.zip",
		server.URL + "/CNPJ/F.K03200$Z.D20312.CNAECSV.zip",
		server.URL + "/CNPJ/F.K03200$Z.D20312.MOTICSV.zip",
		server.URL + "/CNPJ/F.K03200$Z.D20312.MUNICCSV.zip",
		server.URL + "/CNPJ/F.K03200$Z.D20312.NATJUCSV.zip",
		server.URL + "/CNPJ/F.K03200$Z.D20312.PAISCSV.zip",
		server.URL + "/CNPJ/F.K03200$Z.D20312.QUALSCSV.zip",
		server.URL + "/CNPJ/K3241.K03200Y0.D20312.EMPRECSV.zip",
		server.URL + "/CNPJ/K3241.K03200Y0.D20312.ESTABELE.zip",
		server.URL + "/CNPJ/K3241.K03200Y0.D20312.SOCIOCSV.zip",
		server.URL + "/CNPJ/K3241.K03200Y1.D20312.EMPRECSV.zip",
		server.URL + "/CNPJ/K3241.K03200Y1.D20312.ESTABELE.zip",
		server.URL + "/CNPJ/K3241.K03200Y1.D20312.SOCIOCSV.zip",
		server.URL + "/CNPJ/K3241.K03200Y2.D20312.EMPRECSV.zip",
		server.URL + "/CNPJ/K3241.K03200Y2.D20312.ESTABELE.zip",
		server.URL + "/CNPJ/K3241.K03200Y2.D20312.SOCIOCSV.zip",
		server.URL + "/CNPJ/K3241.K03200Y3.D20312.EMPRECSV.zip",
		server.URL + "/CNPJ/K3241.K03200Y3.D20312.ESTABELE.zip",
		server.URL + "/CNPJ/K3241.K03200Y3.D20312.SOCIOCSV.zip",
		server.URL + "/CNPJ/K3241.K03200Y4.D20312.EMPRECSV.zip",
		server.URL + "/CNPJ/K3241.K03200Y4.D20312.ESTABELE.zip",
		server.URL + "/CNPJ/K3241.K03200Y4.D20312.SOCIOCSV.zip",
		server.URL + "/CNPJ/K3241.K03200Y5.D20312.EMPRECSV.zip",
		server.URL + "/CNPJ/K3241.K03200Y5.D20312.ESTABELE.zip",
		server.URL + "/CNPJ/K3241.K03200Y5.D20312.SOCIOCSV.zip",
		server.URL + "/CNPJ/K3241.K03200Y6.D20312.EMPRECSV.zip",
		server.URL + "/CNPJ/K3241.K03200Y6.D20312.ESTABELE.zip",
		server.URL + "/CNPJ/K3241.K03200Y6.D20312.SOCIOCSV.zip",
		server.URL + "/CNPJ/K3241.K03200Y7.D20312.EMPRECSV.zip",
		server.URL + "/CNPJ/K3241.K03200Y7.D20312.ESTABELE.zip",
		server.URL + "/CNPJ/K3241.K03200Y7.D20312.SOCIOCSV.zip",
		server.URL + "/CNPJ/K3241.K03200Y8.D20312.EMPRECSV.zip",
		server.URL + "/CNPJ/K3241.K03200Y8.D20312.ESTABELE.zip",
		server.URL + "/CNPJ/K3241.K03200Y8.D20312.SOCIOCSV.zip",
		server.URL + "/CNPJ/K3241.K03200Y9.D20312.EMPRECSV.zip",
		server.URL + "/CNPJ/K3241.K03200Y9.D20312.ESTABELE.zip",
		server.URL + "/CNPJ/K3241.K03200Y9.D20312.SOCIOCSV.zip"}

	pageParser := RfbPage{server.Client(), server.URL}
	rfbFiles, err := pageParser.GetRfbFiles()

	assert.Nil(t, err)
	assert.Equal(t, time.Date(2022, time.April, 1, 0, 0, 0, 0, time.UTC), rfbFiles.UpdatedAt)
	assert.ElementsMatch(t, files, rfbFiles.FilesUrl)
}

func TestGetRfbFilesWhenIsOutages(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.URL.String(), "/CNPJ")
		rw.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	pageParser := RfbPage{server.Client(), server.URL}
	_, err := pageParser.GetRfbFiles()

	assert.Equal(t, fmt.Sprint(err), "status code error: 404 Not Found")
}

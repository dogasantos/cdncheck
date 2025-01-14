package cdncheck

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

var cidrRegex = regexp.MustCompile(`[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\/[0-9]{1,3}`)

type scraperFunc func(httpClient *http.Client) ([]string, error)
type scraperWithOptionsFunc func(httpClient *http.Client, options *Options) ([]string, error)

// scrapeAzure scrapes Microsoft Azure firewall's CIDR ranges from their datacenter
func scrapeAzure(httpClient *http.Client) ([]string, error) {
	resp, err := httpClient.Get("https://download.microsoft.com/download/0/1/8/018E208D-54F8-44CD-AA26-CD7BC9524A8C/PublicIPs_20200824.xml")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	body := string(data)

	cidrs := cidrRegex.FindAllString(body, -1)
	return cidrs, nil
}

// scrapeCloudFront scrapes CloudFront firewall's CIDR ranges from their API
func scrapeCloudFront(httpClient *http.Client) ([]string, error) {
	resp, err := httpClient.Get("https://d7uri8nf7uskq.cloudfront.net/tools/list-cloudfront-ips")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	body := string(data)

	cidrs := cidrRegex.FindAllString(body, -1)
	return cidrs, nil
}

// scrapeCloudflare scrapes cloudflare firewall's CIDR ranges from their API
func scrapeCloudflare(httpClient *http.Client) ([]string, error) {
	resp, err := httpClient.Get("https://www.cloudflare.com/ips-v4")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	body := string(data)

	cidrs := cidrRegex.FindAllString(body, -1)
	return cidrs, nil
}

// scrapeIncapsula scrapes incapsula firewall's CIDR ranges from their API
func scrapeIncapsula(httpClient *http.Client) ([]string, error) {
	req, err := http.NewRequest(http.MethodPost, "https://my.incapsula.com/api/integration/v1/ips", strings.NewReader("resp_format=text"))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	body := string(data)

	cidrs := cidrRegex.FindAllString(body, -1)
	return cidrs, nil
}

// scrapeAkamai scrapes akamai firewall's CIDR ranges from ipinfo
func scrapeAkamai(httpClient *http.Client, options *Options) ([]string, error) {
	req, err := makeReqWithAuth(http.MethodGet, "https://ipinfo.io/AS12222", "Authorization", "Bearer "+options.IPInfoToken)
	if err != nil {
		return nil, err
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	body := string(data)

	cidrs := cidrRegex.FindAllString(body, -1)
	return cidrs, nil
}

// scrapeSucuri scrapes sucuri firewall's CIDR ranges from ipinfo
func scrapeSucuri(httpClient *http.Client, options *Options) ([]string, error) {
	req, err := makeReqWithAuth(http.MethodGet, "https://ipinfo.io/AS30148", "Authorization", "Bearer "+options.IPInfoToken)
	if err != nil {
		return nil, err
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	body := string(data)

	cidrs := cidrRegex.FindAllString(body, -1)
	return cidrs, nil
}

func scrapeFastly(httpClient *http.Client) ([]string, error) {
	resp, err := httpClient.Get("https://api.fastly.com/public-ip-list")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	body := string(data)

	cidrs := cidrRegex.FindAllString(body, -1)
	return cidrs, nil
}

// scrapeLeaseweb scrapes leaseweb firewall's CIDR ranges from ipinfo
func scrapeLeaseweb(httpClient *http.Client, options *Options) ([]string, error) {
	req, err := makeReqWithAuth(http.MethodGet, "https://ipinfo.io/AS60626", "Authorization", "Bearer "+options.IPInfoToken)
	if err != nil {
		return nil, err
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	body := string(data)

	cidrs := cidrRegex.FindAllString(body, -1)
	return cidrs, nil
}

func scrapeProjectDiscovery(httpClient *http.Client) (map[string][]string, error) {
	resp, err := httpClient.Get("https://cdn.nuclei.sh")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data map[string][]string
	if err := jsoniter.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func makeReqWithAuth(method, URL, headerName, bearerValue string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(headerName, "Bearer "+bearerValue)
	return req, nil
}

// scrapeAzion scrapes Azion firewall's CIDR ranges from ipinfo
func scrapeAzion(httpClient *http.Client, options *Options) ([]string, error) {
	req, err := makeReqWithAuth(http.MethodGet, "https://ipinfo.io/AS52580", "Authorization", "Bearer "+options.IPInfoToken)
	if err != nil {
		return nil, err
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	body := string(data)

	cidrs := cidrRegex.FindAllString(body, -1)
	return cidrs, nil
}



// scrapeStackPath scrapes stackpath firewall's CIDR ranges from ipinfo
func scrapeStackPath(httpClient *http.Client, options *Options) ([]string, error) {
	cidrs := []string{}

	ASNlist :=  [5]string{"AS33438", "AS20446", "AS199156", "AS18607", "AS12989"}

	for _, asn := range ASNlist {
		ipinfourl := fmt.Sprintf("https://ipinfo.io/%s", asn)

		req, err := makeReqWithAuth(http.MethodGet, ipinfourl, "Authorization", "Bearer "+options.IPInfoToken)

		if err != nil {
			return nil, err
		}
		resp, err := httpClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		body := string(data)

		for _, cidr := range cidrRegex.FindAllString(body, -1) {
			cidrs = append(cidrs,cidr)
		}

	
	}

	return cidrs, nil
}


// scrapeEdgeCast scrapes Edgecast firewall's CIDR ranges from ipinfo
func scrapeEdgeCast(httpClient *http.Client, options *Options) ([]string, error) {
	cidrs := []string{}

	ASNlist :=  [3]string{"AS15133", "AS14210", "AS14153"}

	for _, asn := range ASNlist {
		ipinfourl := fmt.Sprintf("https://ipinfo.io/%s", asn)

		req, err := makeReqWithAuth(http.MethodGet, ipinfourl, "Authorization", "Bearer "+options.IPInfoToken)
		if err != nil {
			return nil, err
		}
		resp, err := httpClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		body := string(data)

		for _, cidr := range cidrRegex.FindAllString(body, -1) {
			cidrs = append(cidrs,cidr)
		}
	
	}
	
	return cidrs, nil
}

// scrapeLimeLight scrapes LimeLight firewall's CIDR ranges from ipinfo
func scrapeLimeLight(httpClient *http.Client, options *Options) ([]string, error) {
	cidrs := []string{}

	ASNlist :=  [9]string{"AS60261","AS55429","AS45396", "AS38622", "AS37277", "AS26506", "AS25804", "AS23059", "AS22822" }
	
	for _, asn := range ASNlist {
		ipinfourl := fmt.Sprintf("https://ipinfo.io/%s", asn)

		req, err := makeReqWithAuth(http.MethodGet, ipinfourl, "Authorization", "Bearer "+options.IPInfoToken)
		if err != nil {
			return nil, err
		}
		resp, err := httpClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		body := string(data)

		for _, cidr := range cidrRegex.FindAllString(body, -1) {
			cidrs = append(cidrs,cidr)
		}
	}
	
	return cidrs, nil
}


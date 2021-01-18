package soap

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antchfx/xmlquery"
)

// GetStringFromXML ...
func GetStringFromXML(node *xmlquery.Node, selector string) string {
	childNode := xmlquery.FindOne(node, selector)
	if childNode != nil {
		return childNode.InnerText()
	}
	return ""
}

// GetIntFromXML doc ..
func GetIntFromXML(node *xmlquery.Node, selector string) (int, error) {
	childNode := xmlquery.FindOne(node, selector)
	if childNode != nil {
		value, err := strconv.Atoi(childNode.InnerText())
		if err != nil {
			return 0, fmt.Errorf("Couldn't parse value from selector: %s. [%+v]", selector, childNode)
		}
		return value, nil
	}
	return 0, nil
}

// GetIntFromXML doc ..
func getFloatFromXML(node *xmlquery.Node, selector string) (float64, error) {
	childNode := xmlquery.FindOne(node, selector)
	if childNode != nil {
		value, err := strconv.ParseFloat(childNode.InnerText(), 32)
		if err != nil {
			return 0, fmt.Errorf("Couldn't parse amount from selector: %s. [%+v]", selector, childNode)
		}
		return value, nil
	}
	return 0, nil
}

// CleanXML doc ..
func CleanXML(request string) string {
	request = strings.Replace(request, "<![CDATA[", "", -1)
	request = strings.Replace(request, "]]>", "", -1)
	return request
}

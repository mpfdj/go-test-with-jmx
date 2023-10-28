package tests

import (
	"context"
	"fmt"
	"github.com/newrelic/nrjmx/gojmx"
	"testing"
)

func TestWithJmx(t *testing.T) {

	config := &gojmx.JMXConfig{
		Hostname:         "localhost",
		Port:             5000,
		RequestTimeoutMs: 10000,
	}

	// Connect to JMX endpoint.
	client, err := gojmx.NewClient(context.Background()).Open(config)
	handleError(err)

	// Get the mBean names.
	mBeanNames, err := client.QueryMBeanNames("jaeger.de.miel.mbeans:name=*")

	fmt.Printf("%#v\n", mBeanNames)

	//handleError(err)
	//
	//// Get the Attribute names for each mBeanName.
	//for _, mBeanName := range mBeanNames {
	//	mBeanAttrNames, err := client.GetMBeanAttributeNames(mBeanName)
	//	handleError(err)
	//
	//	// Get the attribute value for each mBeanName and mBeanAttributeName.
	//	jmxAttrs, err := client.GetMBeanAttributes(mBeanName, mBeanAttrNames...)
	//	if err != nil {
	//		fmt.Println(err)
	//		continue
	//	}
	//	for _, attr := range jmxAttrs {
	//		if attr.ResponseType == gojmx.ResponseTypeErr {
	//			fmt.Println(attr.StatusMsg)
	//			continue
	//		}
	//		fmt.Printf("%#v\n", attr)
	//	}
	//}

}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

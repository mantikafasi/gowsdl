// Code generated by gowsdl DO NOT EDIT.

package gen

import (
	"github.com/mantikafasi/goxml"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

var wsdl = `<?xml version="1.0" encoding="utf-8"?>
<wsdl:definitions xmlns:s="http://www.w3.org/2001/XMLSchema"
                  xmlns:tns="http://www.mnb.hu/webservices/"
                  xmlns:soap="http://schemas.xmlsoap.org/wsdl/soap/"
                  xmlns:http="http://schemas.xmlsoap.org/wsdl/http/"
                  targetNamespace="http://www.mnb.hu/webservices/"
                  xmlns:wsdl="http://schemas.xmlsoap.org/wsdl/">
  <wsdl:documentation xmlns:wsdl="http://schemas.xmlsoap.org/wsdl/">MNB curreny exchange rate webservice.</wsdl:documentation>
  <wsdl:types>
    <s:schema elementFormDefault="qualified" targetNamespace="http://www.mnb.hu/webservices/">
      <s:element name="GetInfo">
        <s:complexType>
          <s:sequence>
            <s:element name="Id">
              <s:annotation>
                <s:documentation>comment</s:documentation>
              </s:annotation>
              <s:simpleType>
                <s:restriction base="s:string">
                  <s:minLength value="2"/>
                </s:restriction>
              </s:simpleType>
            </s:element>
          </s:sequence>
        </s:complexType>
      </s:element>
      <s:element name="GetInfoResponse">
        <s:complexType>
          <s:sequence>
            <s:element minOccurs="0" maxOccurs="1" name="GetInfoResult" type="s:string">
                <s:annotation>
                    <s:documentation>this is a comment</s:documentation>
                </s:annotation>
            </s:element>
          </s:sequence>
        </s:complexType>
      </s:element>
      <s:complexType name="ResponseStatus">
        <s:sequence>
          <s:element name="status" minOccurs="0" maxOccurs="unbounded">
            <s:complexType>
              <s:simpleContent>
                <s:extension base="s:string">
                  <s:attribute name="code" use="required">
                    <s:simpleType>
                      <s:restriction base="s:string">
                        <s:enumeration value="UnrecognizedTrimName" />
                        <s:enumeration value="UnusedTrimName" />
                      </s:restriction>
                    </s:simpleType>
                  </s:attribute>
                </s:extension>
              </s:simpleContent>
            </s:complexType>
          </s:element>
        </s:sequence>
        <s:attribute ref="tns:responseCode"/>
      </s:complexType>
      <s:attribute name="responseCode">
        <s:simpleType>
          <s:restriction base="s:string">
            <s:enumeration value="Successful" />
            <s:enumeration value="Unsuccessful" />
            <s:enumeration value="ConditionallySuccessful" />
          </s:restriction>
        </s:simpleType>
      </s:attribute>
    </s:schema>
  </wsdl:types>
  <wsdl:message name="GetInfoSoapIn">
    <wsdl:part name="parameters" element="tns:GetInfo" />
  </wsdl:message>
  <wsdl:message name="GetInfoSoapOut">
    <wsdl:part name="parameters" element="tns:GetInfoResponse" />
  </wsdl:message>
  <wsdl:portType name="MNBArfolyamServiceType">
    <wsdl:operation name="GetInfoSoap">
      <wsdl:input message="tns:GetInfoSoapIn"/>
      <wsdl:output message="tns:GetInfoSoapOut"/>
    </wsdl:operation>
  </wsdl:portType>
  <wsdl:binding name="MNBArfolyamBinding" type="tns:MNBArfolyamServiceType">
    <soap:binding style="document" transport="http://schemas.xmlsoap.org/soap/http" />
    <wsdl:operation name="GetInfoSoap">
      <wsdl:input>
        <soap:body use="literal" />
      </wsdl:input>
      <wsdl:output>
        <soap:body use="literal" />
      </wsdl:output>
    </wsdl:operation>
  </wsdl:binding>
  <wsdl:service name="MNBArfolyamService">
    <wsdl:documentation xmlns:wsdl="http://schemas.xmlsoap.org/wsdl/">MNB curreny exchange rate webservice.</wsdl:documentation>
    <wsdl:port name="MNBArfolyamServiceSoap" binding="tns:MNBArfolyamBinding">
      <soap:address location="http://example.org/" />
    </wsdl:port>
  </wsdl:service>
</wsdl:definitions>
`

var WSDLUndefinedError = errors.New("Server was unable to process request. --> Object reference not set to an instance of an object.")

type SOAPEnvelopeRequest struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    SOAPBodyRequest
}

type SOAPBodyRequest struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	GetInfo *GetInfo `xml:,omitempty`
}

type SOAPEnvelopeResponse struct {
	XMLName    xml.Name `xml:"soap:Envelope"`
	PrefixSoap string   `xml:"xmlns:soap,attr"`
	PrefixXsi  string   `xml:"xmlns:xsi,attr"`
	PrefixXsd  string   `xml:"xmlns:xsd,attr"`

	Body SOAPBodyResponse
}

func NewSOAPEnvelopResponse() *SOAPEnvelopeResponse {
	return &SOAPEnvelopeResponse{
		PrefixSoap: "http://schemas.xmlsoap.org/soap/envelope/",
		PrefixXsd:  "http://www.w3.org/2001/XMLSchema",
		PrefixXsi:  "http://www.w3.org/2001/XMLSchema-instance",
	}
}

type Fault struct {
	XMLName xml.Name `xml:"SOAP-ENV:Fault"`
	Space   string   `xml:"xmlns:SOAP-ENV,omitempty,attr"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

type SOAPBodyResponse struct {
	XMLName xml.Name `xml:"soap:Body"`
	Fault   *Fault   `xml:",omitempty"`

	GetInfo *GetInfoResponse `xml:",omitempty"`
}

func (service *SOAPBodyRequest) GetInfoFunc(request *GetInfo) (*GetInfoResponse, error) {
	return &GetInfoResponse{
		GetInfoResult: "gowsdl, " + request.Id,
	}, nil
}

func (service *SOAPEnvelopeRequest) call(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/xml; charset=utf-8")
	val := reflect.ValueOf(&service.Body).Elem()
	n := val.NumField()
	var field reflect.Value
	var name string
	find := false

	if r.Method == http.MethodGet {
		w.Write([]byte(wsdl))
		return
	}

	resp := NewSOAPEnvelopResponse()
	defer func() {
		if r := recover(); r != nil {
			resp.Body.Fault = &Fault{}
			resp.Body.Fault.Space = "http://schemas.xmlsoap.org/soap/envelope/"
			resp.Body.Fault.Code = "soap:Server"
			resp.Body.Fault.Detail = fmt.Sprintf("%v", r)
			resp.Body.Fault.String = fmt.Sprintf("%v", r)
		}
		xml.NewEncoder(w).Encode(resp)
	}()

	header := r.Header.Get("Content-Type")
	if strings.Index(header, "application/soap+xml") >= 0 {
		panic("Could not find an appropriate Transport Binding to invoke.")
	}

	err := xml.NewDecoder(r.Body).Decode(service)
	if err != nil {
		panic(err)
	}

	for i := 0; i < n; i++ {
		field = val.Field(i)
		name = val.Type().Field(i).Name
		if field.Kind() != reflect.Ptr {
			continue
		}
		if field.IsNil() {
			continue
		}
		if field.IsValid() {
			find = true
			break
		}
	}

	if !find {
		panic(WSDLUndefinedError)
	} else {
		m := val.Addr().MethodByName(name + "Func")
		if !m.IsValid() {
			panic(WSDLUndefinedError)
		}

		vals := m.Call([]reflect.Value{field})
		if vals[1].IsNil() {
			reflect.ValueOf(&resp.Body).Elem().FieldByName(name).Set(vals[0])
		} else {
			panic(vals[1].Interface())
		}
	}

}

func Endpoint(w http.ResponseWriter, r *http.Request) {
	request := SOAPEnvelopeRequest{}
	request.call(w, r)
}

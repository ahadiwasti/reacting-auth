package controllers

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"fmt"

	// "github.com/ahadiwasti/reacting-auth/pkg/api/dto"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/gin-gonic/gin"
)

type ERPController struct {
	BaseController
}

// common
type Envelope struct {
	XMLName xml.Name    `xml:"soapenv:Envelope"`
	Soapenv string      `xml:"xmlns:soapenv,attr"`
	OGl     string      `xml:"xmlns:ogl,attr"`
	Header  xml.Name    `xml:"soapenv:Header"`
	Body    interface{} `xml:"soapenv:Body"`
}

// method 1
type GetPurchaseOrders struct {
	Body string `xml:"ogl:GetPurchaseOrders"`
}

// method 1
func NewEnvelope() *Envelope {
	envelope := &Envelope{
		Soapenv: "http://schemas.xmlsoap.org/soap/envelope/",
		OGl:     "http://172.27.54.62:80/OglobaWebService/",
	}
	return envelope
}

// method 1
func NewRequestEnvelope(body interface{}) *Envelope {
	envelope := NewEnvelope()
	envelope.Body = body
	return envelope
}

// method 1
type EnvelopePO struct {
	XMLName xml.Name    `xml:"Envelope" json:"-"`
	Soap    string      `xml:"xmlns:soap,attr,omitempty" json:"-"`
	SoapEnv string      `xml:"xmlns:soapenv,attr,omitempty" json:"-"`
	OGL     string      `xml:"xmlns:og1,attr,omitempty" json:"-"`
	XSI     string      `xml:"xmlns:xsi,attr,omitempty" json:"-"`
	XSD     string      `xml:"xmlns:xsd,attr,omitempty" json:"-"`
	Body    interface{} `xml:"Body" json:"body"`
}

// method 1
type GetPurchaseOrdersResponse struct {
	Response struct {
		Result struct {
			PurchaseOrderData purchaseOrderData `xml:"PurchaseOrderData" json:"PurchaseOrderData"`
		} `xml:"GetPurchaseOrdersResult" json:"GetPurchaseOrdersResult"`
	} `xml:"GetPurchaseOrdersResponse" json:"GetPurchaseOrdersResponse"`
}

// method 1
type purchaseOrderData struct {
	BporSeq                 string `xml:"BporSeq" json:"bpor_seq"`
	BporOrgID               string `xml:"BporOrgId" json:"bpor_org_id"`
	BporOrganizationID      string `xml:"BporOrganizationId" json:"bpor_organization_id"`
	BporSubinventoryCode    string `xml:"BporSubinventoryCode" json:"bpor_subinventory_code"`
	BporPurchaseOrderNo     string `xml:"BporPurchaseOrderNo" json:"bpor_purchase_order_no"`
	BporSupplierCode        string `xml:"BporSupplierCode" json:"bpor_supplier_code"`
	BporOrderDate           string `xml:"BporOrderDate" json:"bpor_order_date"`
	BporCreationDate        string `xml:"BporCreationDate" json:"bpor_creation_date"`
	BporCreatedUser         string `xml:"BporCreatedUser" json:"bpor_created_user"`
	BporLastUpdatedDate     string `xml:"BporLastUpdatedDate,omitempty" json:"bpor_last_updated_date,omitempty"`
	BporUpdatedUser         string `xml:"BporUpdatedUser,omitempty" json:"bpor_updated_user,omitempty"`
	BporReadFlag            string `xml:"BporReadFlag" json:"bpor_read_flag"`
	BporCurrency            string `xml:"BporCurrency" json:"bpor_currency"`
	Attribute1              string `xml:"Attribute1" json:"attribute_1"`
	PurchaseOrderDetailColl struct {
		PurchaseOrderDetl []purchaseOrderDetl `xml:"PurchaseOrderDetl" json:"purchase_order_detl"`
	} `xml:"PurchaseOrderDetailColl" json:"purchase_order_detail_coll"`
}

// method 1
type purchaseOrderDetl struct {
	BporSeq              string `xml:"BporSeq" json:"bpor_seq"`
	BporOrgID            string `xml:"BporOrgId" json:"bpor_org_id"`
	BporOrganizationID   string `xml:"BporOrganizationId" json:"bpor_organization_id"`
	BporSubinventoryCode string `xml:"BporSubinventoryCode" json:"bpor_subinventory_code"`
	BporPurchaseOrderNo  string `xml:"BporPurchaseOrderNo" json:"bpor_purchase_order_no"`
	BpodSrlNo            string `xml:"BpodSrlNo" json:"bpod_srl_no"`
	BpodItemCode         string `xml:"BpodItemCode" json:"bpod_item_code"`
	BpodOrderQty         string `xml:"BpodOrderQty" json:"bpod_order_qty"`
	BpodCreationDate     string `xml:"BpodCreationDate" json:"bpod_creation_date"`
	BpodCreatedUser      string `xml:"BpodCreatedUser" json:"bpod_created_user"`
	BpodLastUpdatedDate  string `xml:"BpodLastUpdatedDate,omitempty" json:"bpod_last_updated_date"`
	BpodUpdatedUser      string `xml:"BpodUpdatedUser,omitempty" json:"bpod_updated_user"`
	BpodReadFlag         string `xml:"BpodReadFlag" json:"bpod_read_flag"`
}

// 2 method for update
type UpdatePurchaseOrderReceivedBody struct {
	XMLName                     xml.Name `xml:"soapenv:Body"`
	UpdatePurchaseOrderReceived UpdatePurchaseOrderReceived
}
type UpdatePurchaseOrderReceived struct {
	XMLName xml.Name `xml:"ogl:UpdatePurchaseOrderReceived" json:"UpdatePurchaseOrderReceived"`
	BporSeq BporSeq  `xml:"ogl:BporSeq" json:"BporSeq"`
}

type BporSeq struct {
	Int []int64 `xml:"ogl:int" json:"Int"`
}

type UpdatePurchaseOrderReceivedResponseBody struct {
	XMLName                             xml.Name `xml:"soapenv:Body"`
	UpdatePurchaseOrderReceivedResponse UpdatePurchaseOrderReceivedResponse
}
type UpdatePurchaseOrderReceivedResponse struct {
	XMLName                           xml.Name `xml:"UpdatePurchaseOrderReceivedResponse" json:"UpdatePurchaseOrderReceivedResponse"`
	Xmlns                             string   `xml:"xmlns,attr" json:"-"`
	UpdatePurchaseOrderReceivedResult int32    `xml:"UpdatePurchaseOrderReceivedResult" json:"UpdatePurchaseOrderReceivedResult"`
}

// 3 method for insert request header
type InsetPurchaseReceiptsHeaderBody struct {
	XMLName                     xml.Name `xml:"soapenv:Body"`
	InsetPurchaseReceiptsHeader InsetPurchaseReceiptsHeader
}

type InsetPurchaseReceiptsHeader struct {
	XMLName         xml.Name        `xml:"ogl:InsetPurchaseReceiptsHeader" json:"InsetPurchaseReceiptsHeader"`
	PurchaseHeaders PurchaseHeaders `xml:"ogl:PurchaseHeaders"  json:"PurchaseHeaders"`
}
type PurchaseHeaders struct {
	PurchaseReceiptsHead PurchaseReceiptsHead `xml:"ogl:PurchaseReceiptsHead"  json:"PurchaseReceiptsHead"`
}

type PurchaseReceiptsHead struct {
	BporOrgID            int    `xml:"ogl:BporOrgId" json:"bpor_org_id"`
	BporOrganizationID   int    `xml:"ogl:BporOrganizationId" json:"bpor_organization_id"`
	BporSubinventoryCode int    `xml:"ogl:BporSubinventoryCode" json:"bpor_subinventory_code"`
	BporPurchaseOrderNo  int    `xml:"ogl:BporPurchaseOrderNo" json:"bpor_purchase_order_no"`
	BprhPoReceiptNo      int    `xml:"ogl:BprhPoReceiptNo" json:"bprh_po_receipt_no"`
	BprhSupplier         int    `xml:"ogl:BprhSupplier" json:"bprh_supplier"`
	BprhReceiptDate      string `xml:"ogl:BprhReceiptDate" json:"bprh_receipt_date"`
	BprhCreationDate     string `xml:"ogl:BprhCreationDate" json:"bprh_creation_date"`
	BprhCreatedUser      int    `xml:"ogl:BprhCreatedUser" json:"bprh_created_user"`
	BprhLastUpdatedDate  string `xml:"ogl:BprhLastUpdatedDate" json:"bprh_last_updated_user"`
	BprhUpdatedUser      int    `xml:"ogl:BprhUpdatedUser" json:"bprh_updated_user"`
	BprhReadFlag         string `xml:"ogl:BprhReadFlag" json:"bprh_read_flag"`
}

type InsetPurchaseReceiptsHeaderResponseBody struct {
	XMLName                             xml.Name `xml:"soapenv:Body"`
	InsetPurchaseReceiptsHeaderResponse InsetPurchaseReceiptsHeaderResponse
}

type InsetPurchaseReceiptsHeaderResponse struct {
	XMLName                           xml.Name `xml:"ogl:InsetPurchaseReceiptsHeaderResponse" json:"-"`
	Text                              string   `xml:",chardata" json:"-"`
	Xmlns                             string   `xml:"xmlns,attr" json:"-"`
	InsetPurchaseReceiptsHeaderResult string   `xml:"og1:InsetPurchaseReceiptsHeaderResult" json:"InsetPurchaseReceiptsHeaderResult"`
}

//4 method

type InsetPurchaseReceiptsDetailBody struct {
	XMLName                      xml.Name `xml:"soapenv:Body"`
	InsertPurchaseReceiptsDetail InsertPurchaseReceiptsDetail
}

type InsertPurchaseReceiptsDetail struct {
	XMLName              xml.Name             `xml:"ogl:InsetPurchaseReceiptsDetail" json:"InsetPurchaseReceiptsDetail"`
	PurchaseReceiptsDtls PurchaseReceiptsDtls `xml:"ogl:PurchaseReceptsDtls" json:"PurchaseReceptsDtls"`
}

type PurchaseReceiptsDtls struct {
	PurchaseReceiptDetl []PurchaseReceiptDetl `xml:"ogl:PurchaseReceiptDetl" json:"PurchaseReceiptDetl"`
}

type PurchaseReceiptDetl struct {
	BporOrgID            int    `xml:"ogl:BporOrgId" json:"bpor_org_id"`
	BporOrganizationID   int    `xml:"ogl:BporOrganizationId" json:"bpor_organization_id"`
	BporSubinventoryCode int    `xml:"ogl:BporSubinventoryCode" json:"bpor_subinventory_code"`
	BporPurchaseOrderNo  int    `xml:"ogl:BporPurchaseOrderNo" json:"bpor_purchase_order_no"`
	BprdReceiptNumber    int    `xml:"ogl:BprdReceiptNumber" json:"bprd_receipt_number"`
	BprdPoDetlSrlNo      int    `xml:"ogl:BprdPoDetlSrlNo" json:"bprd_po_detl_srl_no"`
	BprdItemCode         string `xml:"ogl:BprdItemCode" json:"bprd_item_code"`
	BprdReceivedQty      int    `xml:"ogl:BprdReceivedQty" json:"bprd_received_qty"`
	BprdRequestedQty     int    `xml:"ogl:BprdRequestedQty" json:"bprd_requested_qty"`
	BprdCreationDate     string `xml:"ogl:BprdCreationDate" json:"bprd_creation_date"`
	BprdCreatedUser      string `xml:"ogl:BprdCreatedUser" json:"bprd_created_user"`
	BprdLastUpdatedDate  string `xml:"ogl:BprdLastUpdatedDate" json:"bprd_last_updated_date"`
	BprdUpdatedUser      int    `xml:"ogl:BprdUpdatedUser" json:"bprd_updated_user"`
	BprdReadFlag         string `xml:"ogl:BprdReadFlag" json:"bprd_read_flag"`
}

type InsetPurchaseReceiptsDetailResponseBody struct {
	XMLName                             xml.Name `xml:"soapenv:Body"`
	InsetPurchaseReceiptsDetailResponse InsetPurchaseReceiptsDetailResponse
}
type InsetPurchaseReceiptsDetailResponse struct {
	XMLName                           xml.Name `xml:"InsetPurchaseReceiptsDetailResponse" json:"InsetPurchaseReceiptsDetailResponse"`
	Text                              string   `xml:",chardata" json:"-"`
	Xmlns                             string   `xml:"xmlns,attr" json:"-"`
	InsetPurchaseReceiptsDetailResult string   `xml:"InsetPurchaseReceiptsDetailResult" json:"InsetPurchaseReceiptsHeaderResult"`
}

//Custome Method
type erpSt struct {
	Url  string `json:"url" uri:"url" form:"url"`
	Body string `json:"body" uri:"body" form:"body"`
}

// method 1
func (g *ERPController) GetPO(c *gin.Context) {
	requestEnvelope := NewRequestEnvelope(&GetPurchaseOrders{})
	output, _ := xml.MarshalIndent(requestEnvelope, "  ", "    ")
	res, err := AxiomAPI(string(output))
	response, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(response))
	if err != nil {
		fmt.Println("callAPI err: ", err)
	}
	defer res.Body.Close()
	txnPO := &EnvelopePO{
		Body: new(GetPurchaseOrdersResponse),
	}
	if err = xml.Unmarshal(response, txnPO); err != nil {
		fmt.Println(err)
	}
	// 1. Marshal to json
	jsonData, err := json.Marshal(*txnPO)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Json data \n", string(jsonData))
	var result map[string]interface{}
	if err := json.Unmarshal(jsonData, &result); err != nil {
		fmt.Println(err)
	}
	resp(c, result)

}

// func (g *ERPController) erpParserfunc(c *gin.Context) {
//     var w http.ResponseWriter = c.Writer
//     var req *http.Request = c.Request
//     var erpStvalue erpSt
//     if g.BindAndValidate(c, &erpStvalue) {
// 		fmt.Fprintf(w, "successfully deleted" + erpStvalue.Url)
//     }
//     return
// }
//method 2
func (g *ERPController) UpdatePO(c *gin.Context) {
	// get the ID and for the soap response
	// this json value plz get from controller
	jsData := []byte(`{"UpdatePurchaseOrderReceived":{"BporSeq":{"Int":"22675"}}}`)
	var data UpdatePurchaseOrderReceived
	json.Unmarshal([]byte(jsData), &data)
	xmlData, err := xml.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(xmlData))
	requestEnvelope := NewRequestEnvelope(string(xmlData))
	output, _ := xml.MarshalIndent(requestEnvelope, "  ", "    ")
	output = bytes.Replace(output, []byte("&gt;"), []byte(">"), -1)
	output = bytes.Replace(output, []byte("&lt;"), []byte("<"), -1)
	fmt.Println(string(output))
	res, err := AxiomAPI(string(output))
	response, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(response))
	if err != nil {
		fmt.Println("callAPI err: ", err)
		return
	}
	defer res.Body.Close()
	//resp.Write(response)
	//return

	txnPO := &EnvelopePO{
		Body: new(UpdatePurchaseOrderReceivedResponseBody),
	}
	if err = xml.Unmarshal(response, txnPO); err != nil {
		fmt.Println(err)
		return
	}
	// 1. Marshal to json
	jsonData, err := json.Marshal(*txnPO)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("Json data \n\n\n", string(jsonData))
	return

}

//method 3
func (g *ERPController) InsertHeaderPO(c *gin.Context) {
	// u should pass from controller
	jsData := []byte(` {"InsetPurchaseReceiptsHeader":{"PurchaseHeaders":{"PurchaseReceiptsHead":{"bpor_org_id":"103","bpor_organization_id":"208","bpor_subinventory_code":"01097","bpor_purchase_order_no":"161041","bprh_po_receipt_no":"3","bprh_supplier":"2125","bprh_receipt_date":"2020-11-04T17:20:04","bprh_creation_date":"2020-11-02T17:20:04","bprh_created_user":"-1",
    "bprh_last_updated_user":"2020-11-02T17:20:04","bprh_updated_user":"-1","bprh_read_flag":"N"}}}}`)
	var data InsetPurchaseReceiptsHeader
	json.Unmarshal([]byte(jsData), &data)
	xmlData, err := xml.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(xmlData))
	requestEnvelope := NewRequestEnvelope(string(xmlData))
	output, _ := xml.MarshalIndent(requestEnvelope, "  ", "    ")
	output = bytes.Replace(output, []byte("&gt;"), []byte(">"), -1)
	output = bytes.Replace(output, []byte("&lt;"), []byte("<"), -1)
	fmt.Println(string(output))
	res, err := AxiomAPI(string(output))
	response, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(response))
	if err != nil {
		fmt.Println("callAPI err: ", err)
		return
	}
	defer res.Body.Close()

	//  resp.Write(response)
	//  return

	txnPO := &EnvelopePO{
		Body: new(InsetPurchaseReceiptsHeaderResponseBody),
	}
	if err = xml.Unmarshal(response, txnPO); err != nil {
		fmt.Println(err)
		return
	}
	// 1. Marshal to json
	jsonData, err := json.Marshal(*txnPO)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("Json data \n\n\n", string(jsonData))
	return

}

//method 4
func (g *ERPController) InsertDetailPO(c *gin.Context) {
	// u should pass ur json from controller
	jsData := []byte(` {"InsetPurchaseReceiptsDetail":{"PurchaseReceiptsDtls":{"PurchaseReceiptDetl":[{"bpor_org_id":103,"bpor_organization_id":208,"bpor_subinventory_code":1097,"bpor_purchase_order_no":161041,"bprd_receipt_number":3,"bprd_po_detl_srl_no":2,"bprd_item_code":"OPDUCOMEV024","bprd_received_qty":20000,"bprd_requested_qty":20000,"bprd_creation_date":"2020-11-04T11:58:04","bprd_created_user":"1671","bprd_last_updated_date":"2020-11-04T11:58:04","bprd_updated_user":1671,"bprd_read_flag":"N"},{"bpor_org_id":104,"bpor_organization_id":208,"bpor_subinventory_code":1097,"bpor_purchase_order_no":161041,"bprd_receipt_number":3,"bprd_po_detl_srl_no":2,"bprd_item_code":"OPDUCOMEV024","bprd_received_qty":20000,"bprd_requested_qty":20000,"bprd_creation_date":"2020-11-04T11:58:04","bprd_created_user":"1671","bprd_last_updated_date":"2020-11-04T11:58:04","bprd_updated_user":1671,"bprd_read_flag":"N"}]}}}
    `)
	var data InsertPurchaseReceiptsDetail
	json.Unmarshal([]byte(jsData), &data)
	xmlData, err := xml.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(xmlData))
	requestEnvelope := NewRequestEnvelope(string(xmlData))
	output, _ := xml.MarshalIndent(requestEnvelope, "  ", "    ")
	output = bytes.Replace(output, []byte("&gt;"), []byte(">"), -1)
	output = bytes.Replace(output, []byte("&lt;"), []byte("<"), -1)
	fmt.Println(string(output))

	res, err := AxiomAPI(string(output))
	response, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(response))
	if err != nil {
		fmt.Println("callAPI err: ", err)
		return
	}
	defer res.Body.Close()
	txnPO := &EnvelopePO{
		Body: new(InsetPurchaseReceiptsDetailResponseBody),
	}
	if err = xml.Unmarshal(response, txnPO); err != nil {
		fmt.Println(err)
		return
	}
	// 1. Marshal to json
	jsonData, err := json.Marshal(*txnPO)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("Json data \n\n\n", string(jsonData))
	return
}

// common
func AxiomAPI(strBody string) (*http.Response, error) {
	// check if the value exists in new system and return the than for soap xml response

	defer func() { fmt.Println("callAPI: completed fetch for service") }()
	req, err := http.NewRequest(http.MethodPost,
		"http://172.27.53.125/OglobaUATTest/OglobaWebService.asmx?wsdl", bytes.NewBuffer([]byte(strings.TrimSpace(strBody))))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "text/xml; charset=ISO-8859-1")

	dump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Printf("%q", dump)

	cfg := &tls.Config{
		InsecureSkipVerify: true,
	}

	tr := &http.Transport{
		TLSClientConfig: cfg,
	}
	client := &http.Client{
		Transport: tr,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(resp)
	return resp, nil
}

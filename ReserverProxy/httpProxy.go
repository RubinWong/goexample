// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"reflect"
// 	"runtime"
// 	"strconv"
// 	"time"

// 	"github.com/coreos/pkg/capnslog"
// 	"golang.org/x/net/http/httpguts"
// )

// func GetGid() (gid uint64) {
// 	b := make([]byte, 64)
// 	b = b[:runtime.Stack(b, false)]
// 	b = bytes.TrimPrefix(b, []byte("goroutine "))
// 	b = b[:bytes.IndexByte(b, ' ')]
// 	n, err := strconv.ParseUint(string(b), 10, 64)
// 	if err != nil {
// 	   panic(err)
// 	}
// 	return n
// }

// type HTTPError struct {
// 	Msg string `json:"message"`
// 	Code int   `json:"code"`
// }

// func (e HTTPError) WriteTo(w http.ResponseWriter) error {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(e.Code)
// 	b, err := json.Marshal(e)
// 	if err != nil {
// 		plog.Panicf("marshal HTTPError should never fail (%v)", err)
// 	}
// 	if _, err := w.Write(b); err != nil {
// 		return err
// 	}
// 	return nil
// }

// type Proxy struct {
// 	Router map[string]func()
// }

// func copyHeader(dst http.Header, src http.Header) {
// 	for k, vv := range src {
// 		for _, v := range vv {
// 			dst.Add(k, v)
// 		}
// 	}
// }

// var (
// 	plog = capnslog.NewPackageLogger("reserv", "proxy/httpproxy")

// 	singleHopHeaders = []string{
// 		"Connection",
// 		"Keep-Alive",
// 		"Proxy-Authenticate",
// 		"Proxy-Authorization",
// 		"Te", // canonicalized version of "TE"
// 		"Trailers",
// 		"Transfer-Encoding",
// 		"Upgrade",
// 	}
// )

// func removeSingleHopHeader(src *http.Request) {
// 	for _, v := range singleHopHeaders {
// 		src.Header.Del(v)
// 	}
// }

// func prepareRequest(src *http.Request) (req *http.Request, err error) {
// 	req = &http.Request{}
// 	err = nil
// 	*req = *src
	
// 	req.Header = make(http.Header)
// 	copyHeader(req.Header, src.Header)
// 	normalizeRequest(req)
// 	removeSingleHopHeader(req)
	
// 	if (src.Body != nil) {
// 		body, err := ioutil.ReadAll(src.Body)
// 		if err != nil {
// 			plog.Error("read body fail")
// 			return nil , err;
// 		}
// 		req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
// 	}
// 	return
// }

// func (m* Proxy)ServeHTTP(rw http.ResponseWriter, clientreq *http.Request) {
// 	fmt.Println("ServeHTTP ", GetGid())
// 	req, err := prepareRequest(clientreq)
// 	if  err != nil {
// 		msg := fmt.Sprintf("failed to read request body: %v ", err)
// 		plog.Debug(msg)
// 		e := &HTTPError{
// 			Msg: msg,
// 			Code: http.StatusInternalServerError,
// 		}
// 		if err = e.WriteTo(rw); err != nil {
// 			plog.Error("answer client fail")
// 		}
// 		return
// 	}

// 	fmt.Println(req);
// }

// func normalizeRequest(req *http.Request) {
// 	req.Close = false
// 	req.Proto = "HTTP/1.1"
// 	req.ProtoMajor = 1
// 	req.ProtoMinor = 1
// }

// const (
// 	intflag int = flags.int
// )

// func main() {
// 	mux := http.NewServeMux()
// 	p := &Proxy{}
// 	mux.Handle("/", p)

// 	server := &http.Server{
// 		Addr: ":1120",
// 		WriteTimeout: time.Second * 3,
// 		Handler: mux,
// 	}

// 	server.ListenAndServe();
// }

// func Trailers(req *http.Request) (trailers map[string][]string) {
// 	trailers = make(map[string][]string)
// 	for _, h := range req.Header[http.CanonicalHeaderKey("Trailer")] {
// 		foreachHeaderElement(h, func(key string) {
// 			key = http.CanonicalHeaderKey(key)
// 			if !httpguts.ValidTrailerHeader(key) {
// 				// Ignore since forbidden by RFC 7230, section 4.1.2.
// 				return
// 			}
// 			trailers[key] = nil
// 		})
// 	}
// 	for _, h := range req.Header {
// 		if _, ok := trailers[h.Key]; ok {
// 			trailers[h.Key] = append(trailers[h.Key], h.Value)
// 		}
// 	}
// 	return
// }

// func HandleTrailers(w http.ResponseWriter, r *http.Request) {
// 	trailers := Trailers(r)
// 	for _, t := range trailers {
// 		for _, v := range t {
// 			w.Header().Add("Trailer", t)
// 			w.Header().Add("Trailer", v)
// 		}
// 	}
// }

// // process htttp request
// func HandleHttpRequest(w http.ResponseWriter, r *http.Request) {
// 	//fmt.Println("HandleHttpRequest ", GetGid())
// 	//fmt.Println(r);
// 	//fmt.Println(r.Header);
// 	//fmt.Println(r.Header.Get("Content-Length"));
// 	//fmt.Println(r.Header.Get("Content-Type"));
// 	//fmt.Println(r.Header.Get("Host"));
// 	//fmt.Println(r.Header.Get("User-Agent"));
// 	//fmt.Println(r.Header.Get("Accept"));
// 	//fmt.Println(r.Header.Get("Accept-Encoding"));
// 	//fmt.Println(r.Header.Get("Accept-Language"));
// 	//fmt.Println(r.Header.Get("Connection"));
// 	//fmt.Println(r.Header.Get("Keep-Alive"));
// 	//fmt.Println(r.Header.Get("Proxy-Authenticate"));
// 	//fmt.Println(r.Header.Get("Proxy-Authorization"));
// 	//fmt.Println(r.Header.Get("Te"));
// 	//fmt.Println(r.Header.Get("Trailers"));
// 	//fmt.Println(r.Header.Get("Transfer-Encoding"));
// 	//fmt.Println(r.Header.Get("Upgrade"));
// 	//fmt.Println(r.Header.Get("Cookie"));
// 	//fmt.Println(r.Header.Get("Origin"));
// 	//fmt.Println(r.Header.Get("Referer"));
// 	//fmt.Println(r.Header.Get("Sec-Fetch-Mode"));
// 	//fmt.Println(r.Header.Get("Sec-Fetch-Site"));
// 	//fmt.Println(r.Header.Get("Sec-Fetch-User"));
// 	//fmt.Println(r.Header.Get("Sec-WebSocket-Key"));
// 	//fmt.Println(r.Header.Get("Sec-WebSocket-Version"));
// 	//fmt.Println(r.Header.Get("Sec-WebSocket-Extensions"));
// 	//fmt.Println(r.Header.Get("Sec-WebSocket-Protocol"));
// 	//fmt.Println(r.Header.Get("Sec-WebSocket-Accept"));
// 	//fmt.Println(r.Header.Get("Sec-WebSocket-Key1"));
// 	//fmt.Println(r.Header.Get("Sec-WebSocket-Key2"));
// 	//fmt.Println(r.Header.Get("Sec-WebSocket-Origin"));
// 	//fmt.Println(r.Header.Get("Sec-WebSocket-Location"));
// 	fmt.Println(r.Header.Get("x-forwarded-for"))
// 	fmt.Println(r.Header.Get("x-forwarded-proto"))
// 	fmt.Println(r.Header.Get("x-forwarded-port"))
// 	fmt.Println(r.Header.Get("x-forwarded-host"))
// 	fmt.Println(r.Header.Get("x-forwarded-server"))
// 	fmt.Println(r.Header.Get("x-forwarded-ssl"))
// 	fmt.Println(r.Header.Get("x-real-ip"))
// }

// func Marshal(v interface{}) ([]byte, error) {
// 	var buf bytes.Buffer
// 	enc := json.NewEncoder(&buf)
// 	err := enc.Encode(v)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return buf.Bytes(), nil
// }

// // redis result to go struct
// func Unmarshal(b []byte, v interface{}) error {
// 	dec := json.NewDecoder(bytes.NewReader(b))
// 	dec.UseNumber()
// 	return dec.Decode(v)
// }

// func SetField(v interface{}, name string, value interface{}) error {
// 	structValue := reflect.ValueOf(v).Elem()
// 	structFieldValue := structValue.FieldByName(name)

// 	if !structFieldValue.IsValid() {
// 		return fmt.Errorf("No such field: %s in obj", name)
// 	}

// 	if !structFieldValue.CanSet() {
// 		return fmt.Errorf("Cannot set %s field value", name)
// 	}

// 	structFieldType := structFieldValue.Type()
// 	val := reflect.ValueOf(value)
// 	if structFieldType != val.Type() {
// 		return errors.New("Provided value type didn't match obj field type")
// 	}

// 	structFieldValue.Set(val)
// 	return nil
// }

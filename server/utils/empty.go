package utils

//
//import (
//	"bufio"
//	"crypto/tls"
//	"encoding/base64"
//	"fmt"
//	"golang.org/x/net/http/httpguts"
//	"golang.org/x/net/idna"
//	"io"
//	"io/ioutil"
//	"mime"
//	"mime/multipart"
//	"net"
//	"net/http/httptrace"
//	"net/textproto"
//	"net/url"
//	"os"
//	"path"
//	"runtime"
//	"sort"
//	"strconv"
//	"strings"
//	"sync"
//	"sync/atomic"
//	"time"
//)
//
//package http
//
//import (
//"context"
//"crypto/tls"
//"encoding/base64"
//"errors"
//"fmt"
//"io"
//"io/ioutil"
//"log"
//"net/url"
//"reflect"
//"sort"
//"strings"
//"sync"
//"time"
//)
//
//// A Client is an HTTP client. Its zero value (DefaultClient) is a
//// usable client that uses DefaultTransport.
////
//// The Client's Transport typically has internal state (cached TCP
//// connections), so Clients should be reused instead of created as
//// needed. Clients are safe for concurrent use by multiple goroutines.
////
//// A Client is higher-level than a RoundTripper (such as Transport)
//// and additionally handles HTTP details such as cookies and
//// redirects.
////
//// When following redirects, the Client will forward all headers set on the
//// initial Request except:
////
//// • when forwarding sensitive headers like "Authorization",
//// "WWW-Authenticate", and "Cookie" to untrusted targets.
//// These headers will be ignored when following a redirect to a domain
//// that is not a subdomain match or exact match of the initial domain.
//// For example, a redirect from "foo.com" to either "foo.com" or "sub.foo.com"
//// will forward the sensitive headers, but a redirect to "bar.com" will not.
////
//// • when forwarding the "Cookie" header with a non-nil cookie Jar.
//// Since each redirect may mutate the state of the cookie jar,
//// a redirect may possibly alter a cookie set in the initial request.
//// When forwarding the "Cookie" header, any mutated cookies will be omitted,
//// with the expectation that the Jar will insert those mutated cookies
//// with the updated values (assuming the origin matches).
//// If Jar is nil, the initial cookies are forwarded without change.
////
//type Client struct {
//	// Transport specifies the mechanism by which individual
//	// HTTP requests are made.
//	// If nil, DefaultTransport is used.
//	Transport RoundTripper
//
//	// CheckRedirect specifies the policy for handling redirects.
//	// If CheckRedirect is not nil, the client calls it before
//	// following an HTTP redirect. The arguments req and via are
//	// the upcoming request and the requests made already, oldest
//	// first. If CheckRedirect returns an error, the Client's Get
//	// method returns both the previous Response (with its Body
//	// closed) and CheckRedirect's error (wrapped in a url.Error)
//	// instead of issuing the Request req.
//	// As a special case, if CheckRedirect returns ErrUseLastResponse,
//	// then the most recent response is returned with its body
//	// unclosed, along with a nil error.
//	//
//	// If CheckRedirect is nil, the Client uses its default policy,
//	// which is to stop after 10 consecutive requests.
//	CheckRedirect func(req *Request, via []*Request) error
//
//	// Jar specifies the cookie jar.
//	//
//	// The Jar is used to insert relevant cookies into every
//	// outbound Request and is updated with the cookie values
//	// of every inbound Response. The Jar is consulted for every
//	// redirect that the Client follows.
//	//
//	// If Jar is nil, cookies are only sent if they are explicitly
//	// set on the Request.
//	Jar CookieJar
//
//	// Timeout specifies a time limit for requests made by this
//	// Client. The timeout includes connection time, any
//	// redirects, and reading the response body. The timer remains
//	// running after Get, Head, Post, or Do return and will
//	// interrupt reading of the Response.Body.
//	//
//	// A Timeout of zero means no timeout.
//	//
//	// The Client cancels requests to the underlying Transport
//	// as if the Request's Context ended.
//	//
//	// For compatibility, the Client will also use the deprecated
//	// CancelRequest method on Transport if found. New
//	// RoundTripper implementations should use the Request's Context
//	// for cancellation instead of implementing CancelRequest.
//	Timeout time.Duration
//}
//
//// DefaultClient is the default Client and is used by Get, Head, and Post.
//var DefaultClient = &Client{}
//
//// RoundTripper is an interface representing the ability to execute a
//// single HTTP transaction, obtaining the Response for a given Request.
////
//// A RoundTripper must be safe for concurrent use by multiple
//// goroutines.
//type RoundTripper interface {
//	// RoundTrip executes a single HTTP transaction, returning
//	// a Response for the provided Request.
//	//
//	// RoundTrip should not attempt to interpret the response. In
//	// particular, RoundTrip must return err == nil if it obtained
//	// a response, regardless of the response's HTTP status code.
//	// A non-nil err should be reserved for failure to obtain a
//	// response. Similarly, RoundTrip should not attempt to
//	// handle higher-level protocol details such as redirects,
//	// authentication, or cookies.
//	//
//	// RoundTrip should not modify the request, except for
//	// consuming and closing the Request's Body. RoundTrip may
//	// read fields of the request in a separate goroutine. Callers
//	// should not mutate or reuse the request until the Response's
//	// Body has been closed.
//	//
//	// RoundTrip must always close the body, including on errors,
//	// but depending on the implementation may do so in a separate
//	// goroutine even after RoundTrip returns. This means that
//	// callers wanting to reuse the body for subsequent requests
//	// must arrange to wait for the Close call before doing so.
//	//
//	// The Request's URL and Header fields must be initialized.
//	RoundTrip(*Request) (*Response, error)
//}
//
//// refererForURL returns a referer without any authentication info or
//// an empty string if lastReq scheme is https and newReq scheme is http.
//func refererForURL(lastReq, newReq *url.URL) string {
//	// https://tools.ietf.org/html/rfc7231#section-5.5.2
//	//   "Clients SHOULD NOT include a Referer header field in a
//	//    (non-secure) HTTP request if the referring page was
//	//    transferred with a secure protocol."
//	if lastReq.Scheme == "https" && newReq.Scheme == "http" {
//		return ""
//	}
//	referer := lastReq.String()
//	if lastReq.User != nil {
//		// This is not very efficient, but is the best we can
//		// do without:
//		// - introducing a new method on URL
//		// - creating a race condition
//		// - copying the URL struct manually, which would cause
//		//   maintenance problems down the line
//		auth := lastReq.User.String() + "@"
//		referer = strings.Replace(referer, auth, "", 1)
//	}
//	return referer
//}
//
//// didTimeout is non-nil only if err != nil.
//func (c *Client) send(req *Request, deadline time.Time) (resp *Response, didTimeout func() bool, err error) {
//	if c.Jar != nil {
//		for _, cookie := range c.Jar.Cookies(req.URL) {
//			req.AddCookie(cookie)
//		}
//	}
//	resp, didTimeout, err = send(req, c.transport(), deadline)
//	if err != nil {
//		return nil, didTimeout, err
//	}
//	if c.Jar != nil {
//		if rc := resp.Cookies(); len(rc) > 0 {
//			c.Jar.SetCookies(req.URL, rc)
//		}
//	}
//	return resp, nil, nil
//}
//
//func (c *Client) deadline() time.Time {
//	if c.Timeout > 0 {
//		return time.Now().Add(c.Timeout)
//	}
//	return time.Time{}
//}
//
//func (c *Client) transport() RoundTripper {
//	if c.Transport != nil {
//		return c.Transport
//	}
//	return DefaultTransport
//}
//
//// send issues an HTTP request.
//// Caller should close resp.Body when done reading from it.
//func send(ireq *Request, rt RoundTripper, deadline time.Time) (resp *Response, didTimeout func() bool, err error) {
//	req := ireq // req is either the original request, or a modified fork
//
//	if rt == nil {
//		req.closeBody()
//		return nil, alwaysFalse, errors.New("http: no Client.Transport or DefaultTransport")
//	}
//
//	if req.URL == nil {
//		req.closeBody()
//		return nil, alwaysFalse, errors.New("http: nil Request.URL")
//	}
//
//	if req.RequestURI != "" {
//		req.closeBody()
//		return nil, alwaysFalse, errors.New("http: Request.RequestURI can't be set in client requests.")
//	}
//
//	// forkReq forks req into a shallow clone of ireq the first
//	// time it's called.
//	forkReq := func() {
//		if ireq == req {
//			req = new(Request)
//			*req = *ireq // shallow clone
//		}
//	}
//
//	// Most the callers of send (Get, Post, et al) don't need
//	// Headers, leaving it uninitialized. We guarantee to the
//	// Transport that this has been initialized, though.
//	if req.Header == nil {
//		forkReq()
//		req.Header = make(Header)
//	}
//
//	if u := req.URL.User; u != nil && req.Header.Get("Authorization") == "" {
//		username := u.Username()
//		password, _ := u.Password()
//		forkReq()
//		req.Header = cloneOrMakeHeader(ireq.Header)
//		req.Header.Set("Authorization", "Basic "+basicAuth(username, password))
//	}
//
//	if !deadline.IsZero() {
//		forkReq()
//	}
//	stopTimer, didTimeout := setRequestCancel(req, rt, deadline)
//
//	resp, err = rt.RoundTrip(req)
//	if err != nil {
//		stopTimer()
//		if resp != nil {
//			log.Printf("RoundTripper returned a response & error; ignoring response")
//		}
//		if tlsErr, ok := err.(tls.RecordHeaderError); ok {
//			// If we get a bad TLS record header, check to see if the
//			// response looks like HTTP and give a more helpful error.
//			// See golang.org/issue/11111.
//			if string(tlsErr.RecordHeader[:]) == "HTTP/" {
//				err = errors.New("http: server gave HTTP response to HTTPS client")
//			}
//		}
//		return nil, didTimeout, err
//	}
//	if !deadline.IsZero() {
//		resp.Body = &cancelTimerBody{
//			stop:          stopTimer,
//			rc:            resp.Body,
//			reqDidTimeout: didTimeout,
//		}
//	}
//	return resp, nil, nil
//}
//
//// timeBeforeContextDeadline reports whether the non-zero Time t is
//// before ctx's deadline, if any. If ctx does not have a deadline, it
//// always reports true (the deadline is considered infinite).
//func timeBeforeContextDeadline(t time.Time, ctx context.Context) bool {
//	d, ok := ctx.Deadline()
//	if !ok {
//		return true
//	}
//	return t.Before(d)
//}
//
//// knownRoundTripperImpl reports whether rt is a RoundTripper that's
//// maintained by the Go team and known to implement the latest
//// optional semantics (notably contexts). The Request is used
//// to check whether this particular request is using an alternate protocol,
//// in which case we need to check the RoundTripper for that protocol.
//func knownRoundTripperImpl(rt RoundTripper, req *Request) bool {
//	switch t := rt.(type) {
//	case *Transport:
//		if altRT := t.alternateRoundTripper(req); altRT != nil {
//			return knownRoundTripperImpl(altRT, req)
//		}
//		return true
//	case *http2Transport, http2noDialH2RoundTripper:
//		return true
//	}
//	// There's a very minor chance of a false positive with this.
//	// Insted of detecting our golang.org/x/net/http2.Transport,
//	// it might detect a Transport type in a different http2
//	// package. But I know of none, and the only problem would be
//	// some temporarily leaked goroutines if the transport didn't
//	// support contexts. So this is a good enough heuristic:
//	if reflect.TypeOf(rt).String() == "*http2.Transport" {
//		return true
//	}
//	return false
//}
//
//// setRequestCancel sets req.Cancel and adds a deadline context to req
//// if deadline is non-zero. The RoundTripper's type is used to
//// determine whether the legacy CancelRequest behavior should be used.
////
//// As background, there are three ways to cancel a request:
//// First was Transport.CancelRequest. (deprecated)
//// Second was Request.Cancel.
//// Third was Request.Context.
//// This function populates the second and third, and uses the first if it really needs to.
//func setRequestCancel(req *Request, rt RoundTripper, deadline time.Time) (stopTimer func(), didTimeout func() bool) {
//	if deadline.IsZero() {
//		return nop, alwaysFalse
//	}
//	knownTransport := knownRoundTripperImpl(rt, req)
//	oldCtx := req.Context()
//
//	if req.Cancel == nil && knownTransport {
//		// If they already had a Request.Context that's
//		// expiring sooner, do nothing:
//		if !timeBeforeContextDeadline(deadline, oldCtx) {
//			return nop, alwaysFalse
//		}
//
//		var cancelCtx func()
//		req.ctx, cancelCtx = context.WithDeadline(oldCtx, deadline)
//		return cancelCtx, func() bool { return time.Now().After(deadline) }
//	}
//	initialReqCancel := req.Cancel // the user's original Request.Cancel, if any
//
//	var cancelCtx func()
//	if oldCtx := req.Context(); timeBeforeContextDeadline(deadline, oldCtx) {
//		req.ctx, cancelCtx = context.WithDeadline(oldCtx, deadline)
//	}
//
//	cancel := make(chan struct{})
//	req.Cancel = cancel
//
//	doCancel := func() {
//		// The second way in the func comment above:
//		close(cancel)
//		// The first way, used only for RoundTripper
//		// implementations written before Go 1.5 or Go 1.6.
//		type canceler interface{ CancelRequest(*Request) }
//		if v, ok := rt.(canceler); ok {
//			v.CancelRequest(req)
//		}
//	}
//
//	stopTimerCh := make(chan struct{})
//	var once sync.Once
//	stopTimer = func() {
//		once.Do(func() {
//			close(stopTimerCh)
//			if cancelCtx != nil {
//				cancelCtx()
//			}
//		})
//	}
//
//	timer := time.NewTimer(time.Until(deadline))
//	var timedOut atomicBool
//
//	go func() {
//		select {
//		case <-initialReqCancel:
//			doCancel()
//			timer.Stop()
//		case <-timer.C:
//			timedOut.setTrue()
//			doCancel()
//		case <-stopTimerCh:
//			timer.Stop()
//		}
//	}()
//
//	return stopTimer, timedOut.isSet
//}
//
//// See 2 (end of page 4) https://www.ietf.org/rfc/rfc2617.txt
//// "To receive authorization, the client sends the userid and password,
//// separated by a single colon (":") character, within a base64
//// encoded string in the credentials."
//// It is not meant to be urlencoded.
//func basicAuth(username, password string) string {
//	auth := username + ":" + password
//	return base64.StdEncoding.EncodeToString([]byte(auth))
//}
//
//// Get issues a GET to the specified URL. If the response is one of
//// the following redirect codes, Get follows the redirect, up to a
//// maximum of 10 redirects:
////
////    301 (Moved Permanently)
////    302 (Found)
////    303 (See Other)
////    307 (Temporary Redirect)
////    308 (Permanent Redirect)
////
//// An error is returned if there were too many redirects or if there
//// was an HTTP protocol error. A non-2xx response doesn't cause an
//// error. Any returned error will be of type *url.Error. The url.Error
//// value's Timeout method will report true if request timed out or was
//// canceled.
////
//// When err is nil, resp always contains a non-nil resp.Body.
//// Caller should close resp.Body when done reading from it.
////
//// Get is a wrapper around DefaultClient.Get.
////
//// To make a request with custom headers, use NewRequest and
//// DefaultClient.Do.
//func Get(url string) (resp *Response, err error) {
//	return DefaultClient.Get(url)
//}
//
//// Get issues a GET to the specified URL. If the response is one of the
//// following redirect codes, Get follows the redirect after calling the
//// Client's CheckRedirect function:
////
////    301 (Moved Permanently)
////    302 (Found)
////    303 (See Other)
////    307 (Temporary Redirect)
////    308 (Permanent Redirect)
////
//// An error is returned if the Client's CheckRedirect function fails
//// or if there was an HTTP protocol error. A non-2xx response doesn't
//// cause an error. Any returned error will be of type *url.Error. The
//// url.Error value's Timeout method will report true if the request
//// timed out.
////
//// When err is nil, resp always contains a non-nil resp.Body.
//// Caller should close resp.Body when done reading from it.
////
//// To make a request with custom headers, use NewRequest and Client.Do.
//func (c *Client) Get(url string) (resp *Response, err error) {
//	req, err := NewRequest("GET", url, nil)
//	if err != nil {
//		return nil, err
//	}
//	return c.Do(req)
//}
//
//func alwaysFalse() bool { return false }
//
//// ErrUseLastResponse can be returned by Client.CheckRedirect hooks to
//// control how redirects are processed. If returned, the next request
//// is not sent and the most recent response is returned with its body
//// unclosed.
//var ErrUseLastResponse = errors.New("net/http: use last response")
//
//// checkRedirect calls either the user's configured CheckRedirect
//// function, or the default.
//func (c *Client) checkRedirect(req *Request, via []*Request) error {
//	fn := c.CheckRedirect
//	if fn == nil {
//		fn = defaultCheckRedirect
//	}
//	return fn(req, via)
//}
//
//// redirectBehavior describes what should happen when the
//// client encounters a 3xx status code from the server
//func redirectBehavior(reqMethod string, resp *Response, ireq *Request) (redirectMethod string, shouldRedirect, includeBody bool) {
//	switch resp.StatusCode {
//	case 301, 302, 303:
//		redirectMethod = reqMethod
//		shouldRedirect = true
//		includeBody = false
//
//		// RFC 2616 allowed automatic redirection only with GET and
//		// HEAD requests. RFC 7231 lifts this restriction, but we still
//		// restrict other methods to GET to maintain compatibility.
//		// See Issue 18570.
//		if reqMethod != "GET" && reqMethod != "HEAD" {
//			redirectMethod = "GET"
//		}
//	case 307, 308:
//		redirectMethod = reqMethod
//		shouldRedirect = true
//		includeBody = true
//
//		// Treat 307 and 308 specially, since they're new in
//		// Go 1.8, and they also require re-sending the request body.
//		if resp.Header.Get("Location") == "" {
//			// 308s have been observed in the wild being served
//			// without Location headers. Since Go 1.7 and earlier
//			// didn't follow these codes, just stop here instead
//			// of returning an error.
//			// See Issue 17773.
//			shouldRedirect = false
//			break
//		}
//		if ireq.GetBody == nil && ireq.outgoingLength() != 0 {
//			// We had a request body, and 307/308 require
//			// re-sending it, but GetBody is not defined. So just
//			// return this response to the user instead of an
//			// error, like we did in Go 1.7 and earlier.
//			shouldRedirect = false
//		}
//	}
//	return redirectMethod, shouldRedirect, includeBody
//}
//
//// urlErrorOp returns the (*url.Error).Op value to use for the
//// provided (*Request).Method value.
//func urlErrorOp(method string) string {
//	if method == "" {
//		return "Get"
//	}
//	return method[:1] + strings.ToLower(method[1:])
//}
//
//// Do sends an HTTP request and returns an HTTP response, following
//// policy (such as redirects, cookies, auth) as configured on the
//// client.
////
//// An error is returned if caused by client policy (such as
//// CheckRedirect), or failure to speak HTTP (such as a network
//// connectivity problem). A non-2xx status code doesn't cause an
//// error.
////
//// If the returned error is nil, the Response will contain a non-nil
//// Body which the user is expected to close. If the Body is not both
//// read to EOF and closed, the Client's underlying RoundTripper
//// (typically Transport) may not be able to re-use a persistent TCP
//// connection to the server for a subsequent "keep-alive" request.
////
//// The request Body, if non-nil, will be closed by the underlying
//// Transport, even on errors.
////
//// On error, any Response can be ignored. A non-nil Response with a
//// non-nil error only occurs when CheckRedirect fails, and even then
//// the returned Response.Body is already closed.
////
//// Generally Get, Post, or PostForm will be used instead of Do.
////
//// If the server replies with a redirect, the Client first uses the
//// CheckRedirect function to determine whether the redirect should be
//// followed. If permitted, a 301, 302, or 303 redirect causes
//// subsequent requests to use HTTP method GET
//// (or HEAD if the original request was HEAD), with no body.
//// A 307 or 308 redirect preserves the original HTTP method and body,
//// provided that the Request.GetBody function is defined.
//// The NewRequest function automatically sets GetBody for common
//// standard library body types.
////
//// Any returned error will be of type *url.Error. The url.Error
//// value's Timeout method will report true if request timed out or was
//// canceled.
//func (c *Client) Do(req *Request) (*Response, error) {
//	return c.do(req)
//}
//
//var testHookClientDoResult func(retres *Response, reterr error)
//
//func (c *Client) do(req *Request) (retres *Response, reterr error) {
//	if testHookClientDoResult != nil {
//		defer func() { testHookClientDoResult(retres, reterr) }()
//	}
//	if req.URL == nil {
//		req.closeBody()
//		return nil, &url.Error{
//			Op:  urlErrorOp(req.Method),
//			Err: errors.New("http: nil Request.URL"),
//		}
//	}
//
//	var (
//		deadline      = c.deadline()
//		reqs          []*Request
//		resp          *Response
//		copyHeaders   = c.makeHeadersCopier(req)
//		reqBodyClosed = false // have we closed the current req.Body?
//
//		// Redirect behavior:
//		redirectMethod string
//		includeBody    bool
//	)
//	uerr := func(err error) error {
//		// the body may have been closed already by c.send()
//		if !reqBodyClosed {
//			req.closeBody()
//		}
//		var urlStr string
//		if resp != nil && resp.Request != nil {
//			urlStr = stripPassword(resp.Request.URL)
//		} else {
//			urlStr = stripPassword(req.URL)
//		}
//		return &url.Error{
//			Op:  urlErrorOp(reqs[0].Method),
//			URL: urlStr,
//			Err: err,
//		}
//	}
//	for {
//		// For all but the first request, create the next
//		// request hop and replace req.
//		if len(reqs) > 0 {
//			loc := resp.Header.Get("Location")
//			if loc == "" {
//				resp.closeBody()
//				return nil, uerr(fmt.Errorf("%d response missing Location header", resp.StatusCode))
//			}
//			u, err := req.URL.Parse(loc)
//			if err != nil {
//				resp.closeBody()
//				return nil, uerr(fmt.Errorf("failed to parse Location header %q: %v", loc, err))
//			}
//			host := ""
//			if req.Host != "" && req.Host != req.URL.Host {
//				// If the caller specified a custom Host header and the
//				// redirect location is relative, preserve the Host header
//				// through the redirect. See issue #22233.
//				if u, _ := url.Parse(loc); u != nil && !u.IsAbs() {
//					host = req.Host
//				}
//			}
//			ireq := reqs[0]
//			req = &Request{
//				Method:   redirectMethod,
//				Response: resp,
//				URL:      u,
//				Header:   make(Header),
//				Host:     host,
//				Cancel:   ireq.Cancel,
//				ctx:      ireq.ctx,
//			}
//			if includeBody && ireq.GetBody != nil {
//				req.Body, err = ireq.GetBody()
//				if err != nil {
//					resp.closeBody()
//					return nil, uerr(err)
//				}
//				req.ContentLength = ireq.ContentLength
//			}
//
//			// Copy original headers before setting the Referer,
//			// in case the user set Referer on their first request.
//			// If they really want to override, they can do it in
//			// their CheckRedirect func.
//			copyHeaders(req)
//
//			// Add the Referer header from the most recent
//			// request URL to the new one, if it's not https->http:
//			if ref := refererForURL(reqs[len(reqs)-1].URL, req.URL); ref != "" {
//				req.Header.Set("Referer", ref)
//			}
//			err = c.checkRedirect(req, reqs)
//
//			// Sentinel error to let users select the
//			// previous response, without closing its
//			// body. See Issue 10069.
//			if err == ErrUseLastResponse {
//				return resp, nil
//			}
//
//			// Close the previous response's body. But
//			// read at least some of the body so if it's
//			// small the underlying TCP connection will be
//			// re-used. No need to check for errors: if it
//			// fails, the Transport won't reuse it anyway.
//			const maxBodySlurpSize = 2 << 10
//			if resp.ContentLength == -1 || resp.ContentLength <= maxBodySlurpSize {
//				io.CopyN(ioutil.Discard, resp.Body, maxBodySlurpSize)
//			}
//			resp.Body.Close()
//
//			if err != nil {
//				// Special case for Go 1 compatibility: return both the response
//				// and an error if the CheckRedirect function failed.
//				// See https://golang.org/issue/3795
//				// The resp.Body has already been closed.
//				ue := uerr(err)
//				ue.(*url.Error).URL = loc
//				return resp, ue
//			}
//		}
//
//		reqs = append(reqs, req)
//		var err error
//		var didTimeout func() bool
//		if resp, didTimeout, err = c.send(req, deadline); err != nil {
//			// c.send() always closes req.Body
//			reqBodyClosed = true
//			if !deadline.IsZero() && didTimeout() {
//				err = &httpError{
//					// TODO: early in cycle: s/Client.Timeout exceeded/timeout or context cancellation/
//					err:     err.Error() + " (Client.Timeout exceeded while awaiting headers)",
//					timeout: true,
//				}
//			}
//			return nil, uerr(err)
//		}
//
//		var shouldRedirect bool
//		redirectMethod, shouldRedirect, includeBody = redirectBehavior(req.Method, resp, reqs[0])
//		if !shouldRedirect {
//			return resp, nil
//		}
//
//		req.closeBody()
//	}
//}
//
//// makeHeadersCopier makes a function that copies headers from the
//// initial Request, ireq. For every redirect, this function must be called
//// so that it can copy headers into the upcoming Request.
//func (c *Client) makeHeadersCopier(ireq *Request) func(*Request) {
//	// The headers to copy are from the very initial request.
//	// We use a closured callback to keep a reference to these original headers.
//	var (
//		ireqhdr  = cloneOrMakeHeader(ireq.Header)
//		icookies map[string][]*Cookie
//	)
//	if c.Jar != nil && ireq.Header.Get("Cookie") != "" {
//		icookies = make(map[string][]*Cookie)
//		for _, c := range ireq.Cookies() {
//			icookies[c.Name] = append(icookies[c.Name], c)
//		}
//	}
//
//	preq := ireq // The previous request
//	return func(req *Request) {
//		// If Jar is present and there was some initial cookies provided
//		// via the request header, then we may need to alter the initial
//		// cookies as we follow redirects since each redirect may end up
//		// modifying a pre-existing cookie.
//		//
//		// Since cookies already set in the request header do not contain
//		// information about the original domain and path, the logic below
//		// assumes any new set cookies override the original cookie
//		// regardless of domain or path.
//		//
//		// See https://golang.org/issue/17494
//		if c.Jar != nil && icookies != nil {
//			var changed bool
//			resp := req.Response // The response that caused the upcoming redirect
//			for _, c := range resp.Cookies() {
//				if _, ok := icookies[c.Name]; ok {
//					delete(icookies, c.Name)
//					changed = true
//				}
//			}
//			if changed {
//				ireqhdr.Del("Cookie")
//				var ss []string
//				for _, cs := range icookies {
//					for _, c := range cs {
//						ss = append(ss, c.Name+"="+c.Value)
//					}
//				}
//				sort.Strings(ss) // Ensure deterministic headers
//				ireqhdr.Set("Cookie", strings.Join(ss, "; "))
//			}
//		}
//
//		// Copy the initial request's Header values
//		// (at least the safe ones).
//		for k, vv := range ireqhdr {
//			if shouldCopyHeaderOnRedirect(k, preq.URL, req.URL) {
//				req.Header[k] = vv
//			}
//		}
//
//		preq = req // Update previous Request with the current request
//	}
//}
//
//func defaultCheckRedirect(req *Request, via []*Request) error {
//	if len(via) >= 10 {
//		return errors.New("stopped after 10 redirects")
//	}
//	return nil
//}
//
//// Post issues a POST to the specified URL.
////
//// Caller should close resp.Body when done reading from it.
////
//// If the provided body is an io.Closer, it is closed after the
//// request.
////
//// Post is a wrapper around DefaultClient.Post.
////
//// To set custom headers, use NewRequest and DefaultClient.Do.
////
//// See the Client.Do method documentation for details on how redirects
//// are handled.
//func Post(url, contentType string, body io.Reader) (resp *Response, err error) {
//	return DefaultClient.Post(url, contentType, body)
//}
//
//// Post issues a POST to the specified URL.
////
//// Caller should close resp.Body when done reading from it.
////
//// If the provided body is an io.Closer, it is closed after the
//// request.
////
//// To set custom headers, use NewRequest and Client.Do.
////
//// See the Client.Do method documentation for details on how redirects
//// are handled.
//func (c *Client) Post(url, contentType string, body io.Reader) (resp *Response, err error) {
//	req, err := NewRequest("POST", url, body)
//	if err != nil {
//		return nil, err
//	}
//	req.Header.Set("Content-Type", contentType)
//	return c.Do(req)
//}
//
//// PostForm issues a POST to the specified URL, with data's keys and
//// values URL-encoded as the request body.
////
//// The Content-Type header is set to application/x-www-form-urlencoded.
//// To set other headers, use NewRequest and DefaultClient.Do.
////
//// When err is nil, resp always contains a non-nil resp.Body.
//// Caller should close resp.Body when done reading from it.
////
//// PostForm is a wrapper around DefaultClient.PostForm.
////
//// See the Client.Do method documentation for details on how redirects
//// are handled.
//func PostForm(url string, data url.Values) (resp *Response, err error) {
//	return DefaultClient.PostForm(url, data)
//}
//
//// PostForm issues a POST to the specified URL,
//// with data's keys and values URL-encoded as the request body.
////
//// The Content-Type header is set to application/x-www-form-urlencoded.
//// To set other headers, use NewRequest and Client.Do.
////
//// When err is nil, resp always contains a non-nil resp.Body.
//// Caller should close resp.Body when done reading from it.
////
//// See the Client.Do method documentation for details on how redirects
//// are handled.
//func (c *Client) PostForm(url string, data url.Values) (resp *Response, err error) {
//	return c.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
//}
//
//// Head issues a HEAD to the specified URL. If the response is one of
//// the following redirect codes, Head follows the redirect, up to a
//// maximum of 10 redirects:
////
////    301 (Moved Permanently)
////    302 (Found)
////    303 (See Other)
////    307 (Temporary Redirect)
////    308 (Permanent Redirect)
////
//// Head is a wrapper around DefaultClient.Head
//func Head(url string) (resp *Response, err error) {
//	return DefaultClient.Head(url)
//}
//
//// Head issues a HEAD to the specified URL. If the response is one of the
//// following redirect codes, Head follows the redirect after calling the
//// Client's CheckRedirect function:
////
////    301 (Moved Permanently)
////    302 (Found)
////    303 (See Other)
////    307 (Temporary Redirect)
////    308 (Permanent Redirect)
//func (c *Client) Head(url string) (resp *Response, err error) {
//	req, err := NewRequest("HEAD", url, nil)
//	if err != nil {
//		return nil, err
//	}
//	return c.Do(req)
//}
//
//// CloseIdleConnections closes any connections on its Transport which
//// were previously connected from previous requests but are now
//// sitting idle in a "keep-alive" state. It does not interrupt any
//// connections currently in use.
////
//// If the Client's Transport does not have a CloseIdleConnections method
//// then this method does nothing.
//func (c *Client) CloseIdleConnections() {
//	type closeIdler interface {
//		CloseIdleConnections()
//	}
//	if tr, ok := c.transport().(closeIdler); ok {
//		tr.CloseIdleConnections()
//	}
//}
//
//// cancelTimerBody is an io.ReadCloser that wraps rc with two features:
//// 1) on Read error or close, the stop func is called.
//// 2) On Read failure, if reqDidTimeout is true, the error is wrapped and
////    marked as net.Error that hit its timeout.
//type cancelTimerBody struct {
//	stop          func() // stops the time.Timer waiting to cancel the request
//	rc            io.ReadCloser
//	reqDidTimeout func() bool
//}
//
//func (b *cancelTimerBody) Read(p []byte) (n int, err error) {
//	n, err = b.rc.Read(p)
//	if err == nil {
//		return n, nil
//	}
//	b.stop()
//	if err == io.EOF {
//		return n, err
//	}
//	if b.reqDidTimeout() {
//		err = &httpError{
//			err:     err.Error() + " (Client.Timeout or context cancellation while reading body)",
//			timeout: true,
//		}
//	}
//	return n, err
//}
//
//func (b *cancelTimerBody) Close() error {
//	err := b.rc.Close()
//	b.stop()
//	return err
//}
//
//func shouldCopyHeaderOnRedirect(headerKey string, initial, dest *url.URL) bool {
//	switch CanonicalHeaderKey(headerKey) {
//	case "Authorization", "Www-Authenticate", "Cookie", "Cookie2":
//		// Permit sending auth/cookie headers from "foo.com"
//		// to "sub.foo.com".
//
//		// Note that we don't send all cookies to subdomains
//		// automatically. This function is only used for
//		// Cookies set explicitly on the initial outgoing
//		// client request. Cookies automatically added via the
//		// CookieJar mechanism continue to follow each
//		// cookie's scope as set by Set-Cookie. But for
//		// outgoing requests with the Cookie header set
//		// directly, we don't know their scope, so we assume
//		// it's for *.domain.com.
//
//		ihost := canonicalAddr(initial)
//		dhost := canonicalAddr(dest)
//		return isDomainOrSubdomain(dhost, ihost)
//	}
//	// All other headers are copied:
//	return true
//}
//
//// isDomainOrSubdomain reports whether sub is a subdomain (or exact
//// match) of the parent domain.
////
//// Both domains must already be in canonical form.
//func isDomainOrSubdomain(sub, parent string) bool {
//	if sub == parent {
//		return true
//	}
//	// If sub is "foo.example.com" and parent is "example.com",
//	// that means sub must end in "."+parent.
//	// Do it without allocating.
//	if !strings.HasSuffix(sub, parent) {
//		return false
//	}
//	return sub[len(sub)-len(parent)-1] == '.'
//}
//
//func stripPassword(u *url.URL) string {
//	_, passSet := u.User.Password()
//	if passSet {
//		return strings.Replace(u.String(), u.User.String()+"@", u.User.Username()+":***@", 1)
//	}
//	return u.String()
//}
//
//
//// Errors used by the HTTP server.
//var (
//	// ErrBodyNotAllowed is returned by ResponseWriter.Write calls
//	// when the HTTP method or response code does not permit a
//	// body.
//	ErrBodyNotAllowed = errors.New("http: request method or response status code does not allow body")
//
//	// ErrHijacked is returned by ResponseWriter.Write calls when
//	// the underlying connection has been hijacked using the
//	// Hijacker interface. A zero-byte write on a hijacked
//	// connection will return ErrHijacked without any other side
//	// effects.
//	ErrHijacked = errors.New("http: connection has been hijacked")
//
//	// ErrContentLength is returned by ResponseWriter.Write calls
//	// when a Handler set a Content-Length response header with a
//	// declared size and then attempted to write more bytes than
//	// declared.
//	ErrContentLength = errors.New("http: wrote more than the declared Content-Length")
//
//	// Deprecated: ErrWriteAfterFlush is no longer returned by
//	// anything in the net/http package. Callers should not
//	// compare errors against this variable.
//	ErrWriteAfterFlush = errors.New("unused")
//)
//
//// A Handler responds to an HTTP request.
////
//// ServeHTTP should write reply headers and data to the ResponseWriter
//// and then return. Returning signals that the request is finished; it
//// is not valid to use the ResponseWriter or read from the
//// Request.Body after or concurrently with the completion of the
//// ServeHTTP call.
////
//// Depending on the HTTP client software, HTTP protocol version, and
//// any intermediaries between the client and the Go server, it may not
//// be possible to read from the Request.Body after writing to the
//// ResponseWriter. Cautious handlers should read the Request.Body
//// first, and then reply.
////
//// Except for reading the body, handlers should not modify the
//// provided Request.
////
//// If ServeHTTP panics, the server (the caller of ServeHTTP) assumes
//// that the effect of the panic was isolated to the active request.
//// It recovers the panic, logs a stack trace to the server error log,
//// and either closes the network connection or sends an HTTP/2
//// RST_STREAM, depending on the HTTP protocol. To abort a handler so
//// the client sees an interrupted response but the server doesn't log
//// an error, panic with the value ErrAbortHandler.
//type Handler interface {
//	ServeHTTP(ResponseWriter, *Request)
//}
//
//// A ResponseWriter interface is used by an HTTP handler to
//// construct an HTTP response.
////
//// A ResponseWriter may not be used after the Handler.ServeHTTP method
//// has returned.
//type ResponseWriter interface {
//	// Header returns the header map that will be sent by
//	// WriteHeader. The Header map also is the mechanism with which
//	// Handlers can set HTTP trailers.
//	//
//	// Changing the header map after a call to WriteHeader (or
//	// Write) has no effect unless the modified headers are
//	// trailers.
//	//
//	// There are two ways to set Trailers. The preferred way is to
//	// predeclare in the headers which trailers you will later
//	// send by setting the "Trailer" header to the names of the
//	// trailer keys which will come later. In this case, those
//	// keys of the Header map are treated as if they were
//	// trailers. See the example. The second way, for trailer
//	// keys not known to the Handler until after the first Write,
//	// is to prefix the Header map keys with the TrailerPrefix
//	// constant value. See TrailerPrefix.
//	//
//	// To suppress automatic response headers (such as "Date"), set
//	// their value to nil.
//	Header() Header
//
//	// Write writes the data to the connection as part of an HTTP reply.
//	//
//	// If WriteHeader has not yet been called, Write calls
//	// WriteHeader(http.StatusOK) before writing the data. If the Header
//	// does not contain a Content-Type line, Write adds a Content-Type set
//	// to the result of passing the initial 512 bytes of written data to
//	// DetectContentType. Additionally, if the total size of all written
//	// data is under a few KB and there are no Flush calls, the
//	// Content-Length header is added automatically.
//	//
//	// Depending on the HTTP protocol version and the client, calling
//	// Write or WriteHeader may prevent future reads on the
//	// Request.Body. For HTTP/1.x requests, handlers should read any
//	// needed request body data before writing the response. Once the
//	// headers have been flushed (due to either an explicit Flusher.Flush
//	// call or writing enough data to trigger a flush), the request body
//	// may be unavailable. For HTTP/2 requests, the Go HTTP server permits
//	// handlers to continue to read the request body while concurrently
//	// writing the response. However, such behavior may not be supported
//	// by all HTTP/2 clients. Handlers should read before writing if
//	// possible to maximize compatibility.
//	Write([]byte) (int, error)
//
//	// WriteHeader sends an HTTP response header with the provided
//	// status code.
//	//
//	// If WriteHeader is not called explicitly, the first call to Write
//	// will trigger an implicit WriteHeader(http.StatusOK).
//	// Thus explicit calls to WriteHeader are mainly used to
//	// send error codes.
//	//
//	// The provided code must be a valid HTTP 1xx-5xx status code.
//	// Only one header may be written. Go does not currently
//	// support sending user-defined 1xx informational headers,
//	// with the exception of 100-continue response header that the
//	// Server sends automatically when the Request.Body is read.
//	WriteHeader(statusCode int)
//}
//
//// The Flusher interface is implemented by ResponseWriters that allow
//// an HTTP handler to flush buffered data to the client.
////
//// The default HTTP/1.x and HTTP/2 ResponseWriter implementations
//// support Flusher, but ResponseWriter wrappers may not. Handlers
//// should always test for this ability at runtime.
////
//// Note that even for ResponseWriters that support Flush,
//// if the client is connected through an HTTP proxy,
//// the buffered data may not reach the client until the response
//// completes.
//type Flusher interface {
//	// Flush sends any buffered data to the client.
//	Flush()
//}
//
//// The Hijacker interface is implemented by ResponseWriters that allow
//// an HTTP handler to take over the connection.
////
//// The default ResponseWriter for HTTP/1.x connections supports
//// Hijacker, but HTTP/2 connections intentionally do not.
//// ResponseWriter wrappers may also not support Hijacker. Handlers
//// should always test for this ability at runtime.
//type Hijacker interface {
//	// Hijack lets the caller take over the connection.
//	// After a call to Hijack the HTTP server library
//	// will not do anything else with the connection.
//	//
//	// It becomes the caller's responsibility to manage
//	// and close the connection.
//	//
//	// The returned net.Conn may have read or write deadlines
//	// already set, depending on the configuration of the
//	// Server. It is the caller's responsibility to set
//	// or clear those deadlines as needed.
//	//
//	// The returned bufio.Reader may contain unprocessed buffered
//	// data from the client.
//	//
//	// After a call to Hijack, the original Request.Body must not
//	// be used. The original Request's Context remains valid and
//	// is not canceled until the Request's ServeHTTP method
//	// returns.
//	Hijack() (net.Conn, *bufio.ReadWriter, error)
//}
//
//// The CloseNotifier interface is implemented by ResponseWriters which
//// allow detecting when the underlying connection has gone away.
////
//// This mechanism can be used to cancel long operations on the server
//// if the client has disconnected before the response is ready.
////
//// Deprecated: the CloseNotifier interface predates Go's context package.
//// New code should use Request.Context instead.
//type CloseNotifier interface {
//	// CloseNotify returns a channel that receives at most a
//	// single value (true) when the client connection has gone
//	// away.
//	//
//	// CloseNotify may wait to notify until Request.Body has been
//	// fully read.
//	//
//	// After the Handler has returned, there is no guarantee
//	// that the channel receives a value.
//	//
//	// If the protocol is HTTP/1.1 and CloseNotify is called while
//	// processing an idempotent request (such a GET) while
//	// HTTP/1.1 pipelining is in use, the arrival of a subsequent
//	// pipelined request may cause a value to be sent on the
//	// returned channel. In practice HTTP/1.1 pipelining is not
//	// enabled in browsers and not seen often in the wild. If this
//	// is a problem, use HTTP/2 or only use CloseNotify on methods
//	// such as POST.
//	CloseNotify() <-chan bool
//}
//
//var (
//	// ServerContextKey is a context key. It can be used in HTTP
//	// handlers with Context.Value to access the server that
//	// started the handler. The associated value will be of
//	// type *Server.
//	ServerContextKey = &contextKey{"http-server"}
//
//	// LocalAddrContextKey is a context key. It can be used in
//	// HTTP handlers with Context.Value to access the local
//	// address the connection arrived on.
//	// The associated value will be of type net.Addr.
//	LocalAddrContextKey = &contextKey{"local-addr"}
//)
//
//// A conn represents the server side of an HTTP connection.
//type conn struct {
//	// server is the server on which the connection arrived.
//	// Immutable; never nil.
//	server *Server
//
//	// cancelCtx cancels the connection-level context.
//	cancelCtx context.CancelFunc
//
//	// rwc is the underlying network connection.
//	// This is never wrapped by other types and is the value given out
//	// to CloseNotifier callers. It is usually of type *net.TCPConn or
//	// *tls.Conn.
//	rwc net.Conn
//
//	// remoteAddr is rwc.RemoteAddr().String(). It is not populated synchronously
//	// inside the Listener's Accept goroutine, as some implementations block.
//	// It is populated immediately inside the (*conn).serve goroutine.
//	// This is the value of a Handler's (*Request).RemoteAddr.
//	remoteAddr string
//
//	// tlsState is the TLS connection state when using TLS.
//	// nil means not TLS.
//	tlsState *tls.ConnectionState
//
//	// werr is set to the first write error to rwc.
//	// It is set via checkConnErrorWriter{w}, where bufw writes.
//	werr error
//
//	// r is bufr's read source. It's a wrapper around rwc that provides
//	// io.LimitedReader-style limiting (while reading request headers)
//	// and functionality to support CloseNotifier. See *connReader docs.
//	r *connReader
//
//	// bufr reads from r.
//	bufr *bufio.Reader
//
//	// bufw writes to checkConnErrorWriter{c}, which populates werr on error.
//	bufw *bufio.Writer
//
//	// lastMethod is the method of the most recent request
//	// on this connection, if any.
//	lastMethod string
//
//	curReq atomic.Value // of *response (which has a Request in it)
//
//	curState struct{ atomic uint64 } // packed (unixtime<<8|uint8(ConnState))
//
//	// mu guards hijackedv
//	mu sync.Mutex
//
//	// hijackedv is whether this connection has been hijacked
//	// by a Handler with the Hijacker interface.
//	// It is guarded by mu.
//	hijackedv bool
//}
//
//func (c *conn) hijacked() bool {
//	c.mu.Lock()
//	defer c.mu.Unlock()
//	return c.hijackedv
//}
//
//// c.mu must be held.
//func (c *conn) hijackLocked() (rwc net.Conn, buf *bufio.ReadWriter, err error) {
//	if c.hijackedv {
//		return nil, nil, ErrHijacked
//	}
//	c.r.abortPendingRead()
//
//	c.hijackedv = true
//	rwc = c.rwc
//	rwc.SetDeadline(time.Time{})
//
//	buf = bufio.NewReadWriter(c.bufr, bufio.NewWriter(rwc))
//	if c.r.hasByte {
//		if _, err := c.bufr.Peek(c.bufr.Buffered() + 1); err != nil {
//			return nil, nil, fmt.Errorf("unexpected Peek failure reading buffered byte: %v", err)
//		}
//	}
//	c.setState(rwc, StateHijacked)
//	return
//}
//
//// This should be >= 512 bytes for DetectContentType,
//// but otherwise it's somewhat arbitrary.
//const bufferBeforeChunkingSize = 2048
//
//// chunkWriter writes to a response's conn buffer, and is the writer
//// wrapped by the response.bufw buffered writer.
////
//// chunkWriter also is responsible for finalizing the Header, including
//// conditionally setting the Content-Type and setting a Content-Length
//// in cases where the handler's final output is smaller than the buffer
//// size. It also conditionally adds chunk headers, when in chunking mode.
////
//// See the comment above (*response).Write for the entire write flow.
//type chunkWriter struct {
//	res *response
//
//	// header is either nil or a deep clone of res.handlerHeader
//	// at the time of res.writeHeader, if res.writeHeader is
//	// called and extra buffering is being done to calculate
//	// Content-Type and/or Content-Length.
//	header Header
//
//	// wroteHeader tells whether the header's been written to "the
//	// wire" (or rather: w.conn.buf). this is unlike
//	// (*response).wroteHeader, which tells only whether it was
//	// logically written.
//	wroteHeader bool
//
//	// set by the writeHeader method:
//	chunking bool // using chunked transfer encoding for reply body
//}
//
//var (
//	crlf       = []byte("\r\n")
//	colonSpace = []byte(": ")
//)
//
//func (cw *chunkWriter) Write(p []byte) (n int, err error) {
//	if !cw.wroteHeader {
//		cw.writeHeader(p)
//	}
//	if cw.res.req.Method == "HEAD" {
//		// Eat writes.
//		return len(p), nil
//	}
//	if cw.chunking {
//		_, err = fmt.Fprintf(cw.res.conn.bufw, "%x\r\n", len(p))
//		if err != nil {
//			cw.res.conn.rwc.Close()
//			return
//		}
//	}
//	n, err = cw.res.conn.bufw.Write(p)
//	if cw.chunking && err == nil {
//		_, err = cw.res.conn.bufw.Write(crlf)
//	}
//	if err != nil {
//		cw.res.conn.rwc.Close()
//	}
//	return
//}
//
//func (cw *chunkWriter) flush() {
//	if !cw.wroteHeader {
//		cw.writeHeader(nil)
//	}
//	cw.res.conn.bufw.Flush()
//}
//
//func (cw *chunkWriter) close() {
//	if !cw.wroteHeader {
//		cw.writeHeader(nil)
//	}
//	if cw.chunking {
//		bw := cw.res.conn.bufw // conn's bufio writer
//		// zero chunk to mark EOF
//		bw.WriteString("0\r\n")
//		if trailers := cw.res.finalTrailers(); trailers != nil {
//			trailers.Write(bw) // the writer handles noting errors
//		}
//		// final blank line after the trailers (whether
//		// present or not)
//		bw.WriteString("\r\n")
//	}
//}
//
//// A response represents the server side of an HTTP response.
//type response struct {
//	conn             *conn
//	req              *Request // request for this response
//	reqBody          io.ReadCloser
//	cancelCtx        context.CancelFunc // when ServeHTTP exits
//	wroteHeader      bool               // reply header has been (logically) written
//	wroteContinue    bool               // 100 Continue response was written
//	wants10KeepAlive bool               // HTTP/1.0 w/ Connection "keep-alive"
//	wantsClose       bool               // HTTP request has Connection "close"
//
//	w  *bufio.Writer // buffers output in chunks to chunkWriter
//	cw chunkWriter
//
//	// handlerHeader is the Header that Handlers get access to,
//	// which may be retained and mutated even after WriteHeader.
//	// handlerHeader is copied into cw.header at WriteHeader
//	// time, and privately mutated thereafter.
//	handlerHeader Header
//	calledHeader  bool // handler accessed handlerHeader via Header
//
//	written       int64 // number of bytes written in body
//	contentLength int64 // explicitly-declared Content-Length; or -1
//	status        int   // status code passed to WriteHeader
//
//	// close connection after this reply.  set on request and
//	// updated after response from handler if there's a
//	// "Connection: keep-alive" response header and a
//	// Content-Length.
//	closeAfterReply bool
//
//	// requestBodyLimitHit is set by requestTooLarge when
//	// maxBytesReader hits its max size. It is checked in
//	// WriteHeader, to make sure we don't consume the
//	// remaining request body to try to advance to the next HTTP
//	// request. Instead, when this is set, we stop reading
//	// subsequent requests on this connection and stop reading
//	// input from it.
//	requestBodyLimitHit bool
//
//	// trailers are the headers to be sent after the handler
//	// finishes writing the body. This field is initialized from
//	// the Trailer response header when the response header is
//	// written.
//	trailers []string
//
//	handlerDone atomicBool // set true when the handler exits
//
//	// Buffers for Date, Content-Length, and status code
//	dateBuf   [len(TimeFormat)]byte
//	clenBuf   [10]byte
//	statusBuf [3]byte
//
//	// closeNotifyCh is the channel returned by CloseNotify.
//	// TODO(bradfitz): this is currently (for Go 1.8) always
//	// non-nil. Make this lazily-created again as it used to be?
//	closeNotifyCh  chan bool
//	didCloseNotify int32 // atomic (only 0->1 winner should send)
//}
//
//// TrailerPrefix is a magic prefix for ResponseWriter.Header map keys
//// that, if present, signals that the map entry is actually for
//// the response trailers, and not the response headers. The prefix
//// is stripped after the ServeHTTP call finishes and the values are
//// sent in the trailers.
////
//// This mechanism is intended only for trailers that are not known
//// prior to the headers being written. If the set of trailers is fixed
//// or known before the header is written, the normal Go trailers mechanism
//// is preferred:
////    https://golang.org/pkg/net/http/#ResponseWriter
////    https://golang.org/pkg/net/http/#example_ResponseWriter_trailers
//const TrailerPrefix = "Trailer:"
//
//// finalTrailers is called after the Handler exits and returns a non-nil
//// value if the Handler set any trailers.
//func (w *response) finalTrailers() Header {
//	var t Header
//	for k, vv := range w.handlerHeader {
//		if strings.HasPrefix(k, TrailerPrefix) {
//			if t == nil {
//				t = make(Header)
//			}
//			t[strings.TrimPrefix(k, TrailerPrefix)] = vv
//		}
//	}
//	for _, k := range w.trailers {
//		if t == nil {
//			t = make(Header)
//		}
//		for _, v := range w.handlerHeader[k] {
//			t.Add(k, v)
//		}
//	}
//	return t
//}
//
//type atomicBool int32
//
//func (b *atomicBool) isSet() bool { return atomic.LoadInt32((*int32)(b)) != 0 }
//func (b *atomicBool) setTrue()    { atomic.StoreInt32((*int32)(b), 1) }
//
//// declareTrailer is called for each Trailer header when the
//// response header is written. It notes that a header will need to be
//// written in the trailers at the end of the response.
//func (w *response) declareTrailer(k string) {
//	k = CanonicalHeaderKey(k)
//	if !httpguts.ValidTrailerHeader(k) {
//		// Forbidden by RFC 7230, section 4.1.2
//		return
//	}
//	w.trailers = append(w.trailers, k)
//}
//
//// requestTooLarge is called by maxBytesReader when too much input has
//// been read from the client.
//func (w *response) requestTooLarge() {
//	w.closeAfterReply = true
//	w.requestBodyLimitHit = true
//	if !w.wroteHeader {
//		w.Header().Set("Connection", "close")
//	}
//}
//
//// needsSniff reports whether a Content-Type still needs to be sniffed.
//func (w *response) needsSniff() bool {
//	_, haveType := w.handlerHeader["Content-Type"]
//	return !w.cw.wroteHeader && !haveType && w.written < sniffLen
//}
//
//// writerOnly hides an io.Writer value's optional ReadFrom method
//// from io.Copy.
//type writerOnly struct {
//	io.Writer
//}
//
//func srcIsRegularFile(src io.Reader) (isRegular bool, err error) {
//	switch v := src.(type) {
//	case *os.File:
//		fi, err := v.Stat()
//		if err != nil {
//			return false, err
//		}
//		return fi.Mode().IsRegular(), nil
//	case *io.LimitedReader:
//		return srcIsRegularFile(v.R)
//	default:
//		return
//	}
//}
//
//// ReadFrom is here to optimize copying from an *os.File regular file
//// to a *net.TCPConn with sendfile.
//func (w *response) ReadFrom(src io.Reader) (n int64, err error) {
//	// Our underlying w.conn.rwc is usually a *TCPConn (with its
//	// own ReadFrom method). If not, or if our src isn't a regular
//	// file, just fall back to the normal copy method.
//	rf, ok := w.conn.rwc.(io.ReaderFrom)
//	regFile, err := srcIsRegularFile(src)
//	if err != nil {
//		return 0, err
//	}
//	if !ok || !regFile {
//		bufp := copyBufPool.Get().(*[]byte)
//		defer copyBufPool.Put(bufp)
//		return io.CopyBuffer(writerOnly{w}, src, *bufp)
//	}
//
//	// sendfile path:
//
//	if !w.wroteHeader {
//		w.WriteHeader(StatusOK)
//	}
//
//	if w.needsSniff() {
//		n0, err := io.Copy(writerOnly{w}, io.LimitReader(src, sniffLen))
//		n += n0
//		if err != nil {
//			return n, err
//		}
//	}
//
//	w.w.Flush()  // get rid of any previous writes
//	w.cw.flush() // make sure Header is written; flush data to rwc
//
//	// Now that cw has been flushed, its chunking field is guaranteed initialized.
//	if !w.cw.chunking && w.bodyAllowed() {
//		n0, err := rf.ReadFrom(src)
//		n += n0
//		w.written += n0
//		return n, err
//	}
//
//	n0, err := io.Copy(writerOnly{w}, src)
//	n += n0
//	return n, err
//}
//
//// debugServerConnections controls whether all server connections are wrapped
//// with a verbose logging wrapper.
//const debugServerConnections = false
//
//// Create new connection from rwc.
//func (srv *Server) newConn(rwc net.Conn) *conn {
//	c := &conn{
//		server: srv,
//		rwc:    rwc,
//	}
//	if debugServerConnections {
//		c.rwc = newLoggingConn("server", c.rwc)
//	}
//	return c
//}
//
//type readResult struct {
//	n   int
//	err error
//	b   byte // byte read, if n == 1
//}
//
//// connReader is the io.Reader wrapper used by *conn. It combines a
//// selectively-activated io.LimitedReader (to bound request header
//// read sizes) with support for selectively keeping an io.Reader.Read
//// call blocked in a background goroutine to wait for activity and
//// trigger a CloseNotifier channel.
//type connReader struct {
//	conn *conn
//
//	mu      sync.Mutex // guards following
//	hasByte bool
//	byteBuf [1]byte
//	cond    *sync.Cond
//	inRead  bool
//	aborted bool  // set true before conn.rwc deadline is set to past
//	remain  int64 // bytes remaining
//}
//
//func (cr *connReader) lock() {
//	cr.mu.Lock()
//	if cr.cond == nil {
//		cr.cond = sync.NewCond(&cr.mu)
//	}
//}
//
//func (cr *connReader) unlock() { cr.mu.Unlock() }
//
//func (cr *connReader) startBackgroundRead() {
//	cr.lock()
//	defer cr.unlock()
//	if cr.inRead {
//		panic("invalid concurrent Body.Read call")
//	}
//	if cr.hasByte {
//		return
//	}
//	cr.inRead = true
//	cr.conn.rwc.SetReadDeadline(time.Time{})
//	go cr.backgroundRead()
//}
//
//func (cr *connReader) backgroundRead() {
//	n, err := cr.conn.rwc.Read(cr.byteBuf[:])
//	cr.lock()
//	if n == 1 {
//		cr.hasByte = true
//		// We were past the end of the previous request's body already
//		// (since we wouldn't be in a background read otherwise), so
//		// this is a pipelined HTTP request. Prior to Go 1.11 we used to
//		// send on the CloseNotify channel and cancel the context here,
//		// but the behavior was documented as only "may", and we only
//		// did that because that's how CloseNotify accidentally behaved
//		// in very early Go releases prior to context support. Once we
//		// added context support, people used a Handler's
//		// Request.Context() and passed it along. Having that context
//		// cancel on pipelined HTTP requests caused problems.
//		// Fortunately, almost nothing uses HTTP/1.x pipelining.
//		// Unfortunately, apt-get does, or sometimes does.
//		// New Go 1.11 behavior: don't fire CloseNotify or cancel
//		// contexts on pipelined requests. Shouldn't affect people, but
//		// fixes cases like Issue 23921. This does mean that a client
//		// closing their TCP connection after sending a pipelined
//		// request won't cancel the context, but we'll catch that on any
//		// write failure (in checkConnErrorWriter.Write).
//		// If the server never writes, yes, there are still contrived
//		// server & client behaviors where this fails to ever cancel the
//		// context, but that's kinda why HTTP/1.x pipelining died
//		// anyway.
//	}
//	if ne, ok := err.(net.Error); ok && cr.aborted && ne.Timeout() {
//		// Ignore this error. It's the expected error from
//		// another goroutine calling abortPendingRead.
//	} else if err != nil {
//		cr.handleReadError(err)
//	}
//	cr.aborted = false
//	cr.inRead = false
//	cr.unlock()
//	cr.cond.Broadcast()
//}
//
//func (cr *connReader) abortPendingRead() {
//	cr.lock()
//	defer cr.unlock()
//	if !cr.inRead {
//		return
//	}
//	cr.aborted = true
//	cr.conn.rwc.SetReadDeadline(aLongTimeAgo)
//	for cr.inRead {
//		cr.cond.Wait()
//	}
//	cr.conn.rwc.SetReadDeadline(time.Time{})
//}
//
//func (cr *connReader) setReadLimit(remain int64) { cr.remain = remain }
//func (cr *connReader) setInfiniteReadLimit()     { cr.remain = maxInt64 }
//func (cr *connReader) hitReadLimit() bool        { return cr.remain <= 0 }
//
//// handleReadError is called whenever a Read from the client returns a
//// non-nil error.
////
//// The provided non-nil err is almost always io.EOF or a "use of
//// closed network connection". In any case, the error is not
//// particularly interesting, except perhaps for debugging during
//// development. Any error means the connection is dead and we should
//// down its context.
////
//// It may be called from multiple goroutines.
//func (cr *connReader) handleReadError(_ error) {
//	cr.conn.cancelCtx()
//	cr.closeNotify()
//}
//
//// may be called from multiple goroutines.
//func (cr *connReader) closeNotify() {
//	res, _ := cr.conn.curReq.Load().(*response)
//	if res != nil && atomic.CompareAndSwapInt32(&res.didCloseNotify, 0, 1) {
//		res.closeNotifyCh <- true
//	}
//}
//
//func (cr *connReader) Read(p []byte) (n int, err error) {
//	cr.lock()
//	if cr.inRead {
//		cr.unlock()
//		if cr.conn.hijacked() {
//			panic("invalid Body.Read call. After hijacked, the original Request must not be used")
//		}
//		panic("invalid concurrent Body.Read call")
//	}
//	if cr.hitReadLimit() {
//		cr.unlock()
//		return 0, io.EOF
//	}
//	if len(p) == 0 {
//		cr.unlock()
//		return 0, nil
//	}
//	if int64(len(p)) > cr.remain {
//		p = p[:cr.remain]
//	}
//	if cr.hasByte {
//		p[0] = cr.byteBuf[0]
//		cr.hasByte = false
//		cr.unlock()
//		return 1, nil
//	}
//	cr.inRead = true
//	cr.unlock()
//	n, err = cr.conn.rwc.Read(p)
//
//	cr.lock()
//	cr.inRead = false
//	if err != nil {
//		cr.handleReadError(err)
//	}
//	cr.remain -= int64(n)
//	cr.unlock()
//
//	cr.cond.Broadcast()
//	return n, err
//}
//
//var (
//	bufioReaderPool   sync.Pool
//	bufioWriter2kPool sync.Pool
//	bufioWriter4kPool sync.Pool
//)
//
//var copyBufPool = sync.Pool{
//	New: func() interface{} {
//		b := make([]byte, 32*1024)
//		return &b
//	},
//}
//
//func bufioWriterPool(size int) *sync.Pool {
//	switch size {
//	case 2 << 10:
//		return &bufioWriter2kPool
//	case 4 << 10:
//		return &bufioWriter4kPool
//	}
//	return nil
//}
//
//func newBufioReader(r io.Reader) *bufio.Reader {
//	if v := bufioReaderPool.Get(); v != nil {
//		br := v.(*bufio.Reader)
//		br.Reset(r)
//		return br
//	}
//	// Note: if this reader size is ever changed, update
//	// TestHandlerBodyClose's assumptions.
//	return bufio.NewReader(r)
//}
//
//func putBufioReader(br *bufio.Reader) {
//	br.Reset(nil)
//	bufioReaderPool.Put(br)
//}
//
//func newBufioWriterSize(w io.Writer, size int) *bufio.Writer {
//	pool := bufioWriterPool(size)
//	if pool != nil {
//		if v := pool.Get(); v != nil {
//			bw := v.(*bufio.Writer)
//			bw.Reset(w)
//			return bw
//		}
//	}
//	return bufio.NewWriterSize(w, size)
//}
//
//func putBufioWriter(bw *bufio.Writer) {
//	bw.Reset(nil)
//	if pool := bufioWriterPool(bw.Available()); pool != nil {
//		pool.Put(bw)
//	}
//}
//
//// DefaultMaxHeaderBytes is the maximum permitted size of the headers
//// in an HTTP request.
//// This can be overridden by setting Server.MaxHeaderBytes.
//const DefaultMaxHeaderBytes = 1 << 20 // 1 MB
//
//func (srv *Server) maxHeaderBytes() int {
//	if srv.MaxHeaderBytes > 0 {
//		return srv.MaxHeaderBytes
//	}
//	return DefaultMaxHeaderBytes
//}
//
//func (srv *Server) initialReadLimitSize() int64 {
//	return int64(srv.maxHeaderBytes()) + 4096 // bufio slop
//}
//
//// wrapper around io.ReadCloser which on first read, sends an
//// HTTP/1.1 100 Continue header
//type expectContinueReader struct {
//	resp       *response
//	readCloser io.ReadCloser
//	closed     bool
//	sawEOF     bool
//}
//
//func (ecr *expectContinueReader) Read(p []byte) (n int, err error) {
//	if ecr.closed {
//		return 0, ErrBodyReadAfterClose
//	}
//	if !ecr.resp.wroteContinue && !ecr.resp.conn.hijacked() {
//		ecr.resp.wroteContinue = true
//		ecr.resp.conn.bufw.WriteString("HTTP/1.1 100 Continue\r\n\r\n")
//		ecr.resp.conn.bufw.Flush()
//	}
//	n, err = ecr.readCloser.Read(p)
//	if err == io.EOF {
//		ecr.sawEOF = true
//	}
//	return
//}
//
//func (ecr *expectContinueReader) Close() error {
//	ecr.closed = true
//	return ecr.readCloser.Close()
//}
//
//// TimeFormat is the time format to use when generating times in HTTP
//// headers. It is like time.RFC1123 but hard-codes GMT as the time
//// zone. The time being formatted must be in UTC for Format to
//// generate the correct format.
////
//// For parsing this time format, see ParseTime.
//const TimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"
//
//// appendTime is a non-allocating version of []byte(t.UTC().Format(TimeFormat))
//func appendTime(b []byte, t time.Time) []byte {
//	const days = "SunMonTueWedThuFriSat"
//	const months = "JanFebMarAprMayJunJulAugSepOctNovDec"
//
//	t = t.UTC()
//	yy, mm, dd := t.Date()
//	hh, mn, ss := t.Clock()
//	day := days[3*t.Weekday():]
//	mon := months[3*(mm-1):]
//
//	return append(b,
//		day[0], day[1], day[2], ',', ' ',
//		byte('0'+dd/10), byte('0'+dd%10), ' ',
//		mon[0], mon[1], mon[2], ' ',
//		byte('0'+yy/1000), byte('0'+(yy/100)%10), byte('0'+(yy/10)%10), byte('0'+yy%10), ' ',
//		byte('0'+hh/10), byte('0'+hh%10), ':',
//		byte('0'+mn/10), byte('0'+mn%10), ':',
//		byte('0'+ss/10), byte('0'+ss%10), ' ',
//		'G', 'M', 'T')
//}
//
//var errTooLarge = errors.New("http: request too large")
//
//// Read next request from connection.
//func (c *conn) readRequest(ctx context.Context) (w *response, err error) {
//	if c.hijacked() {
//		return nil, ErrHijacked
//	}
//
//	var (
//		wholeReqDeadline time.Time // or zero if none
//		hdrDeadline      time.Time // or zero if none
//	)
//	t0 := time.Now()
//	if d := c.server.readHeaderTimeout(); d != 0 {
//		hdrDeadline = t0.Add(d)
//	}
//	if d := c.server.ReadTimeout; d != 0 {
//		wholeReqDeadline = t0.Add(d)
//	}
//	c.rwc.SetReadDeadline(hdrDeadline)
//	if d := c.server.WriteTimeout; d != 0 {
//		defer func() {
//			c.rwc.SetWriteDeadline(time.Now().Add(d))
//		}()
//	}
//
//	c.r.setReadLimit(c.server.initialReadLimitSize())
//	if c.lastMethod == "POST" {
//		// RFC 7230 section 3 tolerance for old buggy clients.
//		peek, _ := c.bufr.Peek(4) // ReadRequest will get err below
//		c.bufr.Discard(numLeadingCRorLF(peek))
//	}
//	req, err := readRequest(c.bufr, keepHostHeader)
//	if err != nil {
//		if c.r.hitReadLimit() {
//			return nil, errTooLarge
//		}
//		return nil, err
//	}
//
//	if !http1ServerSupportsRequest(req) {
//		return nil, badRequestError("unsupported protocol version")
//	}
//
//	c.lastMethod = req.Method
//	c.r.setInfiniteReadLimit()
//
//	hosts, haveHost := req.Header["Host"]
//	isH2Upgrade := req.isH2Upgrade()
//	if req.ProtoAtLeast(1, 1) && (!haveHost || len(hosts) == 0) && !isH2Upgrade && req.Method != "CONNECT" {
//		return nil, badRequestError("missing required Host header")
//	}
//	if len(hosts) > 1 {
//		return nil, badRequestError("too many Host headers")
//	}
//	if len(hosts) == 1 && !httpguts.ValidHostHeader(hosts[0]) {
//		return nil, badRequestError("malformed Host header")
//	}
//	for k, vv := range req.Header {
//		if !httpguts.ValidHeaderFieldName(k) {
//			return nil, badRequestError("invalid header name")
//		}
//		for _, v := range vv {
//			if !httpguts.ValidHeaderFieldValue(v) {
//				return nil, badRequestError("invalid header value")
//			}
//		}
//	}
//	delete(req.Header, "Host")
//
//	ctx, cancelCtx := context.WithCancel(ctx)
//	req.ctx = ctx
//	req.RemoteAddr = c.remoteAddr
//	req.TLS = c.tlsState
//	if body, ok := req.Body.(*body); ok {
//		body.doEarlyClose = true
//	}
//
//	// Adjust the read deadline if necessary.
//	if !hdrDeadline.Equal(wholeReqDeadline) {
//		c.rwc.SetReadDeadline(wholeReqDeadline)
//	}
//
//	w = &response{
//		conn:          c,
//		cancelCtx:     cancelCtx,
//		req:           req,
//		reqBody:       req.Body,
//		handlerHeader: make(Header),
//		contentLength: -1,
//		closeNotifyCh: make(chan bool, 1),
//
//		// We populate these ahead of time so we're not
//		// reading from req.Header after their Handler starts
//		// and maybe mutates it (Issue 14940)
//		wants10KeepAlive: req.wantsHttp10KeepAlive(),
//		wantsClose:       req.wantsClose(),
//	}
//	if isH2Upgrade {
//		w.closeAfterReply = true
//	}
//	w.cw.res = w
//	w.w = newBufioWriterSize(&w.cw, bufferBeforeChunkingSize)
//	return w, nil
//}
//
//// http1ServerSupportsRequest reports whether Go's HTTP/1.x server
//// supports the given request.
//func http1ServerSupportsRequest(req *Request) bool {
//	if req.ProtoMajor == 1 {
//		return true
//	}
//	// Accept "PRI * HTTP/2.0" upgrade requests, so Handlers can
//	// wire up their own HTTP/2 upgrades.
//	if req.ProtoMajor == 2 && req.ProtoMinor == 0 &&
//		req.Method == "PRI" && req.RequestURI == "*" {
//		return true
//	}
//	// Reject HTTP/0.x, and all other HTTP/2+ requests (which
//	// aren't encoded in ASCII anyway).
//	return false
//}
//
//func (w *response) Header() Header {
//	if w.cw.header == nil && w.wroteHeader && !w.cw.wroteHeader {
//		// Accessing the header between logically writing it
//		// and physically writing it means we need to allocate
//		// a clone to snapshot the logically written state.
//		w.cw.header = w.handlerHeader.Clone()
//	}
//	w.calledHeader = true
//	return w.handlerHeader
//}
//
//// maxPostHandlerReadBytes is the max number of Request.Body bytes not
//// consumed by a handler that the server will read from the client
//// in order to keep a connection alive. If there are more bytes than
//// this then the server to be paranoid instead sends a "Connection:
//// close" response.
////
//// This number is approximately what a typical machine's TCP buffer
//// size is anyway.  (if we have the bytes on the machine, we might as
//// well read them)
//const maxPostHandlerReadBytes = 256 << 10
//
//func checkWriteHeaderCode(code int) {
//	// Issue 22880: require valid WriteHeader status codes.
//	// For now we only enforce that it's three digits.
//	// In the future we might block things over 599 (600 and above aren't defined
//	// at https://httpwg.org/specs/rfc7231.html#status.codes)
//	// and we might block under 200 (once we have more mature 1xx support).
//	// But for now any three digits.
//	//
//	// We used to send "HTTP/1.1 000 0" on the wire in responses but there's
//	// no equivalent bogus thing we can realistically send in HTTP/2,
//	// so we'll consistently panic instead and help people find their bugs
//	// early. (We can't return an error from WriteHeader even if we wanted to.)
//	if code < 100 || code > 999 {
//		panic(fmt.Sprintf("invalid WriteHeader code %v", code))
//	}
//}
//
//// relevantCaller searches the call stack for the first function outside of net/http.
//// The purpose of this function is to provide more helpful error messages.
//func relevantCaller() runtime.Frame {
//	pc := make([]uintptr, 16)
//	n := runtime.Callers(1, pc)
//	frames := runtime.CallersFrames(pc[:n])
//	var frame runtime.Frame
//	for {
//		frame, more := frames.Next()
//		if !strings.HasPrefix(frame.Function, "net/http.") {
//			return frame
//		}
//		if !more {
//			break
//		}
//	}
//	return frame
//}
//
//func (w *response) WriteHeader(code int) {
//	if w.conn.hijacked() {
//		caller := relevantCaller()
//		w.conn.server.logf("http: response.WriteHeader on hijacked connection from %s (%s:%d)", caller.Function, path.Base(caller.File), caller.Line)
//		return
//	}
//	if w.wroteHeader {
//		caller := relevantCaller()
//		w.conn.server.logf("http: superfluous response.WriteHeader call from %s (%s:%d)", caller.Function, path.Base(caller.File), caller.Line)
//		return
//	}
//	checkWriteHeaderCode(code)
//	w.wroteHeader = true
//	w.status = code
//
//	if w.calledHeader && w.cw.header == nil {
//		w.cw.header = w.handlerHeader.Clone()
//	}
//
//	if cl := w.handlerHeader.get("Content-Length"); cl != "" {
//		v, err := strconv.ParseInt(cl, 10, 64)
//		if err == nil && v >= 0 {
//			w.contentLength = v
//		} else {
//			w.conn.server.logf("http: invalid Content-Length of %q", cl)
//			w.handlerHeader.Del("Content-Length")
//		}
//	}
//}
//
//// extraHeader is the set of headers sometimes added by chunkWriter.writeHeader.
//// This type is used to avoid extra allocations from cloning and/or populating
//// the response Header map and all its 1-element slices.
//type extraHeader struct {
//	contentType      string
//	connection       string
//	transferEncoding string
//	date             []byte // written if not nil
//	contentLength    []byte // written if not nil
//}
//
//// Sorted the same as extraHeader.Write's loop.
//var extraHeaderKeys = [][]byte{
//	[]byte("Content-Type"),
//	[]byte("Connection"),
//	[]byte("Transfer-Encoding"),
//}
//
//var (
//	headerContentLength = []byte("Content-Length: ")
//	headerDate          = []byte("Date: ")
//)
//
//// Write writes the headers described in h to w.
////
//// This method has a value receiver, despite the somewhat large size
//// of h, because it prevents an allocation. The escape analysis isn't
//// smart enough to realize this function doesn't mutate h.
//func (h extraHeader) Write(w *bufio.Writer) {
//	if h.date != nil {
//		w.Write(headerDate)
//		w.Write(h.date)
//		w.Write(crlf)
//	}
//	if h.contentLength != nil {
//		w.Write(headerContentLength)
//		w.Write(h.contentLength)
//		w.Write(crlf)
//	}
//	for i, v := range []string{h.contentType, h.connection, h.transferEncoding} {
//		if v != "" {
//			w.Write(extraHeaderKeys[i])
//			w.Write(colonSpace)
//			w.WriteString(v)
//			w.Write(crlf)
//		}
//	}
//}
//
//// writeHeader finalizes the header sent to the client and writes it
//// to cw.res.conn.bufw.
////
//// p is not written by writeHeader, but is the first chunk of the body
//// that will be written. It is sniffed for a Content-Type if none is
//// set explicitly. It's also used to set the Content-Length, if the
//// total body size was small and the handler has already finished
//// running.
//func (cw *chunkWriter) writeHeader(p []byte) {
//	if cw.wroteHeader {
//		return
//	}
//	cw.wroteHeader = true
//
//	w := cw.res
//	keepAlivesEnabled := w.conn.server.doKeepAlives()
//	isHEAD := w.req.Method == "HEAD"
//
//	// header is written out to w.conn.buf below. Depending on the
//	// state of the handler, we either own the map or not. If we
//	// don't own it, the exclude map is created lazily for
//	// WriteSubset to remove headers. The setHeader struct holds
//	// headers we need to add.
//	header := cw.header
//	owned := header != nil
//	if !owned {
//		header = w.handlerHeader
//	}
//	var excludeHeader map[string]bool
//	delHeader := func(key string) {
//		if owned {
//			header.Del(key)
//			return
//		}
//		if _, ok := header[key]; !ok {
//			return
//		}
//		if excludeHeader == nil {
//			excludeHeader = make(map[string]bool)
//		}
//		excludeHeader[key] = true
//	}
//	var setHeader extraHeader
//
//	// Don't write out the fake "Trailer:foo" keys. See TrailerPrefix.
//	trailers := false
//	for k := range cw.header {
//		if strings.HasPrefix(k, TrailerPrefix) {
//			if excludeHeader == nil {
//				excludeHeader = make(map[string]bool)
//			}
//			excludeHeader[k] = true
//			trailers = true
//		}
//	}
//	for _, v := range cw.header["Trailer"] {
//		trailers = true
//		foreachHeaderElement(v, cw.res.declareTrailer)
//	}
//
//	te := header.get("Transfer-Encoding")
//	hasTE := te != ""
//
//	// If the handler is done but never sent a Content-Length
//	// response header and this is our first (and last) write, set
//	// it, even to zero. This helps HTTP/1.0 clients keep their
//	// "keep-alive" connections alive.
//	// Exceptions: 304/204/1xx responses never get Content-Length, and if
//	// it was a HEAD request, we don't know the difference between
//	// 0 actual bytes and 0 bytes because the handler noticed it
//	// was a HEAD request and chose not to write anything. So for
//	// HEAD, the handler should either write the Content-Length or
//	// write non-zero bytes. If it's actually 0 bytes and the
//	// handler never looked at the Request.Method, we just don't
//	// send a Content-Length header.
//	// Further, we don't send an automatic Content-Length if they
//	// set a Transfer-Encoding, because they're generally incompatible.
//	if w.handlerDone.isSet() && !trailers && !hasTE && bodyAllowedForStatus(w.status) && header.get("Content-Length") == "" && (!isHEAD || len(p) > 0) {
//		w.contentLength = int64(len(p))
//		setHeader.contentLength = strconv.AppendInt(cw.res.clenBuf[:0], int64(len(p)), 10)
//	}
//
//	// If this was an HTTP/1.0 request with keep-alive and we sent a
//	// Content-Length back, we can make this a keep-alive response ...
//	if w.wants10KeepAlive && keepAlivesEnabled {
//		sentLength := header.get("Content-Length") != ""
//		if sentLength && header.get("Connection") == "keep-alive" {
//			w.closeAfterReply = false
//		}
//	}
//
//	// Check for an explicit (and valid) Content-Length header.
//	hasCL := w.contentLength != -1
//
//	if w.wants10KeepAlive && (isHEAD || hasCL || !bodyAllowedForStatus(w.status)) {
//		_, connectionHeaderSet := header["Connection"]
//		if !connectionHeaderSet {
//			setHeader.connection = "keep-alive"
//		}
//	} else if !w.req.ProtoAtLeast(1, 1) || w.wantsClose {
//		w.closeAfterReply = true
//	}
//
//	if header.get("Connection") == "close" || !keepAlivesEnabled {
//		w.closeAfterReply = true
//	}
//
//	// If the client wanted a 100-continue but we never sent it to
//	// them (or, more strictly: we never finished reading their
//	// request body), don't reuse this connection because it's now
//	// in an unknown state: we might be sending this response at
//	// the same time the client is now sending its request body
//	// after a timeout.  (Some HTTP clients send Expect:
//	// 100-continue but knowing that some servers don't support
//	// it, the clients set a timer and send the body later anyway)
//	// If we haven't seen EOF, we can't skip over the unread body
//	// because we don't know if the next bytes on the wire will be
//	// the body-following-the-timer or the subsequent request.
//	// See Issue 11549.
//	if ecr, ok := w.req.Body.(*expectContinueReader); ok && !ecr.sawEOF {
//		w.closeAfterReply = true
//	}
//
//	// Per RFC 2616, we should consume the request body before
//	// replying, if the handler hasn't already done so. But we
//	// don't want to do an unbounded amount of reading here for
//	// DoS reasons, so we only try up to a threshold.
//	// TODO(bradfitz): where does RFC 2616 say that? See Issue 15527
//	// about HTTP/1.x Handlers concurrently reading and writing, like
//	// HTTP/2 handlers can do. Maybe this code should be relaxed?
//	if w.req.ContentLength != 0 && !w.closeAfterReply {
//		var discard, tooBig bool
//
//		switch bdy := w.req.Body.(type) {
//		case *expectContinueReader:
//			if bdy.resp.wroteContinue {
//				discard = true
//			}
//		case *body:
//			bdy.mu.Lock()
//			switch {
//			case bdy.closed:
//				if !bdy.sawEOF {
//					// Body was closed in handler with non-EOF error.
//					w.closeAfterReply = true
//				}
//			case bdy.unreadDataSizeLocked() >= maxPostHandlerReadBytes:
//				tooBig = true
//			default:
//				discard = true
//			}
//			bdy.mu.Unlock()
//		default:
//			discard = true
//		}
//
//		if discard {
//			_, err := io.CopyN(ioutil.Discard, w.reqBody, maxPostHandlerReadBytes+1)
//			switch err {
//			case nil:
//				// There must be even more data left over.
//				tooBig = true
//			case ErrBodyReadAfterClose:
//				// Body was already consumed and closed.
//			case io.EOF:
//				// The remaining body was just consumed, close it.
//				err = w.reqBody.Close()
//				if err != nil {
//					w.closeAfterReply = true
//				}
//			default:
//				// Some other kind of error occurred, like a read timeout, or
//				// corrupt chunked encoding. In any case, whatever remains
//				// on the wire must not be parsed as another HTTP request.
//				w.closeAfterReply = true
//			}
//		}
//
//		if tooBig {
//			w.requestTooLarge()
//			delHeader("Connection")
//			setHeader.connection = "close"
//		}
//	}
//
//	code := w.status
//	if bodyAllowedForStatus(code) {
//		// If no content type, apply sniffing algorithm to body.
//		_, haveType := header["Content-Type"]
//
//		// If the Content-Encoding was set and is non-blank,
//		// we shouldn't sniff the body. See Issue 31753.
//		ce := header.Get("Content-Encoding")
//		hasCE := len(ce) > 0
//		if !hasCE && !haveType && !hasTE && len(p) > 0 {
//			setHeader.contentType = DetectContentType(p)
//		}
//	} else {
//		for _, k := range suppressedHeaders(code) {
//			delHeader(k)
//		}
//	}
//
//	if !header.has("Date") {
//		setHeader.date = appendTime(cw.res.dateBuf[:0], time.Now())
//	}
//
//	if hasCL && hasTE && te != "identity" {
//		// TODO: return an error if WriteHeader gets a return parameter
//		// For now just ignore the Content-Length.
//		w.conn.server.logf("http: WriteHeader called with both Transfer-Encoding of %q and a Content-Length of %d",
//			te, w.contentLength)
//		delHeader("Content-Length")
//		hasCL = false
//	}
//
//	if w.req.Method == "HEAD" || !bodyAllowedForStatus(code) {
//		// do nothing
//	} else if code == StatusNoContent {
//		delHeader("Transfer-Encoding")
//	} else if hasCL {
//		delHeader("Transfer-Encoding")
//	} else if w.req.ProtoAtLeast(1, 1) {
//		// HTTP/1.1 or greater: Transfer-Encoding has been set to identity, and no
//		// content-length has been provided. The connection must be closed after the
//		// reply is written, and no chunking is to be done. This is the setup
//		// recommended in the Server-Sent Events candidate recommendation 11,
//		// section 8.
//		if hasTE && te == "identity" {
//			cw.chunking = false
//			w.closeAfterReply = true
//		} else {
//			// HTTP/1.1 or greater: use chunked transfer encoding
//			// to avoid closing the connection at EOF.
//			cw.chunking = true
//			setHeader.transferEncoding = "chunked"
//			if hasTE && te == "chunked" {
//				// We will send the chunked Transfer-Encoding header later.
//				delHeader("Transfer-Encoding")
//			}
//		}
//	} else {
//		// HTTP version < 1.1: cannot do chunked transfer
//		// encoding and we don't know the Content-Length so
//		// signal EOF by closing connection.
//		w.closeAfterReply = true
//		delHeader("Transfer-Encoding") // in case already set
//	}
//
//	// Cannot use Content-Length with non-identity Transfer-Encoding.
//	if cw.chunking {
//		delHeader("Content-Length")
//	}
//	if !w.req.ProtoAtLeast(1, 0) {
//		return
//	}
//
//	if w.closeAfterReply && (!keepAlivesEnabled || !hasToken(cw.header.get("Connection"), "close")) {
//		delHeader("Connection")
//		if w.req.ProtoAtLeast(1, 1) {
//			setHeader.connection = "close"
//		}
//	}
//
//	writeStatusLine(w.conn.bufw, w.req.ProtoAtLeast(1, 1), code, w.statusBuf[:])
//	cw.header.WriteSubset(w.conn.bufw, excludeHeader)
//	setHeader.Write(w.conn.bufw)
//	w.conn.bufw.Write(crlf)
//}
//
//// foreachHeaderElement splits v according to the "#rule" construction
//// in RFC 7230 section 7 and calls fn for each non-empty element.
//func foreachHeaderElement(v string, fn func(string)) {
//	v = textproto.TrimString(v)
//	if v == "" {
//		return
//	}
//	if !strings.Contains(v, ",") {
//		fn(v)
//		return
//	}
//	for _, f := range strings.Split(v, ",") {
//		if f = textproto.TrimString(f); f != "" {
//			fn(f)
//		}
//	}
//}
//
//// writeStatusLine writes an HTTP/1.x Status-Line (RFC 7230 Section 3.1.2)
//// to bw. is11 is whether the HTTP request is HTTP/1.1. false means HTTP/1.0.
//// code is the response status code.
//// scratch is an optional scratch buffer. If it has at least capacity 3, it's used.
//func writeStatusLine(bw *bufio.Writer, is11 bool, code int, scratch []byte) {
//	if is11 {
//		bw.WriteString("HTTP/1.1 ")
//	} else {
//		bw.WriteString("HTTP/1.0 ")
//	}
//	if text, ok := statusText[code]; ok {
//		bw.Write(strconv.AppendInt(scratch[:0], int64(code), 10))
//		bw.WriteByte(' ')
//		bw.WriteString(text)
//		bw.WriteString("\r\n")
//	} else {
//		// don't worry about performance
//		fmt.Fprintf(bw, "%03d status code %d\r\n", code, code)
//	}
//}
//
//// bodyAllowed reports whether a Write is allowed for this response type.
//// It's illegal to call this before the header has been flushed.
//func (w *response) bodyAllowed() bool {
//	if !w.wroteHeader {
//		panic("")
//	}
//	return bodyAllowedForStatus(w.status)
//}
//
//// The Life Of A Write is like this:
////
//// Handler starts. No header has been sent. The handler can either
//// write a header, or just start writing. Writing before sending a header
//// sends an implicitly empty 200 OK header.
////
//// If the handler didn't declare a Content-Length up front, we either
//// go into chunking mode or, if the handler finishes running before
//// the chunking buffer size, we compute a Content-Length and send that
//// in the header instead.
////
//// Likewise, if the handler didn't set a Content-Type, we sniff that
//// from the initial chunk of output.
////
//// The Writers are wired together like:
////
//// 1. *response (the ResponseWriter) ->
//// 2. (*response).w, a *bufio.Writer of bufferBeforeChunkingSize bytes
//// 3. chunkWriter.Writer (whose writeHeader finalizes Content-Length/Type)
////    and which writes the chunk headers, if needed.
//// 4. conn.buf, a bufio.Writer of default (4kB) bytes, writing to ->
//// 5. checkConnErrorWriter{c}, which notes any non-nil error on Write
////    and populates c.werr with it if so. but otherwise writes to:
//// 6. the rwc, the net.Conn.
////
//// TODO(bradfitz): short-circuit some of the buffering when the
//// initial header contains both a Content-Type and Content-Length.
//// Also short-circuit in (1) when the header's been sent and not in
//// chunking mode, writing directly to (4) instead, if (2) has no
//// buffered data. More generally, we could short-circuit from (1) to
//// (3) even in chunking mode if the write size from (1) is over some
//// threshold and nothing is in (2).  The answer might be mostly making
//// bufferBeforeChunkingSize smaller and having bufio's fast-paths deal
//// with this instead.
//func (w *response) Write(data []byte) (n int, err error) {
//	return w.write(len(data), data, "")
//}
//
//func (w *response) WriteString(data string) (n int, err error) {
//	return w.write(len(data), nil, data)
//}
//
//// either dataB or dataS is non-zero.
//func (w *response) write(lenData int, dataB []byte, dataS string) (n int, err error) {
//	if w.conn.hijacked() {
//		if lenData > 0 {
//			caller := relevantCaller()
//			w.conn.server.logf("http: response.Write on hijacked connection from %s (%s:%d)", caller.Function, path.Base(caller.File), caller.Line)
//		}
//		return 0, ErrHijacked
//	}
//	if !w.wroteHeader {
//		w.WriteHeader(StatusOK)
//	}
//	if lenData == 0 {
//		return 0, nil
//	}
//	if !w.bodyAllowed() {
//		return 0, ErrBodyNotAllowed
//	}
//
//	w.written += int64(lenData) // ignoring errors, for errorKludge
//	if w.contentLength != -1 && w.written > w.contentLength {
//		return 0, ErrContentLength
//	}
//	if dataB != nil {
//		return w.w.Write(dataB)
//	} else {
//		return w.w.WriteString(dataS)
//	}
//}
//
//func (w *response) finishRequest() {
//	w.handlerDone.setTrue()
//
//	if !w.wroteHeader {
//		w.WriteHeader(StatusOK)
//	}
//
//	w.w.Flush()
//	putBufioWriter(w.w)
//	w.cw.close()
//	w.conn.bufw.Flush()
//
//	w.conn.r.abortPendingRead()
//
//	// Close the body (regardless of w.closeAfterReply) so we can
//	// re-use its bufio.Reader later safely.
//	w.reqBody.Close()
//
//	if w.req.MultipartForm != nil {
//		w.req.MultipartForm.RemoveAll()
//	}
//}
//
//// shouldReuseConnection reports whether the underlying TCP connection can be reused.
//// It must only be called after the handler is done executing.
//func (w *response) shouldReuseConnection() bool {
//	if w.closeAfterReply {
//		// The request or something set while executing the
//		// handler indicated we shouldn't reuse this
//		// connection.
//		return false
//	}
//
//	if w.req.Method != "HEAD" && w.contentLength != -1 && w.bodyAllowed() && w.contentLength != w.written {
//		// Did not write enough. Avoid getting out of sync.
//		return false
//	}
//
//	// There was some error writing to the underlying connection
//	// during the request, so don't re-use this conn.
//	if w.conn.werr != nil {
//		return false
//	}
//
//	if w.closedRequestBodyEarly() {
//		return false
//	}
//
//	return true
//}
//
//func (w *response) closedRequestBodyEarly() bool {
//	body, ok := w.req.Body.(*body)
//	return ok && body.didEarlyClose()
//}
//
//func (w *response) Flush() {
//	if !w.wroteHeader {
//		w.WriteHeader(StatusOK)
//	}
//	w.w.Flush()
//	w.cw.flush()
//}
//
//func (c *conn) finalFlush() {
//	if c.bufr != nil {
//		// Steal the bufio.Reader (~4KB worth of memory) and its associated
//		// reader for a future connection.
//		putBufioReader(c.bufr)
//		c.bufr = nil
//	}
//
//	if c.bufw != nil {
//		c.bufw.Flush()
//		// Steal the bufio.Writer (~4KB worth of memory) and its associated
//		// writer for a future connection.
//		putBufioWriter(c.bufw)
//		c.bufw = nil
//	}
//}
//
//// Close the connection.
//func (c *conn) close() {
//	c.finalFlush()
//	c.rwc.Close()
//}
//
//// rstAvoidanceDelay is the amount of time we sleep after closing the
//// write side of a TCP connection before closing the entire socket.
//// By sleeping, we increase the chances that the client sees our FIN
//// and processes its final data before they process the subsequent RST
//// from closing a connection with known unread data.
//// This RST seems to occur mostly on BSD systems. (And Windows?)
//// This timeout is somewhat arbitrary (~latency around the planet).
//const rstAvoidanceDelay = 500 * time.Millisecond
//
//type closeWriter interface {
//	CloseWrite() error
//}
//
//var _ closeWriter = (*net.TCPConn)(nil)
//
//// closeWrite flushes any outstanding data and sends a FIN packet (if
//// client is connected via TCP), signalling that we're done. We then
//// pause for a bit, hoping the client processes it before any
//// subsequent RST.
////
//// See https://golang.org/issue/3595
//func (c *conn) closeWriteAndWait() {
//	c.finalFlush()
//	if tcp, ok := c.rwc.(closeWriter); ok {
//		tcp.CloseWrite()
//	}
//	time.Sleep(rstAvoidanceDelay)
//}
//
//// validNextProto reports whether the proto is not a blacklisted ALPN
//// protocol name. Empty and built-in protocol types are blacklisted
//// and can't be overridden with alternate implementations.
//func validNextProto(proto string) bool {
//	switch proto {
//	case "", "http/1.1", "http/1.0":
//		return false
//	}
//	return true
//}
//
//func (c *conn) setState(nc net.Conn, state ConnState) {
//	srv := c.server
//	switch state {
//	case StateNew:
//		srv.trackConn(c, true)
//	case StateHijacked, StateClosed:
//		srv.trackConn(c, false)
//	}
//	if state > 0xff || state < 0 {
//		panic("internal error")
//	}
//	packedState := uint64(time.Now().Unix()<<8) | uint64(state)
//	atomic.StoreUint64(&c.curState.atomic, packedState)
//	if hook := srv.ConnState; hook != nil {
//		hook(nc, state)
//	}
//}
//
//func (c *conn) getState() (state ConnState, unixSec int64) {
//	packedState := atomic.LoadUint64(&c.curState.atomic)
//	return ConnState(packedState & 0xff), int64(packedState >> 8)
//}
//
//// badRequestError is a literal string (used by in the server in HTML,
//// unescaped) to tell the user why their request was bad. It should
//// be plain text without user info or other embedded errors.
//type badRequestError string
//
//func (e badRequestError) Error() string { return "Bad Request: " + string(e) }
//
//// ErrAbortHandler is a sentinel panic value to abort a handler.
//// While any panic from ServeHTTP aborts the response to the client,
//// panicking with ErrAbortHandler also suppresses logging of a stack
//// trace to the server's error log.
//var ErrAbortHandler = errors.New("net/http: abort Handler")
//
//// isCommonNetReadError reports whether err is a common error
//// encountered during reading a request off the network when the
//// client has gone away or had its read fail somehow. This is used to
//// determine which logs are interesting enough to log about.
//func isCommonNetReadError(err error) bool {
//	if err == io.EOF {
//		return true
//	}
//	if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
//		return true
//	}
//	if oe, ok := err.(*net.OpError); ok && oe.Op == "read" {
//		return true
//	}
//	return false
//}
//
//// Serve a new connection.
//func (c *conn) serve(ctx context.Context) {
//	c.remoteAddr = c.rwc.RemoteAddr().String()
//	ctx = context.WithValue(ctx, LocalAddrContextKey, c.rwc.LocalAddr())
//	defer func() {
//		if err := recover(); err != nil && err != ErrAbortHandler {
//			const size = 64 << 10
//			buf := make([]byte, size)
//			buf = buf[:runtime.Stack(buf, false)]
//			c.server.logf("http: panic serving %v: %v\n%s", c.remoteAddr, err, buf)
//		}
//		if !c.hijacked() {
//			c.close()
//			c.setState(c.rwc, StateClosed)
//		}
//	}()
//
//	if tlsConn, ok := c.rwc.(*tls.Conn); ok {
//		if d := c.server.ReadTimeout; d != 0 {
//			c.rwc.SetReadDeadline(time.Now().Add(d))
//		}
//		if d := c.server.WriteTimeout; d != 0 {
//			c.rwc.SetWriteDeadline(time.Now().Add(d))
//		}
//		if err := tlsConn.Handshake(); err != nil {
//			// If the handshake failed due to the client not speaking
//			// TLS, assume they're speaking plaintext HTTP and write a
//			// 400 response on the TLS conn's underlying net.Conn.
//			if re, ok := err.(tls.RecordHeaderError); ok && re.Conn != nil && tlsRecordHeaderLooksLikeHTTP(re.RecordHeader) {
//				io.WriteString(re.Conn, "HTTP/1.0 400 Bad Request\r\n\r\nClient sent an HTTP request to an HTTPS server.\n")
//				re.Conn.Close()
//				return
//			}
//			c.server.logf("http: TLS handshake error from %s: %v", c.rwc.RemoteAddr(), err)
//			return
//		}
//		c.tlsState = new(tls.ConnectionState)
//		*c.tlsState = tlsConn.ConnectionState()
//		if proto := c.tlsState.NegotiatedProtocol; validNextProto(proto) {
//			if fn := c.server.TLSNextProto[proto]; fn != nil {
//				h := initALPNRequest{ctx, tlsConn, serverHandler{c.server}}
//				fn(c.server, tlsConn, h)
//			}
//			return
//		}
//	}
//
//	// HTTP/1.x from here on.
//
//	ctx, cancelCtx := context.WithCancel(ctx)
//	c.cancelCtx = cancelCtx
//	defer cancelCtx()
//
//	c.r = &connReader{conn: c}
//	c.bufr = newBufioReader(c.r)
//	c.bufw = newBufioWriterSize(checkConnErrorWriter{c}, 4<<10)
//
//	for {
//		w, err := c.readRequest(ctx)
//		if c.r.remain != c.server.initialReadLimitSize() {
//			// If we read any bytes off the wire, we're active.
//			c.setState(c.rwc, StateActive)
//		}
//		if err != nil {
//			const errorHeaders = "\r\nContent-Type: text/plain; charset=utf-8\r\nConnection: close\r\n\r\n"
//
//			switch {
//			case err == errTooLarge:
//				// Their HTTP client may or may not be
//				// able to read this if we're
//				// responding to them and hanging up
//				// while they're still writing their
//				// request. Undefined behavior.
//				const publicErr = "431 Request Header Fields Too Large"
//				fmt.Fprintf(c.rwc, "HTTP/1.1 "+publicErr+errorHeaders+publicErr)
//				c.closeWriteAndWait()
//				return
//
//			case isUnsupportedTEError(err):
//				// Respond as per RFC 7230 Section 3.3.1 which says,
//				//      A server that receives a request message with a
//				//      transfer coding it does not understand SHOULD
//				//      respond with 501 (Unimplemented).
//				code := StatusNotImplemented
//
//				// We purposefully aren't echoing back the transfer-encoding's value,
//				// so as to mitigate the risk of cross side scripting by an attacker.
//				fmt.Fprintf(c.rwc, "HTTP/1.1 %d %s%sUnsupported transfer encoding", code, StatusText(code), errorHeaders)
//				return
//
//			case isCommonNetReadError(err):
//				return // don't reply
//
//			default:
//				publicErr := "400 Bad Request"
//				if v, ok := err.(badRequestError); ok {
//					publicErr = publicErr + ": " + string(v)
//				}
//
//				fmt.Fprintf(c.rwc, "HTTP/1.1 "+publicErr+errorHeaders+publicErr)
//				return
//			}
//		}
//
//		// Expect 100 Continue support
//		req := w.req
//		if req.expectsContinue() {
//			if req.ProtoAtLeast(1, 1) && req.ContentLength != 0 {
//				// Wrap the Body reader with one that replies on the connection
//				req.Body = &expectContinueReader{readCloser: req.Body, resp: w}
//			}
//		} else if req.Header.get("Expect") != "" {
//			w.sendExpectationFailed()
//			return
//		}
//
//		c.curReq.Store(w)
//
//		if requestBodyRemains(req.Body) {
//			registerOnHitEOF(req.Body, w.conn.r.startBackgroundRead)
//		} else {
//			w.conn.r.startBackgroundRead()
//		}
//
//		// HTTP cannot have multiple simultaneous active requests.[*]
//		// Until the server replies to this request, it can't read another,
//		// so we might as well run the handler in this goroutine.
//		// [*] Not strictly true: HTTP pipelining. We could let them all process
//		// in parallel even if their responses need to be serialized.
//		// But we're not going to implement HTTP pipelining because it
//		// was never deployed in the wild and the answer is HTTP/2.
//		serverHandler{c.server}.ServeHTTP(w, w.req)
//		w.cancelCtx()
//		if c.hijacked() {
//			return
//		}
//		w.finishRequest()
//		if !w.shouldReuseConnection() {
//			if w.requestBodyLimitHit || w.closedRequestBodyEarly() {
//				c.closeWriteAndWait()
//			}
//			return
//		}
//		c.setState(c.rwc, StateIdle)
//		c.curReq.Store((*response)(nil))
//
//		if !w.conn.server.doKeepAlives() {
//			// We're in shutdown mode. We might've replied
//			// to the user without "Connection: close" and
//			// they might think they can send another
//			// request, but such is life with HTTP/1.1.
//			return
//		}
//
//		if d := c.server.idleTimeout(); d != 0 {
//			c.rwc.SetReadDeadline(time.Now().Add(d))
//			if _, err := c.bufr.Peek(4); err != nil {
//				return
//			}
//		}
//		c.rwc.SetReadDeadline(time.Time{})
//	}
//}
//
//func (w *response) sendExpectationFailed() {
//	// TODO(bradfitz): let ServeHTTP handlers handle
//	// requests with non-standard expectation[s]? Seems
//	// theoretical at best, and doesn't fit into the
//	// current ServeHTTP model anyway. We'd need to
//	// make the ResponseWriter an optional
//	// "ExpectReplier" interface or something.
//	//
//	// For now we'll just obey RFC 7231 5.1.1 which says
//	// "A server that receives an Expect field-value other
//	// than 100-continue MAY respond with a 417 (Expectation
//	// Failed) status code to indicate that the unexpected
//	// expectation cannot be met."
//	w.Header().Set("Connection", "close")
//	w.WriteHeader(StatusExpectationFailed)
//	w.finishRequest()
//}
//
//// Hijack implements the Hijacker.Hijack method. Our response is both a ResponseWriter
//// and a Hijacker.
//func (w *response) Hijack() (rwc net.Conn, buf *bufio.ReadWriter, err error) {
//	if w.handlerDone.isSet() {
//		panic("net/http: Hijack called after ServeHTTP finished")
//	}
//	if w.wroteHeader {
//		w.cw.flush()
//	}
//
//	c := w.conn
//	c.mu.Lock()
//	defer c.mu.Unlock()
//
//	// Release the bufioWriter that writes to the chunk writer, it is not
//	// used after a connection has been hijacked.
//	rwc, buf, err = c.hijackLocked()
//	if err == nil {
//		putBufioWriter(w.w)
//		w.w = nil
//	}
//	return rwc, buf, err
//}
//
//func (w *response) CloseNotify() <-chan bool {
//	if w.handlerDone.isSet() {
//		panic("net/http: CloseNotify called after ServeHTTP finished")
//	}
//	return w.closeNotifyCh
//}
//
//func registerOnHitEOF(rc io.ReadCloser, fn func()) {
//	switch v := rc.(type) {
//	case *expectContinueReader:
//		registerOnHitEOF(v.readCloser, fn)
//	case *body:
//		v.registerOnHitEOF(fn)
//	default:
//		panic("unexpected type " + fmt.Sprintf("%T", rc))
//	}
//}
//
//// requestBodyRemains reports whether future calls to Read
//// on rc might yield more data.
//func requestBodyRemains(rc io.ReadCloser) bool {
//	if rc == NoBody {
//		return false
//	}
//	switch v := rc.(type) {
//	case *expectContinueReader:
//		return requestBodyRemains(v.readCloser)
//	case *body:
//		return v.bodyRemains()
//	default:
//		panic("unexpected type " + fmt.Sprintf("%T", rc))
//	}
//}
//
//// The HandlerFunc type is an adapter to allow the use of
//// ordinary functions as HTTP handlers. If f is a function
//// with the appropriate signature, HandlerFunc(f) is a
//// Handler that calls f.
//type HandlerFunc func(ResponseWriter, *Request)
//
//// ServeHTTP calls f(w, r).
//func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
//	f(w, r)
//}
//
//// Helper handlers
//
//// Error replies to the request with the specified error message and HTTP code.
//// It does not otherwise end the request; the caller should ensure no further
//// writes are done to w.
//// The error message should be plain text.
//func Error(w ResponseWriter, error string, code int) {
//	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
//	w.Header().Set("X-Content-Type-Options", "nosniff")
//	w.WriteHeader(code)
//	fmt.Fprintln(w, error)
//}
//
//// NotFound replies to the request with an HTTP 404 not found error.
//func NotFound(w ResponseWriter, r *Request) { Error(w, "404 page not found", StatusNotFound) }
//
//// NotFoundHandler returns a simple request handler
//// that replies to each request with a ``404 page not found'' reply.
//func NotFoundHandler() Handler { return HandlerFunc(NotFound) }
//
//// StripPrefix returns a handler that serves HTTP requests
//// by removing the given prefix from the request URL's Path
//// and invoking the handler h. StripPrefix handles a
//// request for a path that doesn't begin with prefix by
//// replying with an HTTP 404 not found error.
//func StripPrefix(prefix string, h Handler) Handler {
//	if prefix == "" {
//		return h
//	}
//	return HandlerFunc(func(w ResponseWriter, r *Request) {
//		if p := strings.TrimPrefix(r.URL.Path, prefix); len(p) < len(r.URL.Path) {
//			r2 := new(Request)
//			*r2 = *r
//			r2.URL = new(url.URL)
//			*r2.URL = *r.URL
//			r2.URL.Path = p
//			h.ServeHTTP(w, r2)
//		} else {
//			NotFound(w, r)
//		}
//	})
//}
//
//// Redirect replies to the request with a redirect to url,
//// which may be a path relative to the request path.
////
//// The provided code should be in the 3xx range and is usually
//// StatusMovedPermanently, StatusFound or StatusSeeOther.
////
//// If the Content-Type header has not been set, Redirect sets it
//// to "text/html; charset=utf-8" and writes a small HTML body.
//// Setting the Content-Type header to any value, including nil,
//// disables that behavior.
//func Redirect(w ResponseWriter, r *Request, url string, code int) {
//	if u, err := urlpkg.Parse(url); err == nil {
//		// If url was relative, make its path absolute by
//		// combining with request path.
//		// The client would probably do this for us,
//		// but doing it ourselves is more reliable.
//		// See RFC 7231, section 7.1.2
//		if u.Scheme == "" && u.Host == "" {
//			oldpath := r.URL.Path
//			if oldpath == "" { // should not happen, but avoid a crash if it does
//				oldpath = "/"
//			}
//
//			// no leading http://server
//			if url == "" || url[0] != '/' {
//				// make relative path absolute
//				olddir, _ := path.Split(oldpath)
//				url = olddir + url
//			}
//
//			var query string
//			if i := strings.Index(url, "?"); i != -1 {
//				url, query = url[:i], url[i:]
//			}
//
//			// clean up but preserve trailing slash
//			trailing := strings.HasSuffix(url, "/")
//			url = path.Clean(url)
//			if trailing && !strings.HasSuffix(url, "/") {
//				url += "/"
//			}
//			url += query
//		}
//	}
//
//	h := w.Header()
//
//	// RFC 7231 notes that a short HTML body is usually included in
//	// the response because older user agents may not understand 301/307.
//	// Do it only if the request didn't already have a Content-Type header.
//	_, hadCT := h["Content-Type"]
//
//	h.Set("Location", hexEscapeNonASCII(url))
//	if !hadCT && (r.Method == "GET" || r.Method == "HEAD") {
//		h.Set("Content-Type", "text/html; charset=utf-8")
//	}
//	w.WriteHeader(code)
//
//	// Shouldn't send the body for POST or HEAD; that leaves GET.
//	if !hadCT && r.Method == "GET" {
//		body := "<a href=\"" + htmlEscape(url) + "\">" + statusText[code] + "</a>.\n"
//		fmt.Fprintln(w, body)
//	}
//}
//
//var htmlReplacer = strings.NewReplacer(
//	"&", "&amp;",
//	"<", "&lt;",
//	">", "&gt;",
//	// "&#34;" is shorter than "&quot;".
//	`"`, "&#34;",
//	// "&#39;" is shorter than "&apos;" and apos was not in HTML until HTML5.
//	"'", "&#39;",
//)
//
//func htmlEscape(s string) string {
//	return htmlReplacer.Replace(s)
//}
//
//// Redirect to a fixed URL
//type redirectHandler struct {
//	url  string
//	code int
//}
//
//func (rh *redirectHandler) ServeHTTP(w ResponseWriter, r *Request) {
//	Redirect(w, r, rh.url, rh.code)
//}
//
//// RedirectHandler returns a request handler that redirects
//// each request it receives to the given url using the given
//// status code.
////
//// The provided code should be in the 3xx range and is usually
//// StatusMovedPermanently, StatusFound or StatusSeeOther.
//func RedirectHandler(url string, code int) Handler {
//	return &redirectHandler{url, code}
//}
//
//// ServeMux is an HTTP request multiplexer.
//// It matches the URL of each incoming request against a list of registered
//// patterns and calls the handler for the pattern that
//// most closely matches the URL.
////
//// Patterns name fixed, rooted paths, like "/favicon.ico",
//// or rooted subtrees, like "/images/" (note the trailing slash).
//// Longer patterns take precedence over shorter ones, so that
//// if there are handlers registered for both "/images/"
//// and "/images/thumbnails/", the latter handler will be
//// called for paths beginning "/images/thumbnails/" and the
//// former will receive requests for any other paths in the
//// "/images/" subtree.
////
//// Note that since a pattern ending in a slash names a rooted subtree,
//// the pattern "/" matches all paths not matched by other registered
//// patterns, not just the URL with Path == "/".
////
//// If a subtree has been registered and a request is received naming the
//// subtree root without its trailing slash, ServeMux redirects that
//// request to the subtree root (adding the trailing slash). This behavior can
//// be overridden with a separate registration for the path without
//// the trailing slash. For example, registering "/images/" causes ServeMux
//// to redirect a request for "/images" to "/images/", unless "/images" has
//// been registered separately.
////
//// Patterns may optionally begin with a host name, restricting matches to
//// URLs on that host only. Host-specific patterns take precedence over
//// general patterns, so that a handler might register for the two patterns
//// "/codesearch" and "codesearch.google.com/" without also taking over
//// requests for "http://www.google.com/".
////
//// ServeMux also takes care of sanitizing the URL request path and the Host
//// header, stripping the port number and redirecting any request containing . or
//// .. elements or repeated slashes to an equivalent, cleaner URL.
//type ServeMux struct {
//	mu    sync.RWMutex
//	m     map[string]muxEntry
//	es    []muxEntry // slice of entries sorted from longest to shortest.
//	hosts bool       // whether any patterns contain hostnames
//}
//
//type muxEntry struct {
//	h       Handler
//	pattern string
//}
//
//// NewServeMux allocates and returns a new ServeMux.
//func NewServeMux() *ServeMux { return new(ServeMux) }
//
//// DefaultServeMux is the default ServeMux used by Serve.
//var DefaultServeMux = &defaultServeMux
//
//var defaultServeMux ServeMux
//
//// cleanPath returns the canonical path for p, eliminating . and .. elements.
//func cleanPath(p string) string {
//	if p == "" {
//		return "/"
//	}
//	if p[0] != '/' {
//		p = "/" + p
//	}
//	np := path.Clean(p)
//	// path.Clean removes trailing slash except for root;
//	// put the trailing slash back if necessary.
//	if p[len(p)-1] == '/' && np != "/" {
//		// Fast path for common case of p being the string we want:
//		if len(p) == len(np)+1 && strings.HasPrefix(p, np) {
//			np = p
//		} else {
//			np += "/"
//		}
//	}
//	return np
//}
//
//// stripHostPort returns h without any trailing ":<port>".
//func stripHostPort(h string) string {
//	// If no port on host, return unchanged
//	if strings.IndexByte(h, ':') == -1 {
//		return h
//	}
//	host, _, err := net.SplitHostPort(h)
//	if err != nil {
//		return h // on error, return unchanged
//	}
//	return host
//}
//
//// Find a handler on a handler map given a path string.
//// Most-specific (longest) pattern wins.
//func (mux *ServeMux) match(path string) (h Handler, pattern string) {
//	// Check for exact match first.
//	v, ok := mux.m[path]
//	if ok {
//		return v.h, v.pattern
//	}
//
//	// Check for longest valid match.  mux.es contains all patterns
//	// that end in / sorted from longest to shortest.
//	for _, e := range mux.es {
//		if strings.HasPrefix(path, e.pattern) {
//			return e.h, e.pattern
//		}
//	}
//	return nil, ""
//}
//
//// redirectToPathSlash determines if the given path needs appending "/" to it.
//// This occurs when a handler for path + "/" was already registered, but
//// not for path itself. If the path needs appending to, it creates a new
//// URL, setting the path to u.Path + "/" and returning true to indicate so.
//func (mux *ServeMux) redirectToPathSlash(host, path string, u *url.URL) (*url.URL, bool) {
//	mux.mu.RLock()
//	shouldRedirect := mux.shouldRedirectRLocked(host, path)
//	mux.mu.RUnlock()
//	if !shouldRedirect {
//		return u, false
//	}
//	path = path + "/"
//	u = &url.URL{Path: path, RawQuery: u.RawQuery}
//	return u, true
//}
//
//// shouldRedirectRLocked reports whether the given path and host should be redirected to
//// path+"/". This should happen if a handler is registered for path+"/" but
//// not path -- see comments at ServeMux.
//func (mux *ServeMux) shouldRedirectRLocked(host, path string) bool {
//	p := []string{path, host + path}
//
//	for _, c := range p {
//		if _, exist := mux.m[c]; exist {
//			return false
//		}
//	}
//
//	n := len(path)
//	if n == 0 {
//		return false
//	}
//	for _, c := range p {
//		if _, exist := mux.m[c+"/"]; exist {
//			return path[n-1] != '/'
//		}
//	}
//
//	return false
//}
//
//// Handler returns the handler to use for the given request,
//// consulting r.Method, r.Host, and r.URL.Path. It always returns
//// a non-nil handler. If the path is not in its canonical form, the
//// handler will be an internally-generated handler that redirects
//// to the canonical path. If the host contains a port, it is ignored
//// when matching handlers.
////
//// The path and host are used unchanged for CONNECT requests.
////
//// Handler also returns the registered pattern that matches the
//// request or, in the case of internally-generated redirects,
//// the pattern that will match after following the redirect.
////
//// If there is no registered handler that applies to the request,
//// Handler returns a ``page not found'' handler and an empty pattern.
//func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string) {
//
//	// CONNECT requests are not canonicalized.
//	if r.Method == "CONNECT" {
//		// If r.URL.Path is /tree and its handler is not registered,
//		// the /tree -> /tree/ redirect applies to CONNECT requests
//		// but the path canonicalization does not.
//		if u, ok := mux.redirectToPathSlash(r.URL.Host, r.URL.Path, r.URL); ok {
//			return RedirectHandler(u.String(), StatusMovedPermanently), u.Path
//		}
//
//		return mux.handler(r.Host, r.URL.Path)
//	}
//
//	// All other requests have any port stripped and path cleaned
//	// before passing to mux.handler.
//	host := stripHostPort(r.Host)
//	path := cleanPath(r.URL.Path)
//
//	// If the given path is /tree and its handler is not registered,
//	// redirect for /tree/.
//	if u, ok := mux.redirectToPathSlash(host, path, r.URL); ok {
//		return RedirectHandler(u.String(), StatusMovedPermanently), u.Path
//	}
//
//	if path != r.URL.Path {
//		_, pattern = mux.handler(host, path)
//		url := *r.URL
//		url.Path = path
//		return RedirectHandler(url.String(), StatusMovedPermanently), pattern
//	}
//
//	return mux.handler(host, r.URL.Path)
//}
//
//// handler is the main implementation of Handler.
//// The path is known to be in canonical form, except for CONNECT methods.
//func (mux *ServeMux) handler(host, path string) (h Handler, pattern string) {
//	mux.mu.RLock()
//	defer mux.mu.RUnlock()
//
//	// Host-specific pattern takes precedence over generic ones
//	if mux.hosts {
//		h, pattern = mux.match(host + path)
//	}
//	if h == nil {
//		h, pattern = mux.match(path)
//	}
//	if h == nil {
//		h, pattern = NotFoundHandler(), ""
//	}
//	return
//}
//
//// ServeHTTP dispatches the request to the handler whose
//// pattern most closely matches the request URL.
//func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
//	if r.RequestURI == "*" {
//		if r.ProtoAtLeast(1, 1) {
//			w.Header().Set("Connection", "close")
//		}
//		w.WriteHeader(StatusBadRequest)
//		return
//	}
//	h, _ := mux.Handler(r)
//	h.ServeHTTP(w, r)
//}
//
//// Handle registers the handler for the given pattern.
//// If a handler already exists for pattern, Handle panics.
//func (mux *ServeMux) Handle(pattern string, handler Handler) {
//	mux.mu.Lock()
//	defer mux.mu.Unlock()
//
//	if pattern == "" {
//		panic("http: invalid pattern")
//	}
//	if handler == nil {
//		panic("http: nil handler")
//	}
//	if _, exist := mux.m[pattern]; exist {
//		panic("http: multiple registrations for " + pattern)
//	}
//
//	if mux.m == nil {
//		mux.m = make(map[string]muxEntry)
//	}
//	e := muxEntry{h: handler, pattern: pattern}
//	mux.m[pattern] = e
//	if pattern[len(pattern)-1] == '/' {
//		mux.es = appendSorted(mux.es, e)
//	}
//
//	if pattern[0] != '/' {
//		mux.hosts = true
//	}
//}
//
//func appendSorted(es []muxEntry, e muxEntry) []muxEntry {
//	n := len(es)
//	i := sort.Search(n, func(i int) bool {
//		return len(es[i].pattern) < len(e.pattern)
//	})
//	if i == n {
//		return append(es, e)
//	}
//	// we now know that i points at where we want to insert
//	es = append(es, muxEntry{}) // try to grow the slice in place, any entry works.
//	copy(es[i+1:], es[i:])      // Move shorter entries down
//	es[i] = e
//	return es
//}
//
//// HandleFunc registers the handler function for the given pattern.
//func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
//	if handler == nil {
//		panic("http: nil handler")
//	}
//	mux.Handle(pattern, HandlerFunc(handler))
//}
//
//// Handle registers the handler for the given pattern
//// in the DefaultServeMux.
//// The documentation for ServeMux explains how patterns are matched.
//func Handle(pattern string, handler Handler) { DefaultServeMux.Handle(pattern, handler) }
//
//// HandleFunc registers the handler function for the given pattern
//// in the DefaultServeMux.
//// The documentation for ServeMux explains how patterns are matched.
//func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
//	DefaultServeMux.HandleFunc(pattern, handler)
//}
//
//// Serve accepts incoming HTTP connections on the listener l,
//// creating a new service goroutine for each. The service goroutines
//// read requests and then call handler to reply to them.
////
//// The handler is typically nil, in which case the DefaultServeMux is used.
////
//// HTTP/2 support is only enabled if the Listener returns *tls.Conn
//// connections and they were configured with "h2" in the TLS
//// Config.NextProtos.
////
//// Serve always returns a non-nil error.
//func Serve(l net.Listener, handler Handler) error {
//	srv := &Server{Handler: handler}
//	return srv.Serve(l)
//}
//
//// ServeTLS accepts incoming HTTPS connections on the listener l,
//// creating a new service goroutine for each. The service goroutines
//// read requests and then call handler to reply to them.
////
//// The handler is typically nil, in which case the DefaultServeMux is used.
////
//// Additionally, files containing a certificate and matching private key
//// for the server must be provided. If the certificate is signed by a
//// certificate authority, the certFile should be the concatenation
//// of the server's certificate, any intermediates, and the CA's certificate.
////
//// ServeTLS always returns a non-nil error.
//func ServeTLS(l net.Listener, handler Handler, certFile, keyFile string) error {
//	srv := &Server{Handler: handler}
//	return srv.ServeTLS(l, certFile, keyFile)
//}
//
//// A Server defines parameters for running an HTTP server.
//// The zero value for Server is a valid configuration.
//type Server struct {
//	// Addr optionally specifies the TCP address for the server to listen on,
//	// in the form "host:port". If empty, ":http" (port 80) is used.
//	// The service names are defined in RFC 6335 and assigned by IANA.
//	// See net.Dial for details of the address format.
//	Addr string
//
//	Handler Handler // handler to invoke, http.DefaultServeMux if nil
//
//	// TLSConfig optionally provides a TLS configuration for use
//	// by ServeTLS and ListenAndServeTLS. Note that this value is
//	// cloned by ServeTLS and ListenAndServeTLS, so it's not
//	// possible to modify the configuration with methods like
//	// tls.Config.SetSessionTicketKeys. To use
//	// SetSessionTicketKeys, use Server.Serve with a TLS Listener
//	// instead.
//	TLSConfig *tls.Config
//
//	// ReadTimeout is the maximum duration for reading the entire
//	// request, including the body.
//	//
//	// Because ReadTimeout does not let Handlers make per-request
//	// decisions on each request body's acceptable deadline or
//	// upload rate, most users will prefer to use
//	// ReadHeaderTimeout. It is valid to use them both.
//	ReadTimeout time.Duration
//
//	// ReadHeaderTimeout is the amount of time allowed to read
//	// request headers. The connection's read deadline is reset
//	// after reading the headers and the Handler can decide what
//	// is considered too slow for the body. If ReadHeaderTimeout
//	// is zero, the value of ReadTimeout is used. If both are
//	// zero, there is no timeout.
//	ReadHeaderTimeout time.Duration
//
//	// WriteTimeout is the maximum duration before timing out
//	// writes of the response. It is reset whenever a new
//	// request's header is read. Like ReadTimeout, it does not
//	// let Handlers make decisions on a per-request basis.
//	WriteTimeout time.Duration
//
//	// IdleTimeout is the maximum amount of time to wait for the
//	// next request when keep-alives are enabled. If IdleTimeout
//	// is zero, the value of ReadTimeout is used. If both are
//	// zero, there is no timeout.
//	IdleTimeout time.Duration
//
//	// MaxHeaderBytes controls the maximum number of bytes the
//	// server will read parsing the request header's keys and
//	// values, including the request line. It does not limit the
//	// size of the request body.
//	// If zero, DefaultMaxHeaderBytes is used.
//	MaxHeaderBytes int
//
//	// TLSNextProto optionally specifies a function to take over
//	// ownership of the provided TLS connection when an ALPN
//	// protocol upgrade has occurred. The map key is the protocol
//	// name negotiated. The Handler argument should be used to
//	// handle HTTP requests and will initialize the Request's TLS
//	// and RemoteAddr if not already set. The connection is
//	// automatically closed when the function returns.
//	// If TLSNextProto is not nil, HTTP/2 support is not enabled
//	// automatically.
//	TLSNextProto map[string]func(*Server, *tls.Conn, Handler)
//
//	// ConnState specifies an optional callback function that is
//	// called when a client connection changes state. See the
//	// ConnState type and associated constants for details.
//	ConnState func(net.Conn, ConnState)
//
//	// ErrorLog specifies an optional logger for errors accepting
//	// connections, unexpected behavior from handlers, and
//	// underlying FileSystem errors.
//	// If nil, logging is done via the log package's standard logger.
//	ErrorLog *log.Logger
//
//	// BaseContext optionally specifies a function that returns
//	// the base context for incoming requests on this server.
//	// The provided Listener is the specific Listener that's
//	// about to start accepting requests.
//	// If BaseContext is nil, the default is context.Background().
//	// If non-nil, it must return a non-nil context.
//	BaseContext func(net.Listener) context.Context
//
//	// ConnContext optionally specifies a function that modifies
//	// the context used for a new connection c. The provided ctx
//	// is derived from the base context and has a ServerContextKey
//	// value.
//	ConnContext func(ctx context.Context, c net.Conn) context.Context
//
//	disableKeepAlives int32     // accessed atomically.
//	inShutdown        int32     // accessed atomically (non-zero means we're in Shutdown)
//	nextProtoOnce     sync.Once // guards setupHTTP2_* init
//	nextProtoErr      error     // result of http2.ConfigureServer if used
//
//	mu         sync.Mutex
//	listeners  map[*net.Listener]struct{}
//	activeConn map[*conn]struct{}
//	doneChan   chan struct{}
//	onShutdown []func()
//}
//
//func (s *Server) getDoneChan() <-chan struct{} {
//	s.mu.Lock()
//	defer s.mu.Unlock()
//	return s.getDoneChanLocked()
//}
//
//func (s *Server) getDoneChanLocked() chan struct{} {
//	if s.doneChan == nil {
//		s.doneChan = make(chan struct{})
//	}
//	return s.doneChan
//}
//
//func (s *Server) closeDoneChanLocked() {
//	ch := s.getDoneChanLocked()
//	select {
//	case <-ch:
//		// Already closed. Don't close again.
//	default:
//		// Safe to close here. We're the only closer, guarded
//		// by s.mu.
//		close(ch)
//	}
//}
//
//// Close immediately closes all active net.Listeners and any
//// connections in state StateNew, StateActive, or StateIdle. For a
//// graceful shutdown, use Shutdown.
////
//// Close does not attempt to close (and does not even know about)
//// any hijacked connections, such as WebSockets.
////
//// Close returns any error returned from closing the Server's
//// underlying Listener(s).
//func (srv *Server) Close() error {
//	atomic.StoreInt32(&srv.inShutdown, 1)
//	srv.mu.Lock()
//	defer srv.mu.Unlock()
//	srv.closeDoneChanLocked()
//	err := srv.closeListenersLocked()
//	for c := range srv.activeConn {
//		c.rwc.Close()
//		delete(srv.activeConn, c)
//	}
//	return err
//}
//
//// shutdownPollInterval is how often we poll for quiescence
//// during Server.Shutdown. This is lower during tests, to
//// speed up tests.
//// Ideally we could find a solution that doesn't involve polling,
//// but which also doesn't have a high runtime cost (and doesn't
//// involve any contentious mutexes), but that is left as an
//// exercise for the reader.
//var shutdownPollInterval = 500 * time.Millisecond
//
//// Shutdown gracefully shuts down the server without interrupting any
//// active connections. Shutdown works by first closing all open
//// listeners, then closing all idle connections, and then waiting
//// indefinitely for connections to return to idle and then shut down.
//// If the provided context expires before the shutdown is complete,
//// Shutdown returns the context's error, otherwise it returns any
//// error returned from closing the Server's underlying Listener(s).
////
//// When Shutdown is called, Serve, ListenAndServe, and
//// ListenAndServeTLS immediately return ErrServerClosed. Make sure the
//// program doesn't exit and waits instead for Shutdown to return.
////
//// Shutdown does not attempt to close nor wait for hijacked
//// connections such as WebSockets. The caller of Shutdown should
//// separately notify such long-lived connections of shutdown and wait
//// for them to close, if desired. See RegisterOnShutdown for a way to
//// register shutdown notification functions.
////
//// Once Shutdown has been called on a server, it may not be reused;
//// future calls to methods such as Serve will return ErrServerClosed.
//func (srv *Server) Shutdown(ctx context.Context) error {
//	atomic.StoreInt32(&srv.inShutdown, 1)
//
//	srv.mu.Lock()
//	lnerr := srv.closeListenersLocked()
//	srv.closeDoneChanLocked()
//	for _, f := range srv.onShutdown {
//		go f()
//	}
//	srv.mu.Unlock()
//
//	ticker := time.NewTicker(shutdownPollInterval)
//	defer ticker.Stop()
//	for {
//		if srv.closeIdleConns() {
//			return lnerr
//		}
//		select {
//		case <-ctx.Done():
//			return ctx.Err()
//		case <-ticker.C:
//		}
//	}
//}
//
//// RegisterOnShutdown registers a function to call on Shutdown.
//// This can be used to gracefully shutdown connections that have
//// undergone ALPN protocol upgrade or that have been hijacked.
//// This function should start protocol-specific graceful shutdown,
//// but should not wait for shutdown to complete.
//func (srv *Server) RegisterOnShutdown(f func()) {
//	srv.mu.Lock()
//	srv.onShutdown = append(srv.onShutdown, f)
//	srv.mu.Unlock()
//}
//
//// closeIdleConns closes all idle connections and reports whether the
//// server is quiescent.
//func (s *Server) closeIdleConns() bool {
//	s.mu.Lock()
//	defer s.mu.Unlock()
//	quiescent := true
//	for c := range s.activeConn {
//		st, unixSec := c.getState()
//		// Issue 22682: treat StateNew connections as if
//		// they're idle if we haven't read the first request's
//		// header in over 5 seconds.
//		if st == StateNew && unixSec < time.Now().Unix()-5 {
//			st = StateIdle
//		}
//		if st != StateIdle || unixSec == 0 {
//			// Assume unixSec == 0 means it's a very new
//			// connection, without state set yet.
//			quiescent = false
//			continue
//		}
//		c.rwc.Close()
//		delete(s.activeConn, c)
//	}
//	return quiescent
//}
//
//func (s *Server) closeListenersLocked() error {
//	var err error
//	for ln := range s.listeners {
//		if cerr := (*ln).Close(); cerr != nil && err == nil {
//			err = cerr
//		}
//		delete(s.listeners, ln)
//	}
//	return err
//}
//
//// A ConnState represents the state of a client connection to a server.
//// It's used by the optional Server.ConnState hook.
//type ConnState int
//
//const (
//	// StateNew represents a new connection that is expected to
//	// send a request immediately. Connections begin at this
//	// state and then transition to either StateActive or
//	// StateClosed.
//	StateNew ConnState = iota
//
//	// StateActive represents a connection that has read 1 or more
//	// bytes of a request. The Server.ConnState hook for
//	// StateActive fires before the request has entered a handler
//	// and doesn't fire again until the request has been
//	// handled. After the request is handled, the state
//	// transitions to StateClosed, StateHijacked, or StateIdle.
//	// For HTTP/2, StateActive fires on the transition from zero
//	// to one active request, and only transitions away once all
//	// active requests are complete. That means that ConnState
//	// cannot be used to do per-request work; ConnState only notes
//	// the overall state of the connection.
//	StateActive
//
//	// StateIdle represents a connection that has finished
//	// handling a request and is in the keep-alive state, waiting
//	// for a new request. Connections transition from StateIdle
//	// to either StateActive or StateClosed.
//	StateIdle
//
//	// StateHijacked represents a hijacked connection.
//	// This is a terminal state. It does not transition to StateClosed.
//	StateHijacked
//
//	// StateClosed represents a closed connection.
//	// This is a terminal state. Hijacked connections do not
//	// transition to StateClosed.
//	StateClosed
//)
//
//var stateName = map[ConnState]string{
//	StateNew:      "new",
//	StateActive:   "active",
//	StateIdle:     "idle",
//	StateHijacked: "hijacked",
//	StateClosed:   "closed",
//}
//
//func (c ConnState) String() string {
//	return stateName[c]
//}
//
//// serverHandler delegates to either the server's Handler or
//// DefaultServeMux and also handles "OPTIONS *" requests.
//type serverHandler struct {
//	srv *Server
//}
//
//func (sh serverHandler) ServeHTTP(rw ResponseWriter, req *Request) {
//	handler := sh.srv.Handler
//	if handler == nil {
//		handler = DefaultServeMux
//	}
//	if req.RequestURI == "*" && req.Method == "OPTIONS" {
//		handler = globalOptionsHandler{}
//	}
//	handler.ServeHTTP(rw, req)
//}
//
//// ListenAndServe listens on the TCP network address srv.Addr and then
//// calls Serve to handle requests on incoming connections.
//// Accepted connections are configured to enable TCP keep-alives.
////
//// If srv.Addr is blank, ":http" is used.
////
//// ListenAndServe always returns a non-nil error. After Shutdown or Close,
//// the returned error is ErrServerClosed.
//func (srv *Server) ListenAndServe() error {
//	if srv.shuttingDown() {
//		return ErrServerClosed
//	}
//	addr := srv.Addr
//	if addr == "" {
//		addr = ":http"
//	}
//	ln, err := net.Listen("tcp", addr)
//	if err != nil {
//		return err
//	}
//	return srv.Serve(ln)
//}
//
//var testHookServerServe func(*Server, net.Listener) // used if non-nil
//
//// shouldDoServeHTTP2 reports whether Server.Serve should configure
//// automatic HTTP/2. (which sets up the srv.TLSNextProto map)
//func (srv *Server) shouldConfigureHTTP2ForServe() bool {
//	if srv.TLSConfig == nil {
//		// Compatibility with Go 1.6:
//		// If there's no TLSConfig, it's possible that the user just
//		// didn't set it on the http.Server, but did pass it to
//		// tls.NewListener and passed that listener to Serve.
//		// So we should configure HTTP/2 (to set up srv.TLSNextProto)
//		// in case the listener returns an "h2" *tls.Conn.
//		return true
//	}
//	// The user specified a TLSConfig on their http.Server.
//	// In this, case, only configure HTTP/2 if their tls.Config
//	// explicitly mentions "h2". Otherwise http2.ConfigureServer
//	// would modify the tls.Config to add it, but they probably already
//	// passed this tls.Config to tls.NewListener. And if they did,
//	// it's too late anyway to fix it. It would only be potentially racy.
//	// See Issue 15908.
//	return strSliceContains(srv.TLSConfig.NextProtos, http2NextProtoTLS)
//}
//
//// ErrServerClosed is returned by the Server's Serve, ServeTLS, ListenAndServe,
//// and ListenAndServeTLS methods after a call to Shutdown or Close.
//var ErrServerClosed = errors.New("http: Server closed")
//
//// Serve accepts incoming connections on the Listener l, creating a
//// new service goroutine for each. The service goroutines read requests and
//// then call srv.Handler to reply to them.
////
//// HTTP/2 support is only enabled if the Listener returns *tls.Conn
//// connections and they were configured with "h2" in the TLS
//// Config.NextProtos.
////
//// Serve always returns a non-nil error and closes l.
//// After Shutdown or Close, the returned error is ErrServerClosed.
//func (srv *Server) Serve(l net.Listener) error {
//	if fn := testHookServerServe; fn != nil {
//		fn(srv, l) // call hook with unwrapped listener
//	}
//
//	origListener := l
//	l = &onceCloseListener{Listener: l}
//	defer l.Close()
//
//	if err := srv.setupHTTP2_Serve(); err != nil {
//		return err
//	}
//
//	if !srv.trackListener(&l, true) {
//		return ErrServerClosed
//	}
//	defer srv.trackListener(&l, false)
//
//	baseCtx := context.Background()
//	if srv.BaseContext != nil {
//		baseCtx = srv.BaseContext(origListener)
//		if baseCtx == nil {
//			panic("BaseContext returned a nil context")
//		}
//	}
//
//	var tempDelay time.Duration // how long to sleep on accept failure
//
//	ctx := context.WithValue(baseCtx, ServerContextKey, srv)
//	for {
//		rw, err := l.Accept()
//		if err != nil {
//			select {
//			case <-srv.getDoneChan():
//				return ErrServerClosed
//			default:
//			}
//			if ne, ok := err.(net.Error); ok && ne.Temporary() {
//				if tempDelay == 0 {
//					tempDelay = 5 * time.Millisecond
//				} else {
//					tempDelay *= 2
//				}
//				if max := 1 * time.Second; tempDelay > max {
//					tempDelay = max
//				}
//				srv.logf("http: Accept error: %v; retrying in %v", err, tempDelay)
//				time.Sleep(tempDelay)
//				continue
//			}
//			return err
//		}
//		connCtx := ctx
//		if cc := srv.ConnContext; cc != nil {
//			connCtx = cc(connCtx, rw)
//			if connCtx == nil {
//				panic("ConnContext returned nil")
//			}
//		}
//		tempDelay = 0
//		c := srv.newConn(rw)
//		c.setState(c.rwc, StateNew) // before Serve can return
//		go c.serve(connCtx)
//	}
//}
//
//// ServeTLS accepts incoming connections on the Listener l, creating a
//// new service goroutine for each. The service goroutines perform TLS
//// setup and then read requests, calling srv.Handler to reply to them.
////
//// Files containing a certificate and matching private key for the
//// server must be provided if neither the Server's
//// TLSConfig.Certificates nor TLSConfig.GetCertificate are populated.
//// If the certificate is signed by a certificate authority, the
//// certFile should be the concatenation of the server's certificate,
//// any intermediates, and the CA's certificate.
////
//// ServeTLS always returns a non-nil error. After Shutdown or Close, the
//// returned error is ErrServerClosed.
//func (srv *Server) ServeTLS(l net.Listener, certFile, keyFile string) error {
//	// Setup HTTP/2 before srv.Serve, to initialize srv.TLSConfig
//	// before we clone it and create the TLS Listener.
//	if err := srv.setupHTTP2_ServeTLS(); err != nil {
//		return err
//	}
//
//	config := cloneTLSConfig(srv.TLSConfig)
//	if !strSliceContains(config.NextProtos, "http/1.1") {
//		config.NextProtos = append(config.NextProtos, "http/1.1")
//	}
//
//	configHasCert := len(config.Certificates) > 0 || config.GetCertificate != nil
//	if !configHasCert || certFile != "" || keyFile != "" {
//		var err error
//		config.Certificates = make([]tls.Certificate, 1)
//		config.Certificates[0], err = tls.LoadX509KeyPair(certFile, keyFile)
//		if err != nil {
//			return err
//		}
//	}
//
//	tlsListener := tls.NewListener(l, config)
//	return srv.Serve(tlsListener)
//}
//
//// trackListener adds or removes a net.Listener to the set of tracked
//// listeners.
////
//// We store a pointer to interface in the map set, in case the
//// net.Listener is not comparable. This is safe because we only call
//// trackListener via Serve and can track+defer untrack the same
//// pointer to local variable there. We never need to compare a
//// Listener from another caller.
////
//// It reports whether the server is still up (not Shutdown or Closed).
//func (s *Server) trackListener(ln *net.Listener, add bool) bool {
//	s.mu.Lock()
//	defer s.mu.Unlock()
//	if s.listeners == nil {
//		s.listeners = make(map[*net.Listener]struct{})
//	}
//	if add {
//		if s.shuttingDown() {
//			return false
//		}
//		s.listeners[ln] = struct{}{}
//	} else {
//		delete(s.listeners, ln)
//	}
//	return true
//}
//
//func (s *Server) trackConn(c *conn, add bool) {
//	s.mu.Lock()
//	defer s.mu.Unlock()
//	if s.activeConn == nil {
//		s.activeConn = make(map[*conn]struct{})
//	}
//	if add {
//		s.activeConn[c] = struct{}{}
//	} else {
//		delete(s.activeConn, c)
//	}
//}
//
//func (s *Server) idleTimeout() time.Duration {
//	if s.IdleTimeout != 0 {
//		return s.IdleTimeout
//	}
//	return s.ReadTimeout
//}
//
//func (s *Server) readHeaderTimeout() time.Duration {
//	if s.ReadHeaderTimeout != 0 {
//		return s.ReadHeaderTimeout
//	}
//	return s.ReadTimeout
//}
//
//func (s *Server) doKeepAlives() bool {
//	return atomic.LoadInt32(&s.disableKeepAlives) == 0 && !s.shuttingDown()
//}
//
//func (s *Server) shuttingDown() bool {
//	// TODO: replace inShutdown with the existing atomicBool type;
//	// see https://github.com/golang/go/issues/20239#issuecomment-381434582
//	return atomic.LoadInt32(&s.inShutdown) != 0
//}
//
//// SetKeepAlivesEnabled controls whether HTTP keep-alives are enabled.
//// By default, keep-alives are always enabled. Only very
//// resource-constrained environments or servers in the process of
//// shutting down should disable them.
//func (srv *Server) SetKeepAlivesEnabled(v bool) {
//	if v {
//		atomic.StoreInt32(&srv.disableKeepAlives, 0)
//		return
//	}
//	atomic.StoreInt32(&srv.disableKeepAlives, 1)
//
//	// Close idle HTTP/1 conns:
//	srv.closeIdleConns()
//
//	// TODO: Issue 26303: close HTTP/2 conns as soon as they become idle.
//}
//
//func (s *Server) logf(format string, args ...interface{}) {
//	if s.ErrorLog != nil {
//		s.ErrorLog.Printf(format, args...)
//	} else {
//		log.Printf(format, args...)
//	}
//}
//
//// logf prints to the ErrorLog of the *Server associated with request r
//// via ServerContextKey. If there's no associated server, or if ErrorLog
//// is nil, logging is done via the log package's standard logger.
//func logf(r *Request, format string, args ...interface{}) {
//	s, _ := r.Context().Value(ServerContextKey).(*Server)
//	if s != nil && s.ErrorLog != nil {
//		s.ErrorLog.Printf(format, args...)
//	} else {
//		log.Printf(format, args...)
//	}
//}
//
//// ListenAndServe listens on the TCP network address addr and then calls
//// Serve with handler to handle requests on incoming connections.
//// Accepted connections are configured to enable TCP keep-alives.
////
//// The handler is typically nil, in which case the DefaultServeMux is used.
////
//// ListenAndServe always returns a non-nil error.
//func ListenAndServe(addr string, handler Handler) error {
//	server := &Server{Addr: addr, Handler: handler}
//	return server.ListenAndServe()
//}
//
//// ListenAndServeTLS acts identically to ListenAndServe, except that it
//// expects HTTPS connections. Additionally, files containing a certificate and
//// matching private key for the server must be provided. If the certificate
//// is signed by a certificate authority, the certFile should be the concatenation
//// of the server's certificate, any intermediates, and the CA's certificate.
//func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error {
//	server := &Server{Addr: addr, Handler: handler}
//	return server.ListenAndServeTLS(certFile, keyFile)
//}
//
//// ListenAndServeTLS listens on the TCP network address srv.Addr and
//// then calls ServeTLS to handle requests on incoming TLS connections.
//// Accepted connections are configured to enable TCP keep-alives.
////
//// Filenames containing a certificate and matching private key for the
//// server must be provided if neither the Server's TLSConfig.Certificates
//// nor TLSConfig.GetCertificate are populated. If the certificate is
//// signed by a certificate authority, the certFile should be the
//// concatenation of the server's certificate, any intermediates, and
//// the CA's certificate.
////
//// If srv.Addr is blank, ":https" is used.
////
//// ListenAndServeTLS always returns a non-nil error. After Shutdown or
//// Close, the returned error is ErrServerClosed.
//func (srv *Server) ListenAndServeTLS(certFile, keyFile string) error {
//	if srv.shuttingDown() {
//		return ErrServerClosed
//	}
//	addr := srv.Addr
//	if addr == "" {
//		addr = ":https"
//	}
//
//	ln, err := net.Listen("tcp", addr)
//	if err != nil {
//		return err
//	}
//
//	defer ln.Close()
//
//	return srv.ServeTLS(ln, certFile, keyFile)
//}
//
//// setupHTTP2_ServeTLS conditionally configures HTTP/2 on
//// srv and reports whether there was an error setting it up. If it is
//// not configured for policy reasons, nil is returned.
//func (srv *Server) setupHTTP2_ServeTLS() error {
//	srv.nextProtoOnce.Do(srv.onceSetNextProtoDefaults)
//	return srv.nextProtoErr
//}
//
//// setupHTTP2_Serve is called from (*Server).Serve and conditionally
//// configures HTTP/2 on srv using a more conservative policy than
//// setupHTTP2_ServeTLS because Serve is called after tls.Listen,
//// and may be called concurrently. See shouldConfigureHTTP2ForServe.
////
//// The tests named TestTransportAutomaticHTTP2* and
//// TestConcurrentServerServe in server_test.go demonstrate some
//// of the supported use cases and motivations.
//func (srv *Server) setupHTTP2_Serve() error {
//	srv.nextProtoOnce.Do(srv.onceSetNextProtoDefaults_Serve)
//	return srv.nextProtoErr
//}
//
//func (srv *Server) onceSetNextProtoDefaults_Serve() {
//	if srv.shouldConfigureHTTP2ForServe() {
//		srv.onceSetNextProtoDefaults()
//	}
//}
//
//// onceSetNextProtoDefaults configures HTTP/2, if the user hasn't
//// configured otherwise. (by setting srv.TLSNextProto non-nil)
//// It must only be called via srv.nextProtoOnce (use srv.setupHTTP2_*).
//func (srv *Server) onceSetNextProtoDefaults() {
//	if omitBundledHTTP2 || strings.Contains(os.Getenv("GODEBUG"), "http2server=0") {
//		return
//	}
//	// Enable HTTP/2 by default if the user hasn't otherwise
//	// configured their TLSNextProto map.
//	if srv.TLSNextProto == nil {
//		conf := &http2Server{
//			NewWriteScheduler: func() http2WriteScheduler { return http2NewPriorityWriteScheduler(nil) },
//		}
//		srv.nextProtoErr = http2ConfigureServer(srv, conf)
//	}
//}
//
//// TimeoutHandler returns a Handler that runs h with the given time limit.
////
//// The new Handler calls h.ServeHTTP to handle each request, but if a
//// call runs for longer than its time limit, the handler responds with
//// a 503 Service Unavailable error and the given message in its body.
//// (If msg is empty, a suitable default message will be sent.)
//// After such a timeout, writes by h to its ResponseWriter will return
//// ErrHandlerTimeout.
////
//// TimeoutHandler supports the Pusher interface but does not support
//// the Hijacker or Flusher interfaces.
//func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler {
//	return &timeoutHandler{
//		handler: h,
//		body:    msg,
//		dt:      dt,
//	}
//}
//
//// ErrHandlerTimeout is returned on ResponseWriter Write calls
//// in handlers which have timed out.
//var ErrHandlerTimeout = errors.New("http: Handler timeout")
//
//type timeoutHandler struct {
//	handler Handler
//	body    string
//	dt      time.Duration
//
//	// When set, no context will be created and this context will
//	// be used instead.
//	testContext context.Context
//}
//
//func (h *timeoutHandler) errorBody() string {
//	if h.body != "" {
//		return h.body
//	}
//	return "<html><head><title>Timeout</title></head><body><h1>Timeout</h1></body></html>"
//}
//
//func (h *timeoutHandler) ServeHTTP(w ResponseWriter, r *Request) {
//	ctx := h.testContext
//	if ctx == nil {
//		var cancelCtx context.CancelFunc
//		ctx, cancelCtx = context.WithTimeout(r.Context(), h.dt)
//		defer cancelCtx()
//	}
//	r = r.WithContext(ctx)
//	done := make(chan struct{})
//	tw := &timeoutWriter{
//		w:   w,
//		h:   make(Header),
//		req: r,
//	}
//	panicChan := make(chan interface{}, 1)
//	go func() {
//		defer func() {
//			if p := recover(); p != nil {
//				panicChan <- p
//			}
//		}()
//		h.handler.ServeHTTP(tw, r)
//		close(done)
//	}()
//	select {
//	case p := <-panicChan:
//		panic(p)
//	case <-done:
//		tw.mu.Lock()
//		defer tw.mu.Unlock()
//		dst := w.Header()
//		for k, vv := range tw.h {
//			dst[k] = vv
//		}
//		if !tw.wroteHeader {
//			tw.code = StatusOK
//		}
//		w.WriteHeader(tw.code)
//		w.Write(tw.wbuf.Bytes())
//	case <-ctx.Done():
//		tw.mu.Lock()
//		defer tw.mu.Unlock()
//		w.WriteHeader(StatusServiceUnavailable)
//		io.WriteString(w, h.errorBody())
//		tw.timedOut = true
//	}
//}
//
//type timeoutWriter struct {
//	w    ResponseWriter
//	h    Header
//	wbuf bytes.Buffer
//	req  *Request
//
//	mu          sync.Mutex
//	timedOut    bool
//	wroteHeader bool
//	code        int
//}
//
//var _ Pusher = (*timeoutWriter)(nil)
//
//// Push implements the Pusher interface.
//func (tw *timeoutWriter) Push(target string, opts *PushOptions) error {
//	if pusher, ok := tw.w.(Pusher); ok {
//		return pusher.Push(target, opts)
//	}
//	return ErrNotSupported
//}
//
//func (tw *timeoutWriter) Header() Header { return tw.h }
//
//func (tw *timeoutWriter) Write(p []byte) (int, error) {
//	tw.mu.Lock()
//	defer tw.mu.Unlock()
//	if tw.timedOut {
//		return 0, ErrHandlerTimeout
//	}
//	if !tw.wroteHeader {
//		tw.writeHeaderLocked(StatusOK)
//	}
//	return tw.wbuf.Write(p)
//}
//
//func (tw *timeoutWriter) writeHeaderLocked(code int) {
//	checkWriteHeaderCode(code)
//
//	switch {
//	case tw.timedOut:
//		return
//	case tw.wroteHeader:
//		if tw.req != nil {
//			caller := relevantCaller()
//			logf(tw.req, "http: superfluous response.WriteHeader call from %s (%s:%d)", caller.Function, path.Base(caller.File), caller.Line)
//		}
//	default:
//		tw.wroteHeader = true
//		tw.code = code
//	}
//}
//
//func (tw *timeoutWriter) WriteHeader(code int) {
//	tw.mu.Lock()
//	defer tw.mu.Unlock()
//	tw.writeHeaderLocked(code)
//}
//
//// onceCloseListener wraps a net.Listener, protecting it from
//// multiple Close calls.
//type onceCloseListener struct {
//	net.Listener
//	once     sync.Once
//	closeErr error
//}
//
//func (oc *onceCloseListener) Close() error {
//	oc.once.Do(oc.close)
//	return oc.closeErr
//}
//
//func (oc *onceCloseListener) close() { oc.closeErr = oc.Listener.Close() }
//
//// globalOptionsHandler responds to "OPTIONS *" requests.
//type globalOptionsHandler struct{}
//
//func (globalOptionsHandler) ServeHTTP(w ResponseWriter, r *Request) {
//	w.Header().Set("Content-Length", "0")
//	if r.ContentLength != 0 {
//		// Read up to 4KB of OPTIONS body (as mentioned in the
//		// spec as being reserved for future use), but anything
//		// over that is considered a waste of server resources
//		// (or an attack) and we abort and close the connection,
//		// courtesy of MaxBytesReader's EOF behavior.
//		mb := MaxBytesReader(w, r.Body, 4<<10)
//		io.Copy(ioutil.Discard, mb)
//	}
//}
//
//// initALPNRequest is an HTTP handler that initializes certain
//// uninitialized fields in its *Request. Such partially-initialized
//// Requests come from ALPN protocol handlers.
//type initALPNRequest struct {
//	ctx context.Context
//	c   *tls.Conn
//	h   serverHandler
//}
//
//// BaseContext is an exported but unadvertised http.Handler method
//// recognized by x/net/http2 to pass down a context; the TLSNextProto
//// API predates context support so we shoehorn through the only
//// interface we have available.
//func (h initALPNRequest) BaseContext() context.Context { return h.ctx }
//
//func (h initALPNRequest) ServeHTTP(rw ResponseWriter, req *Request) {
//	if req.TLS == nil {
//		req.TLS = &tls.ConnectionState{}
//		*req.TLS = h.c.ConnectionState()
//	}
//	if req.Body == nil {
//		req.Body = NoBody
//	}
//	if req.RemoteAddr == "" {
//		req.RemoteAddr = h.c.RemoteAddr().String()
//	}
//	h.h.ServeHTTP(rw, req)
//}
//
//// loggingConn is used for debugging.
//type loggingConn struct {
//	name string
//	net.Conn
//}
//
//var (
//	uniqNameMu   sync.Mutex
//	uniqNameNext = make(map[string]int)
//)
//
//func newLoggingConn(baseName string, c net.Conn) net.Conn {
//	uniqNameMu.Lock()
//	defer uniqNameMu.Unlock()
//	uniqNameNext[baseName]++
//	return &loggingConn{
//		name: fmt.Sprintf("%s-%d", baseName, uniqNameNext[baseName]),
//		Conn: c,
//	}
//}
//
//func (c *loggingConn) Write(p []byte) (n int, err error) {
//	log.Printf("%s.Write(%d) = ....", c.name, len(p))
//	n, err = c.Conn.Write(p)
//	log.Printf("%s.Write(%d) = %d, %v", c.name, len(p), n, err)
//	return
//}
//
//func (c *loggingConn) Read(p []byte) (n int, err error) {
//	log.Printf("%s.Read(%d) = ....", c.name, len(p))
//	n, err = c.Conn.Read(p)
//	log.Printf("%s.Read(%d) = %d, %v", c.name, len(p), n, err)
//	return
//}
//
//func (c *loggingConn) Close() (err error) {
//	log.Printf("%s.Close() = ...", c.name)
//	err = c.Conn.Close()
//	log.Printf("%s.Close() = %v", c.name, err)
//	return
//}
//
//// checkConnErrorWriter writes to c.rwc and records any write errors to c.werr.
//// It only contains one field (and a pointer field at that), so it
//// fits in an interface value without an extra allocation.
//type checkConnErrorWriter struct {
//	c *conn
//}
//
//func (w checkConnErrorWriter) Write(p []byte) (n int, err error) {
//	n, err = w.c.rwc.Write(p)
//	if err != nil && w.c.werr == nil {
//		w.c.werr = err
//		w.c.cancelCtx()
//	}
//	return
//}
//
//func numLeadingCRorLF(v []byte) (n int) {
//	for _, b := range v {
//		if b == '\r' || b == '\n' {
//			n++
//			continue
//		}
//		break
//	}
//	return
//
//}
//
//func strSliceContains(ss []string, s string) bool {
//	for _, v := range ss {
//		if v == s {
//			return true
//		}
//	}
//	return false
//}
//
//// tlsRecordHeaderLooksLikeHTTP reports whether a TLS record header
//// looks like it might've been a misdirected plaintext HTTP request.
//func tlsRecordHeaderLooksLikeHTTP(hdr [5]byte) bool {
//	switch string(hdr[:]) {
//	case "GET /", "HEAD ", "POST ", "PUT /", "OPTIO":
//		return true
//	}
//	return false
//}
//
//// Copyright 2009 The Go Authors. All rights reserved.
//// Use of this source code is governed by a BSD-style
//// license that can be found in the LICENSE file.
//
//// HTTP Request reading and parsing.
//
//package http
//
//import (
//"bufio"
//"bytes"
//"context"
//"crypto/tls"
//"encoding/base64"
//"errors"
//"fmt"
//"io"
//"io/ioutil"
//"mime"
//"mime/multipart"
//"net"
//"net/http/httptrace"
//"net/textproto"
//"net/url"
//urlpkg "net/url"
//"strconv"
//"strings"
//"sync"
//
//"golang.org/x/net/idna"
//)
//
//const (
//	defaultMaxMemory = 32 << 20 // 32 MB
//)
//
//// ErrMissingFile is returned by FormFile when the provided file field name
//// is either not present in the request or not a file field.
//var ErrMissingFile = errors.New("http: no such file")
//
//// ProtocolError represents an HTTP protocol error.
////
//// Deprecated: Not all errors in the http package related to protocol errors
//// are of type ProtocolError.
//type ProtocolError struct {
//	ErrorString string
//}
//
//func (pe *ProtocolError) Error() string { return pe.ErrorString }
//
//var (
//	// ErrNotSupported is returned by the Push method of Pusher
//	// implementations to indicate that HTTP/2 Push support is not
//	// available.
//	ErrNotSupported = &ProtocolError{"feature not supported"}
//
//	// Deprecated: ErrUnexpectedTrailer is no longer returned by
//	// anything in the net/http package. Callers should not
//	// compare errors against this variable.
//	ErrUnexpectedTrailer = &ProtocolError{"trailer header without chunked transfer encoding"}
//
//	// ErrMissingBoundary is returned by Request.MultipartReader when the
//	// request's Content-Type does not include a "boundary" parameter.
//	ErrMissingBoundary = &ProtocolError{"no multipart boundary param in Content-Type"}
//
//	// ErrNotMultipart is returned by Request.MultipartReader when the
//	// request's Content-Type is not multipart/form-data.
//	ErrNotMultipart = &ProtocolError{"request Content-Type isn't multipart/form-data"}
//
//	// Deprecated: ErrHeaderTooLong is no longer returned by
//	// anything in the net/http package. Callers should not
//	// compare errors against this variable.
//	ErrHeaderTooLong = &ProtocolError{"header too long"}
//
//	// Deprecated: ErrShortBody is no longer returned by
//	// anything in the net/http package. Callers should not
//	// compare errors against this variable.
//	ErrShortBody = &ProtocolError{"entity body too short"}
//
//	// Deprecated: ErrMissingContentLength is no longer returned by
//	// anything in the net/http package. Callers should not
//	// compare errors against this variable.
//	ErrMissingContentLength = &ProtocolError{"missing ContentLength in HEAD response"}
//)
//
//type badStringError struct {
//	what string
//	str  string
//}
//
//func (e *badStringError) Error() string { return fmt.Sprintf("%s %q", e.what, e.str) }
//
//// Headers that Request.Write handles itself and should be skipped.
//var reqWriteExcludeHeader = map[string]bool{
//	"Host":              true, // not in Header map anyway
//	"User-Agent":        true,
//	"Content-Length":    true,
//	"Transfer-Encoding": true,
//	"Trailer":           true,
//}
//
//// A Request represents an HTTP request received by a server
//// or to be sent by a client.
////
//// The field semantics differ slightly between client and server
//// usage. In addition to the notes on the fields below, see the
//// documentation for Request.Write and RoundTripper.
//type Request struct {
//	// Method specifies the HTTP method (GET, POST, PUT, etc.).
//	// For client requests, an empty string means GET.
//	//
//	// Go's HTTP client does not support sending a request with
//	// the CONNECT method. See the documentation on Transport for
//	// details.
//	Method string
//
//	// URL specifies either the URI being requested (for server
//	// requests) or the URL to access (for client requests).
//	//
//	// For server requests, the URL is parsed from the URI
//	// supplied on the Request-Line as stored in RequestURI.  For
//	// most requests, fields other than Path and RawQuery will be
//	// empty. (See RFC 7230, Section 5.3)
//	//
//	// For client requests, the URL's Host specifies the server to
//	// connect to, while the Request's Host field optionally
//	// specifies the Host header value to send in the HTTP
//	// request.
//	URL *url.URL
//
//	// The protocol version for incoming server requests.
//	//
//	// For client requests, these fields are ignored. The HTTP
//	// client code always uses either HTTP/1.1 or HTTP/2.
//	// See the docs on Transport for details.
//	Proto      string // "HTTP/1.0"
//	ProtoMajor int    // 1
//	ProtoMinor int    // 0
//
//	// Header contains the request header fields either received
//	// by the server or to be sent by the client.
//	//
//	// If a server received a request with header lines,
//	//
//	//	Host: example.com
//	//	accept-encoding: gzip, deflate
//	//	Accept-Language: en-us
//	//	fOO: Bar
//	//	foo: two
//	//
//	// then
//	//
//	//	Header = map[string][]string{
//	//		"Accept-Encoding": {"gzip, deflate"},
//	//		"Accept-Language": {"en-us"},
//	//		"Foo": {"Bar", "two"},
//	//	}
//	//
//	// For incoming requests, the Host header is promoted to the
//	// Request.Host field and removed from the Header map.
//	//
//	// HTTP defines that header names are case-insensitive. The
//	// request parser implements this by using CanonicalHeaderKey,
//	// making the first character and any characters following a
//	// hyphen uppercase and the rest lowercase.
//	//
//	// For client requests, certain headers such as Content-Length
//	// and Connection are automatically written when needed and
//	// values in Header may be ignored. See the documentation
//	// for the Request.Write method.
//	Header Header
//
//	// Body is the request's body.
//	//
//	// For client requests, a nil body means the request has no
//	// body, such as a GET request. The HTTP Client's Transport
//	// is responsible for calling the Close method.
//	//
//	// For server requests, the Request Body is always non-nil
//	// but will return EOF immediately when no body is present.
//	// The Server will close the request body. The ServeHTTP
//	// Handler does not need to.
//	Body io.ReadCloser
//
//	// GetBody defines an optional func to return a new copy of
//	// Body. It is used for client requests when a redirect requires
//	// reading the body more than once. Use of GetBody still
//	// requires setting Body.
//	//
//	// For server requests, it is unused.
//	GetBody func() (io.ReadCloser, error)
//
//	// ContentLength records the length of the associated content.
//	// The value -1 indicates that the length is unknown.
//	// Values >= 0 indicate that the given number of bytes may
//	// be read from Body.
//	//
//	// For client requests, a value of 0 with a non-nil Body is
//	// also treated as unknown.
//	ContentLength int64
//
//	// TransferEncoding lists the transfer encodings from outermost to
//	// innermost. An empty list denotes the "identity" encoding.
//	// TransferEncoding can usually be ignored; chunked encoding is
//	// automatically added and removed as necessary when sending and
//	// receiving requests.
//	TransferEncoding []string
//
//	// Close indicates whether to close the connection after
//	// replying to this request (for servers) or after sending this
//	// request and reading its response (for clients).
//	//
//	// For server requests, the HTTP server handles this automatically
//	// and this field is not needed by Handlers.
//	//
//	// For client requests, setting this field prevents re-use of
//	// TCP connections between requests to the same hosts, as if
//	// Transport.DisableKeepAlives were set.
//	Close bool
//
//	// For server requests, Host specifies the host on which the
//	// URL is sought. For HTTP/1 (per RFC 7230, section 5.4), this
//	// is either the value of the "Host" header or the host name
//	// given in the URL itself. For HTTP/2, it is the value of the
//	// ":authority" pseudo-header field.
//	// It may be of the form "host:port". For international domain
//	// names, Host may be in Punycode or Unicode form. Use
//	// golang.org/x/net/idna to convert it to either format if
//	// needed.
//	// To prevent DNS rebinding attacks, server Handlers should
//	// validate that the Host header has a value for which the
//	// Handler considers itself authoritative. The included
//	// ServeMux supports patterns registered to particular host
//	// names and thus protects its registered Handlers.
//	//
//	// For client requests, Host optionally overrides the Host
//	// header to send. If empty, the Request.Write method uses
//	// the value of URL.Host. Host may contain an international
//	// domain name.
//	Host string
//
//	// Form contains the parsed form data, including both the URL
//	// field's query parameters and the PATCH, POST, or PUT form data.
//	// This field is only available after ParseForm is called.
//	// The HTTP client ignores Form and uses Body instead.
//	Form url.Values
//
//	// PostForm contains the parsed form data from PATCH, POST
//	// or PUT body parameters.
//	//
//	// This field is only available after ParseForm is called.
//	// The HTTP client ignores PostForm and uses Body instead.
//	PostForm url.Values
//
//	// MultipartForm is the parsed multipart form, including file uploads.
//	// This field is only available after ParseMultipartForm is called.
//	// The HTTP client ignores MultipartForm and uses Body instead.
//	MultipartForm *multipart.Form
//
//	// Trailer specifies additional headers that are sent after the request
//	// body.
//	//
//	// For server requests, the Trailer map initially contains only the
//	// trailer keys, with nil values. (The client declares which trailers it
//	// will later send.)  While the handler is reading from Body, it must
//	// not reference Trailer. After reading from Body returns EOF, Trailer
//	// can be read again and will contain non-nil values, if they were sent
//	// by the client.
//	//
//	// For client requests, Trailer must be initialized to a map containing
//	// the trailer keys to later send. The values may be nil or their final
//	// values. The ContentLength must be 0 or -1, to send a chunked request.
//	// After the HTTP request is sent the map values can be updated while
//	// the request body is read. Once the body returns EOF, the caller must
//	// not mutate Trailer.
//	//
//	// Few HTTP clients, servers, or proxies support HTTP trailers.
//	Trailer Header
//
//	// RemoteAddr allows HTTP servers and other software to record
//	// the network address that sent the request, usually for
//	// logging. This field is not filled in by ReadRequest and
//	// has no defined format. The HTTP server in this package
//	// sets RemoteAddr to an "IP:port" address before invoking a
//	// handler.
//	// This field is ignored by the HTTP client.
//	RemoteAddr string
//
//	// RequestURI is the unmodified request-target of the
//	// Request-Line (RFC 7230, Section 3.1.1) as sent by the client
//	// to a server. Usually the URL field should be used instead.
//	// It is an error to set this field in an HTTP client request.
//	RequestURI string
//
//	// TLS allows HTTP servers and other software to record
//	// information about the TLS connection on which the request
//	// was received. This field is not filled in by ReadRequest.
//	// The HTTP server in this package sets the field for
//	// TLS-enabled connections before invoking a handler;
//	// otherwise it leaves the field nil.
//	// This field is ignored by the HTTP client.
//	TLS *tls.ConnectionState
//
//	// Cancel is an optional channel whose closure indicates that the client
//	// request should be regarded as canceled. Not all implementations of
//	// RoundTripper may support Cancel.
//	//
//	// For server requests, this field is not applicable.
//	//
//	// Deprecated: Set the Request's context with NewRequestWithContext
//	// instead. If a Request's Cancel field and context are both
//	// set, it is undefined whether Cancel is respected.
//	Cancel <-chan struct{}
//
//	// Response is the redirect response which caused this request
//	// to be created. This field is only populated during client
//	// redirects.
//	Response *Response
//
//	// ctx is either the client or server context. It should only
//	// be modified via copying the whole Request using WithContext.
//	// It is unexported to prevent people from using Context wrong
//	// and mutating the contexts held by callers of the same request.
//	ctx context.Context
//}
//
//// Context returns the request's context. To change the context, use
//// WithContext.
////
//// The returned context is always non-nil; it defaults to the
//// background context.
////
//// For outgoing client requests, the context controls cancellation.
////
//// For incoming server requests, the context is canceled when the
//// client's connection closes, the request is canceled (with HTTP/2),
//// or when the ServeHTTP method returns.
//func (r *Request) Context() context.Context {
//	if r.ctx != nil {
//		return r.ctx
//	}
//	return context.Background()
//}
//
//// WithContext returns a shallow copy of r with its context changed
//// to ctx. The provided ctx must be non-nil.
////
//// For outgoing client request, the context controls the entire
//// lifetime of a request and its response: obtaining a connection,
//// sending the request, and reading the response headers and body.
////
//// To create a new request with a context, use NewRequestWithContext.
//// To change the context of a request, such as an incoming request you
//// want to modify before sending back out, use Request.Clone. Between
//// those two uses, it's rare to need WithContext.
//func (r *Request) WithContext(ctx context.Context) *Request {
//	if ctx == nil {
//		panic("nil context")
//	}
//	r2 := new(Request)
//	*r2 = *r
//	r2.ctx = ctx
//	r2.URL = cloneURL(r.URL) // legacy behavior; TODO: try to remove. Issue 23544
//	return r2
//}
//
//// Clone returns a deep copy of r with its context changed to ctx.
//// The provided ctx must be non-nil.
////
//// For an outgoing client request, the context controls the entire
//// lifetime of a request and its response: obtaining a connection,
//// sending the request, and reading the response headers and body.
//func (r *Request) Clone(ctx context.Context) *Request {
//	if ctx == nil {
//		panic("nil context")
//	}
//	r2 := new(Request)
//	*r2 = *r
//	r2.ctx = ctx
//	r2.URL = cloneURL(r.URL)
//	if r.Header != nil {
//		r2.Header = r.Header.Clone()
//	}
//	if r.Trailer != nil {
//		r2.Trailer = r.Trailer.Clone()
//	}
//	if s := r.TransferEncoding; s != nil {
//		s2 := make([]string, len(s))
//		copy(s2, s)
//		r2.TransferEncoding = s
//	}
//	r2.Form = cloneURLValues(r.Form)
//	r2.PostForm = cloneURLValues(r.PostForm)
//	r2.MultipartForm = cloneMultipartForm(r.MultipartForm)
//	return r2
//}
//
//// ProtoAtLeast reports whether the HTTP protocol used
//// in the request is at least major.minor.
//func (r *Request) ProtoAtLeast(major, minor int) bool {
//	return r.ProtoMajor > major ||
//		r.ProtoMajor == major && r.ProtoMinor >= minor
//}
//
//// UserAgent returns the client's User-Agent, if sent in the request.
//func (r *Request) UserAgent() string {
//	return r.Header.Get("User-Agent")
//}
//
//// Cookies parses and returns the HTTP cookies sent with the request.
//func (r *Request) Cookies() []*Cookie {
//	return readCookies(r.Header, "")
//}
//
//// ErrNoCookie is returned by Request's Cookie method when a cookie is not found.
//var ErrNoCookie = errors.New("http: named cookie not present")
//
//// Cookie returns the named cookie provided in the request or
//// ErrNoCookie if not found.
//// If multiple cookies match the given name, only one cookie will
//// be returned.
//func (r *Request) Cookie(name string) (*Cookie, error) {
//	for _, c := range readCookies(r.Header, name) {
//		return c, nil
//	}
//	return nil, ErrNoCookie
//}
//
//// AddCookie adds a cookie to the request. Per RFC 6265 section 5.4,
//// AddCookie does not attach more than one Cookie header field. That
//// means all cookies, if any, are written into the same line,
//// separated by semicolon.
//func (r *Request) AddCookie(c *Cookie) {
//	s := fmt.Sprintf("%s=%s", sanitizeCookieName(c.Name), sanitizeCookieValue(c.Value))
//	if c := r.Header.Get("Cookie"); c != "" {
//		r.Header.Set("Cookie", c+"; "+s)
//	} else {
//		r.Header.Set("Cookie", s)
//	}
//}
//
//// Referer returns the referring URL, if sent in the request.
////
//// Referer is misspelled as in the request itself, a mistake from the
//// earliest days of HTTP.  This value can also be fetched from the
//// Header map as Header["Referer"]; the benefit of making it available
//// as a method is that the compiler can diagnose programs that use the
//// alternate (correct English) spelling req.Referrer() but cannot
//// diagnose programs that use Header["Referrer"].
//func (r *Request) Referer() string {
//	return r.Header.Get("Referer")
//}
//
//// multipartByReader is a sentinel value.
//// Its presence in Request.MultipartForm indicates that parsing of the request
//// body has been handed off to a MultipartReader instead of ParseMultipartForm.
//var multipartByReader = &multipart.Form{
//	Value: make(map[string][]string),
//	File:  make(map[string][]*multipart.FileHeader),
//}
//
//// MultipartReader returns a MIME multipart reader if this is a
//// multipart/form-data or a multipart/mixed POST request, else returns nil and an error.
//// Use this function instead of ParseMultipartForm to
//// process the request body as a stream.
//func (r *Request) MultipartReader() (*multipart.Reader, error) {
//	if r.MultipartForm == multipartByReader {
//		return nil, errors.New("http: MultipartReader called twice")
//	}
//	if r.MultipartForm != nil {
//		return nil, errors.New("http: multipart handled by ParseMultipartForm")
//	}
//	r.MultipartForm = multipartByReader
//	return r.multipartReader(true)
//}
//
//func (r *Request) multipartReader(allowMixed bool) (*multipart.Reader, error) {
//	v := r.Header.Get("Content-Type")
//	if v == "" {
//		return nil, ErrNotMultipart
//	}
//	d, params, err := mime.ParseMediaType(v)
//	if err != nil || !(d == "multipart/form-data" || allowMixed && d == "multipart/mixed") {
//		return nil, ErrNotMultipart
//	}
//	boundary, ok := params["boundary"]
//	if !ok {
//		return nil, ErrMissingBoundary
//	}
//	return multipart.NewReader(r.Body, boundary), nil
//}
//
//// isH2Upgrade reports whether r represents the http2 "client preface"
//// magic string.
//func (r *Request) isH2Upgrade() bool {
//	return r.Method == "PRI" && len(r.Header) == 0 && r.URL.Path == "*" && r.Proto == "HTTP/2.0"
//}
//
//// Return value if nonempty, def otherwise.
//func valueOrDefault(value, def string) string {
//	if value != "" {
//		return value
//	}
//	return def
//}
//
//// NOTE: This is not intended to reflect the actual Go version being used.
//// It was changed at the time of Go 1.1 release because the former User-Agent
//// had ended up on a blacklist for some intrusion detection systems.
//// See https://codereview.appspot.com/7532043.
//const defaultUserAgent = "Go-http-client/1.1"
//
//// Write writes an HTTP/1.1 request, which is the header and body, in wire format.
//// This method consults the following fields of the request:
////	Host
////	URL
////	Method (defaults to "GET")
////	Header
////	ContentLength
////	TransferEncoding
////	Body
////
//// If Body is present, Content-Length is <= 0 and TransferEncoding
//// hasn't been set to "identity", Write adds "Transfer-Encoding:
//// chunked" to the header. Body is closed after it is sent.
//func (r *Request) Write(w io.Writer) error {
//	return r.write(w, false, nil, nil)
//}
//
//// WriteProxy is like Write but writes the request in the form
//// expected by an HTTP proxy. In particular, WriteProxy writes the
//// initial Request-URI line of the request with an absolute URI, per
//// section 5.3 of RFC 7230, including the scheme and host.
//// In either case, WriteProxy also writes a Host header, using
//// either r.Host or r.URL.Host.
//func (r *Request) WriteProxy(w io.Writer) error {
//	return r.write(w, true, nil, nil)
//}
//
//// errMissingHost is returned by Write when there is no Host or URL present in
//// the Request.
//var errMissingHost = errors.New("http: Request.Write on Request with no Host or URL set")
//
//// extraHeaders may be nil
//// waitForContinue may be nil
//func (r *Request) write(w io.Writer, usingProxy bool, extraHeaders Header, waitForContinue func() bool) (err error) {
//	trace := httptrace.ContextClientTrace(r.Context())
//	if trace != nil && trace.WroteRequest != nil {
//		defer func() {
//			trace.WroteRequest(httptrace.WroteRequestInfo{
//				Err: err,
//			})
//		}()
//	}
//
//	// Find the target host. Prefer the Host: header, but if that
//	// is not given, use the host from the request URL.
//	//
//	// Clean the host, in case it arrives with unexpected stuff in it.
//	host := cleanHost(r.Host)
//	if host == "" {
//		if r.URL == nil {
//			return errMissingHost
//		}
//		host = cleanHost(r.URL.Host)
//	}
//
//	// According to RFC 6874, an HTTP client, proxy, or other
//	// intermediary must remove any IPv6 zone identifier attached
//	// to an outgoing URI.
//	host = removeZone(host)
//
//	ruri := r.URL.RequestURI()
//	if usingProxy && r.URL.Scheme != "" && r.URL.Opaque == "" {
//		ruri = r.URL.Scheme + "://" + host + ruri
//	} else if r.Method == "CONNECT" && r.URL.Path == "" {
//		// CONNECT requests normally give just the host and port, not a full URL.
//		ruri = host
//		if r.URL.Opaque != "" {
//			ruri = r.URL.Opaque
//		}
//	}
//	if stringContainsCTLByte(ruri) {
//		return errors.New("net/http: can't write control character in Request.URL")
//	}
//	// TODO: validate r.Method too? At least it's less likely to
//	// come from an attacker (more likely to be a constant in
//	// code).
//
//	// Wrap the writer in a bufio Writer if it's not already buffered.
//	// Don't always call NewWriter, as that forces a bytes.Buffer
//	// and other small bufio Writers to have a minimum 4k buffer
//	// size.
//	var bw *bufio.Writer
//	if _, ok := w.(io.ByteWriter); !ok {
//		bw = bufio.NewWriter(w)
//		w = bw
//	}
//
//	_, err = fmt.Fprintf(w, "%s %s HTTP/1.1\r\n", valueOrDefault(r.Method, "GET"), ruri)
//	if err != nil {
//		return err
//	}
//
//	// Header lines
//	_, err = fmt.Fprintf(w, "Host: %s\r\n", host)
//	if err != nil {
//		return err
//	}
//	if trace != nil && trace.WroteHeaderField != nil {
//		trace.WroteHeaderField("Host", []string{host})
//	}
//
//	// Use the defaultUserAgent unless the Header contains one, which
//	// may be blank to not send the header.
//	userAgent := defaultUserAgent
//	if r.Header.has("User-Agent") {
//		userAgent = r.Header.Get("User-Agent")
//	}
//	if userAgent != "" {
//		_, err = fmt.Fprintf(w, "User-Agent: %s\r\n", userAgent)
//		if err != nil {
//			return err
//		}
//		if trace != nil && trace.WroteHeaderField != nil {
//			trace.WroteHeaderField("User-Agent", []string{userAgent})
//		}
//	}
//
//	// Process Body,ContentLength,Close,Trailer
//	tw, err := newTransferWriter(r)
//	if err != nil {
//		return err
//	}
//	err = tw.writeHeader(w, trace)
//	if err != nil {
//		return err
//	}
//
//	err = r.Header.writeSubset(w, reqWriteExcludeHeader, trace)
//	if err != nil {
//		return err
//	}
//
//	if extraHeaders != nil {
//		err = extraHeaders.write(w, trace)
//		if err != nil {
//			return err
//		}
//	}
//
//	_, err = io.WriteString(w, "\r\n")
//	if err != nil {
//		return err
//	}
//
//	if trace != nil && trace.WroteHeaders != nil {
//		trace.WroteHeaders()
//	}
//
//	// Flush and wait for 100-continue if expected.
//	if waitForContinue != nil {
//		if bw, ok := w.(*bufio.Writer); ok {
//			err = bw.Flush()
//			if err != nil {
//				return err
//			}
//		}
//		if trace != nil && trace.Wait100Continue != nil {
//			trace.Wait100Continue()
//		}
//		if !waitForContinue() {
//			r.closeBody()
//			return nil
//		}
//	}
//
//	if bw, ok := w.(*bufio.Writer); ok && tw.FlushHeaders {
//		if err := bw.Flush(); err != nil {
//			return err
//		}
//	}
//
//	// Write body and trailer
//	err = tw.writeBody(w)
//	if err != nil {
//		if tw.bodyReadError == err {
//			err = requestBodyReadError{err}
//		}
//		return err
//	}
//
//	if bw != nil {
//		return bw.Flush()
//	}
//	return nil
//}
//
//// requestBodyReadError wraps an error from (*Request).write to indicate
//// that the error came from a Read call on the Request.Body.
//// This error type should not escape the net/http package to users.
//type requestBodyReadError struct{ error }
//
//func idnaASCII(v string) (string, error) {
//	// TODO: Consider removing this check after verifying performance is okay.
//	// Right now punycode verification, length checks, context checks, and the
//	// permissible character tests are all omitted. It also prevents the ToASCII
//	// call from salvaging an invalid IDN, when possible. As a result it may be
//	// possible to have two IDNs that appear identical to the user where the
//	// ASCII-only version causes an error downstream whereas the non-ASCII
//	// version does not.
//	// Note that for correct ASCII IDNs ToASCII will only do considerably more
//	// work, but it will not cause an allocation.
//	if isASCII(v) {
//		return v, nil
//	}
//	return idna.Lookup.ToASCII(v)
//}
//
//// cleanHost cleans up the host sent in request's Host header.
////
//// It both strips anything after '/' or ' ', and puts the value
//// into Punycode form, if necessary.
////
//// Ideally we'd clean the Host header according to the spec:
////   https://tools.ietf.org/html/rfc7230#section-5.4 (Host = uri-host [ ":" port ]")
////   https://tools.ietf.org/html/rfc7230#section-2.7 (uri-host -> rfc3986's host)
////   https://tools.ietf.org/html/rfc3986#section-3.2.2 (definition of host)
//// But practically, what we are trying to avoid is the situation in
//// issue 11206, where a malformed Host header used in the proxy context
//// would create a bad request. So it is enough to just truncate at the
//// first offending character.
//func cleanHost(in string) string {
//	if i := strings.IndexAny(in, " /"); i != -1 {
//		in = in[:i]
//	}
//	host, port, err := net.SplitHostPort(in)
//	if err != nil { // input was just a host
//		a, err := idnaASCII(in)
//		if err != nil {
//			return in // garbage in, garbage out
//		}
//		return a
//	}
//	a, err := idnaASCII(host)
//	if err != nil {
//		return in // garbage in, garbage out
//	}
//	return net.JoinHostPort(a, port)
//}
//
//// removeZone removes IPv6 zone identifier from host.
//// E.g., "[fe80::1%en0]:8080" to "[fe80::1]:8080"
//func removeZone(host string) string {
//	if !strings.HasPrefix(host, "[") {
//		return host
//	}
//	i := strings.LastIndex(host, "]")
//	if i < 0 {
//		return host
//	}
//	j := strings.LastIndex(host[:i], "%")
//	if j < 0 {
//		return host
//	}
//	return host[:j] + host[i:]
//}
//
//// ParseHTTPVersion parses an HTTP version string.
//// "HTTP/1.0" returns (1, 0, true).
//func ParseHTTPVersion(vers string) (major, minor int, ok bool) {
//	const Big = 1000000 // arbitrary upper bound
//	switch vers {
//	case "HTTP/1.1":
//		return 1, 1, true
//	case "HTTP/1.0":
//		return 1, 0, true
//	}
//	if !strings.HasPrefix(vers, "HTTP/") {
//		return 0, 0, false
//	}
//	dot := strings.Index(vers, ".")
//	if dot < 0 {
//		return 0, 0, false
//	}
//	major, err := strconv.Atoi(vers[5:dot])
//	if err != nil || major < 0 || major > Big {
//		return 0, 0, false
//	}
//	minor, err = strconv.Atoi(vers[dot+1:])
//	if err != nil || minor < 0 || minor > Big {
//		return 0, 0, false
//	}
//	return major, minor, true
//}
//
//func validMethod(method string) bool {
//	/*
//	     Method         = "OPTIONS"                ; Section 9.2
//	                    | "GET"                    ; Section 9.3
//	                    | "HEAD"                   ; Section 9.4
//	                    | "POST"                   ; Section 9.5
//	                    | "PUT"                    ; Section 9.6
//	                    | "DELETE"                 ; Section 9.7
//	                    | "TRACE"                  ; Section 9.8
//	                    | "CONNECT"                ; Section 9.9
//	                    | extension-method
//	   extension-method = token
//	     token          = 1*<any CHAR except CTLs or separators>
//	*/
//	return len(method) > 0 && strings.IndexFunc(method, isNotToken) == -1
//}
//
//// NewRequest wraps NewRequestWithContext using the background context.
//func NewRequest(method, url string, body io.Reader) (*Request, error) {
//	return NewRequestWithContext(context.Background(), method, url, body)
//}
//
//// NewRequestWithContext returns a new Request given a method, URL, and
//// optional body.
////
//// If the provided body is also an io.Closer, the returned
//// Request.Body is set to body and will be closed by the Client
//// methods Do, Post, and PostForm, and Transport.RoundTrip.
////
//// NewRequestWithContext returns a Request suitable for use with
//// Client.Do or Transport.RoundTrip. To create a request for use with
//// testing a Server Handler, either use the NewRequest function in the
//// net/http/httptest package, use ReadRequest, or manually update the
//// Request fields. For an outgoing client request, the context
//// controls the entire lifetime of a request and its response:
//// obtaining a connection, sending the request, and reading the
//// response headers and body. See the Request type's documentation for
//// the difference between inbound and outbound request fields.
////
//// If body is of type *bytes.Buffer, *bytes.Reader, or
//// *strings.Reader, the returned request's ContentLength is set to its
//// exact value (instead of -1), GetBody is populated (so 307 and 308
//// redirects can replay the body), and Body is set to NoBody if the
//// ContentLength is 0.
//func NewRequestWithContext(ctx context.Context, method, url string, body io.Reader) (*Request, error) {
//	if method == "" {
//		// We document that "" means "GET" for Request.Method, and people have
//		// relied on that from NewRequest, so keep that working.
//		// We still enforce validMethod for non-empty methods.
//		method = "GET"
//	}
//	if !validMethod(method) {
//		return nil, fmt.Errorf("net/http: invalid method %q", method)
//	}
//	if ctx == nil {
//		return nil, errors.New("net/http: nil Context")
//	}
//	u, err := urlpkg.Parse(url)
//	if err != nil {
//		return nil, err
//	}
//	rc, ok := body.(io.ReadCloser)
//	if !ok && body != nil {
//		rc = ioutil.NopCloser(body)
//	}
//	// The host's colon:port should be normalized. See Issue 14836.
//	u.Host = removeEmptyPort(u.Host)
//	req := &Request{
//		ctx:        ctx,
//		Method:     method,
//		URL:        u,
//		Proto:      "HTTP/1.1",
//		ProtoMajor: 1,
//		ProtoMinor: 1,
//		Header:     make(Header),
//		Body:       rc,
//		Host:       u.Host,
//	}
//	if body != nil {
//		switch v := body.(type) {
//		case *bytes.Buffer:
//			req.ContentLength = int64(v.Len())
//			buf := v.Bytes()
//			req.GetBody = func() (io.ReadCloser, error) {
//				r := bytes.NewReader(buf)
//				return ioutil.NopCloser(r), nil
//			}
//		case *bytes.Reader:
//			req.ContentLength = int64(v.Len())
//			snapshot := *v
//			req.GetBody = func() (io.ReadCloser, error) {
//				r := snapshot
//				return ioutil.NopCloser(&r), nil
//			}
//		case *strings.Reader:
//			req.ContentLength = int64(v.Len())
//			snapshot := *v
//			req.GetBody = func() (io.ReadCloser, error) {
//				r := snapshot
//				return ioutil.NopCloser(&r), nil
//			}
//		default:
//			// This is where we'd set it to -1 (at least
//			// if body != NoBody) to mean unknown, but
//			// that broke people during the Go 1.8 testing
//			// period. People depend on it being 0 I
//			// guess. Maybe retry later. See Issue 18117.
//		}
//		// For client requests, Request.ContentLength of 0
//		// means either actually 0, or unknown. The only way
//		// to explicitly say that the ContentLength is zero is
//		// to set the Body to nil. But turns out too much code
//		// depends on NewRequest returning a non-nil Body,
//		// so we use a well-known ReadCloser variable instead
//		// and have the http package also treat that sentinel
//		// variable to mean explicitly zero.
//		if req.GetBody != nil && req.ContentLength == 0 {
//			req.Body = NoBody
//			req.GetBody = func() (io.ReadCloser, error) { return NoBody, nil }
//		}
//	}
//
//	return req, nil
//}
//
//// BasicAuth returns the username and password provided in the request's
//// Authorization header, if the request uses HTTP Basic Authentication.
//// See RFC 2617, Section 2.
//func (r *Request) BasicAuth() (username, password string, ok bool) {
//	auth := r.Header.Get("Authorization")
//	if auth == "" {
//		return
//	}
//	return parseBasicAuth(auth)
//}
//
//// parseBasicAuth parses an HTTP Basic Authentication string.
//// "Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==" returns ("Aladdin", "open sesame", true).
//func parseBasicAuth(auth string) (username, password string, ok bool) {
//	const prefix = "Basic "
//	// Case insensitive prefix match. See Issue 22736.
//	if len(auth) < len(prefix) || !strings.EqualFold(auth[:len(prefix)], prefix) {
//		return
//	}
//	c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
//	if err != nil {
//		return
//	}
//	cs := string(c)
//	s := strings.IndexByte(cs, ':')
//	if s < 0 {
//		return
//	}
//	return cs[:s], cs[s+1:], true
//}
//
//// SetBasicAuth sets the request's Authorization header to use HTTP
//// Basic Authentication with the provided username and password.
////
//// With HTTP Basic Authentication the provided username and password
//// are not encrypted.
////
//// Some protocols may impose additional requirements on pre-escaping the
//// username and password. For instance, when used with OAuth2, both arguments
//// must be URL encoded first with url.QueryEscape.
//func (r *Request) SetBasicAuth(username, password string) {
//	r.Header.Set("Authorization", "Basic "+basicAuth(username, password))
//}
//
//// parseRequestLine parses "GET /foo HTTP/1.1" into its three parts.
//func parseRequestLine(line string) (method, requestURI, proto string, ok bool) {
//	s1 := strings.Index(line, " ")
//	s2 := strings.Index(line[s1+1:], " ")
//	if s1 < 0 || s2 < 0 {
//		return
//	}
//	s2 += s1 + 1
//	return line[:s1], line[s1+1 : s2], line[s2+1:], true
//}
//
//var textprotoReaderPool sync.Pool
//
//func newTextprotoReader(br *bufio.Reader) *textproto.Reader {
//	if v := textprotoReaderPool.Get(); v != nil {
//		tr := v.(*textproto.Reader)
//		tr.R = br
//		return tr
//	}
//	return textproto.NewReader(br)
//}
//
//func putTextprotoReader(r *textproto.Reader) {
//	r.R = nil
//	textprotoReaderPool.Put(r)
//}
//
//// ReadRequest reads and parses an incoming request from b.
////
//// ReadRequest is a low-level function and should only be used for
//// specialized applications; most code should use the Server to read
//// requests and handle them via the Handler interface. ReadRequest
//// only supports HTTP/1.x requests. For HTTP/2, use golang.org/x/net/http2.
//func ReadRequest(b *bufio.Reader) (*Request, error) {
//	return readRequest(b, deleteHostHeader)
//}
//
//// Constants for readRequest's deleteHostHeader parameter.
//const (
//	deleteHostHeader = true
//	keepHostHeader   = false
//)
//
//func readRequest(b *bufio.Reader, deleteHostHeader bool) (req *Request, err error) {
//	tp := newTextprotoReader(b)
//	req = new(Request)
//
//	// First line: GET /index.html HTTP/1.0
//	var s string
//	if s, err = tp.ReadLine(); err != nil {
//		return nil, err
//	}
//	defer func() {
//		putTextprotoReader(tp)
//		if err == io.EOF {
//			err = io.ErrUnexpectedEOF
//		}
//	}()
//
//	var ok bool
//	req.Method, req.RequestURI, req.Proto, ok = parseRequestLine(s)
//	if !ok {
//		return nil, &badStringError{"malformed HTTP request", s}
//	}
//	if !validMethod(req.Method) {
//		return nil, &badStringError{"invalid method", req.Method}
//	}
//	rawurl := req.RequestURI
//	if req.ProtoMajor, req.ProtoMinor, ok = ParseHTTPVersion(req.Proto); !ok {
//		return nil, &badStringError{"malformed HTTP version", req.Proto}
//	}
//
//	// CONNECT requests are used two different ways, and neither uses a full URL:
//	// The standard use is to tunnel HTTPS through an HTTP proxy.
//	// It looks like "CONNECT www.google.com:443 HTTP/1.1", and the parameter is
//	// just the authority section of a URL. This information should go in req.URL.Host.
//	//
//	// The net/rpc package also uses CONNECT, but there the parameter is a path
//	// that starts with a slash. It can be parsed with the regular URL parser,
//	// and the path will end up in req.URL.Path, where it needs to be in order for
//	// RPC to work.
//	justAuthority := req.Method == "CONNECT" && !strings.HasPrefix(rawurl, "/")
//	if justAuthority {
//		rawurl = "http://" + rawurl
//	}
//
//	if req.URL, err = url.ParseRequestURI(rawurl); err != nil {
//		return nil, err
//	}
//
//	if justAuthority {
//		// Strip the bogus "http://" back off.
//		req.URL.Scheme = ""
//	}
//
//	// Subsequent lines: Key: value.
//	mimeHeader, err := tp.ReadMIMEHeader()
//	if err != nil {
//		return nil, err
//	}
//	req.Header = Header(mimeHeader)
//
//	// RFC 7230, section 5.3: Must treat
//	//	GET /index.html HTTP/1.1
//	//	Host: www.google.com
//	// and
//	//	GET http://www.google.com/index.html HTTP/1.1
//	//	Host: doesntmatter
//	// the same. In the second case, any Host line is ignored.
//	req.Host = req.URL.Host
//	if req.Host == "" {
//		req.Host = req.Header.get("Host")
//	}
//	if deleteHostHeader {
//		delete(req.Header, "Host")
//	}
//
//	fixPragmaCacheControl(req.Header)
//
//	req.Close = shouldClose(req.ProtoMajor, req.ProtoMinor, req.Header, false)
//
//	err = readTransfer(req, b)
//	if err != nil {
//		return nil, err
//	}
//
//	if req.isH2Upgrade() {
//		// Because it's neither chunked, nor declared:
//		req.ContentLength = -1
//
//		// We want to give handlers a chance to hijack the
//		// connection, but we need to prevent the Server from
//		// dealing with the connection further if it's not
//		// hijacked. Set Close to ensure that:
//		req.Close = true
//	}
//	return req, nil
//}
//
//// MaxBytesReader is similar to io.LimitReader but is intended for
//// limiting the size of incoming request bodies. In contrast to
//// io.LimitReader, MaxBytesReader's result is a ReadCloser, returns a
//// non-EOF error for a Read beyond the limit, and closes the
//// underlying reader when its Close method is called.
////
//// MaxBytesReader prevents clients from accidentally or maliciously
//// sending a large request and wasting server resources.
//func MaxBytesReader(w ResponseWriter, r io.ReadCloser, n int64) io.ReadCloser {
//	return &maxBytesReader{w: w, r: r, n: n}
//}
//
//type maxBytesReader struct {
//	w   ResponseWriter
//	r   io.ReadCloser // underlying reader
//	n   int64         // max bytes remaining
//	err error         // sticky error
//}
//
//func (l *maxBytesReader) Read(p []byte) (n int, err error) {
//	if l.err != nil {
//		return 0, l.err
//	}
//	if len(p) == 0 {
//		return 0, nil
//	}
//	// If they asked for a 32KB byte read but only 5 bytes are
//	// remaining, no need to read 32KB. 6 bytes will answer the
//	// question of the whether we hit the limit or go past it.
//	if int64(len(p)) > l.n+1 {
//		p = p[:l.n+1]
//	}
//	n, err = l.r.Read(p)
//
//	if int64(n) <= l.n {
//		l.n -= int64(n)
//		l.err = err
//		return n, err
//	}
//
//	n = int(l.n)
//	l.n = 0
//
//	// The server code and client code both use
//	// maxBytesReader. This "requestTooLarge" check is
//	// only used by the server code. To prevent binaries
//	// which only using the HTTP Client code (such as
//	// cmd/go) from also linking in the HTTP server, don't
//	// use a static type assertion to the server
//	// "*response" type. Check this interface instead:
//	type requestTooLarger interface {
//		requestTooLarge()
//	}
//	if res, ok := l.w.(requestTooLarger); ok {
//		res.requestTooLarge()
//	}
//	l.err = errors.New("http: request body too large")
//	return n, l.err
//}
//
//func (l *maxBytesReader) Close() error {
//	return l.r.Close()
//}
//
//func copyValues(dst, src url.Values) {
//	for k, vs := range src {
//		dst[k] = append(dst[k], vs...)
//	}
//}
//
//func parsePostForm(r *Request) (vs url.Values, err error) {
//	if r.Body == nil {
//		err = errors.New("missing form body")
//		return
//	}
//	ct := r.Header.Get("Content-Type")
//	// RFC 7231, section 3.1.1.5 - empty type
//	//   MAY be treated as application/octet-stream
//	if ct == "" {
//		ct = "application/octet-stream"
//	}
//	ct, _, err = mime.ParseMediaType(ct)
//	switch {
//	case ct == "application/x-www-form-urlencoded":
//		var reader io.Reader = r.Body
//		maxFormSize := int64(1<<63 - 1)
//		if _, ok := r.Body.(*maxBytesReader); !ok {
//			maxFormSize = int64(10 << 20) // 10 MB is a lot of text.
//			reader = io.LimitReader(r.Body, maxFormSize+1)
//		}
//		b, e := ioutil.ReadAll(reader)
//		if e != nil {
//			if err == nil {
//				err = e
//			}
//			break
//		}
//		if int64(len(b)) > maxFormSize {
//			err = errors.New("http: POST too large")
//			return
//		}
//		vs, e = url.ParseQuery(string(b))
//		if err == nil {
//			err = e
//		}
//	case ct == "multipart/form-data":
//		// handled by ParseMultipartForm (which is calling us, or should be)
//		// TODO(bradfitz): there are too many possible
//		// orders to call too many functions here.
//		// Clean this up and write more tests.
//		// request_test.go contains the start of this,
//		// in TestParseMultipartFormOrder and others.
//	}
//	return
//}
//
//// ParseForm populates r.Form and r.PostForm.
////
//// For all requests, ParseForm parses the raw query from the URL and updates
//// r.Form.
////
//// For POST, PUT, and PATCH requests, it also reads the request body, parses it
//// as a form and puts the results into both r.PostForm and r.Form. Request body
//// parameters take precedence over URL query string values in r.Form.
////
//// If the request Body's size has not already been limited by MaxBytesReader,
//// the size is capped at 10MB.
////
//// For other HTTP methods, or when the Content-Type is not
//// application/x-www-form-urlencoded, the request Body is not read, and
//// r.PostForm is initialized to a non-nil, empty value.
////
//// ParseMultipartForm calls ParseForm automatically.
//// ParseForm is idempotent.
//func (r *Request) ParseForm() error {
//	var err error
//	if r.PostForm == nil {
//		if r.Method == "POST" || r.Method == "PUT" || r.Method == "PATCH" {
//			r.PostForm, err = parsePostForm(r)
//		}
//		if r.PostForm == nil {
//			r.PostForm = make(url.Values)
//		}
//	}
//	if r.Form == nil {
//		if len(r.PostForm) > 0 {
//			r.Form = make(url.Values)
//			copyValues(r.Form, r.PostForm)
//		}
//		var newValues url.Values
//		if r.URL != nil {
//			var e error
//			newValues, e = url.ParseQuery(r.URL.RawQuery)
//			if err == nil {
//				err = e
//			}
//		}
//		if newValues == nil {
//			newValues = make(url.Values)
//		}
//		if r.Form == nil {
//			r.Form = newValues
//		} else {
//			copyValues(r.Form, newValues)
//		}
//	}
//	return err
//}
//
//// ParseMultipartForm parses a request body as multipart/form-data.
//// The whole request body is parsed and up to a total of maxMemory bytes of
//// its file parts are stored in memory, with the remainder stored on
//// disk in temporary files.
//// ParseMultipartForm calls ParseForm if necessary.
//// After one call to ParseMultipartForm, subsequent calls have no effect.
//func (r *Request) ParseMultipartForm(maxMemory int64) error {
//	if r.MultipartForm == multipartByReader {
//		return errors.New("http: multipart handled by MultipartReader")
//	}
//	if r.Form == nil {
//		err := r.ParseForm()
//		if err != nil {
//			return err
//		}
//	}
//	if r.MultipartForm != nil {
//		return nil
//	}
//
//	mr, err := r.multipartReader(false)
//	if err != nil {
//		return err
//	}
//
//	f, err := mr.ReadForm(maxMemory)
//	if err != nil {
//		return err
//	}
//
//	if r.PostForm == nil {
//		r.PostForm = make(url.Values)
//	}
//	for k, v := range f.Value {
//		r.Form[k] = append(r.Form[k], v...)
//		// r.PostForm should also be populated. See Issue 9305.
//		r.PostForm[k] = append(r.PostForm[k], v...)
//	}
//
//	r.MultipartForm = f
//
//	return nil
//}
//
//// FormValue returns the first value for the named component of the query.
//// POST and PUT body parameters take precedence over URL query string values.
//// FormValue calls ParseMultipartForm and ParseForm if necessary and ignores
//// any errors returned by these functions.
//// If key is not present, FormValue returns the empty string.
//// To access multiple values of the same key, call ParseForm and
//// then inspect Request.Form directly.
//func (r *Request) FormValue(key string) string {
//	if r.Form == nil {
//		r.ParseMultipartForm(defaultMaxMemory)
//	}
//	if vs := r.Form[key]; len(vs) > 0 {
//		return vs[0]
//	}
//	return ""
//}
//
//// PostFormValue returns the first value for the named component of the POST,
//// PATCH, or PUT request body. URL query parameters are ignored.
//// PostFormValue calls ParseMultipartForm and ParseForm if necessary and ignores
//// any errors returned by these functions.
//// If key is not present, PostFormValue returns the empty string.
//func (r *Request) PostFormValue(key string) string {
//	if r.PostForm == nil {
//		r.ParseMultipartForm(defaultMaxMemory)
//	}
//	if vs := r.PostForm[key]; len(vs) > 0 {
//		return vs[0]
//	}
//	return ""
//}
//
//// FormFile returns the first file for the provided form key.
//// FormFile calls ParseMultipartForm and ParseForm if necessary.
//func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error) {
//	if r.MultipartForm == multipartByReader {
//		return nil, nil, errors.New("http: multipart handled by MultipartReader")
//	}
//	if r.MultipartForm == nil {
//		err := r.ParseMultipartForm(defaultMaxMemory)
//		if err != nil {
//			return nil, nil, err
//		}
//	}
//	if r.MultipartForm != nil && r.MultipartForm.File != nil {
//		if fhs := r.MultipartForm.File[key]; len(fhs) > 0 {
//			f, err := fhs[0].Open()
//			return f, fhs[0], err
//		}
//	}
//	return nil, nil, ErrMissingFile
//}
//
//func (r *Request) expectsContinue() bool {
//	return hasToken(r.Header.get("Expect"), "100-continue")
//}
//
//func (r *Request) wantsHttp10KeepAlive() bool {
//	if r.ProtoMajor != 1 || r.ProtoMinor != 0 {
//		return false
//	}
//	return hasToken(r.Header.get("Connection"), "keep-alive")
//}
//
//func (r *Request) wantsClose() bool {
//	if r.Close {
//		return true
//	}
//	return hasToken(r.Header.get("Connection"), "close")
//}
//
//func (r *Request) closeBody() {
//	if r.Body != nil {
//		r.Body.Close()
//	}
//}
//
//func (r *Request) isReplayable() bool {
//	if r.Body == nil || r.Body == NoBody || r.GetBody != nil {
//		switch valueOrDefault(r.Method, "GET") {
//		case "GET", "HEAD", "OPTIONS", "TRACE":
//			return true
//		}
//		// The Idempotency-Key, while non-standard, is widely used to
//		// mean a POST or other request is idempotent. See
//		// https://golang.org/issue/19943#issuecomment-421092421
//		if r.Header.has("Idempotency-Key") || r.Header.has("X-Idempotency-Key") {
//			return true
//		}
//	}
//	return false
//}
//
//// outgoingLength reports the Content-Length of this outgoing (Client) request.
//// It maps 0 into -1 (unknown) when the Body is non-nil.
//func (r *Request) outgoingLength() int64 {
//	if r.Body == nil || r.Body == NoBody {
//		return 0
//	}
//	if r.ContentLength != 0 {
//		return r.ContentLength
//	}
//	return -1
//}
//
//// requestMethodUsuallyLacksBody reports whether the given request
//// method is one that typically does not involve a request body.
//// This is used by the Transport (via
//// transferWriter.shouldSendChunkedRequestBody) to determine whether
//// we try to test-read a byte from a non-nil Request.Body when
//// Request.outgoingLength() returns -1. See the comments in
//// shouldSendChunkedRequestBody.
//func requestMethodUsuallyLacksBody(method string) bool {
//	switch method {
//	case "GET", "HEAD", "DELETE", "OPTIONS", "PROPFIND", "SEARCH":
//		return true
//	}
//	return false
//}
//
//// requiresHTTP1 reports whether this request requires being sent on
//// an HTTP/1 connection.
//func (r *Request) requiresHTTP1() bool {
//	return hasToken(r.Header.Get("Connection"), "upgrade") &&
//		strings.EqualFold(r.Header.Get("Upgrade"), "websocket")
//}

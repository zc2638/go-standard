package htmlTemp

func HeaderHtml(host, backUri string) string {
	return `
<!doctype html>
<html xmlns=http://www.w3.org/1999/xhtml>
<meta charset=utf-8>
<script src="https://cdn.mathjax.org/mathjax/latest/MathJax.js?config=TeX-AMS_CHTML-full"></script>
<style>@font-face {
        font-family: octicons-link;
        src: url(data:font/woff;charset=utf-8;base64,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) format('woff')
    }

    .octicon {
        overflow: hidden;
        vertical-align: text-bottom
    }

    .boxed-group {
        width: 980px;
        margin-right: auto;
        margin-left: auto;
        position: relative;
        margin-top: 30px;
        margin-bottom: 30px;
        border-radius: 3px;
        font-family: -apple-system, BlinkMacSystemFont, segoe ui, Helvetica, Arial, sans-serif, apple color emoji, segoe ui emoji, segoe ui symbol;
        -ms-text-size-adjust: 100%;
        -webkit-text-size-adjust: 100%;
        line-height: 1.5;
        color: #24292e;
        font-size: 16px;
        line-height: 1.5;
        word-wrap: break-word
    }

    .boxed-group h3 {
        display: block;
        padding: 9px 10px 10px;
        margin: 0;
        font-size: 14px;
        line-height: 17px;
        background-color: #f6f8fa;
        border: 1px solid rgba(27, 31, 35, .15);
        border-bottom: 0;
        border-radius: 3px 3px 0 0;
        font-weight: 600
    }

    .markdown-body {
        padding: 45px;
        word-wrap: break-word;
        background-color: #fff;
        border: 1px solid #ddd;
        border-bottom-right-radius: 3px;
        border-bottom-left-radius: 3px
    }

    .markdown-body .pl-c {
        color: #6a737d
    }

    .markdown-body .pl-c1, .markdown-body .pl-s .pl-v {
        color: #005cc5
    }

    .markdown-body .pl-e, .markdown-body .pl-en {
        color: #6f42c1
    }

    .markdown-body .pl-smi, .markdown-body .pl-s .pl-s1 {
        color: #24292e
    }

    .markdown-body .pl-ent {
        color: #22863a
    }

    .markdown-body .pl-k {
        color: #d73a49
    }

    .markdown-body .pl-s, .markdown-body .pl-pds, .markdown-body .pl-s .pl-pse .pl-s1, .markdown-body .pl-sr, .markdown-body .pl-sr .pl-cce, .markdown-body .pl-sr .pl-sre, .markdown-body .pl-sr .pl-sra {
        color: #032f62
    }

    .markdown-body .pl-v, .markdown-body .pl-smw {
        color: #e36209
    }

    .markdown-body .pl-bu {
        color: #b31d28
    }

    .markdown-body .pl-ii {
        color: #fafbfc;
        background-color: #b31d28
    }

    .markdown-body .pl-c2 {
        color: #fafbfc;
        background-color: #d73a49
    }

    .markdown-body .pl-c2::before {
        content: "^M"
    }

    .markdown-body .pl-sr .pl-cce {
        font-weight: 700;
        color: #22863a
    }

    .markdown-body .pl-ml {
        color: #735c0f
    }

    .markdown-body .pl-mh, .markdown-body .pl-mh .pl-en, .markdown-body .pl-ms {
        font-weight: 700;
        color: #005cc5
    }

    .markdown-body .pl-mi {
        font-style: italic;
        color: #24292e
    }

    .markdown-body .pl-mb {
        font-weight: 700;
        color: #24292e
    }

    .markdown-body .pl-md {
        color: #b31d28;
        background-color: #ffeef0
    }

    .markdown-body .pl-mi1 {
        color: #22863a;
        background-color: #f0fff4
    }

    .markdown-body .pl-mc {
        color: #e36209;
        background-color: #ffebda
    }

    .markdown-body .pl-mi2 {
        color: #f6f8fa;
        background-color: #005cc5
    }

    .markdown-body .pl-mdr {
        font-weight: 700;
        color: #6f42c1
    }

    .markdown-body .pl-ba {
        color: #586069
    }

    .markdown-body .pl-sg {
        color: #959da5
    }

    .markdown-body .pl-corl {
        text-decoration: underline;
        color: #032f62
    }

    .markdown-body .octicon {
        display: inline-block;
        vertical-align: text-top;
        fill: currentColor
    }

    .markdown-body a {
        background-color: transparent
    }

    .markdown-body a:active, .markdown-body a:hover {
        outline-width: 0
    }

    .markdown-body strong {
        font-weight: inherit
    }

    .markdown-body strong {
        font-weight: bolder
    }

    .markdown-body h1 {
        font-size: 2em;
        margin: .67em 0
    }

    .markdown-body img {
        border-style: none
    }

    .markdown-body code, .markdown-body kbd, .markdown-body pre {
        font-family: monospace, monospace;
        font-size: 1em
    }

    .markdown-body hr {
        box-sizing: content-box;
        height: 0;
        overflow: visible
    }

    .markdown-body input {
        font: inherit;
        margin: 0
    }

    .markdown-body input {
        overflow: visible
    }

    .markdown-body [type=checkbox] {
        box-sizing: border-box;
        padding: 0
    }

    .markdown-body * {
        box-sizing: border-box
    }

    .markdown-body input {
        font-family: inherit;
        font-size: inherit;
        line-height: inherit
    }

    .markdown-body a {
        color: #0366d6;
        text-decoration: none
    }

    .markdown-body a:hover {
        text-decoration: underline
    }

    .markdown-body strong {
        font-weight: 600
    }

    .markdown-body hr {
        height: 0;
        margin: 15px 0;
        overflow: hidden;
        background: 0 0;
        border: 0;
        border-bottom: 1px solid #dfe2e5
    }

    .markdown-body hr::before {
        display: table;
        content: ""
    }

    .markdown-body hr::after {
        display: table;
        clear: both;
        content: ""
    }

    .markdown-body table {
        border-spacing: 0;
        border-collapse: collapse
    }

    .markdown-body td, .markdown-body th {
        padding: 0
    }

    .markdown-body h1, .markdown-body h2, .markdown-body h3, .markdown-body h4, .markdown-body h5, .markdown-body h6 {
        margin-top: 0;
        margin-bottom: 0
    }

    .markdown-body h1 {
        font-size: 32px;
        font-weight: 600
    }

    .markdown-body h2 {
        font-size: 24px;
        font-weight: 600
    }

    .markdown-body h3 {
        font-size: 20px;
        font-weight: 600
    }

    .markdown-body h4 {
        font-size: 16px;
        font-weight: 600
    }

    .markdown-body h5 {
        font-size: 14px;
        font-weight: 600
    }

    .markdown-body h6 {
        font-size: 12px;
        font-weight: 600
    }

    .markdown-body p {
        margin-top: 0;
        margin-bottom: 10px
    }

    .markdown-body blockquote {
        margin: 0
    }

    .markdown-body ul, .markdown-body ol {
        padding-left: 0;
        margin-top: 0;
        margin-bottom: 0
    }

    .markdown-body ol ol, .markdown-body ul ol {
        list-style-type: lower-roman
    }

    .markdown-body ul ul ol, .markdown-body ul ol ol, .markdown-body ol ul ol, .markdown-body ol ol ol {
        list-style-type: lower-alpha
    }

    .markdown-body dd {
        margin-left: 0
    }

    .markdown-body code {
        font-family: sfmono-regular, Consolas, liberation mono, Menlo, Courier, monospace;
        font-size: 12px
    }

    .markdown-body pre {
        margin-top: 0;
        margin-bottom: 0;
        font-family: sfmono-regular, Consolas, liberation mono, Menlo, Courier, monospace;
        font-size: 12px
    }

    .markdown-body .octicon {
        vertical-align: text-bottom
    }

    .markdown-body .pl-0 {
        padding-left: 0 !important
    }

    .markdown-body .pl-1 {
        padding-left: 4px !important
    }

    .markdown-body .pl-2 {
        padding-left: 8px !important
    }

    .markdown-body .pl-3 {
        padding-left: 16px !important
    }

    .markdown-body .pl-4 {
        padding-left: 24px !important
    }

    .markdown-body .pl-5 {
        padding-left: 32px !important
    }

    .markdown-body .pl-6 {
        padding-left: 40px !important
    }

    .markdown-body::before {
        display: table;
        content: ""
    }

    .markdown-body::after {
        display: table;
        clear: both;
        content: ""
    }

    .markdown-body > *:first-child {
        margin-top: 0 !important
    }

    .markdown-body > *:last-child {
        margin-bottom: 0 !important
    }

    .markdown-body a:not([href]) {
        color: inherit;
        text-decoration: none
    }

    .markdown-body .anchor {
        float: left;
        padding-right: 4px;
        margin-left: -20px;
        line-height: 1
    }

    .markdown-body .anchor:focus {
        outline: 0
    }

    .markdown-body p, .markdown-body blockquote, .markdown-body ul, .markdown-body ol, .markdown-body dl, .markdown-body table, .markdown-body pre {
        margin-top: 0;
        margin-bottom: 16px
    }

    .markdown-body hr {
        height: .25em;
        padding: 0;
        margin: 24px 0;
        background-color: #e1e4e8;
        border: 0
    }

    .markdown-body blockquote {
        padding: 0 1em;
        color: #6a737d;
        border-left: .25em solid #dfe2e5
    }

    .markdown-body blockquote > :first-child {
        margin-top: 0
    }

    .markdown-body blockquote > :last-child {
        margin-bottom: 0
    }

    .markdown-body kbd {
        display: inline-block;
        padding: 3px 5px;
        font-size: 11px;
        line-height: 10px;
        color: #444d56;
        vertical-align: middle;
        background-color: #fafbfc;
        border: solid 1px #c6cbd1;
        border-bottom-color: #959da5;
        border-radius: 3px;
        box-shadow: inset 0 -1px 0 #959da5
    }

    .markdown-body h1, .markdown-body h2, .markdown-body h3, .markdown-body h4, .markdown-body h5, .markdown-body h6 {
        margin-top: 24px;
        margin-bottom: 16px;
        font-weight: 600;
        line-height: 1.25
    }

    .markdown-body h1 .octicon-link, .markdown-body h2 .octicon-link, .markdown-body h3 .octicon-link, .markdown-body h4 .octicon-link, .markdown-body h5 .octicon-link, .markdown-body h6 .octicon-link {
        color: #1b1f23;
        vertical-align: middle;
        visibility: hidden
    }

    .markdown-body h1:hover .anchor, .markdown-body h2:hover .anchor, .markdown-body h3:hover .anchor, .markdown-body h4:hover .anchor, .markdown-body h5:hover .anchor, .markdown-body h6:hover .anchor {
        text-decoration: none
    }

    .markdown-body h1:hover .anchor .octicon-link, .markdown-body h2:hover .anchor .octicon-link, .markdown-body h3:hover .anchor .octicon-link, .markdown-body h4:hover .anchor .octicon-link, .markdown-body h5:hover .anchor .octicon-link, .markdown-body h6:hover .anchor .octicon-link {
        visibility: visible
    }

    .markdown-body h1 {
        padding-bottom: .3em;
        font-size: 2em;
        border-bottom: 1px solid #eaecef
    }

    .markdown-body h2 {
        padding-bottom: .3em;
        font-size: 1.5em;
        border-bottom: 1px solid #eaecef
    }

    .markdown-body h3 {
        font-size: 1.25em
    }

    .markdown-body h4 {
        font-size: 1em
    }

    .markdown-body h5 {
        font-size: .875em
    }

    .markdown-body h6 {
        font-size: .85em;
        color: #6a737d
    }

    .markdown-body ul, .markdown-body ol {
        padding-left: 2em
    }

    .markdown-body ul ul, .markdown-body ul ol, .markdown-body ol ol, .markdown-body ol ul {
        margin-top: 0;
        margin-bottom: 0
    }

    .markdown-body li {
        word-wrap: break-all
    }

    .markdown-body li > p {
        margin-top: 16px
    }

    .markdown-body li + li {
        margin-top: .25em
    }

    .markdown-body dl {
        padding: 0
    }

    .markdown-body dl dt {
        padding: 0;
        margin-top: 16px;
        font-size: 1em;
        font-style: italic;
        font-weight: 600
    }

    .markdown-body dl dd {
        padding: 0 16px;
        margin-bottom: 16px
    }

    .markdown-body table {
        display: block;
        width: 100%;
        overflow: auto
    }

    .markdown-body table th {
        font-weight: 600
    }

    .markdown-body table th, .markdown-body table td {
        padding: 6px 13px;
        border: 1px solid #dfe2e5
    }

    .markdown-body table tr {
        background-color: #fff;
        border-top: 1px solid #c6cbd1
    }

    .markdown-body table tr:nth-child(2n) {
        background-color: #f6f8fa
    }

    .markdown-body img {
        max-width: 100%;
        box-sizing: content-box;
        background-color: #fff
    }

    .markdown-body img[align=right] {
        padding-left: 20px
    }

    .markdown-body img[align=left] {
        padding-right: 20px
    }

    .markdown-body code {
        padding: .2em .4em;
        margin: 0;
        font-size: 85%;
        background-color: rgba(27, 31, 35, .05);
        border-radius: 3px
    }

    .markdown-body pre {
        word-wrap: normal
    }

    .markdown-body pre > code {
        padding: 0;
        margin: 0;
        font-size: 100%;
        word-break: normal;
        white-space: pre;
        background: 0 0;
        border: 0
    }

    .markdown-body .highlight {
        margin-bottom: 16px
    }

    .markdown-body .highlight pre {
        margin-bottom: 0;
        word-break: normal
    }

    .markdown-body .highlight pre, .markdown-body pre {
        padding: 16px;
        overflow: auto;
        font-size: 85%;
        line-height: 1.45;
        background-color: #f6f8fa;
        border-radius: 3px
    }

    .markdown-body pre code {
        display: inline;
        max-width: auto;
        padding: 0;
        margin: 0;
        overflow: visible;
        line-height: inherit;
        word-wrap: normal;
        background-color: transparent;
        border: 0
    }

    .markdown-body .full-commit .btn-outline:not(:disabled):hover {
        color: #005cc5;
        border-color: #005cc5
    }

    .markdown-body kbd {
        display: inline-block;
        padding: 3px 5px;
        font: 11px sfmono-regular, Consolas, liberation mono, Menlo, Courier, monospace;
        line-height: 10px;
        color: #444d56;
        vertical-align: middle;
        background-color: #fafbfc;
        border: solid 1px #d1d5da;
        border-bottom-color: #c6cbd1;
        border-radius: 3px;
        box-shadow: inset 0 -1px 0 #c6cbd1
    }

    .markdown-body :checked + .radio-label {
        position: relative;
        z-index: 1;
        border-color: #0366d6
    }

    .markdown-body .task-list-item {
        list-style-type: none
    }

    .markdown-body .task-list-item + .task-list-item {
        margin-top: 3px
    }

    .markdown-body .task-list-item input {
        margin: 0 .2em .25em -1.6em;
        vertical-align: middle
    }

    .markdown-body hr {
        border-bottom-color: #eee
    }</style>
<title></title>
<body>
<div class=boxed-group>
    <h3>
        <svg class="octicon octicon-book" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true">
            <path fill-rule="evenodd"
                  d="M3 5h4v1H3V5zm0 3h4V7H3v1zm0 2h4V9H3v1zm11-5h-4v1h4V5zm0 2h-4v1h4V7zm0 2h-4v1h4V9zm2-6v9c0 .55-.45 1-1 1H9.5l-1 1-1-1H2c-.55 0-1-.45-1-1V3c0-.55.45-1 1-1h5.5l1 1 1-1H15c.55 0 1 .45 1 1zm-8 .5L7.5 3H2v9h6V3.5zm7-.5H9.5l-.5.5V12h6V3z"></path>
        </svg>
		
		<span><a href="` + host + `">返回目录</a></span>
		<span><a href="` + backUri + `">返回上一级</a></span>
    </h3>
    <div class="markdown-body article">
`
}

const Footer = `</div></div></body>`

var PackList = []string{
	"archive/tar", "archive/zip",
	"bufio",
	"bytes",
	"compress/bzip2", "compress/flate", "compress/gzip", "compress/lzw", "compress/zlib",
	"container/heap", "container/list", "container/ring",
	"context",
	"crypto/aes", "crypto/cipher", "crypto/des", "crypto/dsa", "crypto/ecdsa", "crypto/elliptic", "crypto/hmac", "crypto/md5", "crypto/rand", "crypto/rc4", "crypto/rsa", "crypto/sha1", "crypto/sha256", "crypto/sha512", "crypto/subtle", "crypto/tls", "crypto/x509",
	"database/sql",
	"encoding/ascii85", "encoding/asn1", "encoding/base32", "encoding/base64", "encoding/binary", "encoding/csv", "encoding/gob", "encoding/hex", "encoding/json", "encoding/pem", "encoding/xml",
	"errors",
	"expvar",
	"flag",
	"fmt",
	"hash/adler32", "hash/crc32", "hash/crc64", "hash/fnv",
	"html", "html/template",
	"image",
	"index/suffixarray",
	"io", "io/ioutil",
	"log", "log/syslog",
	"math", "math/big",
	"mime",
	"net", "net/http", "net/mail", "net/rpc", "net/smtp", "net/textproto", "net/url",
	"os", "os/exec", "os/signal", "os/user",
	"path", "path/filepath",
	"plugin",
	"reflect",
	"regexp",
	"runtime", "runtime/debug", "runtime/pprof", "runtime/trace",
	"sort",
	"strconv",
	"strings",
	"sync", "sync/atomic",
	"text", "text/scanner", "text/tabwriter", "text/template",
	"time",
	"unicode", "unicode/utf8", "unicode/utf16",
	"unsafe",
}

func MenuList(host string) string {

	var list string
	for _, pack := range PackList {
		list += `<a href="` + host + `/` + pack + `">` + pack + `</a><br/>`
	}
	return list
}

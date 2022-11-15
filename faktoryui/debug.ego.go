// Generated by ego.
// DO NOT EDIT

//line debug.ego:1

package faktoryui

import "fmt"
import "html"
import "io"
import "context"

import (
	"net/http"
	"runtime"

	"github.com/contribsys/faktory/client"
)

func ego_debug(w io.Writer, req *http.Request) {
	stats := ctx(req).Store().Stats(req.Context())
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	rdata, rtt := redis_info(req)

//line debug.ego:17
	_, _ = io.WriteString(w, "\n")
//line debug.ego:17
	ego_layout(w, req, func() {
//line debug.ego:18
		_, _ = io.WriteString(w, "\n\n<h3>")
//line debug.ego:19
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Debugging"))))
//line debug.ego:19
		_, _ = io.WriteString(w, "</h3>\n<div class=\"table-responsive\">\n  <table class=\"error table table-bordered table-striped table-light\">\n    <tbody>\n      <tr>\n        <th>")
//line debug.ego:24
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Locale"))))
//line debug.ego:24
		_, _ = io.WriteString(w, "</th>\n        <td>\n          <select name=\"locales\" id=\"faktory_locale\">\n            ")
//line debug.ego:27
		sortedLocaleNames(req, func(locale string, current bool) {
//line debug.ego:28
			_, _ = io.WriteString(w, "\n              <option name=\"")
//line debug.ego:28
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(locale)))
//line debug.ego:28
			_, _ = io.WriteString(w, "\" ")
//line debug.ego:28
			if current {
//line debug.ego:28
				_, _ = io.WriteString(w, " selected ")
//line debug.ego:28
			}
//line debug.ego:28
			_, _ = io.WriteString(w, " > ")
//line debug.ego:28
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(locale)))
//line debug.ego:28
			_, _ = io.WriteString(w, " </option>\n            ")
//line debug.ego:29
		})
//line debug.ego:30
		_, _ = io.WriteString(w, "\n          </select>\n          <span style=\"font-size: small\">\n            Want to help us improve the translations?\n            <a href=\"https://github.com/contribsys/faktory/tree/master/webui/static/locales\">Submit a PR</a>.\n          </span>\n        </td>\n    </tr>\n    <tr>\n      <th>")
//line debug.ego:38
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Version"))))
//line debug.ego:38
		_, _ = io.WriteString(w, "</th>\n      <td>")
//line debug.ego:39
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(client.Name)))
//line debug.ego:39
		_, _ = io.WriteString(w, " ")
//line debug.ego:39
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(client.Version)))
//line debug.ego:39
		_, _ = io.WriteString(w, "</td>\n    </tr>\n    <tr>\n    <th>")
//line debug.ego:42
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Data Location"))))
//line debug.ego:42
		_, _ = io.WriteString(w, "</th>\n      <td>")
//line debug.ego:43
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(stats["name"])))
//line debug.ego:43
		_, _ = io.WriteString(w, "</td>\n    </tr>\n    <tr>\n      <th>")
//line debug.ego:46
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Runtime"))))
//line debug.ego:46
		_, _ = io.WriteString(w, "</th>\n      <td>Goroutines: ")
//line debug.ego:47
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(runtime.NumGoroutine())))
//line debug.ego:47
		_, _ = io.WriteString(w, ", CPUs: ")
//line debug.ego:47
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(runtime.NumCPU())))
//line debug.ego:47
		_, _ = io.WriteString(w, "</td>\n    </tr>\n    <tr>\n      <th>")
//line debug.ego:50
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Memory"))))
//line debug.ego:50
		_, _ = io.WriteString(w, "</th>\n      <td>\n        Alloc (KB): ")
//line debug.ego:52
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(m.Alloc/1024)))
//line debug.ego:52
		_, _ = io.WriteString(w, "<br/>\n        Live Objects: ")
//line debug.ego:53
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(m.Mallocs-m.Frees)))
//line debug.ego:54
		_, _ = io.WriteString(w, "\n        ")
//line debug.ego:54
		if amt := client.RssKb(); amt != 0 {
//line debug.ego:55
			_, _ = io.WriteString(w, "\n        <br/>RSS: ")
//line debug.ego:55
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(displayRss(amt))))
//line debug.ego:56
			_, _ = io.WriteString(w, "\n        ")
//line debug.ego:56
		}
//line debug.ego:57
		_, _ = io.WriteString(w, "\n      </td>\n    </tr>\n    <tr>\n      <th>")
//line debug.ego:60
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "GC"))))
//line debug.ego:60
		_, _ = io.WriteString(w, "</th>\n      <td>\n        PauseTotal (µs): ")
//line debug.ego:62
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(m.PauseTotalNs/1000)))
//line debug.ego:62
		_, _ = io.WriteString(w, "<br/>\n        NumGC: ")
//line debug.ego:63
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(m.NumGC)))
//line debug.ego:64
		_, _ = io.WriteString(w, "\n      </td>\n    </tr>\n    <tr>\n      <th>\n        ")
//line debug.ego:68
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Redis RTT"))))
//line debug.ego:69
		_, _ = io.WriteString(w, "\n        <a href=\"https://github.com/contribsys/faktory/wiki/Storage#rtt\"><span class=\"info-circle\" title=\"Click to learn more about RTT\">?</span></a>\n      </th>\n      <td class=\"bg-")
//line debug.ego:71
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(category_for_rtt(rtt))))
//line debug.ego:71
		_, _ = io.WriteString(w, "\">\n        ")
//line debug.ego:72
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(rtt)))
//line debug.ego:72
		_, _ = io.WriteString(w, " µs\n      </td>\n    </tr>\n  </tbody>\n</table>\n</div>\n\n<h3>")
//line debug.ego:79
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Redis Info"))))
//line debug.ego:79
		_, _ = io.WriteString(w, "</h3>\n<pre>\n")
//line debug.ego:81
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(rdata)))
//line debug.ego:82
		_, _ = io.WriteString(w, "\n</pre>\n\n<h3>")
//line debug.ego:84
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Disk Usage"))))
//line debug.ego:84
		_, _ = io.WriteString(w, "</h3>\n<pre>\n<code>&gt; df -h</code>\n")
//line debug.ego:87
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(df_h())))
//line debug.ego:88
		_, _ = io.WriteString(w, "\n</pre>\n\n")
//line debug.ego:90
	})
//line debug.ego:91
	_, _ = io.WriteString(w, "\n")
//line debug.ego:91
}

var _ fmt.Stringer
var _ io.Reader
var _ context.Context
var _ = html.EscapeString

package main

import (
	"html/template"
	"log"
	"path"
)

func loadTemplates(dir string) *template.Template {
	if dir == "" {
		return getTemplates()
	}
	log.Printf("using custom template directory %q", dir)
	t, err := template.New("").ParseFiles(path.Join(dir, "sign_in.html"), path.Join(dir, "error.html"))
	if err != nil {
		log.Fatalf("failed parsing template %s", err)
	}
	return t
}

func getTemplates() *template.Template {
	t, err := template.New("foo").Parse(`{{define "sign_in.html"}}
<!DOCTYPE html>
<html lang="en" charset="utf-8">
<head>
	<title>Sign In</title>
	<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
	<style>
	body {
		font-family: "Helvetica Neue",Helvetica,Arial,sans-serif;
		font-size: 14px;
		line-height: 1.42857143;
		color: #333;
		background: #f0f0f0;
	}
	.signin {
		display:block;
		margin:20px auto;
		max-width:400px;
		background: #fff;
		border:1px solid #ccc;
		border-radius: 10px;
		padding: 20px;
	}
	.center {
		text-align:center;
	}
	.btn {
		color: #fff;
		background-color: #e20074;
		border: 1px solid;
		border-color: rgba(0,0,0,.3);
		-webkit-border-radius: 4;
		-moz-border-radius: 4;
		border-radius: 4px;
		font-size: 14px;
		padding: 6px 12px;
	  	text-decoration: none;
		cursor: pointer;
	}

	.btn:hover {
		background-color: #d1006c;
		border-color: rgba(0,0,0,.3);
		ext-decoration: none;
	}
	label {
		display: inline-block;
		max-width: 100%;
		margin-bottom: 5px;
		font-weight: 700;
	}
	input {
		display: block;
		width: 100%;
		height: 34px;
		padding: 6px 12px;
		font-size: 14px;
		line-height: 1.42857143;
		color: #555;
		background-color: #fff;
		background-image: none;
		border: 1px solid #ccc;
		border-radius: 4px;
		-webkit-box-shadow: inset 0 1px 1px rgba(0,0,0,.075);
		box-shadow: inset 0 1px 1px rgba(0,0,0,.075);
		-webkit-transition: border-color ease-in-out .15s,-webkit-box-shadow ease-in-out .15s;
		-o-transition: border-color ease-in-out .15s,box-shadow ease-in-out .15s;
		transition: border-color ease-in-out .15s,box-shadow ease-in-out .15s;
		margin:0;
		box-sizing: border-box;
	}
	footer {
		display:block;
		font-size:10px;
		color:#aaa;
		text-align:center;
		margin-bottom:10px;
	}
	footer a {
		display:inline-block;
		height:25px;
		line-height:25px;
		color:#aaa;
		text-decoration:underline;
	}
	footer a:hover {
		color:#aaa;
	}
	</style>
</head>
<body>
	<div class="signin center">
	<center><svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" width="215px" height="214px" viewBox="0 0 215 214" enable-background="new 0 0 215 214" xml:space="preserve">  <image id="image0" width="215" height="214" x="0" y="0" xlink:href="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAANcAAADWCAYAAABCIRcQAAAABGdBTUEAALGPC/xhBQAAACBjSFJN
    AAB6JgAAgIQAAPoAAACA6AAAdTAAAOpgAAA6mAAAF3CculE8AAAABmJLR0QA/wD/AP+gvaeTAAAA
    CXBIWXMAABcSAAAXEgFnn9JSAACAAElEQVR42uydd3gUZdfGf7N9N713EgJJ6L1K7wo28FVERfyw
    994RxI4FO6Ii2MXeeOlNmvQOgSSEAAmQhPRsnZ2Z74/NLAmm0ZLgy31dewV2Z2dnZ+eec55T7iNM
    nTqVi7iIizj30DT2AVzERfxboWvsA/hfx5QpU4RqnhZOe0ceKKc+MXXqVOVMdnQRZw/holt4/nEK
    gYRKf0/9d02v1YbK5JEr/V+p5d/AReKdb1wk13lABZlOJYqm0t/aHuo2VPo/lfYFVQmlVHrIlf6e
    +pCq2aYK6S6S7dziIrnOEpWsUmXrU5ksOkBbzUN9XlPxV3/Kw2CxWHSA1s/PT6vT6eAk6WS3201Z
    WZkESDabzQ24APGUh8RJYkmAu+L/7lP+rxJP/esl3UXCnTkukusMcIplUklUmTQ6PATRnfIwAwEJ
    CQn+CQkJAS1btgxq06ZNeExMTIi/v3+Ar6+vr9ls9vXx8fExGo2+BoNBryiKxmAwmDQajSDLsucD
    NRoURVHcbrej4q/LZrOV2+12q9VqLbdardbi4uKS48ePn9i7d29+RkZGUVZWVmlmZmYxUArY8BBJ
    xEMw9aH+XyWjavXgItlOGxfJVU9UIpRqmVQy6QEDHvKof42AX1BQUMjgwYMj27dvH5WSktIsLi4u
    NiIiIjogICDY19c30Gg0WjQajfZ8HreiKJLD4bBbrdbi0tLSgry8vGOHDx/O3r9//5Fdu3YdXbly
    5fH8/PwTQDngpKrlc3GSeFXcyoskqxsXyVUDTnH31LVPZffNUOlhAoISExPDr7rqquY9evRITk5O
    bhkVFdU8MDAw1Gw2+1b3GZIkYbVaKSoqorCw0Ps3Ly+P9u3b069fvxqPLysri5dffhmdTkdgYCC+
    vr74+/vj5+dHYGAg4eHhREdHExERgdlsrnYfTqezvLi4uCAnJycrPT09Y9OmTWl//vlnZlpaWh5Q
    CDg46W66Kv371PXbRYtWDS6SqxpUEKvymkm1SkZOEsovIiIi+pprrmkxaNCg9m3atGkXHR2dGBAQ
    ECwIQpX8YVlZGbm5uRw+fJjDhw9z6NAhDh06RE5ODsXFxTidTlwuFy6XC1EUcTgctGvXjp9++omg
    oKBqj1EURe69915++uknjEYjiqKg0WjQarVotVoMBgNmsxlfX18iIiKIiooiMjKS2NhYkpKSSEpK
    IiIioso+FUVRysrKio4fP35g9+7de/76669dP//8c0ZOTs5RPO6kSjAnVa2a1328SLKTuEiuClRy
    +yoHGFRCGfFYp4B27do1u+222zr16NGjc3x8fKvw8PA4XUW0QcXBgwfZu3cvu3fvZu/eveTk5HDi
    xAnKyspwOp3IsowgCOh0OjQaDYIg/OPhcrmYMmUK//d//1fjMe/YsYMbbrgBl8ulrsMAUBTF+5Bl
    GUmSkCTJS0Cz2UxoaCjNmjUjKSmJlJQU7yMgIMC7f1mW3fn5+dmZmZn7N2/evP2LL77YtmXLlkNA
    MR6r5uQk0SoHUC66jfyPk+sU108NRqiEMlX89W/VqlXCzTff3Gno0KE9W7Ro0SE4OLjKLT8jI4Nt
    27axe/dudu/ezaFDhygoKMBms6HX670kUh+CUHeO2O1207x5c37++ecqF/ypeOihh/j1118xmUz1
    +s6VCed2uxFFEa1Wi5+fH5GRkbRr145u3brRrVs3WrVqReX7RnFxcV56evqulStXbvj888+37d27
    9xBQgodgDk66kao1U6ZOnSrX68D+hfifJVcl108llRFPNM8EWICoxx9/vNPIkSP7dejQoXtwcHCk
    +l673c7OnTtZv349f//9N3v37qWkpMR7oer1erRaLRrN2VWXOZ1OJk+ezMSJE2vcZseOHdx00004
    HI4z/jyVcG63G5fLhVarxd/fn4SEBDp37kyPHj3o1asX4eHh3vcUFhbmpaambvrvf/+7+o033tjh
    drtz8EQh7RWPKtbsf9GS/c+R6xRSqVbKjIdQ/n369Em+6667+vbq1atvQkJCa9XlKy8v5++//2bt
    2rVs2rSJAwcOUFpa6l3fqFbpXEIURVq0aMH3339PcHBwjds98sgj/PTTTzUGLk4XiqIgSRKiKOJ2
    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
    E/Hx8fTu3ZuNGzeSlZVFeno6w4cPx9fXl27dupGZmcnevXvPqkVFTVRLkkRcXBxjxozh8ccf5957
    71UruC+I3JNGoyEqKorevXtz1VVX0aNHD4xGI7m5uRQWFiJJ0jltvFQUBZ1Ox7PPPktSUhKSJPHC
    Cy/w7bffYrFYMBgMaLVaJEliwIABtG/fnrKyMn799dd6C+6ohN26dSv9+vUjOjrap1u3bi3nzZu3
    s7CwsJRKBDufAY7zSq5KstJqACN64cKFD7dv3753YWEh99xzD/v37z+tyguTycTkyZOJiopi9erV
    3HXXXSxbtozly5ezevVqWrZsSWxsLCEhIbRp04Zly5aRmZlJYWEhI0aMQKvV0qtXLzZu3Mjx48dP
    e82lkkpRFHr06MH999/P008/zRVXXEFcXFyTTUrXB+qooGHDhjF8+HDCw8MpLS0lNzfXO7frbEkm
    SRIJCQk8+OCDmEwmVq5cybRp06oMuRBFkfDwcB5++GGCgoLIyMjgp59+8notai6stmPR6/Xk5eWx
    e/duhg8fTmRkZGjv3r3DP/30022cVPeVVq5cKZ8vDpy3GPQp7qA/EPz555+P7d279whJknjxxRfZ
    sWPHaSnDKoqCv78/ycnJAKxdu5bS0lICAwMxmUzs3LmT++67j/Xr1wPQrVs3xo4di06n47fffuPr
    r78GPJHI4ODg08pxKYqCzWZDFEV69OjB22+/zVdffcUNN9xQZSjcvwUJCQncf//9zJ07l7fffpse
    PXp43d+zyQ0qikJISIg3kLNo0aIqo4/UyOrEiRO90cP169djs9kQBMGbwlDzgrXBbDazadMmJk+e
    jNvtpnv37gPnzp07HgjFc00aAd35yn+dzwSP6g76AQGPPvrooLFjx94CnkbHX3755Ywklyuf1JSU
    FG9uR62czs3NZcqUKZw4cQLwFO2qQ7vfffdd5syZw8SJE1m5cmW93UKn04nT6aRbt2689dZbfPXV
    V4wePbrRJaMbAv7+/owePZqvvvqK1157jbZt22K3271F0WcCNegDVElROBwOHA4HN998s3eiS25u
    Lr/++qu35Ak8AyeeeOIJZFmuF8H++OMPPvnkEwCuvvrqG55//vkReAJrfni0Wc4LD86LW1ipAsMX
    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
    C7yKUQbADwhp06ZNh8WLF78SExPTYvfu3dx2220cP378vHYlnw+oqr9ms5m2bdsycuRIhg4dSosW
    J0ffICmIO47iWJaOY1kG4o6jSLllyIoDT95Mh6d2WlNpjNDZks7jLCjevxIgeYitN6OLDcDQPQ7T
    8GSMA1ugaxFS5d379+9n8eLFLFiwgH379uFyuTCZTE02AFUTXC4XISEhzJo1i06dOpGbm3voiiuu
    eHbTpk1bqRodPG/V2g1BLgHPFWQEAoCQW265ZeC777471d/fP3jFihU88MADWK3WJu8iVgdJknC5
    XMiyTFRUFP369WPYsGF06dqFiPCqAqFydgmutHzEXccQdx1H3HMc6UABUpEdxe1xzxRvoFU4hXCn
    ku6kHr7nPTInI5c6NAY9mnA/dMmhGNpFomsXib5dJPqkMDShVbt/jx8/zubNm1m0aBFr164lNzcX
    rVZ71oW2jQW3243RaOStt97isssuo7y8vOTJJ598YcaMGUvwtO17awfPhzuo4ryTC7wEUzuXA4HQ
    55577spnnnnmaZPJZPr55595+umnkWX5grtDqlATlA6HA6PRSExMDB06dKBXr1706NmDFklJ6DWn
    fDe3gnSsFOloCVJ2Ce4jxUjZxUhHSpDyylHsIopNRHG6UcqcIHlusoJRCxYDgkWPYNSj8TOijfRF
    ExeILjYQbWwAurhAtFH+aCJ8/3msQGZmJhvWr2fdunVs376d7Oxsr5Vq6oGm2qDWsb7wwguMGzcO
    p9MpvvXWW68/++yzv+BZZxVxHtdZldEg5AIvwXR49OaDgLCZM2fedOutt96n0+m0M2fO5PXXX/fO
    Rr6QIcuyV7dBq9USGBhIq1ataNu2La1atSIlJYX4hASCg4Lq2BHglkBSQJS80mGCOgBc/VsHD4pL
    SsjJziY9PZ20tDR27drF7t27KSgo8N7ldTrdv+a8P/HEE9x9993Isqx89dVXH99yyy1zgDw8xLJx
    DqswakODkQuqBDgsQDAQ+euvv9519dVXjwd4+eWXmTlz5lk11DU1qOpCqmCqRqMhICCAqKgomjdv
    TkpKCtHR0URHRxMZFUlERAR+AQHohNP7/m5ZorSklLzcXI4dO8axY8c4evQo6enpHDx4kOPHj1Nc
    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"></image>
    </svg></center>

	<form method="GET" action="{{.ProxyPrefix}}/start">
	<input type="hidden" name="rd" value="{{.Redirect}}">
	{{ if .SignInMessage }}
	<p>{{.SignInMessage}}</p>
	{{ end}}
	<button type="submit" class="btn">Sign in with a {{.ProviderName}} Account</button><br/>
	</form>
	</div>

	{{ if .CustomLogin }}
	<div class="signin">
	<form method="POST" action="{{.ProxyPrefix}}/sign_in">
		<input type="hidden" name="rd" value="{{.Redirect}}">
		<label for="username">Username:</label><input type="text" name="username" id="username" size="10"><br/>
		<label for="password">Password:</label><input type="password" name="password" id="password" size="10"><br/>
		<button type="submit" class="btn">Sign In</button>
	</form>
	</div>
	{{ end }}
	<script>
		if (window.location.hash) {
			(function() {
				var inputs = document.getElementsByName('rd');
				for (var i = 0; i < inputs.length; i++) {
					inputs[i].value += window.location.hash;
				}
			})();
		}
	</script>
	<footer>
	{{ if eq .Footer "-" }}
	{{ else if eq .Footer ""}}
	Secured with <a href="https://github.com/openshift/oauth-proxy#oauth2_proxy">OAuth2 Proxy</a> version {{.Version}}
	{{ else }}
	{{.Footer}}
	{{ end }}
	</footer>
</body>
</html>
{{end}}`)
	if err != nil {
		log.Fatalf("failed parsing template %s", err)
	}

	t, err = t.Parse(`{{define "error.html"}}
<!DOCTYPE html>
<html lang="en" charset="utf-8">
<head>
	<title>{{.Title}}</title>
	<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
</head>
<body>
	<h2>{{.Title}}</h2>
	<p>{{.Message}}</p>
	<hr>
	<p><a href="{{.ProxyPrefix}}/sign_in">Sign In</a></p>
</body>
</html>{{end}}`)
	if err != nil {
		log.Fatalf("failed parsing template %s", err)
	}
	return t
}

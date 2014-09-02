go-httprequest
===============
A http request abstraction. Why? I kept having to write this layer over and over again. Sadly, I'm tired of having N versions/variations in various projects. I pray for this be the one http abstraction to rule them all ... for my projects.

Setup / Dependencies
--------------------
 $ go get github.com/mreiferson/go-httpclient
 $ go get github.com/heatxsink/go-httprequest

Documentation
-------------
If you would like to use http basic auth with your http requests use the function UseBasicAuth(bool) in conjunction with setting Username and Password public variables on the HttpRequest struct.

Also if you would like to use proxy your http requests use the function UseProxy(bool) and set the ProxyUrl public variable on the HttpRequest struct.

For everything else please take a look at httprequest_test.go.
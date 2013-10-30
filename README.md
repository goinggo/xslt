# Xslt

Copyright 2013 Ardan Studios. All rights reserved.  
Use of this source code is governed by a BSD-style license that can be found in the LICENSE handle.

This program provides an sample to learn how a process a xslt stylesheet against an xml document.
	
Ardan Studios  
12973 SW 112 ST, Suite 153  
Miami, FL 33186  
bill@ardanstudios.com

GoingGo.net Post:
http://www.goinggo.net/2013/10/using-xslt-with-go.html

	-- Get, build and install the code
	export GOPATH=$HOME/goinggo
	go get github.com/goinggo/xslt
	
	-- Run the code
	cd $GOPATH/src/github.com/goinggo/xslt
	go build
	./xslt stylesheet.xslt deals.xml
// Copyright 1999-2017 Alibaba Group.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/sirupsen/logrus"

	. "github.com/duyanghao/eagle/proxy/global"
	_ "github.com/duyanghao/eagle/proxy/initializer"
)

func main() {

	runtime.GOMAXPROCS(4)

	logrus.Infof("start proxy param: %+v", G_CommandLine)

	logrus.Infof("launch proxy on port: %d", G_CommandLine.Port)

	var err error

	if G_UseHttps {
		err = http.ListenAndServeTLS(fmt.Sprintf(":%d", G_CommandLine.Port), G_CommandLine.CertFile, G_CommandLine.KeyFile, nil)
	} else {
		err = http.ListenAndServe(fmt.Sprintf(":%d", G_CommandLine.Port), nil)

	}

	if err != nil {
		fmt.Printf("%v", err)
		logrus.Fatal(err)
	}
}

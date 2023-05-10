// The MIT License (MIT)
//
// Copyright (c) 2020 Fiber
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//
// This file may have been modified by CloudWeGo authors. All CloudWeGo
// Modifications are Copyright 2022 CloudWeGo Authors.

package main

import (
	"context"
	"data1/shiba/keyauth"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"net/http"
	"os"
)

const BEARER string = "innovationlabchaogeyangconggejinjinjiewuyuanqingchenxiaolu"

func main() {
	err := os.Setenv("BEARER", BEARER)
	if err != nil {
		return
	}

	h := server.Default()
	h.Use(keyauth.New(
		keyauth.WithContextKey("Bearer"),
		//keyauth.WithKeyLookUp("query:token", ""),
		/*
			the default @ option.go:

				authScheme: "Bearer",
				contextKey: "token",
				keyLookup:  "header:" + consts.HeaderAuthorization,
		*/

		// The middleware is skipped when true is returned.
		keyauth.WithFilter(func(c context.Context, ctx *app.RequestContext) bool {
			//always return true so always skip...
			// always return false so never skip the mw(middleware)
			return false
		}),

		// It may be used to validate key.
		// If returns false or err != nil, then errorHandler is used.
		// Below is the errorHandler
		keyauth.WithValidator(func(ctx context.Context, requestContext *app.RequestContext, s string) (bool, error) {
			//s is the extracted token
			if s == os.Getenv("BEARER") {
				return true, nil
			} else {
				return false, keyauth.ErrMissingOrMalformedAPIKey
			}
		}),

		// It may be used to define a custom error.
		keyauth.WithErrorHandler(func(ctx context.Context, requestContext *app.RequestContext, err error) {
			requestContext.AbortWithMsg("reached WithErrorHandler", http.StatusBadRequest)
		}),
	))
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		value, _ := ctx.Get("Bearer")
		ctx.JSON(consts.StatusOK, utils.H{"pong you the bearer": value})
	})
	h.Spin()
}

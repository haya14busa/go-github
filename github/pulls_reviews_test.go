// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestPullRequestReviewService_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/pulls/1/reviews", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	reviews, _, err := client.PullRequestReview.List("o", "r", 1)
	if err != nil {
		t.Errorf("PullRequestReview.List returned error: %v", err)
	}

	want := []*PullRequestReview{{ID: Int(1)}, {ID: Int(2)}}
	if !reflect.DeepEqual(reviews, want) {
		t.Errorf("PullRequestReview.List returned %+v, want %+v", reviews, want)
	}
}

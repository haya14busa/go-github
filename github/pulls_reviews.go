// Copyright 2016 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"fmt"
	"time"
)

// PullRequestReviewService handles communication with the pull request reviews
// related methods of the GitHub API.
//
// GitHub API docs: http://developer.github.com/v3/pulls/reviews
type PullRequestReviewService service

// PullRequestReview represents a review of a pull request.
type PullRequestReview struct {
	ID          *int       `json:"id,omitempty"`
	User        *User      `json:"user,omitempty"`
	Body        *string    `json:"body,omitempty"`
	SubmittedAt *time.Time `json:"submitted_at,omitempty"`

	// State can be "approved", "rejected", or "commented".
	State *string `json:"state,omitempty"`
}

// List the pull requests reveiws for the specified repository.
//
// GitHub API docs: https://developer.github.com/v3/pulls/reviews/#list-reviews-on-a-pull-request
func (s *PullRequestReviewService) List(owner string, repo string, number int) ([]*PullRequestReview, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/pulls/%d/reviews", owner, repo, number)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", mediaTypePullRequestReview)

	reviews := new([]*PullRequestReview)
	resp, err := s.client.Do(req, reviews)
	if err != nil {
		return nil, resp, err
	}

	return *reviews, resp, err
}

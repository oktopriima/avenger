/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 05/07/21, 15:09
 * Copyright (c) 2021
 */

package main

import "testing"

func Test_testError(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := testError(); err != nil {
				t.Errorf("Error found %v", err)
			}

			t.Log("success")
		})
	}
}

package main

import "testing"

func Test_main(t *testing.T) {
	t.Run("It should not panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("panicked")
			}
		}()
	})
}

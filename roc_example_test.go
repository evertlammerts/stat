// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stat_test

import (
	"fmt"

	"github.com/gonum/stat"
)

func ExampleROC_unweighted() {
	y := []float64{0, 3, 5, 6, 7.5, 8}
	classes := []bool{true, false, true, false, false, false}
	weights := []float64{4, 1, 6, 3, 2, 2}

	tpr, fpr := stat.ROC(0, y, classes, weights)
	fmt.Printf("true  positive rate: %v\n", tpr)
	fmt.Printf("false positive rate: %v\n", fpr)

	// Output:
	// true  positive rate: [0 0.4 0.4 1 1 1 1]
	// false positive rate: [0 0 0.125 0.125 0.5 0.75 1]
}

func ExampleROC_weighted() {
	y := []float64{0, 3, 5, 6, 7.5, 8}
	classes := []bool{true, false, true, false, false, false}

	tpr, fpr := stat.ROC(0, y, classes, nil)
	fmt.Printf("true  positive rate: %v\n", tpr)
	fmt.Printf("false positive rate: %v\n", fpr)

	// Output:
	// true  positive rate: [0 0.5 0.5 1 1 1 1]
	// false positive rate: [0 0 0.25 0.25 0.5 0.75 1]
}
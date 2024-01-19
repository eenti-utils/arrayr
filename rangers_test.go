package arrayr

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/eenti-utils/typr"
)

func TestRange_1(t *testing.T) {

	var testResult int

	assrtEqual(t, 0, testResult)

	sum := func(_ int, e int) (r typr.Op) {
		testResult += e
		return
	}

	Range(sum, From(1, 2, 3))

	assrtEqual(t, 6, testResult)
}

func TestRange_2(t *testing.T) {
	var testResult []string

	assrtEqual(t, 0, len(testResult))

	addToList := func(_ int, e string) (r typr.Op) {
		testResult = append(testResult, e)
		return
	}

	Range(addToList, From("apple", "banana", "coconut"))
	assrtEqual(t, 3, len(testResult))
	assrtEqual(t, []string{"apple", "banana", "coconut"}, testResult)
}

func TestRange_3(t *testing.T) {
	var testResult []string

	assrtEqual(t, 0, len(testResult))

	addToList := func(_ int, e string) (r typr.Op) {
		testResult = append(testResult, e)
		return
	}

	Range(addToList, From("apple", "banana", "coconut"), RangeOpts[string]{Step: -1}) // negative step means range in reverse order
	assrtEqual(t, 3, len(testResult))
	assrtEqual(t, []string{"coconut", "banana", "apple"}, testResult)
}

func TestRange_4(t *testing.T) {
	var testResult []string

	assrtEqual(t, 0, len(testResult))

	addToList := func(_ int, e string) (r typr.Op) {
		testResult = append(testResult, e)
		return
	}

	testFilter := func(s string) bool {
		return strings.Contains(s, "a")
	}

	testValidator := func(s string) (e error) {
		if strings.HasPrefix(s, "c") {
			return
		}
		e = fmt.Errorf("invalid element [ %s ]", s)
		return
	}

	Range(addToList, From("apple", "banana", "coconut"), RangeOpts[string]{Step: -1, FilterElements: testFilter, ValidateElements: testValidator}) // nothing should get through...
	assrtEqual(t, 0, len(testResult))
}

func TestRange_5(t *testing.T) {
	var testResult []string

	assrtEqual(t, 0, len(testResult))

	addToList := func(_ int, e string) (r typr.Op) {
		testResult = append(testResult, e)
		return
	}

	testFilter := func(s string) bool {
		return strings.Contains(s, "a")
	}

	Range(addToList, From("apple", "banana", "coconut"), RangeOpts[string]{Step: -1, FilterElements: testFilter})
	assrtEqual(t, 2, len(testResult))
	assrtEqual(t, []string{"banana", "apple"}, testResult)
}

func TestRange_6(t *testing.T) {
	var testResult []string

	assrtEqual(t, 0, len(testResult))

	addToList := func(_ int, e string) (r typr.Op) {
		testResult = append(testResult, e)
		return
	}

	testValidator := func(s string) (e error) {
		if strings.HasPrefix(s, "c") {
			return
		}
		e = fmt.Errorf("invalid element [ %s ]", s)
		return
	}

	Range(addToList, From("apple", "banana", "coconut"), RangeOpts[string]{Step: -1, ValidateElements: testValidator})
	assrtEqual(t, 1, len(testResult))
	assrtEqual(t, []string{"coconut"}, testResult)
}

func TestRange_7(t *testing.T) {
	var testResult []string

	assrtEqual(t, 0, len(testResult))

	addToList := func(_ int, e string) (r typr.Op) {
		testResult = append(testResult, e)
		return
	}

	Range(addToList, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}, RangeOpts[string]{Step: -2})
	assrtEqual(t, 5, len(testResult))
	assrtEqual(t, []string{"10", "8", "6", "4", "2"}, testResult)
}

func TestRange_Concurrently(t *testing.T) {
	var testResult []string

	assrtEqual(t, 0, len(testResult))

	ch := make(chan string)

	go func() {
		/***********************************************
		 * pop items off of the channel one at a time, *
		 * and add them to testResult                  *
		 * the items will show up, at some point ...   *
		 ***********************************************/
		var item string
		itemOK := true
		for itemOK {
			if item, itemOK = <-ch; itemOK {
				testResult = append(testResult, item)
			}
		}
	}()

	addToListSafely := func(_ int, e string) (r typr.Op) {
		// just send whatever's received to the channel
		// it's the safest bet ...
		ch <- e
		return
	}

	options := RangeOpts[string]{
		Step:         -2,   // submit every other element to the Ranger function in "reverse order", starting with the last element
		Concurrently: true, // process elements concurrently
	}

	Range(
		addToListSafely,	// the Ranger function
		[]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},	// the array
		options,
	)

	close(ch)	// finished with the channel

	time.Sleep(1 * time.Second)	// tic toc tic
	// *hopefully* the last elements have been added to testResult, by now 
	// with time, and concurrency, it's always a coin-toss

	assrtEqual(t, 5, len(testResult))
	testResultStr := strings.Join(testResult, " ") + " "
	assrtTrue(t, strings.Contains(testResultStr, "2 "))
	assrtTrue(t, strings.Contains(testResultStr, "4 "))
	assrtTrue(t, strings.Contains(testResultStr, "6 "))
	assrtTrue(t, strings.Contains(testResultStr, "8 "))
	assrtTrue(t, strings.Contains(testResultStr, "10 "))

	t.Logf("testResult == %v", testResult)
}

package cache

import (
	"londonDaily/trains"
	"reflect"
	"testing"
)

func TestCache(t *testing.T) {
	initScruct := struct{ testing int }{testing: 10}
	t.Run("Test New", func(t *testing.T) {
		c := New(10, initScruct)
		if c.cacheExpire != 10 {
			t.Errorf("Expected cacheExpire to be 10, got %d", c.cacheExpire)
		}
		if c.Result != initScruct {
			t.Errorf("Expected result to be 'response', got %s", c.Result)
		}
	})

	t.Run("Test isValid", func(t *testing.T) {
		c := New(10, struct{}{})
		if c.isValid() {
			t.Error("Expected cache to be invalid")
		}
	})

	t.Run("Test update", func(t *testing.T) {
		c := New(10, initScruct)

		trainStruct := trains.TrainResponse{
			Filter:   nil,
			Services: nil,
		}

		c.Update(trainStruct)
		result, ok := c.Result.(trains.TrainResponse)
		if !ok {
			t.Errorf("Expected result to be of type trains.TrainResponse, got %T", c.Result)
		} else if !reflect.DeepEqual(result, trainStruct) {
			t.Errorf("Expected result to be %v, got %v", trainStruct, result)
		}

		if !c.isValid() {
			t.Error("Expected cache to be valid")
		}
	})
}

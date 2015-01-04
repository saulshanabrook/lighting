package core

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func assertNoError(err error) {
	So(err, ShouldEqual, nil)
	return
}

type dummyState struct {
	dimmers DimmerMap
}

func (d dummyState) DimmersAt() (dimmers DimmerMap, err error) {
	dimmers = d.dimmers
	return
}

func TestStateStack(t *testing.T) {

	Convey("Given no states", t, func() {

		Convey("It should return an empty state", func() {
			dimmers, err := StateStack{[]State{}}.DimmersAt()
			assertNoError(err)
			So(dimmers, ShouldResemble, DimmerMap{})
		})

	})

	Convey("For one state", t, func() {

		Convey("Should return an equivalent state", func() {
			dimmers, err := StateStack{
				[]State{
					dummyState{DimmerMap{1: 0}},
				}}.DimmersAt()
			assertNoError(err)
			So(dimmers, ShouldResemble, DimmerMap{1: 0})
		})

	})

	Convey("For multiple states", t, func() {

		Convey("Should combine those states", func() {
			dimmers, err := StateStack{
				[]State{
					dummyState{DimmerMap{1: 0}},
					dummyState{DimmerMap{2: 0}},
				}}.DimmersAt()
			assertNoError(err)
			So(dimmers, ShouldResemble, DimmerMap{1: 0, 2: 0})
		})

		Convey("Later one should take precedence", func() {
			dimmers, err := StateStack{
				[]State{
					dummyState{DimmerMap{1: 0}},
					dummyState{DimmerMap{1: 100}},
				}}.DimmersAt()
			assertNoError(err)
			So(dimmers, ShouldResemble, DimmerMap{1: 100})
		})

	})

}

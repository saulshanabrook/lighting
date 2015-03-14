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

func TestDimmer(t *testing.T) {

	Convey("When it is valid", t, func() {

		Convey("DimmersAt should be just that dimmer", func() {
			dimmers, err := Dimmer{1, 50}.DimmersAt()
			assertNoError(err)
			So(dimmers, ShouldResemble, DimmerMap{1: 50})
		})
	})

	Convey("When level is too high", t, func() {

		Convey("Should return an error", func() {
			_, err := Dimmer{1, 200}.DimmersAt()
			So(err, ShouldNotBeNil)

		})

	})

	Convey("When level is too low", t, func() {

		Convey("Should return an error", func() {
			_, err := Dimmer{1, -1}.DimmersAt()
			So(err, ShouldNotBeNil)

		})

	})
}

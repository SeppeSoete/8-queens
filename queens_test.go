package main

import(
    "testing"
    "github.com/SeppeSoete/8-queens/board"
)

func TestHasQueen(t *testing.T) {
    for r := 0; r < 8; r++ {
        for c := 0; c < 8; c++ {
            var b board.Board
            b = b.PlaceQueen(r,c)
            if !b.HasQueen(r,c) {
                t.Errorf("HasQueen does not find a queen on row %d and col %d.",r,c)
            }
            for rt := 0; rt < 8; rt++ {
                for ct := 0; ct < 8; ct++ {
                    if rt != r && ct != c {
                        if b.HasQueen(rt,ct){
                            t.Error("HasQueen found a queen that doesn't exist.",r,c)
                        }
                    }
                }
            }
        }
    }
}

func TestToString(t *testing.T) {
    {
        var b board.Board
        b = b.PlaceQueen(0,0)
        expected := "Q|X|X|X|X|X|X|X|\n"
        expected += "----------------\n"
        expected += "X|X|X|X|X|X|X|X|\n"
        expected += "----------------\n"
        expected += "X|X|X|X|X|X|X|X|\n"
        expected += "----------------\n"
        expected += "X|X|X|X|X|X|X|X|\n"
        expected += "----------------\n"
        expected += "X|X|X|X|X|X|X|X|\n"
        expected += "----------------\n"
        expected += "X|X|X|X|X|X|X|X|\n"
        expected += "----------------\n"
        expected += "X|X|X|X|X|X|X|X|\n"
        expected += "----------------\n"
        expected += "X|X|X|X|X|X|X|X|\n"
        expected += "----------------\n"
        if expected != b.ToString() {
            t.Error("Output is not as expected")
        }
    }
    {
        var b board.Board
        b = b.PlaceQueen(0,7)
        expected := "X|X|X|X|X|X|X|Q|\n"
        expected += "----------------\n"
        expected += "X|X|X|X|X|X|X|X|\n"
        expected += "----------------\n"
        expected += "X|X|X|X|X|X|X|X|\n"
        expected += "----------------\n"
        expected += "X|X|X|X|X|X|X|X|\n"
        expected += "----------------\n"
        expected += "X|X|X|X|X|X|X|X|\n"
        expected += "----------------\n"
        expected += "X|X|X|X|X|X|X|X|\n"
        expected += "----------------\n"
        expected += "X|X|X|X|X|X|X|X|\n"
        expected += "----------------\n"
        expected += "X|X|X|X|X|X|X|X|\n"
        expected += "----------------\n"
        if expected != b.ToString() {
            t.Error("Output is not as expected")
        }
    }
    {
        var b board.Board
        b = b.PlaceQueen(2,3)
        expected := "X|X|X|X|X|X|X|X|\n"
        expected += "----------------\n"
        expected += "X|X|X|X|X|X|X|X|\n"
        expected += "----------------\n"
        expected += "X|X|X|Q|X|X|X|X|\n"
        expected += "----------------\n"
        expected += "X|X|X|X|X|X|X|X|\n"
        expected += "----------------\n"
        expected += "X|X|X|X|X|X|X|X|\n"
        expected += "----------------\n"
        expected += "X|X|X|X|X|X|X|X|\n"
        expected += "----------------\n"
        expected += "X|X|X|X|X|X|X|X|\n"
        expected += "----------------\n"
        expected += "X|X|X|X|X|X|X|X|\n"
        expected += "----------------\n"
        if expected != b.ToString() {
            t.Error("Output is not as expected")
        }
    }
}
func TestPlaceQueen(t *testing.T) {
    var b board.Board
    b = b.PlaceQueen(0,0)
    if !b.HasQueen(0,0) {
        t.Error("HasQueen does not find the queen placed by PlaceQueen.")
    }
}

func TestEquals(t *testing.T) {
    var a,b board.Board
    a = a.PlaceQueen(1,2)
    b = b.PlaceQueen(1,2)
    if !b.Equals(a) || !a.Equals(b) {
        t.Error("Equals returns false on two identical boards.")
    }
}

func TestCanPlaceQueen(t *testing.T) {
    //no conflicts
    {
        var b board.Board
        b = b.PlaceQueen(0,0)
        if !b.CanPlaceQueen(1,5){
            t.Error("TestCanPlaceQueen gives a false negative with a queen on (0,0) for coordinates (1,5).\n")
        }
        if !b.CanPlaceQueen(3,1){
            t.Error("TestCanPlaceQueen gives a false negative with a queen on (0,0) for coordinates (3,1).\n")
        }
        if b.CanPlaceQueen(0,0){
            t.Error("You can place a queen where there already is one.")
        }
    }

    //diagonal conflict
    {
        var b board.Board
        b = b.PlaceQueen(3,3)
        if b.CanPlaceQueen(2,2) ||
            b.CanPlaceQueen(4,4) ||
            b.CanPlaceQueen(4,2) ||
            b.CanPlaceQueen(2,4) {
            t.Error("TestCanPlaceQueen gives a false positive for diagonals.")
        }
    }
    //Horizontal conflict
    for r := 0; r < 8; r++ {
        for c := 0; c < 8; c++ {
            var b board.Board
            b = b.PlaceQueen(r,c)
            for i:= 0; i < 8; i++ {
                if b.CanPlaceQueen(r, i){
                    t.Error("TestCanPlaceQueen gives a false positive for the same row.")
                }
            }
        }
    }

    //Vertical conflict
    for r := 0; r < 8; r++ {
        for c := 0; c < 8; c++ {
            var b board.Board
            b = b.PlaceQueen(r,c)
            for i:= 0; i < 8; i++ {
                if b.CanPlaceQueen(i, c){
                    t.Error("TestCanPlaceQueen gives a false positive for the same collumn.")
                }
            }
        }
    }

}

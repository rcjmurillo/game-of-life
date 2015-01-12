## Game of Life

Simple version of the [Game of Life](http://en.wikipedia.org/wiki/Conway%27s_Game_of_Life)
created to improve my knowledge of the Go language.

This version is terminal oriented, every generation is printed to stdout.

## How to use it

Just run:

        go run main.go

Or import the game package:

        import "github.com/rcjmurillo/game-of-life/game"

## Some patterns

        // Pattern 1
        g := game.NewGame(15, 15)
        g.Seed([]game.Point{{6, 7}, {7, 5}, {7, 6}, {7, 8}, {7, 9}, {8, 7}})

        // Pattern 2
        g := game.NewGame(20, 20)
        g.Seed([]game.Point{
            {7, 7}, {7, 8}, {7, 9}, {7, 10}, {7, 11}, {7, 12},
            {8, 7}, {8, 12},
            {9, 7}, {9, 8}, {9, 9}, {9, 10}, {9, 11}, {9, 12},
            {10, 7}, {10, 12},
            {11, 7}, {11, 8}, {11, 9}, {11, 10}, {11, 11}, {11, 12},
            {12, 7}, {12, 12},
            {13, 7}, {13, 8}, {13, 9}, {13, 10}, {13, 11}, {13, 12},
        })

        // Pattern 3
        g := game.NewGame(21, 21)
        g.Seed([]game.Point{
            {7, 11}, {8, 10}, {8, 12},
            {9, 9}, {9, 13},
            {10, 9}, {10, 13},
            {11, 9}, {11, 13},
            {13, 11}, {12, 10}, {12, 12},
        })

        // Pattern 4
        g := game.NewGame(41, 41)
        g.Seed([]game.Point{
            {17, 21}, {18, 20}, {18, 22}, {19, 21},
            {19, 19}, {19, 23},
            {20, 19}, {20, 23},
            {21, 19}, {21, 23},
            {23, 21}, {22, 20}, {22, 22}, {21, 21},
        })

## License

Game of Life is released under the [MIT License](http://www.opensource.org/licenses/MIT).

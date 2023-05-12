Simple Golang Vote
==================

This is a simple "fullstack" voting webapp that I create to learn golang. 

## App description
The feature should be pretty simple:

1. Assume this website only handle a single poll. For example, polling about your favorite fruits.
2. The data would be stored as a SQLite table. The table named 'vote' would store the actual voted value.
3. It would consist of a single page: the form to pick your vote, and to display the resulting votes.
4. For the frontend side, it should be a simple HTML page, styled with [MVP.css](https://andybrewer.github.io/mvp/).
5. For the backend, it would use [Echo Framework](https://echo.labstack.com/).

## How to use

1. [Install go](https://go.dev/doc/install).
2. Clone this repo, cd to the folder
3. run `go run main.go`. 

## Future challenge
- User auth. Detect if the user has already voted. The simple implementation maybe to just store the IP of the user
- Multiple voting campaign, editable voting choices
- Multiple answer on a single vote

## License
MIT License

Copyright © 2023 Abdurrahman Shofy Adianto

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

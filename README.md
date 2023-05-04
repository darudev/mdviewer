# Markdown Viewer

Simple tool to parse and render Markdown files. And then make the HTML output available to view on a local web server.


### Install and usage

Requires Golang compiler.

Clone repo.

Install dependencies

    go mod tidy

Run

    go run . <path-to-file.md>

or build a binary and run

    go build .

    ./mdviewer <path-to-file.md>

Browse to http://localhost:8888

Now you can edit the markdown and just refresh browser to see changes.


### Change port etc.

Modify code and rebuild. :)


### Security 

NONE! Only for local usage!


# `make_blog` - static blog generator

This Git repository contains a static blog generator written for my personal
website.

## Building

No external libraries are needed -- just the Go standard library. Run `go build`
to build the program, then `./make_blog` (no arguments) to generate the static
site from the templates and articles in `tmpl/` and `articles/`.

The program regenerates all articles on every build. Unless your blog has
thousands of articles, full builds should take a few seconds at most, so there
is no need for incremental builds.

## Usage

Put articles in `articles/`. The format is `${DATE}_${IND}_${SLUG}.html`, where
`DATE` is the date of publication in ISO 8601 format, `IND` is a two-character
string used to order articles that were published on the same date, and `SLUG`
is a "slug" name used in the article's URL.

There are three template files in `tmpl/`, which generate the HTML index, RSS
index, and individual articles.

## License

Copyright (c) 2024 Charles Hood <chood@chood.net>

This program is free software: you can redistribute it and/or modify it under
the terms of the GNU General Public License as published by the Free Software
Foundation, either version 3 of the License, or (at your option) any later
version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY
WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A
PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with
this program. If not, see <https://www.gnu.org/licenses/>.

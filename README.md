![TMSU](http://tmsu.org/images/tmsu.png)

[ ![Build status](https://codeship.com/projects/ad81c060-5f0e-0132-2efc-3643fcd47fc7/status?branch=master)](https://codeship.com/projects/51490)

Overview
========

TMSU is a tool for tagging your files. It provides a simple command-line utility
for applying tags and a virtual filesystem to give you a tag-based view of your
files from any other program.

TMSU does not alter your files in any way: they remain unchanged on disk, or on
the network, wherever your put them. TMSU maintains its own database and you
simply gain an additional view, which you can mount where you like, based upon
the tags you set up.

Usage
=====

You can tag a file by specifying the file and the list of tags to apply:

    $ tmsu tag banana.jpg fruit art year=2015

Or you can apply tags to multiple files:

    $ tmsu tag --tags="fruit still-life art" banana.jpg apple.png

You can query for files with or without particular tags:

    $ tmsu files fruit and not still-life
    
A subcommand overview and detail on how to use each subcommand is available via the
integrated help:

    $ tmsu help
    $ tmsu help tags

Documentation is maintained online on the wiki:

  * <https://github.com/oniony/TMSU/wiki>

Installing
==========

Binary builds for a limited number of architectures and operating system
combinations are available:

  * <https://github.com/oniony/TMSU/releases>

(If you would rather build from the source code then please see `COMPILING.md`
in the root of the repository.)

You will need to ensure that both FUSE and Sqlite3 are installed for the
program to function. These packages are typically available with your
operating system's package management system.

1. Install the binary

    Copy the program binary. The location may be different for your operating
    system:

        $ sudo cp bin/tmsu /usr/bin

2. Optional: Zsh completion

    Copy the Zsh completion file to the Zsh site-functions directory:

        $ cp misc/zsh/_tmsu /usr/share/zsh/site-functions

About
=====

TMSU itself is written and maintained by [Paul Ruane](mailto:Paul Ruane <paul@tmsu.org>).

The creation of TMSU is motivation in itself, but if you should feel inclinded
to make a small gift via Bitcoin then it shall be gratefully received:

  * `1TMSU5TL3Yj6AGP7Wq6uahTfkTSX2nWvM`

TMSU is written in Go: <http://www.golang.org/>

Much of the functionality the program provides is made possible by the FUSE and
Sqlite3 libraries, their Go bindings and the Go language standard library.

  * Website: <http://tmsu.org/>
  * Project: <https://github.com/oniony/TMSU/>
  * Wiki: <https://github.com/oniony/TMSU/wiki>
  * Issue tracker: <https://github.com/oniony/TMSU/issues>
  * Mailing list: <http://groups.google.com/group/tmsu>

Release Notes
=============

v0.6 (in development)
----

  * Added --force option to 'tag' command to allow tagging of missing or
    permission denied paths and broken symlinks.
  * 'imply' now creates tags if necessary (and 'autoCreateTags' is set).
  * Performance improvements to the virtual filesystem.
  * Fixed 'too many SQL variables' when merging tags applied to lots of files.
  * Added --name option to 'tags' to force printing of name even if there is
    only a single file argument, which is useful when using xargs.

v0.5.2
------

  * Fixed bug where concurrent access to the virtual filesystem would cause
    a runtime panic.

v0.5.1
------

  * Fixed bug with database initialization when .tmsu directory does not
    already exist.

v0.5.0
------

  *Note: This release has some important changes, including the renaming of
  some options, the introduction of local databases and a switch from absolute
  to relative paths in the database. Please read the following release notes
  carefully.*

  * The --untagged option on the 'files' and 'status' subcommands has been
    replaced by a new 'untagged' subcommand, which should be more intuitive.
  * The --all option on the 'files', 'tags' and 'values' subcommands has been
    removed. These commands now list the full set of files/tags/values when run
    without arguments. For the 'tags' subcommand this replaces the previous
    behaviour of listing tags for the files in the working directory: use 'tmsu
    tags *' for approximately the previous behaviour.
  * The 'repair' subcommand --pretend short option has changed from -p to -P (so
    that -p can be recycled for --path).
  * The 'repair' subcommand's argument now specify paths to search for moved
    files and no longer limit how much of the database is repaired. A new --path
    argument is provided for reducing the repair to a portion of the database.
  * A new --manual option on the 'repair' subcommand allows targetted repair of
    moved files or directories.
  * The exclamation mark character (!) is no longer permitted within a tag or
    value name. Please rename tags using the 'rename' command. (Value names will
    need to be updated manually using the Sqlite3 tooling.)
  * Added --colour option to the 'tags' subcommand to highlight implied tags.
  * 'tag' subcommand will, by default, no longer explicitly apply tags that are
    already implied (unless the new --explicit option is specified).
  * Added subcommand aliases, e.g. 'query' for 'files'.
  * It is now possible to tag a broken symbolic link: instead of an error this
    will now be reported as a warning.
  * It is now possible to remove tags with values via the VFS.
  * 'tag' subcommand can tag multiple files with different tags by reading from
    standard input by passing an argument of '-'.
  * TMSU will now automatically use a local database in .tmsu/db in working
    directory or any parent. The new 'init' subcommand allows a new local
    database to be initialized. See [Switching Databases](https://github.com/oniony/TMSU/wiki/Switching%20Databases).
  * Paths are now stored relative to the .tmsu directory's parent rather than as
    absolute paths. This allows a branch of the filesystem to be moved around,
    shared or archived whilst preserving the tagging information. Existing
    absolute paths can be converted by running a manual repair:

        tmsu repair --manual / /

  * Added 'config' subcommand to view and amend settings.
  * The 'help' subcommand now wraps textual output to fit the terminal.
  * Rudimentary Microsoft Windows support (no virtual filesystem yet).
  * TMSU can now be built without the Makefile.
  * Bug fixes.

v0.4.3
------

  * Fixed unit-test problems.

v0.4.2
------

  * Fixed bug where 'dynamic:MD5' and 'dynamic:SHA1' fingerprint algorithms
    were actually using SHA256.

v0.4.1
------

  * Tag values are now shown as directories in the virtual filesystem.

v0.4.0
------

  *Note: This release changes the database schema to facilitate tag values. To
  upgrade your existing v0.3.0 database please run the following:*

    $ cp ~/.tmsu/default.db ~/.tmsu/default.db.backup
    $ sqlite3 -init misc/db-upgrade/0.3_to_0.4.0.sql ~/.tmsu/default.db .q

  * Added support for tag values, e.g. 'tmsu tag song.mp3 country=uk' and the
    querying of files based upon these values, e.g. 'year > 2000'.
  * 'tags' and 'values' subcommands now tabulate output, by default, when run
    from terminal.
  * Added ability to configure which fingerprint algorithm to use.
  * Implied tags now calculated on-the-fly when the database is queried. This
    results in a (potentially) smaller database and ability to have updates to
    the implied tags affect previously tagged files.
  * Added --explicit option to 'files' and 'tags' subcommands to show only
    explicit tags (omitting any implied tags).
  * Added --path option to 'files' subcommand to retrieve just those files
    matching or under the path specified.
  * Added --untagged option to 'files' subcommand which, when combined with
    --path, will also include untagged files from the filesystem at the
    specified path.
  * Removed the --recursive option from the 'files' subcommand which was flawed:
    use 'tmsu files query | xargs find' instead.
  * Added ability to configure whether new tags and values are automatically
    created or not or a per-database basis.
  * Added --unmodified option to 'repair' subcommand to force the recalculation
    of fingerprints of unmodified files.
  * Renamed --force option of 'repair' subcommand to --remove.
  * Added support for textual comparison operators: 'eq', 'ne', 'lt', 'gt',
    'le' and 'ge', which do not need escaping unlike '<', '>', &c.
  * Improved Zsh completion with respect to tag values.
  * Significant performance improvements.
  * Removed support for '-' operator: use 'not' instead.
  * Bug fixes.

v0.3.0
------

  *Note: This release changes what tag names are allowed. To ensure the tag
  names in your existing databases are still valid, please run the following
  script:*

    $ cp ~/.tmsu/default.db ~/.tmsu/default.db.backup
    $ sqlite3 -init misc/db-upgrade/clean_tag_names.sql ~/.tmsu/default.db

  * Added support for file queries, e.g. 'fish and chips and (mushy-peas or
    ketchup)'.
  * Added support for file query directories in the virtual filesystem.
  * Added global option --database for specifying database location.
  * Added ability to rename and delete tags via the virtual filesystem.
  * 'tag' subcommand now allows tags to be created up front.
  * 'copy' and 'imply' subcommands now support multiple destination tags.
  * Improved 'stats' subcommand.
  * Added man page.
  * Added script to allow the virtual filesystem to be mounted via the
    system 'mount' command or on startup via the fstab.
  * Bug fixes.

v0.2.2
------

  * Fixed virtual filesystem.

v0.2.1
------

  * Fixed bug where excluding multiple tags would return incorrect results.
  * Fixed Go 1.1 compilation problems. 

v0.2.0
------

  * Added support for tag implications, e.g. tag 'a' implies 'b'. New 'imply'
    subcommand for managing these.
  * Added --force option to 'repair' subcommand to remove missing files (and
    associated taggings) from the database.
  * Added --from option to 'tag' subcommand to allow tags to copied from one
    file to another. e.g. 'tmsu tag -f a b' will apply file b's tags to file a.
    ('tag -r -f a a' will recursively retag a directory's contents.)
  * Added --directory option to 'status' subcommand to stop it recursively
    processing directory contents.
  * Added --print0 option to 'files' subcommand to allow use with xargs.
  * Added --count option to 'tags' and 'files' subcommand to list tag/file count
    rather than names.
  * Bug fixes and unit-test improvements.

- - -

Copyright 2011-2015 Paul Ruane

Copying and distribution of this file, with or without modification,
are permitted in any medium without royalty provided the copyright
notice and this notice are preserved.  This file is offered as-is,
without any warranty.

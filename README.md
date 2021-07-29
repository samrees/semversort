# semversort

This project is an easy CLI utility for bash/shell/etc to deal with semantic versions.

Its the dumbest, least-required-maintenance possible wrapper around https://github.com/Masterminds/semver/,
a decent library.

If you see this hasn't been updated in multiple years, thats ok, at the time of writing this its 87 lines.
If you need it to do something extra, pull requests welcome.

This only exists because more people dont adopt: https://samver.org/ which is sortable by standard unix `sort`.

## Usage


```
Usage of ./semversort:
  -constraint string
    	list versions greatest to least, if versions pass given constraint.
  -greatest
    	display the greatest version for a given list
  -least
    	display the least version for a given list
```

example:

`echo -e "1.2.3\n3.4.5" | ./semversort -constraint='>=0.0.0'`

more info on semver constraints:

https://github.com/Masterminds/semver/tree/v3.1.1#basic-comparisons

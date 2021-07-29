# semversort

This project is an easy CLI utility for bash/shell/etc to deal with semantic versions.

Its the dumbest, least-required-maintenance possible wrapper around https://github.com/Masterminds/semver/,
a decent library.

If you see this hasn't been updated in multiple years, thats ok, at the time of writing this its 87 lines.
If you need it to do something extra, pull requests welcome.

This only exists because more people dont adopt: https://samver.org/ which is sortable by standard unix `sort`.

## Usage

sorts versions you pass to it over a pipe.

```
Usage of ./semversort:
  -constraint string
    	list versions only if versions pass given constraint
  -greatest
    	display the greatest version for a given list
  -least
    	display the least version for a given list
  -reverse
    	lists verions greatest to least
```

examples:

```
[~]> echo -e "1.2.3\n4.5.6\n2.9.100+woot\n0.3.1"
1.2.3
4.5.6
2.9.100+woot
0.3.1
[~]> echo -e "1.2.3\n4.5.6\n2.9.100+woot\n0.3.1" | ./semversort
0.3.1
1.2.3
2.9.100+woot
4.5.6
[~]> echo -e "1.2.3\n4.5.6\n2.9.100+woot\n0.3.1" | ./semversort -constraint='>=2.0.0'
2.9.100+woot
4.5.6
[~]> echo -e "1.2.3\n4.5.6\n2.9.100+woot\n0.3.1" | ./semversort -reverse
4.5.6
2.9.100+woot
1.2.3
0.3.1
[~]> echo -e "1.2.3\n4.5.6\n2.9.100+woot\n0.3.1" | ./semversort -greatest
4.5.6
```

more info on semver constraints:

https://github.com/Masterminds/semver/tree/v3.1.1#basic-comparisons
